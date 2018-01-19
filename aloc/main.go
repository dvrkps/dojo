package main

import "fmt"

func main() {
	var s []int
	show(s)

	s = cAdd(s, 1, 2, 3)
	pAdd(&s, 11, 22, 33)
	show(s)
}

func cAdd(s []int, v ...int) []int {
	return append(s, v...)
}

func pAdd(s *[]int, v ...int) {
	*s = append(*s, v...)
}

func show(s []int) {
	fmt.Printf("%v\nlen: %v cap: %v\n%p\n\n",
		s,
		len(s),
		cap(s),
		s)
}
