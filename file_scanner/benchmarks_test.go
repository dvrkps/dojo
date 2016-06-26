package main

import (
	"io"
	"testing"
)

var resultPersons Persons

func benchmarkScan(b *testing.B, rows int, fn func(io.Reader) Persons) {
	f := fakeReader(rows)
	b.ResetTimer()
	var r Persons
	for n := 0; n < b.N; n++ {
		r = fn(f)
	}
	resultPersons = r
}

func BenchmarkScanStringNew(b *testing.B) {
	benchmarkScan(b, 10000, scanString)
}

func BenchmarkScanBytesNew(b *testing.B) {
	benchmarkScan(b, 10000, scanBytes)
}
