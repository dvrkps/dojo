package log

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestLog(t *testing.T) {
	tests := []struct {
		name    string
		buffer  io.Writer
		verbose bool
		action  func(log *Log)
		want    string
	}{
		{
			name:    "nil",
			buffer:  nil,
			verbose: false,
			action: func(log *Log) {
				log.Printf("Printf %v", 42)
			},
			want: "",
		},
		{
			name:    "Printf",
			buffer:  &bytes.Buffer{},
			verbose: false,
			action: func(log *Log) {
				log.Printf("Printf %v", 42)
			},
			want: "prefix: Printf 42\n",
		},
		{
			name:    "Verbosef verbose",
			buffer:  &bytes.Buffer{},
			verbose: true,
			action: func(log *Log) {
				log.Verbosef("Verbosef verbose %v", 42)
			},
			want: "prefix: Verbosef verbose 42\n",
		},
		{
			name:    "Verbosef",
			buffer:  &bytes.Buffer{},
			verbose: false,
			action: func(log *Log) {
				log.Verbosef("Verbosef %v", 42)
			},
			want: "",
		},
	}
	for _, tt := range tests {
		lgr := New(tt.buffer, "prefix: ")
		lgr.SetVerbose(tt.verbose)
		tt.action(&lgr)
		buf, ok := tt.buffer.(fmt.Stringer)

		if ok {
			got := buf.String()
			if got != tt.want {
				t.Errorf("%s: = %q; want %q",
					tt.name, got, tt.want)
			}
		}
	}
}
