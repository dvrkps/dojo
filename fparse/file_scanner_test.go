package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func fakeReader(size int) io.Reader {
	var i int
	var buf bytes.Buffer
	for i < size {
		i++
		s := fmt.Sprintf("%d,name%d,%d\n", i, i, size-i)
		_, _ = buf.WriteString(s)
	}
	return bytes.NewBuffer(buf.Bytes())
}

func testScan(t *testing.T, fn func(io.Reader) Persons) {
	const noRows = 10
	f := fakeReader(noRows)
	ps := fn(f)
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

func TestScanString(t *testing.T) {
	testScan(t, scanString)
}

func TestScanBytes(t *testing.T) {
	testScan(t, scanBytes)
}
