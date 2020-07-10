package minimal

import (
	"bytes"
	"testing"
)

func TestLog(t *testing.T) {
	tests := []struct {
		name    string
		verbose bool
		action  func(log *Log)
		want    string
	}{
		{
			name:    "F",
			verbose: false,
			action: func(log *Log) {
				log.F("F %v", 42)
			},
			want: "prefix: F 42\n",
		},
		{
			name:    "Vf verbose",
			verbose: true,
			action: func(log *Log) {
				log.Vf("Vf verbose %v", 42)
			},
			want: "prefix: Vf verbose 42\n",
		},
		{
			name:    "Vf",
			verbose: false,
			action: func(log *Log) {
				log.Debugf("Vf %v", 42)
			},
			want: "",
		},
	}
	for _, tt := range tests {
		log, buf := testLogSetup(tt.verbose)
		tt.action(log)
		got := buf.String()
		if got != tt.want {
			t.Errorf("%s: = %q; want %q",
				tt.name, got, tt.want)
		}
	}
}

func testLogSetup(verbose bool) (*Log, *bytes.Buffer) {
	var buf = &bytes.Buffer{}
	l := New(buf, "prefix")
	l.SetVerbose(verbose)
	return l, buf
}
