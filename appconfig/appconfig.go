package main

import (
	"fmt"
	"io"
	"os"
)

type appConfig struct {
	osargs []string
	stdout io.Writer
	stderr io.Writer
}

func runApp(cfg *appConfig) int {
	fmt.Println(cfg.osargs)
	return 0
}

func main() {
	code := runApp(&appConfig{
		osargs: os.Args,
		stdout: os.Stdout,
		stderr: os.Stderr,
	})

	fmt.Println(code)
}
