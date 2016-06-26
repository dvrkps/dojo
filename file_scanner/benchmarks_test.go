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

const benchRows = 1000

func BenchmarkScanString(b *testing.B) {
	benchmarkScan(b, benchRows, scanString)
}

func BenchmarkScanBytes(b *testing.B) {
	benchmarkScan(b, benchRows, scanBytes)
}

func BenchmarkScanConcurrently(b *testing.B) {
	benchmarkScan(b, benchRows, scanConcurrently)
}
