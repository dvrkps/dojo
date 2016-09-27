package main

import (
	"io"
	"log"
	"os"
)

func main() {
	os.Exit(run())
}

const exitOk = iota

func run() int {
	return exitOk
}

// App represent application and
// holds configuration.
type App struct {
	osargs []string
	stdout io.Writer
	logger *log.Logger
}
