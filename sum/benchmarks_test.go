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
