package main

import (
	"fmt"
)

func main() {
	a := gen("a", 3)
	b := gen("b", 5)
	out := merge(a, b)
	for i := range out {
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
			ch <- payload(name, i)
		}
		close(ch)
	}()
	return ch
}

func merge(a, b chan string) chan string {
	out := make(chan string)
	go func() {
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				out <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				out <- v
			}
		}
		close(out)
	}()
	return out
}
