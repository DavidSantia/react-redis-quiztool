package quiztool

import (
	"net"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"
)

type Handler func(*WSClient, WSRedisCommand)

type WSClient struct {
	SessionId string
	wshost    string
	socket    *websocket.Conn
	send      chan WSMessage
	stop      chan bool
	active    bool
	router    map[string]Handler
	redisWrap *Wrapper
	redisSock net.Conn
}

type WSRedisCommand struct {
	Command string  `json:"command"`
	Args    string  `json:"args"`
}

type WSMessage struct {
	Command string         `json:"command"`
	Data    WSRedisCommand `json:"data"`
}

type Wrapper struct {
	Debug   bool
	Command string
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
	redisConn redis.Conn
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
