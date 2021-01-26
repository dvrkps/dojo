package main

import "testing"

func TestIsIt(t *testing.T) {
	tt := &testing.T{}

	const fail = true

	IsIt(tt, fail)

	if tt.Failed() {
		t.Fail() //Error("puklo")
	}

}
