package main

import (
	"fmt"
	"os"

	rrq "github.com/DavidSantia/react-redis-quiztool"
)

func main() {

	// debug: 0 = off, 1 = on, 2 = verbose
	var debug int = 0
	qzt := rrq.New(debug)

	err := qzt.ConnectMain()
	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}
	defer qzt.RedisConn.Close()

	var filename string = "data/plant-quiz.csv"
	err = qzt.Parse(filename)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}

	qzt.MapRecords()

	err = qzt.StoreQuiz()
	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
