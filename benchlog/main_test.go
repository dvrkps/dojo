package main

import "testing"

func BenchNil(b *testing.B) {
	for n := 0; n < b.N; n++ {
	}
}
