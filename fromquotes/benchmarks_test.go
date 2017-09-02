package main

import "testing"

var result []string

func benchCases() map[string][]string {
	m := map[string][]string{
		"short": short(),
		"long":  long(),
	}
	return m
}

func benchFunc(b *testing.B, fn func([]string) []string, in []string) {
	var r []string
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = fn(in)
	}
	result = r
}
