package main

import "testing"

func benchQuery() map[string]string {
	return map[string]string{"sv": "value of sv"}
}

func BenchmarkWihRegexp(b *testing.B) {
}
