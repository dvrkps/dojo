package log

import (
	"io"
	"log"
)

// Log contains loggers for info and error level.
type Log struct {
	verbose bool
	ol      *log.Logger
	el      *log.Logger
}

// New creates log.
func New(stdout, stderr io.Writer) *Log {
	l := Log{}
	if stdout != nil {
		l.ol = log.New(stdout, "", 0)
	}
	if stderr != nil {
		l.el = log.New(stderr, "", 0)
	}
	return &l
}

// Logf writes formated message.
func (l *Log) Logf(format string, v ...interface{}) {
	outf(l.ol, format, v...)
}

// Errorf writes formated error message.
func (l *Log) Errorf(format string, v ...interface{}) {
	outf(l.el, format, v...)
}

func outf(l *log.Logger, format string, v ...interface{}) {
	if l != nil {
		l.Printf(format, v...)
	}
}
