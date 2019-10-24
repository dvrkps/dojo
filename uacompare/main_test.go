package main

import (
	"io/ioutil"
	"testing"
)

func TestRun(t *testing.T) {
	var (
		debug  = true
		stdout = ioutil.Discard // os.Stdout // &bytes.Buffer{}
		stderr = ioutil.Discard // os.Stderr      // &bytes.Buffer{}
	)

	if got := run(testDataPath, debug, stdout, stderr); got != exitOK {
		t.Errorf("run(%q, ...) got %v, want %v", testDataPath, got, exitOK)
	}
}
