package main

func main() {}

func one(data []int) int {
	var sum int
	for _, i := range data {
		sum += i
	}
	return sum
}
