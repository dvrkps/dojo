package main

import (
	"fmt"
	"time"
)

func main() {
	a := gen("a", 3)
	for i := range a {
		println(i)
	}
}

const (
	defaultName = "empty"
	minLimit    = 1
	maxLimit    = 10
)

func payload(name string, max int) string {
	if name == "" {
		name = defaultName
	}
	if max < minLimit {
		max = maxLimit
	}
	return fmt.Sprintf("%s%d", name, max)
}

func gen(name string, max int) chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i <= max; i++ {
			time.Sleep(1e9)
			ch <- payload(name, i)
		}
		close(ch)
	}()
	return ch
}
