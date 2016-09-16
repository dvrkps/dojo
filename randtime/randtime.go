package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	min := 1 * time.Nanosecond
	max := 5 * time.Hour
	rt := randDurationInRange(min, max)
	fmt.Println(rt)
}

func randDurationInRange(minDur, maxDur time.Duration) time.Duration {

	min := int64(minDur)
	max := int64(maxDur)

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	x := r.Int63n(max-min) + min

	return time.Duration(x)
}
