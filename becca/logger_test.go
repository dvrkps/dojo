package main

import (
	"bytes"
	"testing"
)

func testSetup() (*Logger, *bytes.Buffer, *bytes.Buffer) {
	var (
		bufOut = &bytes.Buffer{}
		bufErr = &bytes.Buffer{}
	)

	return NewLogger(bufOut, bufErr), bufOut, bufErr
}

func TestNewLogger(t *testing.T) {
	l := NewLogger(nil, nil)
	if l == nil {
		t.Error("NewLogger() == nil")
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
		l, bufOut, _ := testSetup()
		testLevel(l.Info, tt.format, tt.args)
		got := bufOut.String()
		if got != tt.want {
			t.Errorf("Info(%q, %v) = %q; want %q",
				tt.format, tt.args, got, tt.want)
		}
	}
}

func TestLoggerErr(t *testing.T) {
	tests := []struct {
		format string
		args   []interface{}
		want   string
	}{
		{
			format: "err",
			want:   "err\n",
		},
		{
			format: "%v %v %v",
			args:   []interface{}{"err", 42, 3.14},
			want:   "err 42 3.14\n",
		},
	}

	for _, tt := range tests {
		l, _, bufErr := testSetup()
		testLevel(l.Err, tt.format, tt.args)
		got := bufErr.String()
		if got != tt.want {
			t.Errorf("Err(%q, %v) = %q; want %q",
				tt.format, tt.args, got, tt.want)
		}
	}
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
