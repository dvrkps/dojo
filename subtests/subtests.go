package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("1+2 = %v", sum(1, 2))
}
