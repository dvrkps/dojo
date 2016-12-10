// +build testtwo

package main

import "testing"

func TestTwoAdd(t *testing.T) {
	tests := []int{21, 22, 23}
	testAdd(t, value, tests)
}
