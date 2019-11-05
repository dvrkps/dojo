package log

import (
	"bytes"
	"testing"
)

type testLogger struct {
	lgr    *Log
	bufOut *bytes.Buffer
	bufErr *bytes.Buffer
}

func testNew() testLogger {
	var (
		bufOut bytes.Buffer
		bufErr bytes.Buffer
	)
	l := testLogger{
		lgr:    New(&bufOut, &bufErr),
		bufOut: &bufOut,
		bufErr: &bufErr,
	}
	return l
}

const (
	infoLevel  = "Infof"
	errorLevel = "Errorf"
)

func TestLogger(t *testing.T) {

	tests := []struct {
		lgr    testLogger
		level  string
		format string
		args   []interface{}
		want   string
	}{
		{
			lgr:    testNew(),
			level:  infoLevel,
			format: "%v %v %v",
			args:   []interface{}{"info", 42, 3.14},
			want:   "info 42 3.14\n",
		},
		{
			lgr:    testNew(),
			level:  errorLevel,
			format: "%v %v %v",
			args:   []interface{}{"error", 42, 3.14},
			want:   "error 42 3.14\n",
		},
	}

	for _, tt := range tests {
		testLevel(t, &tt.lgr, tt.level, tt.format, tt.args, tt.want)
	}
}

func testLevel(t *testing.T, tl *testLogger, level string, format string, args []interface{}, want string) {
	t.Helper()
	var got string
	switch level {
	case infoLevel:
		tl.lgr.Infof(format, args...)
		got = tl.bufOut.String()
	case errorLevel:
		tl.lgr.Errorf(format, args...)
		got = tl.bufErr.String()
	default:
		t.Errorf("invalid level %q", level)
	}
	if got != want {
		t.Errorf("%v(%q, %v) = %q; want %q",
			level, format, args, got, want)
	}
}
