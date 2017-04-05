package main

import (
	"fmt"
	"os"

	rrq "github.com/DavidSantia/react-redis-quiztool"
)

func main() {

	// debug: 0 = off, 1 = on, 2 = verbose
	var debug int = 1
	qzt := rrq.New(debug)

	err := qzt.ConnectRedisSocket()
	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}

	err = qzt.InitiateRedisWebocket()
	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
