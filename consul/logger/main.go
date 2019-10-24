package logger

import (
	"io"
	"log"
)

// Logger holds all loggers.
type Logger struct {
	verbose bool
	il      *log.Logger
	dl      *log.Logger
	el      *log.Logger
}

// New creates logger.
func New(verbose bool, stdout, stderr io.Writer) *Logger {
	l := &Logger{
		verbose: verbose,
	}
	if stdout != nil {
		l.il = log.New(stdout, "", 0)
		l.dl = log.New(stdout, "", 0)
	}
	if stderr != nil {
		l.el = log.New(stderr, "", 0)
	}
	return l
}

// Info logs info level messages.
func (l *Logger) Info(format string, v ...interface{}) {
	l.output(l.il, format, v...)
}

// Debug logs debug level messages.
func (l *Logger) Debug(format string, v ...interface{}) {
	if !l.verbose {
		return
	}
	l.output(l.dl, format, v...)
}

// Error logs error level messages.
func (l *Logger) Error(format string, v ...interface{}) {
	l.output(l.el, format, v...)
}

// SetVerbose sets verbose logging.
func (l *Logger) SetVerbose(v bool) {
	l.verbose = v
}

func (l *Logger) output(lgr *log.Logger, format string, v ...interface{}) {
	if lgr == nil {
		return
	}
	if len(v) < 1 {
		lgr.Print(format)
		return
	}
	lgr.Printf(format, v...)
}
