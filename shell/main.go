package main

import (
	"io"
	"log"
	"os"
)

func main() {
	lgr := log.New(os.Stderr, "", 0)

	err := run(os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		lgr.Printf("run: %v", err)
		os.Exit(1)
	}
}

func run(stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	return nil
}
