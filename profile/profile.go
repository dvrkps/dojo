package profile

import (
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const maxValue = 99

func strData1(n int) io.Reader {
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
