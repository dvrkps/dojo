package minimal

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
