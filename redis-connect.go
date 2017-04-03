package quiztool

import (
	"log"
	"net"
	"os"

	"github.com/garyburd/redigo/redis"
)

func getServer() (server string) {
	var host, port string

	host = os.Getenv("REDIS_HOST")
	if len(host) == 0 {
		host = "localhost"
	}
	port = os.Getenv("REDIS_PORT")
	if len(port) == 0 {
		port = "6379"
	}

	return host + ":" + port
}

// Connect using socket only
func (qzt *QuizApp) ConnectLight() (err error) {

	server := getServer()

	if qzt.Debug > 0 {
		log.Printf("DEBUG Connecting to Redis server at %s\n", server)
	}

	qzt.RedisSock, err = net.Dial("tcp", server)
	if err != nil {
		return err
	}

	qzt.Buf = make([]byte, 8192)
	return
}

// Connect using Redis package
func (qzt *QuizApp) ConnectMain() (err error) {

	server := getServer()

	if qzt.Debug > 0 {
		log.Printf("DEBUG Connecting to Redis server at %s\n", server)
	}

	qzt.RedisConn, err = redis.Dial("tcp", server)
	if err != nil {
		return err
	}

	return
}
