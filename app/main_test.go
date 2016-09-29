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

func loggerTest(t *testing.T, funcName string, buf *bytes.Buffer, want string) {
	got := buf.String()
	if got != want {
		t.Errorf("%s(...) = %q; want %q",
			funcName, got, want)
	}
}

func TestApp_Log(t *testing.T) {
	a, _, buf := mockApp()
	a.Log("text", 12)
	loggerTest(t, "Log", buf, "text12\n")
}

func TestApp_Logf(t *testing.T) {
	a, _, buf := mockApp()
	a.Logf("%d %s", 46, "text")
	loggerTest(t, "Logf", buf, "46 text\n")
}

func TestApp_Logln(t *testing.T) {
	a, _, buf := mockApp()
	a.Logln(23, "a", "text")
	loggerTest(t, "Logln", buf, "23 a text\n")
}

/*
func TestApp_Printf(t *testing.T) {
	a, out, buf := mockApp()
	a.Printf("%d %s", 46, " text")
	loggerTest(t, "Printf", buf, "46 text\n")
}
*/
