package main

import (
	"bytes"
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
	writeTest(t, "Log", gotOut, gotLog, "text12\n", "a")
}

func TestApp_Logf(t *testing.T) {
	a, _, buf := mockApp()
	a.Logf("%d %s", 46, "text")
	writerTest(t, "Logf", buf, "46 text\n")
}

func TestApp_Logln(t *testing.T) {
	a, _, buf := mockApp()
	a.Logln(23, "a", "text")
	writerTest(t, "Logln", buf, "23 a text\n")
}

func TestApp_Printf(t *testing.T) {
	a, out, _ := mockApp()
	a.Printf("%d %s", 46, "text")
	writerTest(t, "Printf", out, "46 text")
}
