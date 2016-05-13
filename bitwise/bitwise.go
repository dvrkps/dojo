package main

import "fmt"

func main() {
	bitwiseAnd()
}

func bitwiseAnd() {
	a := byte(25)
	b := byte(37)

	// bit is true only if both source(a,b) bit are 1(true).
	c := a & b

	prt("a", a)
	prt("b", b)
	prt("a & b", c)
}

func prt(s string, b byte) {
	fmt.Printf("%-10s: %08b\n", s, b)
}
