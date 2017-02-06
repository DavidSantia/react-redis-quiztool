package quiztool

import (
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
	"fmt"
)

func (qzt *QuizApp) ConnectDatastore() (err error) {

	host := os.Getenv("REDIS_HOST"); if len(host) == 0 {
		host = "localhost"
	}
	port := os.Getenv("REDIS_PORT"); if len(port) == 0 {
		port = "6379"
	}
	server := host + ":" + port

	if qzt.Debug > 0 {
		log.Printf("DEBUG Connecting to Redis server at %s\n", server)
	}

	qzt.RedisConn, err = redis.Dial("tcp", server)
	if err != nil {
		return err
	}

	return nil
}

func (qzt *QuizApp) StoreQuiz() (err error) {

	quiz := fmt.Sprintf("quiz:%d", qzt.Quiz.Number)
	log.Printf("Storing %s \"%s\" [%d rows, %d columns]\n", quiz, qzt.Quiz.Title, qzt.Quiz.Rows, qzt.Quiz.Columns)

	// Store quiz information
	_, err = qzt.RedisConn.Do("HMSET", quiz, "title", qzt.Quiz.Title,
		"rows", qzt.Quiz.Rows, "columns", qzt.Quiz.Columns)
	if err != nil {
		return err
	}

	// Store quiz questions
	for i := 0; i < qzt.Quiz.Rows; i++ {
		question := fmt.Sprintf("%s:q:%d", quiz, i+1)
		if qzt.Debug > 0 {
			log.Printf("DEBUG Store %#v\n", redis.Args{question}.AddFlat(qzt.Quiz.Map[i]))
		}
		_, err = qzt.RedisConn.Do("HMSET", redis.Args{question}.AddFlat(qzt.Quiz.Map[i])...)
		if err != nil {
			return err
		}
	}

	log.Printf("Data loaded into Redis\n")
	return nil
}
