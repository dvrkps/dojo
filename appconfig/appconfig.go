package main

import (
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

func runApp(cfg *appConfig) int {
	if cfg == nil {
		return 1
	}
	if len(cfg.osargs) < 1 {
		return 1
	}
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
