package quiztool

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func (qzt *QuizApp) Parse(filename string) (err error) {

	log.Printf("Parsing CSV: \"%s\"\n", filename)
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

	// length includes header
	if len(qzt.Records) <= 1 {
		return fmt.Errorf("No records in CSV file %s\n", qzt.Filename)
	}

	return nil
}
