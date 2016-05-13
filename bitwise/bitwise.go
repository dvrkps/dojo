package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	bitwiseAnd()
}

func bitwiseAnd() {
	fmt.Println("\nbitwiseAND")
	a := bin("101")
	b := bin("011")

	// bit is 1 only if
	// sources(a,b) bit are 1.
	c := a & b

	prt("a", a)
	prt("b", b)
	prt("a & b", c)
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
