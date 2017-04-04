package quiztool

import (
	"fmt"
	"log"
)

func (wsclient *WSClient) WriteRedis(cmd, data string) (err error) {
	wsclient.redisWrap.Command = cmd
	command := cmd + " " + data
	if wsclient.redisWrap.Debug {
		log.Printf("VERBOSE Sending command: %s\n", command)
	}

	var n int
	b := []byte(command + "\r\n")

	// Write socket
	n, err = wsclient.redisSock.Write(b)
	if err != nil {
		log.Printf("Error writing Redis socket: %v\n", err)
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
		log.Printf("Error reading from Redis socket: %v\n", err)
		return
	}

	if wsclient.redisWrap.Debug {
		log.Printf("VERBOSE Receiving Redis result: %s\n", data)
	}
	return
}