package main

import (
	"io"
	"log"
	"os"
)

func main() {
	os.Exit(run(NewApp()))
}

const exitOk = iota

func run(app *App) int {
	return exitOk
}

// App represent application and
// holds configuration.
type App struct {
	osargs []string
	stdout io.Writer
	logger *log.Logger
}

func NewApp() *App {
	return &App{
		osargs: os.Args,
		stdout: os.Stdout,
		logger: log.New(os.Stderr, "", 0)}
}
