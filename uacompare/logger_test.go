package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewLogger(t *testing.T) {
	lgr := newLogger(true, nil, nil)

	if lgr.dl == nil {
		t.Error("nil debug")
	}
}

func testSetupLogger(t *testing.T, debug bool) (*Logger, func(string, string)) {
	var (
		bo = &bytes.Buffer{}
		be = &bytes.Buffer{}
	)

	lgr := newLogger(debug, bo, be)

	return lgr, testLoggerHelper(t, bo, be)
}

func testLoggerHelper(t *testing.T, bo, be fmt.Stringer) func(string, string) {
	return func(wantOut, wantErr string) {
		t.Helper()
		if got := bo.String(); got != wantOut {
			t.Errorf("stdout: got %q, want %q", got, wantOut)
		}
		if got := be.String(); got != wantErr {
			t.Errorf("stderr: got %q, want %q", got, wantErr)
		}
	}
}

func TestLoggerFalseDebug(t *testing.T) {
	lgr, check := testSetupLogger(t, false)
	lgr.Debug("message")
	check("", "")
}

func TestLoggerDebug(t *testing.T) {
	lgr, check := testSetupLogger(t, true)
	lgr.Debug("message")
	check("message\n", "")
}

func TestLoggerDebugf(t *testing.T) {
	lgr, check := testSetupLogger(t, true)
	lgr.Debugf("%s-%d-%t", "message", 1, true)
	check("message-1-true\n", "")
}

func TestLoggerInfo(t *testing.T) {
	lgr, check := testSetupLogger(t, true)
	lgr.Info("message")
	check("message\n", "")
}

func TestLoggerInfof(t *testing.T) {
	lgr, check := testSetupLogger(t, true)
	lgr.Infof("%s-%d-%t", "message", 1, true)
	check("message-1-true\n", "")
}

func TestLoggerError(t *testing.T) {
	lgr, check := testSetupLogger(t, true)
	lgr.Error("message")
	check("", "message\n")
}

func TestLoggerErrorf(t *testing.T) {
	lgr, check := testSetupLogger(t, true)
	lgr.Errorf("%s-%d-%t", "message", 1, true)
	check("", "message-1-true\n")
}
