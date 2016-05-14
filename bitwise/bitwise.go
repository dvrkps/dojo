package main

import "fmt"

func main() {
	operators()
}

func operators() {

	pf := func(s string, vals ...int) {
		const (
			row = "%04b = %d\n"
		)

		fmt.Println(s)
		for _, i := range vals {
			fmt.Printf(row, i, i)
		}
		fmt.Println()
	}

	// Use bitwise AND & to get the bits
	// that are in 3 AND 6
	pf("and", 3, 6, 3&6)

	// Use bitwise OR | to get the bits
	// that are in 3 OR 6
	pf("or", 3, 6, 3|6)

	// Use bitwise XOR ^ to get the bits
	// that are in 3 OR 6 BUT NOT BOTH
	pf("xor", 3, 6, 3^6)

	// Use bit clear AND NOT &^ to get the bits
	// that are in 3 AND NOT 6 (order matters)
	pf("and not", 3, 6, 3&^6)
}
