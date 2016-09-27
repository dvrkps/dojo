package main

import "os"

func main() {
	os.Exit(run())
}

const exitOk = iota

func run() int {
	return exitOk
}
