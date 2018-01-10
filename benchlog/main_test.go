package main

import "testing"

func BenchmarkNil(b *testing.B) {
	l := nilLog{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l.Print("nillog")
	}
}
