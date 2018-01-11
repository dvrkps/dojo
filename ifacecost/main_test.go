package main

import "testing"

var result []string

func BenchmarkType(b *testing.B) {
	w := words{sentence: "aaa bbb ccc"}
	var r []string
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = typeSplit(w)
	}
	result = r
}

func BenchmarkInterface(b *testing.B) {
	w := words{sentence: "aaa bbb ccc"}
	var r []string
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = ifaceSplit(w)
	}
	result = r
}
