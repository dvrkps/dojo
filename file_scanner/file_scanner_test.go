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

func TestScanString(t *testing.T) {
	f := fakeReader(10)
	ps := scanString(f)
	fmt.Println(ps)
}

func TestScanBytes(t *testing.T) {
	f := fakeReader(10)
	ps := scanBytes(f)
	fmt.Println(ps)
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
