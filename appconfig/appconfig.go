package main

import (
	"io"
	"os"
)

type appConfig struct {
	osargs []string
	stdout io.Writer
	stderr io.Writer
}

func runApp(cfg *appConfig) int {
	if len(cfg.osargs) < 2 {
		return 1
	}
	return 0
}

func main() {
	code := runApp(&appConfig{
		osargs: os.Args,
		stdout: os.Stdout,
		stderr: os.Stderr,
	})

	os.Exit(code)
}
