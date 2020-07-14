package main

import "fmt"

type kind int

const (
	zero kind = iota
	one
	two
	three
	last = three
)

func main() {
	for i := zero - 1; i <= last+1; i++ {
		fmt.Printf("% v: %v\n", i, setSomething(i))
	}
}

func setSomething(k kind) error {
	if k < zero {
		return fmt.Errorf("%v < zero", k)
	}

	if k > last {
		return fmt.Errorf("%v > %v", k, last)
	}

	return nil
}
