package quiztool

import (
	"log"
	"strings"
	"testing"
)

var qzt *QuizApp

func TestConnect(t *testing.T) {
	var err error

	// debug: 0 = off, 1 = on, 2 = verbose
	var debug int = 1
	qzt = New(debug)

	// Passing Test
	err = qzt.ConnectLight()
	if err != nil {
		t.Errorf("Expected new quiztool, got: %v\n", err)
	}
}

func TestCommandsThatFail(t *testing.T) {
	var err error
	var cmd, data string

	cmd = "HGET"
	data = "quiz:1 titlen"
	err = qzt.WriteSocket(cmd, data)
	if err != nil {
		t.Errorf("Expected successful write, got: %v\n", err)
	}

	data, err = qzt.ReadSocket()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Printf("Got expected error result: %q\n", err)
		} else {
			t.Errorf("Expected not found error, got error: %v\n", err)
		}
	} else {
		t.Errorf("Expected not found error, got result: %s\n", data)
	}

	cmd = "BLAH"
	data = ""
	err = qzt.WriteSocket(cmd, data)
	if err != nil {
		t.Errorf("Expected successful write, got: %v\n", err)
	}
	data, err = qzt.ReadSocket()
	if err != nil {
		if strings.Contains(err.Error(), "unknown command") {
			log.Printf("Got expected error result: %q\n", err)
		} else {
			t.Errorf("Expected unknown command error, got error: %v\n", err)
		}
	} else {
		t.Errorf("Expected unknown command error, got result: %s\n", data)
	}
}

func TestCommandsThatPass(t *testing.T) {
	var err error
	var cmd, data string

	cmd = "HGET"
	data = "quiz:1 title"
	err = qzt.WriteSocket(cmd, data)
	if err != nil {
		t.Errorf("Expected successful write, got: %v\n", err)
	}

	data, err = qzt.ReadSocket()
	if err != nil {
		t.Errorf("Expected successful read, got result: %s\n", data)
	} else {
		if strings.Contains(data, "All") {
			log.Printf("Got expected result: %s\n", data)
		} else {
			t.Errorf("Expected to contain 'All', got result: %s\n", data)
		}
	}

	cmd = "HGETALL"
	data = "quiz:1"
	err = qzt.WriteSocket(cmd, data)
	if err != nil {
		t.Errorf("Expected successful write, got: %v\n", err)
	}

	data, err = qzt.ReadSocket()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if err != nil {
		t.Errorf("Expected successful read, got result: %s\n", data)
	} else {
		if strings.Contains(data, "title") && strings.Contains(data, "All") {
			log.Printf("Got expected result: %s\n", data)
		} else {
			t.Errorf("Expected to contain 'title' and 'All', got result: %s\n", data)
		}
	}

	cmd = "HMGET"
	data = "quiz:1 title junk categories"
	err = qzt.WriteSocket(cmd, data)
	if err != nil {
		t.Errorf("Expected successful write, got: %v\n", err)
	}

	data, err = qzt.ReadSocket()
	if err != nil {
		t.Errorf("Expected successful read, got result: %s\n", data)
	} else {
		if strings.Contains(data, "All") && strings.Contains(data, "null") {
			log.Printf("Got expected result: %s\n", data)
		} else {
			t.Errorf("Expected to contain 'All' and 'null', got result: %s\n", data)
		}
	}
}
