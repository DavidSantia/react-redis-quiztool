package main

import (
	"fmt"
	"os"

	rrq "github.com/DavidSantia/react-redis-quiztool"
)

func main() {

	// debug: 0 = off, 1 = on, 2 = verbose
	var debug int = 2
	qzt := rrq.New(debug)

	fmt.Printf("Starting ConnectSockets\n")

	err := qzt.ConnectSockets()
	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
