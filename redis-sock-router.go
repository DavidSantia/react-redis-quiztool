package quiztool

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// When HTTP port accessed, upgrade to Websocket
func serveWs(wsclient *WSClient, w http.ResponseWriter, r *http.Request) {
	var attempts int = 3
	var err error

	for i := 0; i < attempts; i++ {
		wsclient.socket, err = upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Failed Websocket: %v\n", err)
		} else {
			break
		}
	}
	if err != nil {
		// Exit on repeated fails
		return
	}

	log.Printf("<Websocket Ready>\n")

	// Start Websocket write thread
	go wsclient.WriteWS()

	// Use this thread for read
	wsclient.ReadWS()
	return
}

func (wsclient *WSClient) Route(msg WSMessage) {
	if wsclient.redisWrap.Debug {
		log.Printf("VERBOSE Route message %q\n", msg)
	}

	if msg.Command == "" {
		log.Printf("<Command is null>\n")
		return
	}

	// Lookup handler function
	route, found := wsclient.router[msg.Command]
	if found {
		route(wsclient, msg.Data)
	} else {
		log.Printf("<Unrecognized command %s>\n", msg.Command)
	}
}

func (wsclient *WSClient) Handle(command string, handler Handler) {
	wsclient.router[command] = handler
}

func (wsclient *WSClient) AddRoutes() {
	wsclient.Handle("start", subscribeRedis)
	wsclient.Handle("stop", unsubscribeRedis)
	wsclient.Handle("do", writeRedis)
}
