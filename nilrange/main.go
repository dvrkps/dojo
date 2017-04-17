package main

func main() {
	sum := one([]int{1, 2, 3})
	println(sum)
}

func one(data []int) int {
	var sum int
	for _, i := range data {
		sum += i
	}
	return sum
}
