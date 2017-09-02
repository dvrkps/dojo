package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	in  []string
	out []string
}

func testCases() map[string]testCase {
	m := map[string]testCase{
		"short": {
			in:  short(),
			out: []string{".hi", "My name is Omar", "\"123\""},
		},
		"long": {
			in:  long(),
			out: []string{".hi I'm the real Slim ShadyMy name is Omar", "hello", "world", "\"123\"", "a"},
		},
	}
	return m
}

var result []string

func TestOriginal(t *testing.T) {
	for name, tc := range testCases() {
		got := original(tc.in)
		want := tc.out
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s\t: original(%v) = %v; want %v", name, tc.in, got, want)
		}
	}
}
