package logger

import (
	"bytes"
	"testing"
)

func testSetup(verbose bool) (*Logger, *bytes.Buffer, *bytes.Buffer) {
	var (
		bufOut = &bytes.Buffer{}
		bufErr = &bytes.Buffer{}
	)

	return New(verbose, bufOut, bufErr), bufOut, bufErr
}

func testLevel(
	fn func(format string, v ...interface{}),
	format string,
	args []interface{}) {
	if len(args) < 1 {
		fn(format)
		return
	}
	fn(format, args...)
}

func TestNew(t *testing.T) {
	l := New(false, nil, nil)
	if l == nil {
		t.Error("New() == nil")
	}
	if l.verbose {
		t.Error("verbose == true")
	}
	if l.il != nil {
		t.Error("il != nil")
	}
	if l.dl != nil {
		t.Error("dl != nil")
	}
	if l.el != nil {
		t.Error("el != nil")
	}
	l.Info("try to write to nil logger")
}

func TestSetVerbose(t *testing.T) {
	l := New(false, nil, nil)
	const want = true
	l.SetVerbose(want)
	got := l.verbose
	if got != want {
		t.Errorf("SetVerbose(%v) == %v; want %v", want, got, want)
	}
}

func TestLoggerInfo(t *testing.T) {
	tests := []struct {
		format string
		args   []interface{}
		want   string
	}{
		{
			format: "info",
			want:   "info\n",
		},
		{
			format: "%v %v %v",
			args:   []interface{}{"info", 42, 3.14},
			want:   "info 42 3.14\n",
		},
	}

	for _, tt := range tests {
		l, bufOut, _ := testSetup(false)
		testLevel(l.Info, tt.format, tt.args)
		got := bufOut.String()
		if got != tt.want {
			t.Errorf("Info(%q, %v) = %q; want %q",
				tt.format, tt.args, got, tt.want)
		}
	}
}

func TestLoggerDebug(t *testing.T) {
	tests := []struct {
		verbose bool
		format  string
		args    []interface{}
		want    string
	}{
		{
			verbose: true,
			format:  "debug",
			want:    "debug\n",
		},
		{
			verbose: true,
			format:  "%v %v %v",
			args:    []interface{}{"debug", 42, 3.14},
			want:    "debug 42 3.14\n",
		},
		{
			format: "debug",
			want:   "",
		},
	}

	for _, tt := range tests {
		l, bufOut, _ := testSetup(tt.verbose)
		testLevel(l.Debug, tt.format, tt.args)
		got := bufOut.String()
		if got != tt.want {
			t.Errorf("verbose:%v Debug(%q, %v) = %q; want %q",
				tt.verbose, tt.format, tt.args, got, tt.want)
		}
	}
}

func TestLoggerError(t *testing.T) {
	tests := []struct {
		format string
		args   []interface{}
		want   string
	}{
		{
			format: "error",
			want:   "error\n",
		},
		{
			format: "%v %v %v",
			args:   []interface{}{"error", 42, 3.14},
			want:   "error 42 3.14\n",
		},
	}

	for _, tt := range tests {
		l, _, bufErr := testSetup(false)
		testLevel(l.Error, tt.format, tt.args)
		got := bufErr.String()
		if got != tt.want {
			t.Errorf("Error(%q, %v) = %q; want %q",
				tt.format, tt.args, got, tt.want)
		}
	}
}
