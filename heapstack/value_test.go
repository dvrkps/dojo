package main

import (
	"strconv"
	"testing"
)

func TestValueID(t *testing.T) {
	tests := []int{
		5,
	}
	const newID = 42
	vid := newValueID(newID)
	if vid.id != newID {
		t.Errorf("newValue(%v) = %v; want %v", newID, vid.id, newID)
	}
	for _, id := range tests {
		vid.setID(id)
		got := vid.string()
		want := strconv.Itoa(id)
		if got != want {
			t.Errorf("id = %v: string() = %q; want %q", vid.id, got, want)
		}
	}
}
