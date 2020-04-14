package main

import (
	"testing"

	"github.com/dvrkps/dojo/fileload"
)

func TestParse(t *testing.T) {
	fileload.TestParse(t, parse)
}

func BenchmarkParse(b *testing.B) {
	var data fileload.Data
	b.Run("99", func(b *testing.B) {
		data = fileload.BenchParse(b, parse, fileload.Rows99)
		_ = len(data)
	})
}
