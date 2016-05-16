package main

import (
	"fmt"
	"os"
)

func main() {

	evenOdd()

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

func evenOdd() {

	vals := [...]struct {
		n    int
		want bool
	}{
		{n: 2, want: true},
		{n: 5, want: false},
	}

	for _, v := range vals {
		bit := isEvenBitwise(v.n)
		rem := isEvenRemainder(v.n)
		if bit != rem {
			fmt.Println("bitwise != remainder")
			os.Exit(1)
		}
		if got, want := bit, v.want; got != want {
			fmt.Printf("isEvenBitwise(%d) = %v; want %v\n",
				v.n, got, want)
		}
		if got, want := rem, v.want; got != want {
			fmt.Printf("isEvenRemainder(%d) = %v; want %v\n",
				v.n, got, want)
		}
	}

}

func isEvenBitwise(i int) bool {
	v := (i & 1) == 0
	return v
}

func isEvenRemainder(i int) bool {
	v := (i % 2) == 0
	return v
}
