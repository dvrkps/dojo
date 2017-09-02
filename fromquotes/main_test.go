package main

import (
	"fmt"
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

func testFunc(t *testing.T, fname string, fn func([]string) []string) {
	for name, tc := range testCases() {
		name = fmt.Sprintf("%s(%s)", fname, name)
		t.Run(name, func(t *testing.T) {
			got := fn(tc.in)
			want := tc.out
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v; want %v", got, want)
			}
		})
	}
}

func TestOriginal(t *testing.T) {
	testFunc(t, "original", original)
}
