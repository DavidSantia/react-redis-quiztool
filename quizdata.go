package quiztool

import (
	"net"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"
)

type WSClient struct {
	SessionId string
	wshost    string
	socket    *websocket.Conn
	redisWrap *Wrapper
	redisSock net.Conn
}

type WSMessage struct {
	Command string `json:"command"`
	Data    string `json:"data"`
}

type Wrapper struct {
	Debug   bool
	msg     WSMessage
	keyPair bool
	buf     []byte
	bufPtr  int
	bufLen  int
}

type QuizApp struct {
	Debug     int
	Filename  string
	Columns   int
	Records   [][]string
	Quiz      *Quiz
	RedisConn redis.Conn
	wsclient  *WSClient
}

type Question map[string]string

type Quiz struct {
	Number       int
	Title        string
	Rows         int
	Columns      int
	Categories   []string
	CatQuestions [][]Question
}

func New(debug int) *QuizApp {

	// Set to number of fields in CSV header
	const columns = 11

	return &QuizApp{
		Debug:   debug,
		Columns: columns,
	}
}
