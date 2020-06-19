package main

import (
	"fmt"
)

type Numbers struct {
	vals []int
}

func (n *Numbers) Add(v int) {
	n.vals = append(n.vals, v)
}

func (n Numbers) Average() float64 {
	sum := 0.0
	for _, num := range n.vals {
		sum += float64(num)
	}
	return sum / float64(len(n.vals))
}

func main() {
	var numbers Numbers
	numbers.Add(200)
	numbers.Add(43)
	numbers.Add(-6)
	fmt.Println(numbers.Average())
}
