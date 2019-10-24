package main

import (
	"reflect"
	"testing"
)

func TestPure(t *testing.T) {
	const max = 10
	want := testPrimes(max)
	got := pure(max)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("pure(%d) = %v; want %v", max, got, want)
	}
}
