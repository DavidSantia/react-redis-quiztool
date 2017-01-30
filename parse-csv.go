package quiztool

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type QuizData struct {
	Debug    int
	Filename string
	Records  [][]string
}

func New(debug int) *QuizData {

	return &QuizData{Debug: debug}
}

func (qzt *QuizData) Parse(filename string) (err error) {

	log.Printf("Parsing CSV %s\n", filename)
	qzt.Filename = filename

	// Open file to parse
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	// Parse CSV header and data
	r := csv.NewReader(bufio.NewReader(file))
	qzt.Records, err = r.ReadAll()
	if err != nil {
		return err
	}
	file.Close()

	if qzt.Debug > 0 {
		log.Printf("DEBUG CSV header\n%#v\n", qzt.Records[0])
	}

	// length includes header
	total := len(qzt.Records)
	if total <= 1 {
		return fmt.Errorf("No records in CSV file %s\n", qzt.Filename)
	}

	if qzt.Debug > 0 {
		log.Printf("DEBUG Records to process: %d\n", total-1)
	}

	for i := 1; i < total; i++ {
		log.Printf("Record %d: %#v\n", i, qzt.Records[i])
	}

	return nil
}
