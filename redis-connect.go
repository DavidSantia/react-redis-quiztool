package quiztool

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/garyburd/redigo/redis"
)

func getServer() (redisHostPort, wsHostPort string) {
	var redisHost, redisPort, wsHost, wsPort string

	redisHost = os.Getenv("REDIS_HOST")
	if len(redisHost) == 0 {
		redisHost = "localhost"
	}
	redisPort = os.Getenv("REDIS_PORT")
	if len(redisPort) == 0 {
		redisPort = "6379"
	}

	wsHost = os.Getenv("WEBSOCKET_HOST")
	if len(wsHost) == 0 {
		wsHost = "localhost"
	}
	wsPort = os.Getenv("WEBSOCKET_PORT")
	if len(wsPort) == 0 {
		wsPort = "4000"
	}

	redisHostPort = redisHost + ":" + redisPort
	wsHostPort = wsHost + ":" +wsPort
	return
}

// Connect using socket only
func (qzt *QuizApp) ConnectSockets() (err error) {

	redisHostPort, wsHostPort := getServer()
	if qzt.Debug > 0 {
		log.Printf("DEBUG Connecting to Redis server at %s\n", redisHostPort)
	}

	qzt.wsclient = &WSClient{
		wshost:    wsHostPort,
		router:    make(map[string]Handler),
		redisWrap: &Wrapper{
			Debug: qzt.Debug > 1,
			buf:   make([]byte, 8192),
		},
		active:    false,
	}
	qzt.wsclient.AddRoutes()

	// Connect to Redis on TCP port
	qzt.wsclient.redisSock, err = net.Dial("tcp", redisHostPort)
	if err != nil {
		log.Printf("Redis Socket error: %v\n", err)
		return
	}

	// When HTTP port accessed, upgrade to Websocket
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(qzt.wsclient, w, r)
	})

	if qzt.Debug > 0 {
		log.Printf("DEBUG Initiating ListenAndServe at %s\n", wsHostPort)
	}
	err = http.ListenAndServe(wsHostPort, nil)
	if err != nil {
		log.Printf("WebSocket %q error: %v\n", wsHostPort, err)
		return
	}

	return
}

// Connect using Redis package
func (qzt *QuizApp) ConnectMain() (err error) {

	redisHostPort, _ := getServer()
	if qzt.Debug > 0 {
		log.Printf("DEBUG Connecting to Redis server at %s\n", redisHostPort)
	}

	qzt.redisConn, err = redis.Dial("tcp", redisHostPort)
	if err != nil {
		return err
	}

	return
}
