package quiztool

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (wrp *Wrapper) NextLF() bool {
	if wrp.Debug {
		log.Printf("VERBOSE Doing NextLF: bufPtr = %d bufLen = %d\n", wrp.bufPtr, wrp.bufLen)
	}

	// find next LF
	for {
		if wrp.bufPtr == wrp.bufLen-1 {
			return false
		}
		if wrp.buf[wrp.bufPtr] == '\r' && wrp.buf[wrp.bufPtr+1] == '\n' {
			wrp.bufPtr++
			if wrp.bufPtr < wrp.bufLen {
				wrp.bufPtr++
			}
			if wrp.Debug {
				log.Printf("VERBOSE NextLF returned bufPtr=%d\n", wrp.bufPtr)
			}
			return true
		}
		wrp.bufPtr++
	}
}

func (wrp *Wrapper) GetBulkString() (s string, err error) {
	if wrp.Debug {
		log.Printf("VERBOSE GetBulkString processing: %q\n", wrp.buf[wrp.bufPtr:wrp.bufLen])
	}

	// Locate size
	i := wrp.bufPtr
	if !wrp.NextLF() {
		err = fmt.Errorf("Bulk string size EOLN not found: %s", wrp.buf[wrp.bufPtr:wrp.bufLen-2])
		return
	}
	l, err := strconv.Atoi(string(wrp.buf[i+1 : wrp.bufPtr-2]))
	if l == -1 {
		err = fmt.Errorf("%s Bulk string not found", wrp.msg.Command)
		s = "null"
		return
	}

	if wrp.Debug {
		log.Printf("VERBOSE GetBulkString expected length %d\n", l)
	}

	// Locate value
	i = wrp.bufPtr
	if !wrp.NextLF() {
		err = fmt.Errorf("Bulk string value EOLN not found: %s", wrp.buf[wrp.bufPtr:wrp.bufLen-2])
		return
	}

	// Make sure Bulk String length matches RESP size value
	if l != wrp.bufPtr-i-2 {
		err = fmt.Errorf("Expected length %d != actual length %d for Bulk string %s",
			l, wrp.bufPtr-i, wrp.buf[i:wrp.bufPtr-2])
		return
	}

	// Convert to JSON (string will be in "")
	s = fmt.Sprintf("%q", wrp.buf[i:wrp.bufPtr-2])
	if wrp.Debug {
		log.Printf("VERBOSE GetBulkString returned: %s\n", s)
	}
	return
}

func (wrp *Wrapper) ParseBuf() (s string, err error) {
	if wrp.Debug {
		log.Printf("VERBOSE Parsebuf processing: %q\n", wrp.buf[wrp.bufPtr:wrp.bufLen])
	}

	resp_type := wrp.buf[wrp.bufPtr]

	if resp_type == '+' {
		// Simple string, convert to JSON (string will be in "")
		s = fmt.Sprintf("%q", wrp.buf[wrp.bufPtr+1:wrp.bufLen-2])
	} else if resp_type == ':' {
		// Integer, convert to JSON (string)
		s = fmt.Sprintf("%s", wrp.buf[wrp.bufPtr+1:wrp.bufLen-2])
	} else if resp_type == '$' {
		// Bulk string
		s, err = wrp.GetBulkString()
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("Unrecognized response data: %s", wrp.buf[wrp.bufPtr+1:wrp.bufLen-2])
		return
	}

	if wrp.Debug {
		log.Printf("VERBOSE Parsebuf returned: %s\n", s)
	}
	return
}

func (wrp *Wrapper) GetKeyPairOrArray() (s string, err error) {
	var method = "Array"
	if wrp.keyPair {
		method = "KeyPair"
	}
	if wrp.Debug {
		log.Printf("VERBOSE Get%s processing: %q\n", method, wrp.buf[wrp.bufPtr:wrp.bufLen])
	}

	// Locate size
	i := wrp.bufPtr
	if !wrp.NextLF() {
		err = fmt.Errorf("Get%s EOLN not found: %s", method, wrp.buf[wrp.bufPtr:wrp.bufLen])
		return
	}
	l, err := strconv.Atoi(string(wrp.buf[i+1 : wrp.bufPtr-2]))
	if l == -1 || l == 0 {
		err = fmt.Errorf("%s %s not found", wrp.msg.Command, method)
		s = "null"
		return
	}

	if wrp.Debug {
		log.Printf("VERBOSE Get%s expected length %d\n", method, l)
	}

	// Iterate through values
	arr := make([]string, l)
	var elem string

	for a := 0; a < l; a++ {
		// Make sure there are more entries to parse
		if wrp.bufPtr == wrp.bufLen {
			err = fmt.Errorf("Get%s expected length %d != actual length %d", method, l, a+1)
			return
		}
		elem, err = wrp.ParseBuf()

		if wrp.keyPair {
			// Form { elem1 : elem2 }, ...
			if a%2 == 0 {
				arr[a] = elem + ":"
			} else {
				arr[a] = elem + ","
			}
		} else {
			arr[a] = elem + ","
		}

	}

	// Join array of key-pairs, or values
	var data string
	if wrp.keyPair {
		data = "{"
	} else {
		data = "["
	}
	data = data + strings.Join(arr, "")
	// Drop trailing ','
	if wrp.keyPair {
		s = data[:len(data)-1] + "}"
	} else {
		s = data[:len(data)-1] + "]"
	}
	if wrp.Debug {
		log.Printf("VERBOSE Get%s returned: %s\n", method, s)
	}
	return
}

func (wrp *Wrapper) ParseSocket() (data string, err error) {

	wrp.bufPtr = 0
	resp_type := wrp.buf[0]

	// Array data is list if HMGET
	wrp.keyPair = wrp.msg.Command != "HMGET"

	// Check for Redis error
	if resp_type == '-' {
		err = fmt.Errorf("%s", wrp.buf[1:wrp.bufLen-2])
		return
	}

	// Check for Array
	if resp_type == '*' {
		data, err = wrp.GetKeyPairOrArray()
		return
	}

	// Check for Simple string, Integer or Bulk string
	data, err = wrp.ParseBuf()
	return
}
