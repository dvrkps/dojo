package main

import (
	"io"
	"os"
	"path"

	"github.com/dvrkps/aha/log"
)

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}

const (
	exitOk   = 0
	exitErr  = 1
	exitUser = 2
)

func run(args []string, stdout, stderr io.Writer) int {
	cf, err := newCmdFlags(args, stderr)
	if err != nil {
		return exitUser
	}

	lgr := log.New(stdout, stderr)

	if cf.version {
		name := path.Base(args[0])
		lgr.Logf("%s v%s", name, commandVersion)
		return exitOk
	}

	area, err := NewArea(cf.m2, cf.ral, cf.chv)
	if err != nil {
		lgr.Errorf("area: %v", err)
		return exitErr
	}

	lgr.Logf("%v", area)
	return exitOk
}
