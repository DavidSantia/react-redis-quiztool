package quiztool

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (qzt *QuizApp) NextLF() (bool) {
	if qzt.Debug > 1 {
		log.Printf("Doing NextLF: BufPtr = %d BufLen = %d\n", qzt.BufPtr, qzt.BufLen)
	}

	// find next LF
	for {
		if qzt.BufPtr == qzt.BufLen - 1 {
			return false
		}
		if qzt.Buf[qzt.BufPtr] == '\r' && qzt.Buf[qzt.BufPtr+1] == '\n' {
			qzt.BufPtr++
			if qzt.BufPtr < qzt.BufLen {
				qzt.BufPtr++
			}
			if qzt.Debug > 1 {
				log.Printf("NextLF returned BufPtr=%d\n", qzt.BufPtr)
			}
			return true
		}
		qzt.BufPtr++
	}
}

func (qzt *QuizApp) GetBulkString() (s string, err error) {
	if qzt.Debug > 1 {
		log.Printf("GetBulkString processing: %q\n", qzt.Buf[qzt.BufPtr:qzt.BufLen])
	}

	// Locate size
	i := qzt.BufPtr
	if !qzt.NextLF() {
		err = fmt.Errorf("Bulk string size EOLN not found: %s", qzt.Buf[qzt.BufPtr:qzt.BufLen-2])
		return
	}
	l, err := strconv.Atoi(string(qzt.Buf[i+1:qzt.BufPtr-2]))
	if qzt.Debug > 1 {
		log.Printf("GetBulkString expected length %d\n", l)
	}

	// Locate value
	i = qzt.BufPtr
	if !qzt.NextLF() {
		err = fmt.Errorf("Bulk string value EOLN not found: %s", qzt.Buf[qzt.BufPtr:qzt.BufLen-2])
		return
	}

	// Make sure Bulk String length matches RESP size value
	if l != qzt.BufPtr - i - 2 {
		err = fmt.Errorf("Expected length %d != actual length %d for Bulk string %s",
			l, qzt.BufPtr - i, qzt.Buf[i:qzt.BufPtr-2])
		return
	}
	s = fmt.Sprintf("%q", qzt.Buf[i:qzt.BufPtr-2])
	if qzt.Debug > 1 {
		log.Printf("GetBulkString returned: %s\n", s)
	}
	return
}

func (qzt *QuizApp) ParseBuf() (s string, err error) {
	if qzt.Debug > 1 {
		log.Printf("ParseBuf processing: %q\n", qzt.Buf[qzt.BufPtr:qzt.BufLen])
	}

	resp_type := qzt.Buf[qzt.BufPtr]

	if resp_type == '+' {
		// Simple string
		s = fmt.Sprintf("%q", qzt.Buf[qzt.BufPtr+1:qzt.BufLen-2])
	} else if resp_type == ':' {
		// Integer
		s = fmt.Sprintf("%s", qzt.Buf[qzt.BufPtr+1:qzt.BufLen-2])
	} else if resp_type == '$' {
		// Bulk string
		s, err = qzt.GetBulkString()
	} else {
		err = fmt.Errorf("Unrecognized response data: %s", qzt.Buf[qzt.BufPtr+1:qzt.BufLen-2])
		return
	}

	if qzt.Debug > 1 {
		log.Printf("ParseBuf returned: %s\n", s)
	}
	return
}

func (qzt *QuizApp) GetArray() (s string, err error) {
	if qzt.Debug > 1 {
		log.Printf("GetArray processing: %q\n", qzt.Buf[qzt.BufPtr:qzt.BufLen])
	}

	// Locate size
	i := qzt.BufPtr
	if !qzt.NextLF() {
		err = fmt.Errorf("Array EOLN not found: %s", qzt.Buf[qzt.BufPtr:qzt.BufLen])
		return
	}
	l, err := strconv.Atoi(string(qzt.Buf[i+1:qzt.BufPtr-2]))
	if qzt.Debug > 1 {
		log.Printf("GetArray expected length %d\n", l)
	}

	// Iterate through values
	arr := make([]string, l)
	var elem string

	for a := 0; a < l; a++ {
		// Make sure there are more entries to parse
		if qzt.BufPtr == qzt.BufLen {
			err = fmt.Errorf("Array expected length %d != actual length %d", l, a+1)
			return
		}

		// Form { elem1 : elem2 }, ...
		elem, err = qzt.ParseBuf()
		if a % 2 == 0 {
			arr[a] = "{" + elem + ":"
		} else {
			arr[a] = elem + "},"
		}
	}

	// Join array of values
	data := "[" + strings.Join(arr, "")
	// Drop trailing ','
	s = data[:len(data)-1] + "]"
	if qzt.Debug > 1 {
		log.Printf("GetArray returned: %s\n", s)
	}
	return
}

func (qzt *QuizApp) ParseSocket() (data string, err error) {
	// Read socket
	qzt.BufLen, err = qzt.RedisSock.Read(qzt.Buf)
	if err != nil {
		return
	}
	if qzt.BufLen == 0 {
		err = fmt.Errorf("Empty read from Redis socket")
		return
	}

	qzt.BufPtr = 0
	resp_type := qzt.Buf[0]

	// Check for Redis error
	if resp_type == '-' {
		err = fmt.Errorf("Redis protocol: %s", qzt.Buf[1:qzt.BufLen])
		return
	}

	// Check for Array
	if resp_type == '*' {
		data, err = qzt.GetArray()
		return
	}

	// Check for Simple string, Integer or Bulk string
	data, err = qzt.ParseBuf()
	return
}
