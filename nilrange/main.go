package main

func main() {}

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
