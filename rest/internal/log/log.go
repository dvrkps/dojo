// Package log provides support for using
// standard and error loggers.
package log

import (
	"io"
	"log"
)

// Log contains standard and error loggers.
type Log struct {
	verbose bool
	ol      *log.Logger
	el      *log.Logger
}

// New creates log.
func New(stdout, stderr io.Writer) *Log {
	var l Log
	if stdout != nil {
		l.ol = log.New(stdout, "", 0)
	}

	if stderr != nil {
		l.el = log.New(stderr, "", 0)
	}

	return &l
}

// Infof writes formated info message.
func (l *Log) Infof(format string, v ...interface{}) {
	outf(l.ol, format, v...)
}

// Debugf writes formated debug message.
func (l *Log) Debugf(format string, v ...interface{}) {
	if l.verbose {
		outf(l.ol, format, v...)
	}
}

// Errorf writes formated error message.
func (l *Log) Errorf(format string, v ...interface{}) {
	outf(l.el, format, v...)
}

// Verbose sets verbose output.
func (l *Log) Verbose() {
	l.verbose = true
}

func outf(l *log.Logger, format string, v ...interface{}) {
	if l != nil {
		l.Printf(format, v...)
	}
}
