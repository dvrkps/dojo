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

func TestApp_Log(t *testing.T) {
	var buf bytes.Buffer

	a := &App{logger: log.New(&buf, "", 0)}

	a.Log("text", 12)

	got := buf.String()
	want := "text12\n"

	if got != want {
		t.Errorf("Log(...) = %q; want %q", got, want)
	}
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
