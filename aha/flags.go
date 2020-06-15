package main

import (
	"flag"
	"io"
)

type cmdFlags struct {
	version bool
	m2      float64
	ral     float64
	chv     float64
}

func newCmdFlags(args []string, w io.Writer) (cmdFlags, error) {
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)

	fs.SetOutput(w)

	var cf cmdFlags

	fs.BoolVar(&cf.version, "version", false, "show command version")
	fs.Float64Var(&cf.m2, "m2", 0, "square meter value")
	fs.Float64Var(&cf.ral, "ral", 0, "ral value")
	fs.Float64Var(&cf.chv, "chv", 0, "square hvat value")

	err := fs.Parse(args[1:])

	return cf, err
}
