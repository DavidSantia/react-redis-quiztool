package quiztool

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"fmt"
)

func (wsclient *WSClient) Respond() {
	if wsclient.redisWrap.Debug {
		log.Printf("VERBOSE Starting Respond\n")
	}

	var msg WSMessage
	for {
		data, err := wsclient.ReadRedis()
		if err != nil {
			msg = WSMessage{Command: "error", Data: fmt.Sprintf("%q", err.Error())}
			log.Printf("<redis: %s>\n", err.Error())
		} else {
			msg = WSMessage{Command: "success", Data: data}
		}
		wsclient.WriteWS(msg)
	}
}

func (wsclient *WSClient) WriteWS(msg WSMessage) {
	if wsclient.redisWrap.Debug {
		log.Printf("VERBOSE WriteWS %v\n", msg)
	}

	// Create JSON format (msg.Data is already JSON)
	b := []byte(fmt.Sprintf("{command: %q, data: %s}", msg.Command, msg.Data))

	if err := wsclient.socket.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Printf("<websocket: write error %v>\n", err)
		return
	}
}

func (wsclient *WSClient) ReadWS() {
	if wsclient.redisWrap.Debug {
		log.Printf("VERBOSE Starting ReadWS\n")
	}

	var msg WSMessage
	for {
		err := wsclient.socket.ReadJSON(&msg)
		if err != nil {
			// Check for closed session
			if c, ok := err.(*websocket.CloseError); ok {
				log.Printf("<%s>\n", c)
				wsclient.socket.Close()
				return
			}
			log.Printf("<message: parse error %v>\n", err)
			continue
		}
		if wsclient.redisWrap.Debug {
			log.Printf("VERBOSE Route message %q\n", msg)
		}

		if msg.Command == "" {
			log.Printf("<message: command is null>\n")
			continue
		}

		wsclient.redisWrap.msg = msg
		wsclient.WriteRedis()
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// When HTTP port accessed, upgrade to Websocket
func serveWs(wsclient *WSClient, w http.ResponseWriter, r *http.Request) {
	var err error

	wsclient.socket, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("<failed upgrade %v>\n", err)
		return
	}

	log.Printf("<websocket: ready>\n")

	// Do responses in separate thread
	go wsclient.Respond()

	// Use this thread for read
	wsclient.ReadWS()
	return
}