// Package davelog implements a minimal logging package.
// Inspired by article of Dave Cheney.
// https://dave.cheney.net/2015/11/05/lets-talk-about-logging
package davelog

import (
	"io"
	"log"
)

// A Log represents an active logger.
type Log struct {
	output  *log.Logger
	verbose bool
}

// New creates logger.
func New(w io.Writer, verbose bool) *Log {
	l := Log{
		output:  log.New(w, "", 0),
		verbose: verbose,
	}
	return &l
}

func (l *Log) logf(format string, v ...interface{}) {
	if l.output == nil {
		return
	}
	l.output.Printf(format, v...)
}

// Debugf logs debug messages.
func (l *Log) Debugf(format string, v ...interface{}) {
	if !l.verbose {
		return
	}
	l.logf(format, v...)
}
