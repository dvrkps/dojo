package log

import (
	"io"
	"log"
)

// Log contains loggers for info, debug and error level.
type Log struct {
	verbose bool
	ol      *log.Logger
	el      *log.Logger
}

// New creates log.
func New(verbose bool, stdout, stderr io.Writer) *Log {
	l := Log{
		verbose: verbose,
	}
	if stdout != nil {
		l.ol = log.New(stdout, "", 0)
	}
	if stderr != nil {
		l.el = log.New(stderr, "", 0)
	}
	return &l
}

// Infof writes formated info message to log.
func (l *Log) Infof(format string, v ...interface{}) {
	outf(l.ol, format, v...)
}

// Debugf writes formated debug message to log.
func (l *Log) Debugf(format string, v ...interface{}) {
	if l.verbose {
		outf(l.ol, format, v...)
	}
}

// Errorf writes formated error message to the log.
func (l *Log) Errorf(format string, v ...interface{}) {
	outf(l.el, format, v...)
}

func outf(l *log.Logger, format string, v ...interface{}) {
	if l != nil {
		l.Printf(format, v...)
	}
}
