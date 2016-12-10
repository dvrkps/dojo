package main

import "fmt"

var value = 0

func main() {
	fmt.Println("value: ", value)
	fmt.Println("add: ", add(2))
}

func add(i int) int {
	return value + i
}
