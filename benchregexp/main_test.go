package main

import (
	"regexp"
	"testing"
)

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

func BenchmarkOneIsValid(b *testing.B) {
	var r bool
	key := "sv101"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = isValid(key, "sv")
	}
	result = r
}

func BenchmarkOneRegexp(b *testing.B) {
	var r bool
	key := "sv101"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r, _ = regexp.MatchString("(?i)^sv[0-9]{1,3}$", key)
	}
	result = r
}
