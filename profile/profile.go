package profile

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const maxValue = 99

// Data returns random values.
func Data(n int) io.Reader {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		v := randValue(maxValue)
		_, err := fmt.Fprintf(&buf, "%d\n", v)
		if err != nil {
			log.Printf("Fprintf: %v", err)
		}
	}
	return &buf
}

// StrData returns  random values.
func StrData(n int) io.Reader {
	vs := []string{}
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
