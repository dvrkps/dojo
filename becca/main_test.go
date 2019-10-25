package main

import (
	"os"
	"testing"
	"time"
)

func TestNewFilename(t *testing.T) {
	now := time.Date(2018, time.October, 10, 11, 22, 33, 44, time.UTC)
	got := newFilename(now)
	const want = "becca-20181010-112233.tar.gz"
	if got != want {
		t.Errorf("newFilename(...) = %v; want %v", got, want)
	}
}

func TestRunCmd(t *testing.T) {
	if got := runCmd("echo", ""); got != nil {
		t.Errorf("runCmd(echo '') = %v; want nil", got)
	}
	if got := runCmd("invalidPath"); got == nil {
		t.Errorf("runCmd(invalidPath) = %v; want <error>", got)
	}
}

func TestValidArgs(t *testing.T) {
	// working directory
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() error: %v", err)
	}
	// source
	s, d := "invalidpath", wd
	if got := validArgs(s, d); got == nil {
		t.Errorf("validArgs(%v, %v) = %v; want nil", s, d, got)
	}
	// dest
	s, d = wd, "invalidpath"
	if got := validArgs(s, d); got == nil {
		t.Errorf("validArgs(%v, %v) = %v; want nil", s, d, got)
	}
	// dest subdirectory
	s, d = wd, wd+"/.git"
	if got := validArgs(s, d); got == nil {
		t.Errorf("validArgs(%v, %v) = %v; want nil", s, d, got)
	}
}

func TestValidPath(t *testing.T) {
	// working directory
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() error: %v", err)
	}
	// valid
	if got := validPath(wd); got != nil {
		t.Errorf("validPath(%v) = %v; want nil", wd, got)
	}
	// not exists
	if got := validPath("not exists"); got == nil {
		t.Errorf("validPath(not exists) = %v; want error", got)
	}
	// not directory
	if got := validPath(wd + "/main_test.go"); got == nil {
		t.Errorf("validPath(%v/main_test.go) = %v; want error", wd, got)
	}
	// not absolute path
	if got := validPath("../becca"); got == nil {
		t.Errorf("validPath(../becca) = %v; want error", got)
	}
}
