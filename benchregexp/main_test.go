package main

import "testing"

func benchQuery() map[string]string {
	return map[string]string{"sv": "value of sv"}
}

var result bool

func BenchmarkWithRegexp(b *testing.B) {
	var r bool
	q := benchQuery()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = withRegexp(q)
	}
	result = r
}

func BenchmarkNoooRegexp(b *testing.B) {
	var r bool
	q := benchQuery()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = noooRegexp(q)
	}
	result = r
}
