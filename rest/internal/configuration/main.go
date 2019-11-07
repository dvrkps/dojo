// Package configuration provides support for using
// command line arguments for configuration.
package configuration

import (
	"flag"
	"io"
)

// Configuration holds configuration data.
type Configuration struct {
	Debug bool
	Port  int
}

const (
	defaultFlagDebug = false
	defaultFlagPort  = 8080
)

// New creates configuration from flags.
func New(args []string, w io.Writer) (Configuration, error) {
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	fs.SetOutput(w)

	var c Configuration

	fs.BoolVar(&c.Debug, "debug", defaultFlagDebug, "verbose output")
	fs.IntVar(&c.Port, "port", defaultFlagPort, "server port number")

	err := fs.Parse(args[1:])

	return c, err
}
