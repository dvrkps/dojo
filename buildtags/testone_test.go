// +build testone

package main

import "testing"

func TestOneAdd(t *testing.T) {
	tests := []int{11, 12, 13}
	testAdd(t, value, tests)
}
