package fileload

import (
	"bytes"
	"fmt"
	"testing"
)

// BenchParse benchmark parse function.
func BenchParse(b *testing.B, fn ParseFunc, rows int) Data {
	var d Data
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		d = Data{}
		r := newReader(rows)
		b.StartTimer()
		err := fn(r, &d)
		if err != nil {
			b.Fatal(err)
		}
	}

	return d
}

// TestParse is test helper.
func TestParse(t *testing.T, fn ParseFunc) {
	var d Data

	r := newReader(Rows9)

	err := fn(r, &d)
	if err != nil {
		t.Fatalf("fn: %v", err)
	}

	if len(d) != Rows9 {
		t.Fatalf("len(Data) = %v; want %v", len(d), Rows9)
	}

	for i := range d {
		r := d[i]

		if r.ID != i {
			t.Fatalf("id = %v; want %v", r.ID, i)
		}

		var buf bytes.Buffer
		_, err := fmt.Fprintf(&buf, "Row %v", r.ID)
		if err != nil {
			t.Fatalf("buf: %v", err)
		}

		want := buf.Bytes()

		if !bytes.Equal(r.Name, want) {
			t.Fatalf("name = %q; want %q", r.Name, want)
		}
	}
}
