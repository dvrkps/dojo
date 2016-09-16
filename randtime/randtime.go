package main

import (
	"math/rand"
	"time"
)

func main() {
}

func randInRange(min, max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(max-min) + min
}
