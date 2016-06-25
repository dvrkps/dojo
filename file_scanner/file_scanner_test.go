package main

import (
	"bytes"
	"fmt"
	"io"
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
	const fakeRows = 10
	f := fakeReader(fakeRows)
	ps := fn(f)
	fmt.Println(ps)
}

func TestScanString(t *testing.T) {
	testScan(t, scanString)
}

func TestScanBytes(t *testing.T) {
	testScan(t, scanBytes)
}

var resultPersons Persons

func BenchmarkScanString(b *testing.B) {
	f := fakeReader(10000)
	b.ResetTimer()
	var r Persons
	for n := 0; n < b.N; n++ {
		r = scanString(f)
	}
	resultPersons = r
}

func BenchmarkScanBytes(b *testing.B) {
	f := fakeReader(10000)
	b.ResetTimer()
	var r Persons
	for n := 0; n < b.N; n++ {
		r = scanBytes(f)
	}
	resultPersons = r
}
