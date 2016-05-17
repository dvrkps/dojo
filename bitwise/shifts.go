package main

import "fmt"

func shifts() {
	pf := func(b byte) {
		const (
			row = "%08b = %d\n"
		)

		fmt.Printf(row, b, b)
	}

	var b byte

	fmt.Println("left shift")
	var i uint
	for i < 8 {
		b = 1 << i
		pf(b)
		i++
	}

	fmt.Println("\nright shift")
	for b > 0 {
		pf(b)
		b = b >> 1
	}

}
