package main

import "fmt"

func main() {
	v := []int{42, 43}
	fmt.Printf("%p\n", &v)
	println(&v)
	println(v)
	fmt.Printf("%p\n", v)
	fmt.Printf("0: %p\n", &v[0])
	fmt.Printf("1: %p\n", &v[1])
}
