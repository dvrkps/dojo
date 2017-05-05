package main

import (
	"fmt"
	"testing"
)

func TestInitMap(t *testing.T) {
	m := make(map[int]int)
	n := 10
	initMap(m, n)
	got := len(m)
	if got != n {
		t.Errorf("initMap: len = %d; want %d", got, n)
	}
	fmt.Println(m)
}
