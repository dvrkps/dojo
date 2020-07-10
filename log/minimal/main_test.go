package minimal

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
				log.F("F %v", 42)
			},
			want: "",
		},
		{
			name:    "F",
			buffer:  &bytes.Buffer{},
			verbose: false,
			action: func(log *Log) {
				log.F("F %v", 42)
			},
			want: "prefix: F 42\n",
		},
		{
			name:    "Vf verbose",
			buffer:  &bytes.Buffer{},
			verbose: true,
			action: func(log *Log) {
				log.Vf("Vf verbose %v", 42)
			},
			want: "prefix: Vf verbose 42\n",
		},
		{
			name:    "Vf",
			buffer:  &bytes.Buffer{},
			verbose: false,
			action: func(log *Log) {
				log.Vf("Vf %v", 42)
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
