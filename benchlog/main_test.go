package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkNil(b *testing.B) {
	l := nilLog{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l.Print("nillog")
	}
}

func BenchmarkDiscard(b *testing.B) {
	l := normalLog{log: log.New(ioutil.Discard, "", 0)}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l.Print("discardlog")
	}
}

func BenchmarkNormal(b *testing.B) {
	buf := &bytes.Buffer{}
	l := normalLog{log: log.New(buf, "", 0)}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l.Print("normal")
	}
}
