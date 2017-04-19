package sum

import "testing"

var result bool

func benchmarkHelper(b *testing.B, fn func(int, ...int) bool, sum int, nums ...int) {
	var r bool
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = fn(sum, nums...)
	}
	result = r
}

var benchCases = []struct {
	name string
	fn   func(int, ...int) bool
	sum  int
	nums []int
}{
	{name: "basic(8,1,2,3,9)",
		fn:   basic,
		sum:  8,
		nums: []int{1, 2, 3, 9}},
	{name: "basic(8,1,2,4,4)",
		fn:   basic,
		sum:  8,
		nums: []int{1, 2, 4, 4}},
	{name: "better(8,1,2,3,9)",
		fn:   better,
		sum:  8,
		nums: []int{1, 2, 3, 9}},
	{name: "better(8,1,2,4,4)",
		fn:   better,
		sum:  8,
		nums: []int{1, 2, 4, 4}},
	{name: "linear(8,1,2,3,9)",
		fn:   linear,
		sum:  8,
		nums: []int{1, 2, 3, 9}},
	{name: "linear(8,1,2,4,4)",
		fn:   linear,
		sum:  8,
		nums: []int{1, 2, 4, 4}},
	{name: "complements(8,1,2,3,9)",
		fn:   complements,
		sum:  8,
		nums: []int{1, 2, 3, 9}},
	{name: "complements(8,1,2,4,4)",
		fn:   complements,
		sum:  8,
		nums: []int{1, 2, 4, 4}},
}

func BenchmarkAll(b *testing.B) {
	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			benchmarkHelper(b, bc.fn, bc.sum, bc.nums...)
		})
	}
}
