package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	os.Exit(run(1))
}

const (
	exitOk = iota
	exitErr
)

func run(in int) int {

	if in < 1 {
		log.Printf("%d < 1", in)
		return exitErr
	}

	fmt.Printf("in = %d\n", in)

	return exitOk
}
