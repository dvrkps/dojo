package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type appConfig struct {
	osargs []string
	stdout io.Writer
	stderr io.Writer
	logger *log.Logger
}

func (ac *appConfig) write(w io.Writer, a ...interface{}) {
	_, err := fmt.Fprint(w, a...)
	if err != nil {
		ac.logger.Print(err)
	}
}

func runApp(cfg *appConfig) int {
	if cfg == nil {
		return 1
	}
	if len(cfg.osargs) < 1 {
		return 1
	}
	// cfg.logger.Print("eto")
	return 0
}

func main() {
	os.Exit(runApp(&appConfig{
		osargs: os.Args,
		stdout: os.Stdout,
		stderr: os.Stderr,
		logger: log.New(os.Stderr, "", 0),
	}))
}
