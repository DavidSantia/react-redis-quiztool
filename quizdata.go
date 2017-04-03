package quiztool

import (
	"net"

	"github.com/garyburd/redigo/redis"
)

type Wrapper struct {
	Debug     bool
	Command   string
	KeyPair   bool
	Buf       []byte
	BufPtr    int
	BufLen    int
}

type QuizApp struct {
	Debug     int
	Filename  string
	Columns   int
	Records   [][]string
	Quiz      *Quiz
	RedisConn redis.Conn
	RedisSock net.Conn
	RedisWrap *Wrapper
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

	return &QuizApp{Debug: debug, Columns: columns}
}
