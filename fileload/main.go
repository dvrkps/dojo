package fileload

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

//Rows holds number of rows.
const (
	Rows9    = 9
	Rows99   = 99
	Rows999  = 999
	Rows9999 = 9999
)

// ParseFunc is parse function.
type ParseFunc func(r io.Reader, d *Data) error

// Run starts parse function.
func Run(fn ParseFunc, rows int) int {
	const (
		exitOK  = 0
		exitErr = 1
	)

	r := newReader(rows)

	d := make(Data, 0, rows)

	err := fn(r, &d)
	if err != nil {
		log.Printf("parse: %v", err)

		return exitErr
	}

	got := len(d)
	if got != rows {
		log.Printf("len(d) = %d; want %d", got, rows)

		return exitErr
	}

	return exitOK
}

const rowFormat = "%v:Row %v"

func newReader(rows int) io.Reader {
	var buf bytes.Buffer

	const msg = rowFormat + "\n"

	for i := 0; i < rows; i++ {
		_, _ = fmt.Fprintf(&buf, msg, i, i)
	}

	return &buf
}
