package main

import (
	"bytes"
	"errors"
	"log"
	"testing"
)

func TestRun(t *testing.T) {
	got := run(NewApp())
	if got != 0 {
		t.Error("boom")
	}
}

func TestNewApp(t *testing.T) {
	if NewApp() == nil {
		t.Error("NewApp() = <nil>")
	}
}

func mockApp() (*App, *bytes.Buffer, *bytes.Buffer) {
	var o, l bytes.Buffer
	a := &App{
		stdout: &o,
		logger: log.New(&l, "", 0),
	}
	return a, &o, &l
}

func writeTest(t *testing.T, funcName string, bufOut, bufLog *bytes.Buffer, wantOut, wantLog string) {
	gotOut := bufOut.String()
	gotLog := bufLog.String()
	if gotOut != wantOut || gotLog != wantLog {
		t.Errorf("%s(...) = out: %q log: %q; want out:%q, log: %q",
			funcName, gotOut, gotLog, wantOut, wantLog)
	}
}

func TestApp_Log(t *testing.T) {
	a, gotOut, gotLog := mockApp()
	a.Log("text", 12)
	writeTest(t, "Log", gotOut, gotLog, "", "text12\n")
}

func TestApp_Logf(t *testing.T) {
	a, gotOut, gotLog := mockApp()
	a.Logf("%d %s", 46, "text")
	writeTest(t, "Logf", gotOut, gotLog, "", "46 text\n")
}

func TestApp_Logln(t *testing.T) {
	a, gotOut, gotLog := mockApp()
	a.Logln(23, "a", "text")
	writeTest(t, "Logln", gotOut, gotLog, "", "23 a text\n")
}

func TestApp_Printf(t *testing.T) {
	a, gotOut, gotLog := mockApp()
	a.Printf("%d %s", 46, "text")
	writeTest(t, "Printf", gotOut, gotLog, "46 text", "")
}

type mockErrWriter struct {
}

func (w *mockErrWriter) Write(p []byte) (int, error) {
	return 0, errors.New("write error")
}

func TestApp_output_error(t *testing.T) {
	a, gotOut, gotLog := mockApp()
	a.stdout = &mockErrWriter{}
	a.Printf("%s", "something")
	writeTest(t, "output", gotOut, gotLog, "", "write error\n")
}
