package main

import (
	"fmt"
	"io"
	"os"

	"github.com/dvrkps/dojo/rest/internal/configuration"
	"github.com/dvrkps/dojo/rest/internal/log"
)

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}

const (
	exitOk = 0
	//exitErr  = 1
	exitUser = 2
)

func run(args []string, stdout, stderr io.Writer) int {
	config, err := configuration.New(args, stderr)
	if err != nil {
		fmt.Fprintf(stderr, "configuration: %v", err)
		return exitUser
	}

	log := log.New(config.Verbose, stdout, stderr)

	log.Infof("%v", "info")
	log.Debugf("%v", "debug")
	log.Errorf("%v", "err")

	return exitOk
}
