package quiztool

import (
	"fmt"
	"log"
)

func (wsclient *WSClient) WriteRedis() (err error) {
	var n int
	command := wsclient.redisWrap.msg.Command + " " + wsclient.redisWrap.msg.Data
	b := []byte(command + "\r\n")

	if wsclient.redisWrap.Debug {
		log.Printf("VERBOSE Sending command: %s\n", command)
	}

	// Write socket
	n, err = wsclient.redisSock.Write(b)
	if err != nil {
		log.Printf("Error writing Redis: %v\n", err)
		return
	}
	if n != len(b) {
		err = fmt.Errorf("Redis socket wrote too few bytes %d (of %d)", n, len(b))
		return
	}

	return
}

func (wsclient *WSClient) ReadRedis() (data string, err error) {

	// Read socket
	wsclient.redisWrap.bufLen, err = wsclient.redisSock.Read(wsclient.redisWrap.buf)
	if err != nil {
		return
	}
	if wsclient.redisWrap.bufLen == 0 {
		err = fmt.Errorf("Empty read from Redis socket")
		return
	}

	data, err = wsclient.redisWrap.ParseSocket()
	if err != nil {
		log.Printf("Error from Redis: %v\n", err)
		return
	}

	if wsclient.redisWrap.Debug {
		log.Printf("VERBOSE Receiving Redis result: %s\n", data)
	}
	return
}