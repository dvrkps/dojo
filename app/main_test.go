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

func mockAppLogger() (*App, *bytes.Buffer) {
	var buf bytes.Buffer
	a := &App{logger: log.New(&buf, "", 0)}
	return a, &buf
}

func loggerTest(t *testing.T, funcName string, buf *bytes.Buffer, want string) {

	got := buf.String()

	if got != want {
		t.Errorf("%s(...) = %q; want %q",
			funcName, got, want)
	}
}

func TestApp_Log(t *testing.T) {
	a, buf := mockAppLogger()
	a.Log("text", 12)
	loggerTest(t, "Log", buf, "text12\n")
}

func TestApp_Logf(t *testing.T) {
	var buf bytes.Buffer

	a := &App{logger: log.New(&buf, "", 0)}

	a.Logf("%d %s", 46, "text")

	got := buf.String()
	want := "46 text\n"

	if got != want {
		t.Errorf("Logf(...) = %q; want %q", got, want)
	}
}
