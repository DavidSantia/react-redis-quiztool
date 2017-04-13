package quiztool

import (
	"log"
	"strings"
	"testing"
)

var qzt1 *QuizApp

func TestConnectRedis(t *testing.T) {
	var err error

	// debug: 0 = off, 1 = on, 2 = verbose
	var debug int = 1
	qzt1 = New(debug)

	// Connect Test
	err = qzt1.ConnectRedis()
	if err != nil {
		t.Errorf("Expected Redis connection, got: %v\n", err)
	}
}

func TestParse(t *testing.T) {
	var err error
	var filename string = "data/plant-quiz.csv"

	// Parse Test
	err = qzt1.Parse(filename)
	if err != nil {
		t.Errorf("Expected CSV parsed, got: %v\n", err)
	}
}

func TestMapRecords(t *testing.T) {

	if qzt1.Quiz != nil {
		t.Errorf("Expected initial Quiz=nil got: %v\n", qzt1.Quiz)
	}
	
	qzt1.MapRecords()
	if qzt1.Quiz == nil {
		t.Errorf("Expected Quiz got nil\n")
	}

	if qzt1.Quiz.Number != 1 {
		t.Errorf("Expected Quiz Number 1, got: %d\n", qzt1.Quiz.Number)
	}

	if strings.Contains(qzt1.Quiz.Title, "All") {
		log.Printf("Got expected title: %s\n", qzt1.Quiz.Title)
	} else {
		t.Errorf("Expected to contain 'All', got result: %s\n", qzt1.Quiz.Title)
	}
}
