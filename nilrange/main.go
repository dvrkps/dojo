package main

import "os"

func main() { os.Exit(run()) }

func run() int {
	all := []int{
		one(nil),
		one([]int{1, 2, 3}),
		two(nil),
		two([]int{1, 2, 3}),
		three(nil),
		three([]int{1, 2, 3}),
	}
	var sum int
	for _, i := range all {
		sum += i
	}
	return 0
}

func one(data []int) int {
	var sum int
	for _, i := range data {
		sum += i
	}
	return sum
}

func two(data []int) int {
	var sum int
	if data == nil {
		return sum
	}
	for _, i := range data {
		sum += i
	}
	return sum
}

func three(data []int) int {
	if data == nil {
		return 0
	}
	var sum int
	for _, i := range data {
		sum += i
	}
	return sum
}
