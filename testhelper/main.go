package testhelper

// Sum sums numbers.
func Sum(n ...int) int {
	var s int
	for _, i := range n {
		s += i
	}
	return s
}
