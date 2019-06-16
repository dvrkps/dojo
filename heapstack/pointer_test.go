package main

import (
	"strconv"
	"testing"
)

func TestPointerID(t *testing.T) {
	tests := []int{
		5,
	}
	const newID = 42
	pid := newPointerID(newID)
	if pid.id != newID {
		t.Errorf("newPointerID(%v) = %v; want %v", newID, pid.id, newID)
	}
	for _, id := range tests {
		pid.setID(id)
		got := pid.string()
		want := strconv.Itoa(id)
		if got != want {
			t.Errorf("id = %v: string() = %q; want %q", pid.id, got, want)
		}
	}
}
