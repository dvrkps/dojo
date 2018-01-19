package main

import "testing"

var result []int

func BenchmarkCAdd(b *testing.B) {
	var r []int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = cAdd(r, 1)
		r = cAdd(r, 1, 2, 3)
	}
	result = r
}

func BenchmarkPAdd(b *testing.B) {
	var r []int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		pAdd(&r, 1)
		pAdd(&r, 1, 2, 3)
	}
	result = r
}
