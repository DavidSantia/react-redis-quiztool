package quiztool

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (wrp *Wrapper) NextLF() (bool) {
	if wrp.Debug {
		log.Printf("VERBOSE Doing NextLF: BufPtr = %d BufLen = %d\n", wrp.BufPtr, wrp.BufLen)
	}

	// find next LF
	for {
		if wrp.BufPtr == wrp.BufLen - 1 {
			return false
		}
		if wrp.Buf[wrp.BufPtr] == '\r' && wrp.Buf[wrp.BufPtr+1] == '\n' {
			wrp.BufPtr++
			if wrp.BufPtr < wrp.BufLen {
				wrp.BufPtr++
			}
			if wrp.Debug {
				log.Printf("NextLF returned BufPtr=%d\n", wrp.BufPtr)
			}
			return true
		}
		wrp.BufPtr++
	}
}

func (wrp *Wrapper) GetBulkString() (s string, err error) {
	if wrp.Debug {
		log.Printf("VERBOSE GetBulkString processing: %q\n", wrp.Buf[wrp.BufPtr:wrp.BufLen])
	}

	// Locate size
	i := wrp.BufPtr
	if !wrp.NextLF() {
		err = fmt.Errorf("Bulk string size EOLN not found: %s", wrp.Buf[wrp.BufPtr:wrp.BufLen-2])
		return
	}
	l, err := strconv.Atoi(string(wrp.Buf[i+1:wrp.BufPtr-2]))
	if l == -1 {
		err = fmt.Errorf("%s Bulk string not found", wrp.Command)
		s = "null"
		return
	}

	if wrp.Debug {
		log.Printf("VERBOSE GetBulkString expected length %d\n", l)
	}

	// Locate value
	i = wrp.BufPtr
	if !wrp.NextLF() {
		err = fmt.Errorf("Bulk string value EOLN not found: %s", wrp.Buf[wrp.BufPtr:wrp.BufLen-2])
		return
	}

	// Make sure Bulk String length matches RESP size value
	if l != wrp.BufPtr - i - 2 {
		err = fmt.Errorf("Expected length %d != actual length %d for Bulk string %s",
			l, wrp.BufPtr - i, wrp.Buf[i:wrp.BufPtr-2])
		return
	}
	s = fmt.Sprintf("%q", wrp.Buf[i:wrp.BufPtr-2])
	if wrp.Debug {
		log.Printf("VERBOSE GetBulkString returned: %s\n", s)
	}
	return
}

func (wrp *Wrapper) ParseBuf() (s string, err error) {
	if wrp.Debug {
		log.Printf("VERBOSE ParseBuf processing: %q\n", wrp.Buf[wrp.BufPtr:wrp.BufLen])
	}

	resp_type := wrp.Buf[wrp.BufPtr]

	if resp_type == '+' {
		// Simple string
		s = fmt.Sprintf("%q", wrp.Buf[wrp.BufPtr+1:wrp.BufLen-2])
	} else if resp_type == ':' {
		// Integer
		s = fmt.Sprintf("%s", wrp.Buf[wrp.BufPtr+1:wrp.BufLen-2])
	} else if resp_type == '$' {
		// Bulk string
		s, err = wrp.GetBulkString()
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("Unrecognized response data: %s", wrp.Buf[wrp.BufPtr+1:wrp.BufLen-2])
		return
	}

	if wrp.Debug {
		log.Printf("VERBOSE ParseBuf returned: %s\n", s)
	}
	return
}

func (wrp *Wrapper) GetArray() (s string, err error) {
	if wrp.Debug {
		log.Printf("VERBOSE GetArray processing: %q\n", wrp.Buf[wrp.BufPtr:wrp.BufLen])
	}

	// Locate size
	i := wrp.BufPtr
	if !wrp.NextLF() {
		err = fmt.Errorf("Array EOLN not found: %s", wrp.Buf[wrp.BufPtr:wrp.BufLen])
		return
	}
	l, err := strconv.Atoi(string(wrp.Buf[i+1:wrp.BufPtr-2]))
	if l == -1 {
		err = fmt.Errorf("%s Array not found", wrp.Command)
		s = "null"
		return
	}

	if wrp.Debug {
		log.Printf("VERBOSE GetArray expected length %d\n", l)
	}

	// Iterate through values
	arr := make([]string, l)
	var elem string

	for a := 0; a < l; a++ {
		// Make sure there are more entries to parse
		if wrp.BufPtr == wrp.BufLen {
			err = fmt.Errorf("Array expected length %d != actual length %d", l, a+1)
			return
		}
		elem, err = wrp.ParseBuf()

		if wrp.KeyPair {
			// Form { elem1 : elem2 }, ...
			if a % 2 == 0 {
				arr[a] = "{" + elem + ":"
			} else {
				arr[a] = elem + "},"
			}
		} else {
			arr[a] = elem + ","
		}

	}

	// Join array of values
	data := "[" + strings.Join(arr, "")
	// Drop trailing ','
	s = data[:len(data)-1] + "]"
	if wrp.Debug {
		log.Printf("VERBOSE GetArray returned: %s\n", s)
	}
	return
}

func (wrp *Wrapper) ParseSocket() (data string, err error) {

	wrp.BufPtr = 0
	resp_type := wrp.Buf[0]

	// Array data is list if HMGET
	wrp.KeyPair = wrp.Command != "HMGET"

	// Check for Redis error
	if resp_type == '-' {
		err = fmt.Errorf("Redis protocol: %s", wrp.Buf[1:wrp.BufLen-2])
		return
	}

	// Check for Array
	if resp_type == '*' {
		data, err = wrp.GetArray()
		return
	}

	// Check for Simple string, Integer or Bulk string
	data, err = wrp.ParseBuf()
	return
}
