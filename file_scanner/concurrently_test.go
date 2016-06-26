package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestScanConcurrently(t *testing.T) {
	const noRows = 10
	f := fakeReader(noRows)
	ps := scanConcurrently(f)
	fmt.Println(ps)
	if got := len(ps); got != noRows {
		t.Errorf("len(Persons) = %v; want %v", got, noRows)
	}
	want := Person{
		ID:   4,
		Name: "name4",
		Age:  noRows - 4}
	if got := ps[3]; !reflect.DeepEqual(got, want) {
		t.Errorf("p[3] = %v; want %v", got, want)
	}
}
