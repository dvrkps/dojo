package main

import (
	"strconv"
	"testing"
)

func TestClientSetAndString(t *testing.T) {
	tests := []int{
		5,
	}
	const newID = 42
	c := newClientCopy(newID)
	if c.id != newID {
		t.Errorf("newClientCopy(%v) = %v; want %v", newID, c.id, newID)
	}
	for _, id := range tests {
		c.setID(id)
		got := c.string()
		want := strconv.Itoa(id)
		if got != want {
			t.Errorf("id = %v: string() = %q; want %q", c.id, got, want)
		}
	}
}
