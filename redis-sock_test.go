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

	// Connect Test
	err = qzt.ConnectRedisSocket()
	if err != nil {
		t.Errorf("Expected new quiztool, got: %v\n", err)
	}
}

func TestCommandsThatFail(t *testing.T) {
	var err error
	var data string
	wsc := qzt.wsclient

	// Tests the returned errors

	wsc.redisWrap.msg.Command = "HGET"
	wsc.redisWrap.msg.Data = "quiz:1 titlen"
	err = wsc.WriteRedis()
	if err != nil {
		t.Errorf("Expected successful write, got: %v\n", err)
	}

	data, err = wsc.ReadRedis()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Printf("Got expected error result: %q\n", err)
		} else {
			t.Errorf("Expected not found error, got error: %v\n", err)
		}
	} else {
		t.Errorf("Expected not found error, got result: %s\n", data)
	}

	wsc.redisWrap.msg.Command = "BLAH"
	wsc.redisWrap.msg.Data = ""
	err = wsc.WriteRedis()
	if err != nil {
		t.Errorf("Expected successful write, got: %v\n", err)
	}
	data, err = wsc.ReadRedis()
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
	var data string
	wsc := qzt.wsclient

	// Passing Tests

	wsc.redisWrap.msg.Command = "HGET"
	wsc.redisWrap.msg.Data = "quiz:1 title"
	err = wsc.WriteRedis()
	if err != nil {
		t.Errorf("Expected successful write, got: %v\n", err)
	}

	data, err = wsc.ReadRedis()
	if err != nil {
		t.Errorf("Expected successful read, got result: %s\n", data)
	} else {
		if strings.Contains(data, "All") {
			log.Printf("Got expected result: %s\n", data)
		} else {
			t.Errorf("Expected to contain 'All', got result: %s\n", data)
		}
	}

	wsc.redisWrap.msg.Command = "HGETALL"
	wsc.redisWrap.msg.Data = "quiz:1"
	err = wsc.WriteRedis()
	if err != nil {
		t.Errorf("Expected successful write, got: %v\n", err)
	}

	data, err = wsc.ReadRedis()
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

	wsc.redisWrap.msg.Command = "HMGET"
	wsc.redisWrap.msg.Data = "quiz:1 title junk categories"
	err = wsc.WriteRedis()
	if err != nil {
		t.Errorf("Expected successful write, got: %v\n", err)
	}

	data, err = wsc.ReadRedis()
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
