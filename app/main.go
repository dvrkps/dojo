package main

import (
	"fmt"
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

// Log prints to logger like log.Print.
func (a *App) Log(v ...interface{}) {
	a.logger.Print(v...)
}

// Logf prints to logger like log.Printf.
func (a *App) Logf(format string, v ...interface{}) {
	a.logger.Printf(format, v...)
}

// Logln prints to logger like log.Println.
func (a *App) Logln(v ...interface{}) {
	a.logger.Println(v...)
}

const (
	fnPrint = iota
	fnPrintf
	fnPrintln
)

func (a *App) output(typ int, format string, v ...interface{}) {
	var err error
	switch typ {
	case fnPrint:
		_, err = fmt.Fprint(a.stdout, v...)

	case fnPrintf:
		_, err = fmt.Fprintf(a.stdout, format, v...)
	case fnPrintln:
		_, err = fmt.Fprintln(a.stdout, v...)
	}
	if err != nil {
		a.logger.Print(err)
	}
}

// Print prints to stdout like fmt.Print.
func (a *App) Print(v ...interface{}) {
	a.output(fnPrint, "", v...)
}

// Printf prints to stdout like fmt.Printf.
func (a *App) Printf(format string, v ...interface{}) {
	a.output(fnPrintf, format, v...)
}
