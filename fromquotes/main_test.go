package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	in  []string
	out []string
}

func testCases() []testCase {
	r := []testCase{
		{
			in:  short(),
			out: []string{".hi", "My name is Omar", "\"123\""},
		},
		{
			in:  long(),
			out: []string{".hi I'm the real Slim ShadyMy name is Omar", "hello", "world", "\"123\"", "a"},
		},
	}
	return r
}

var result []string

func TestOriginal(t *testing.T) {
	for _, tc := range testCases() {
		got := original(tc.in)
		want := tc.out
		if !reflect.DeepEqual(got, want) {
			t.Errorf("original(%v) = %v; want %v", got, want)
		}
	}
}
