package main

import "fmt"

var value = 0

func main() {
	fmt.Println(value)
}

func add(i int) int {
	return value + i
}
