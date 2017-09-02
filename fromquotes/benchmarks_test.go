package main

import (
	"fmt"
	"testing"
)

var result []string

func benchCases() map[string]func([]string) []string {
	m := map[string]func([]string) []string{
		"original": original,
		"better":   better,
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

func Benchmark(b *testing.B) {
	inputs := map[string][]string{
		"short": short(),
		"long":  long(),
	}
	for bname, fn := range benchCases() {
		for iname, in := range inputs {
			name := fmt.Sprintf("%s(%s)", bname, iname)
			b.Run(name, func(b *testing.B) {
				benchFunc(b, fn, in)
			})
		}

	}
}
