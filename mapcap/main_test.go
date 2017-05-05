package main

import (
	"fmt"
	"testing"
)

func TestInitMap(t *testing.T) {
	m := make(map[int]int)
	n := 10
	initMap(m, n)
	got := len(m)
	if got != n {
		t.Errorf("initMap: len = %d; want %d", got, n)
	}
}

var result map[int]int

var benchs = []int{30000, 100000} //, 1000, 10000}

func BenchmarkMaps(b *testing.B) {
	var name string
	for _, bb := range benchs {

		name = fmt.Sprintf("base(%d)", bb)
		b.Run(name, func(b *testing.B) {
			m := make(map[int]int)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				initMap(m, bb)
			}
			result = m
		})

		name = fmt.Sprintf("cap(%d)", bb)
		b.Run(name, func(b *testing.B) {
			m := make(map[int]int, bb)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				initMap(m, bb)
			}
			result = m
		})

	}
}
