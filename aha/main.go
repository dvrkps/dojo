package main

import (
	"io"
	"log"
	"os"
	"path"
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

	lgr := log.New(stderr, "", 0)

	out := log.New(stdout, "", 0)

	if cf.version {
		name := path.Base(args[0])
		lgr.Printf("%s v%s", name, commandVersion)
		return exitOk
	}

	area, err := NewArea(cf.m2, cf.ral, cf.chv)
	if err != nil {
		lgr.Printf("area: %v", err)
		return exitErr
	}

	out.Printf("%v", area)
	return exitOk
}
