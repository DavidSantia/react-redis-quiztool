package quiztool

import (
	"log"

	"github.com/gorilla/websocket"
	"fmt"
)

func (wsclient *WSClient) startRespond() {

	log.Printf("** In startRespond\n")
	// channel for responses
	wschannel := make(chan WSMessage)
	go func() {
		var msg WSMessage
		msg.Command = "reply"
		for {
			log.Printf("** startRespond for loop\n")
			data, err := wsclient.ReadRedis()
			if err != nil {
				msg.Data = WSRedisCommand{Command: "error", Args: err.Error()}
			} else {
				msg.Data = WSRedisCommand{Command: "success", Args: data}
			}
			wschannel <- msg
		}
	}()

	for {
		select {
		case msg := <-wschannel:
			wsclient.send <- msg
		case <-wsclient.stop:
			log.Printf("<Disconnect>\n")
			return
		}
	}
}

func writeRedis(wsclient *WSClient, redisCommand WSRedisCommand) {
	if !wsclient.active {
		log.Printf("<Session not started, ignoring do>\n")
		return
	}
	if redisCommand.Command == "" {
		log.Printf("<Redis command is null>\n")
		return
	}

	args := fmt.Sprintf("%v", redisCommand.Args)
	if wsclient.redisWrap.Debug {
		log.Printf("Writing Redis with %s %s",  redisCommand.Command, args)
	}

	wsclient.WriteRedis(redisCommand.Command, args)
}

func subscribeRedis(wsclient *WSClient, none WSRedisCommand) {
	log.Printf("Start session %s\n", wsclient.SessionId)
	wsclient.active = true

	// Do responses in separate thread
	go wsclient.startRespond()
}

func unsubscribeRedis(wsclient *WSClient, none WSRedisCommand) {
	wsclient.active = false
	wsclient.stop <- true
}

func (wsclient *WSClient) WriteWS() {
	for msg := range wsclient.send {
		if wsclient.redisWrap.Debug {
			log.Printf("VERBOSE WriteWS %v\n", msg)
		}

		if err := wsclient.socket.WriteJSON(msg); err != nil {
			log.Printf("<Socket write error %v>\n", err)
			break
		}
	}
	wsclient.socket.Close()
}

func (wsclient *WSClient) ReadWS() {
	if wsclient.redisWrap.Debug {
		log.Printf("VERBOSE In function ReadWS\n")
	}

	var msg WSMessage
	for {
		err := wsclient.socket.ReadJSON(&msg)
		if err != nil {
			// Check for closed session
			if c, ok := err.(*websocket.CloseError); ok {
				log.Printf("<Socket connection closed (code %d)>\n", c)
				wsclient.stop <- true
				wsclient.socket.Close()
				return
			}
			log.Printf("<Message parse error %v>\n", err)
			continue
		}
		wsclient.Route(msg)
	}
}