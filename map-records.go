package quiztool

import (
	"log"
)

func (qzt *QuizApp) MapRecords() {

	// Quiz.Number is there if we want to increment when QuizTitle changes

	qzt.Quiz = &Quiz{
		// Assume for now there is just 1 quiz
		Number: 1,
		// Quiz Title is in First column of any record, use first
		Title: qzt.Records[1][0],
		// Quiz Rows does not include CSV header, so length(Records)-1
		Rows: len(qzt.Records)-1,
		// Quiz Columns does not include title, so Columns-1
		Columns: qzt.Columns-1,
	}

	if qzt.Debug > 0 {
		log.Printf("DEBUG Quiz: %#v\n", qzt.Quiz)
	}

	log.Printf("Records to process: %d\n", qzt.Quiz.Rows)
	qzt.Quiz.Map = make([]map[string]string, qzt.Quiz.Rows)

	// Convert CSV array of strings to Map, for Quiz question data
	for i := 0; i < qzt.Quiz.Rows; i++ {
		qzt.Quiz.Map[i] = make(map[string]string)
		for j := 1; j <= qzt.Quiz.Columns; j++ {
			// Add non-null data to Map
			if len(qzt.Records[i+1][j]) != 0 {
				qzt.Quiz.Map[i][qzt.Records[0][j]] = qzt.Records[i+1][j]
			}
		}
	}
}