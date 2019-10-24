package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	d := flag.String("data", "", "path to data file")
	v := flag.Bool("v", false, "verbose output")
	flag.Parse()
	os.Exit(run(*d, *v, os.Stdout, os.Stderr))
}

const (
	exitOK  = 0
	exitErr = 1
)

func run(path string, verbose bool, stdout, stderr io.Writer) int {
	lgr := newLogger(verbose, stdout, stderr)

	data, err := loadData(lgr, path)
	if err != nil {
		lgr.Errorf("load raw data: %v", err)
		return exitErr
	}

	max := len(data)

	avocet := newUAPackage(avocetName, max)
	mileusna := newUAPackage(mileusnaName, max)
	varstr := newUAPackage(varstrName, max)
	xojoc := newUAPackage(xojocName, max)

	for _, ua := range data {
		lgr.Debugf("\n%s", ua)
		showUserAgent(lgr, avocet, ua, avocetParse)
		showUserAgent(lgr, mileusna, ua, mileusnaParse)
		showUserAgent(lgr, varstr, ua, varstrParse)
		showUserAgent(lgr, xojoc, ua, xojocParse)
	}
	showTotals(lgr, avocet, mileusna, varstr, xojoc)

	return exitOK
}

func showUserAgent(lgr *Logger, pkg *uaPackage, ua string, fn func(*uaPackage, string) *uaResult) {
	r := fn(pkg, ua)
	lgr.Debugf("%s", pkg.Result(r))
}

func showTotals(lgr *Logger, pkgs ...*uaPackage) {
	lgr.Info("\n\nTotal:")
	for _, pkg := range pkgs {
		lgr.Infof("%v", pkg.Total())
	}
}
