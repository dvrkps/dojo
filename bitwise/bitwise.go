package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// bitwiseAND()

	// bitwiseOR()

	operators()
}

func bitwiseAND() {
	fmt.Println("\nbitwiseAND")
	a := bin("101")
	b := bin("011")

	// both bit must be 1
	c := a & b

	prt("a", a)
	prt("b", b)
	prt("a & b", c)
}

func bitwiseOR() {
	fmt.Println("\nbitwiseOR")
	a := bin("101")
	b := bin("011")

	// at least one bit must be 1.
	c := a | b

	prt("a", a)
	prt("b", b)
	prt("a | b", c)
}

// bin converts binary string to integer.
func bin(s string) int64 {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func prt(s string, i int64) {
	fmt.Printf("%-10s: %8b = %d\n", s, i, i)
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
	// that are in 1 OR 4
	pf("or", 3, 6, 3|6)

	// Use bitwise XOR ^ to get the bits
	// that are in 3 OR 6 BUT NOT BOTH
	pf("xor", 3, 6, 3^6)

	// Use bit clear AND NOT &^ to get the bits
	// that are in 3 AND NOT 6 (order matters)
	pf("and not", 3, 6, 3&^6)
}
