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

func TestLogger(t *testing.T) {
	for _, tt := range loggerTests() {
		var got string

		lgr := testNew(tt.verbose)

		switch tt.level {
		case infoLevel:
			lgr.out.Infof(tt.format, tt.args...)
			got = lgr.bufOut.String()
		case debugLevel:
			lgr.out.Debugf(tt.format, tt.args...)
			got = lgr.bufOut.String()
		case errorLevel:
			lgr.out.Errorf(tt.format, tt.args...)
			got = lgr.bufErr.String()
		default:
			t.Errorf("invalid level %q", tt.level)
		}

		if got != tt.want {
			t.Errorf("%v(%q, %v) = %q; want %q",
				tt.level, tt.format, tt.args, got, tt.want)
		}
	}
}

type testLogger struct {
	out    *Log
	bufOut *bytes.Buffer
	bufErr *bytes.Buffer
}

func testNew(verbose bool) testLogger {
	var (
		bufOut bytes.Buffer
		bufErr bytes.Buffer
	)

	l := testLogger{
		out:    New(&bufOut, &bufErr),
		bufOut: &bufOut,
		bufErr: &bufErr,
	}

	if verbose {
		l.out.Verbose()
	}

	return l
}

type loggerTest struct {
	verbose bool
	level   string
	format  string
	args    []interface{}
	want    string
}

func loggerTests() []loggerTest {
	tests := []loggerTest{
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
