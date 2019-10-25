package main

import (
	"io"
	"io/ioutil"
	"log"
)

// Logger holds all loggers.
type Logger struct {
	il *log.Logger
	el *log.Logger
}

// NewLogger creates logger.
func NewLogger(stdout, stderr io.Writer) *Logger {
	if stdout == nil {
		stdout = ioutil.Discard
	}
	if stderr == nil {
		stderr = ioutil.Discard
	}
	l := Logger{
		il: log.New(stdout, "", 0),
		el: log.New(stderr, "", 0),
	}
	return &l
}

// Info logs info level messages.
func (l *Logger) Info(format string, v ...interface{}) {
	l.il.Printf(format, v...)
}

// Err logs error level messages.
func (l *Logger) Err(format string, v ...interface{}) {
	l.el.Printf(format, v...)
}
