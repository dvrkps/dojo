package main

import (
	"fmt"
	"io"
	"os"
)

type AppConfig struct {
	stdout io.Writer
	stderr io.Writer
	osargs []string
}

func main() {
	ac := AppConfig{
		stdout: os.Stdout,
		stderr: os.Stderr,
		osargs: os.Args,
	}

	code := realMain(ac)

	fmt.Println(code)
}

func realMain(app AppConfig) int {
	return 0
}
