package main

import "testing"

var resultPersons Persons

func BenchmarkScanString(b *testing.B) {
	f := fakeReader(10000)
	b.ResetTimer()
	var r Persons
	for n := 0; n < b.N; n++ {
		r = scanString(f)
	}
	resultPersons = r
}

func BenchmarkScanBytes(b *testing.B) {
	f := fakeReader(10000)
	b.ResetTimer()
	var r Persons
	for n := 0; n < b.N; n++ {
		r = scanBytes(f)
	}
	resultPersons = r
}
