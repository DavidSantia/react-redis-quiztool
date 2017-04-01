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
		Rows: len(qzt.Records) - 1,
		// Quiz Columns does not include title, so Columns-1
		Columns: qzt.Columns - 1,
	}

	if qzt.Debug > 0 {
		log.Printf("DEBUG Quiz: %#v\n", qzt.Quiz)
	}

	log.Printf("Records to process: %d\n", qzt.Quiz.Rows)
	qzt.Quiz.CatQuestions = make([][]Question, 1)

	// lookup for category list index
	lookup_index := make(map[string]int)

	// Compile list of Categories (start at 1 to skip CSV header)
	for i := 1; i <= qzt.Quiz.Rows; i++ {
		category := qzt.Records[i][1]

		// Did we find a new category?
		_, exist := lookup_index[category]
		if !exist {
			index := len(qzt.Quiz.Categories)
			lookup_index[category] = index

			// Compiling list of distinct categories
			qzt.Quiz.Categories = append(qzt.Quiz.Categories, category)
		}
	}

	// Initialize array of questions for each category
	qzt.Quiz.CatQuestions = make([][]Question, len(qzt.Quiz.Categories))

	// Map Quiz question data and store by Category
	for i := 1; i <= qzt.Quiz.Rows; i++ {
		category := qzt.Records[i][1]
		question := make(Question)

		// Convert Quiz question data to Map
		for j := 1; j <= qzt.Quiz.Columns; j++ {
			if len(qzt.Records[i][j]) != 0 {
				// Add non-null data to Map
				question[qzt.Records[0][j]] = qzt.Records[i][j]
			}
		}

		// Insert Question by category
		index := lookup_index[category]
		qzt.Quiz.CatQuestions[index] = append(qzt.Quiz.CatQuestions[index], question)
	}
}
