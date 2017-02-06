package quiztool

import (
	"github.com/garyburd/redigo/redis"
)

type QuizApp struct {
	Debug     int
	Filename  string
	Columns   int
	Records   [][]string
	Quiz      *Quiz
	RedisConn redis.Conn
}

type Quiz struct {
	Number  int
	Title   string
	Rows    int
	Columns int
	Map     []map[string]string
}

func New(debug int) *QuizApp {

	// Set to number of fields in CSV header
	const columns = 11

	return &QuizApp{Debug: debug, Columns: columns}
}
