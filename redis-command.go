package quiztool

import (
	"fmt"
	"log"
)

func (qzt *QuizApp) WriteSocket(cmd, data string) (err error) {
	qzt.RedisWrap.Command = cmd
	command := cmd + " " + data
	if qzt.Debug > 0 {
		log.Printf("Sending command: %s\n", command)
	}

	var n int
	b := []byte(command + "\r\n")

	// Write socket
	n, err = qzt.RedisSock.Write(b)
	if err != nil {
		return
	}
	if n != len(b) {
		err = fmt.Errorf("Redis socket wrote too few bytes %d (of %d)", n, len(b))
		return
	}

	return
}

func (qzt *QuizApp) ReadSocket() (data string, err error) {
	wrp := qzt.RedisWrap
	if qzt.Debug > 0 {
		log.Printf("Reading result\n")
	}

	// Read socket
	wrp.BufLen, err = qzt.RedisSock.Read(wrp.Buf)
	if err != nil {
		return
	}
	if wrp.BufLen == 0 {
		err = fmt.Errorf("Empty read from Redis socket")
		return
	}

    data, err = wrp.ParseSocket()
    return
}