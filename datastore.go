package quiztool

import (
	"log"
	"os"

	"fmt"
	"github.com/garyburd/redigo/redis"
)

func (qzt *QuizApp) ConnectDatastore() (err error) {

	host := os.Getenv("REDIS_HOST")
	if len(host) == 0 {
		host = "localhost"
	}
	port := os.Getenv("REDIS_PORT")
	if len(port) == 0 {
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

	qtag := fmt.Sprintf("quiz:%d", qzt.Quiz.Number)
	log.Printf("Storing %s \"%s\" [%d rows, %d columns, %d categories]\n",
		qtag, qzt.Quiz.Title, qzt.Quiz.Rows, qzt.Quiz.Columns, len(qzt.Quiz.Categories))

	// Store quiz information
	_, err = qzt.RedisConn.Do("HMSET", qtag, "title", qzt.Quiz.Title,
		"questions", qzt.Quiz.Rows, "categories", len(qzt.Quiz.Categories))
	if err != nil {
		return err
	}

	// Store quiz questions
	for i, category := range qzt.Quiz.Categories {
		ctag := fmt.Sprintf("%s:c:%d", qtag, i+1)
		log.Printf("Storing category %d: %s, %d questions\n",
			i+1, category, len(qzt.Quiz.CatQuestions[i]))

		_, err = qzt.RedisConn.Do("HMSET", ctag, "category", category,
			"questions", len(qzt.Quiz.CatQuestions[i]))
		if err != nil {
			return err
		}

		for j, question := range qzt.Quiz.CatQuestions[i] {
			qqtag := fmt.Sprintf("%s:q:%d", ctag, j+1)
			if qzt.Debug > 0 {
				log.Printf("DEBUG Store %#v\n", redis.Args{qqtag}.AddFlat(question))
			}
			_, err = qzt.RedisConn.Do("HMSET", redis.Args{qqtag}.AddFlat(question)...)
			if err != nil {
				return err
			}
		}
	}

	log.Printf("Data loaded into Redis\n")
	return nil
}
