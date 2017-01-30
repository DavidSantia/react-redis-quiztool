# react-redis-quiztool
## Quiz Tool Demo using React, Go and Redis

To use this package, start with:
```sh
go get github.com/DavidSantia/react-redis-quiztool
```

Next, you will need a CSV file containing youre quiz data.

#### Example
A sample plant quiz is included.

1. Make a directory for your project
2. Save the file [plant-quiz.csv](https://raw.githubusercontent.com/DavidSantia/react-redis-quiztool/master/plant-quiz.csv) into your directory
3. Create a main.go that calls New, Parse and DoQuiz like follows:
```go
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

	var filename string = "plant-quiz.csv"
	err := qzt.Parse(filename)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(0)
	}

	os.Exit(1)
}
```
4. Run it:  **go run main.go**
