package main

import (
	"testing"

	"github.com/dvrkps/dojo/fileload"
)

var data fileload.Data

func BenchmarkRun(b *testing.B) {
	var (
		d   fileload.Data
		err error
	)

	for n := 0; n < b.N; n++ {
		d, err = run()
		if err != nil {
			b.Fatal(err)
		}
	}

	data = d
}
