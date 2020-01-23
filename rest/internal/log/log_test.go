package log

import (
	"bytes"
	"testing"
)

const (
	infoLevel  = "Infof"
	debugLevel = "Debugf"
	errorLevel = "Errorf"
)

func TestLog(t *testing.T) {
	for _, tt := range logTests() {
		var got string

		l, bufOut, bufErr := newTestLog(tt.verbose)

		switch tt.level {
		case infoLevel:
			l.Infof(tt.format, tt.args...)

			got = bufOut.String()
		case debugLevel:
			l.Debugf(tt.format, tt.args...)

			got = bufOut.String()
		case errorLevel:
			l.Errorf(tt.format, tt.args...)

			got = bufErr.String()
		default:
			t.Errorf("invalid level %q", tt.level)
		}

		if got != tt.want {
			t.Errorf("%v(%q, %v) = %q; want %q",
				tt.level, tt.format, tt.args, got, tt.want)
		}
	}
}

func newTestLog(verbose bool) (*Log, *bytes.Buffer, *bytes.Buffer) {
	var (
		bufOut bytes.Buffer
		bufErr bytes.Buffer
	)

	l := New(&bufOut, &bufErr)

	if verbose {
		l.Verbose()
	}

	return l, &bufOut, &bufErr
}

type logTest struct {
	verbose bool
	level   string
	format  string
	args    []interface{}
	want    string
}

func logTests() []logTest {
	tests := []logTest{
		{
			verbose: false,
			level:   infoLevel,
			format:  "%v %v %v",
			args:    []interface{}{"info", 42, 3.14},
			want:    "info 42 3.14\n",
		},
		{
			verbose: true,
			level:   debugLevel,
			format:  "%v %v %v",
			args:    []interface{}{"debug", 42, 3.14},
			want:    "debug 42 3.14\n",
		},
		{
			verbose: false,
			level:   debugLevel,
			format:  "%v %v %v",
			args:    []interface{}{"debug", 42, 3.14},
			want:    "",
		},
		{
			verbose: false,
			level:   errorLevel,
			format:  "%v %v %v",
			args:    []interface{}{"error", 42, 3.14},
			want:    "error 42 3.14\n",
		},
	}

	return tests
}
