package profile

import (
	"bytes"
	"io"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const maxValue = 99

// Data1 returns random values.
func Data1(n int) io.Reader {
	var buf bytes.Buffer
	var v int
	var dst []byte
	var err error
	for i := 0; i < n; i++ {
		v = randValue(maxValue)
		dst = dst[:0]
		dst = strconv.AppendInt(dst, int64(v), 10)
		dst = append(dst, '\n')
		_, err = buf.Write(dst)
		if err != nil {
			log.Printf("Fprintf: %v", err)
		}
	}
	return &buf
}

// Data2 returns random values.
func Data2(n int) io.Reader {
	var buf bytes.Buffer
	var v int
	var err error
	for i := 0; i < n; i++ {
		v = randValue(maxValue)
		_, err = buf.WriteString(strconv.Itoa(v))
		if err != nil {
			log.Printf("WriteString: %v", err)
		}
		_, err = buf.WriteString("\n")
		if err != nil {
			log.Printf("WriteString: %v", err)
		}
	}
	return &buf
}

// StrData1 returns random values.
func StrData1(n int) io.Reader {
	vs := []string{}
	for i := 0; i < n; i++ {
		v := randValue(maxValue)
		vs = append(vs, strconv.Itoa(v))
	}
	all := strings.Join(vs, "\n")
	return strings.NewReader(all)
}

// StrData2 returns random values.
func StrData2(n int) io.Reader {
	vs := make([]string, 0, n)
	for i := 0; i < n; i++ {
		v := randValue(maxValue)
		vs = append(vs, strconv.Itoa(v))
	}
	all := strings.Join(vs, "\n")
	return strings.NewReader(all)
}

func randValue(max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(max)
}
