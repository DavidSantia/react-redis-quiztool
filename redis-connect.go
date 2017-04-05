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

// Connect to Redis using socket
func (qzt *QuizApp) ConnectRedisSocket() (err error) {

	redisHostPort, wsHostPort := getServer()
	if qzt.Debug > 0 {
		log.Printf("DEBUG Connecting to Redis socket at %s\n", redisHostPort)
	}

	qzt.wsclient = &WSClient{
		wshost:    wsHostPort,
		redisWrap: &Wrapper{
			Debug: qzt.Debug > 1,
			buf:   make([]byte, 8192),
		},
	}

	// Connect to Redis on TCP port
	qzt.wsclient.redisSock, err = net.Dial("tcp", redisHostPort)
	if err != nil {
		log.Printf("Redis Socket error: %v\n", err)
		return
	}
	return
}

// Initiate Websocket for messaging Redis
func (qzt *QuizApp) InitiateRedisWebocket() (err error) {

	if qzt.wsclient == nil {
		log.Printf("Need to call ConnectRedisSocket() first\n")
		return
	}

	// Handler that upgrades to websocket
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(qzt.wsclient, w, r)
	})

	if qzt.Debug > 0 {
		log.Printf("DEBUG Initiating ListenAndServe for websocket messaging at %s\n", qzt.wsclient.wshost)
	}

	// Listen and Serve HTTP (upgrades to websocket)
	err = http.ListenAndServe(qzt.wsclient.wshost, nil)
	if err != nil {
		log.Printf("WebSocket %q error: %v\n", qzt.wsclient.wshost, err)
		return
	}
	return
}

// Connect using Redis package
func (qzt *QuizApp) ConnectRedis() (err error) {

	redisHostPort, _ := getServer()
	if qzt.Debug > 0 {
		log.Printf("DEBUG Connecting to Redis server at %s\n", redisHostPort)
	}

	qzt.RedisConn, err = redis.Dial("tcp", redisHostPort)
	if err != nil {
		return err
	}

	return
}
