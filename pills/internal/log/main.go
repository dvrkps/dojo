package log

import (
	"io"
	"log"
)

// A Log represents an logger.
type Log struct {
	output  *log.Logger
	verbose bool
}

// New creates logger.
func New(w io.Writer, prefix string) Log {
	l := Log{}
	if w != nil {
		l.output = log.New(w, prefix, 0)
	}

	return l
}

// SetVerbose sets log verbosity.
func (l *Log) SetVerbose(verbose bool) {
	l.verbose = verbose
}

// Printf logs formatted message.
func (l *Log) Printf(format string, v ...interface{}) {
	l.logf(format, v...)
}

// Verbosef logs formatted verbose message.
func (l *Log) Verbosef(format string, v ...interface{}) {
	if !l.verbose {
		return
	}

	l.logf(format, v...)
}

func (l *Log) logf(format string, v ...interface{}) {
	if l.output == nil {
		return
	}

	l.output.Printf(format, v...)
}
