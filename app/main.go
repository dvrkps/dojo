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

// NewApp create application.
func NewApp() *App {
	return &App{
		osargs: os.Args,
		stdout: os.Stdout,
		logger: log.New(os.Stderr, "", 0)}
}

// Logf prints to logger.
func (a *App) Logf(format string, v ...interface{}) {
	a.logger.Printf(format, v...)
}
