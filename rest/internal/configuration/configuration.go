// Package configuration provides support for using
// command line arguments for configuration.
package configuration

import (
	"flag"
	"io"
)

// Configuration holds configuration data.
type Configuration struct {
	Addr    string
	Verbose bool
}

const (
	defaultFlagVerbose = false
	defaultFlagAddr    = "localhost:8080"
)

// New creates configuration from flags.
func New(args []string, w io.Writer) (*Configuration, error) {
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	fs.SetOutput(w)

	var c Configuration

	fs.StringVar(&c.Addr, "addr", defaultFlagAddr, "TCP address to listen")
	fs.BoolVar(&c.Verbose, "v", defaultFlagVerbose, "verbose output")

	err := fs.Parse(args[1:])

	return &c, err
}
