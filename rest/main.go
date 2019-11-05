package main

import (
	"io"
	"os"
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
	return exitOk
}
