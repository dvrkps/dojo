package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	bitwiseAND()

	bitwiseOR()
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
