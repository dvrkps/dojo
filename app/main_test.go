package main

import "testing"

func TestRun(t *testing.T) {
	got := run()
	if got != 0 {
		t.Error("boom")
	}
}
