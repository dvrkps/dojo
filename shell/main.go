package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	lgr := log.New(os.Stderr, "", 0)

	err := run(os.Stdin, os.Stdout)
	if err != nil {
		lgr.Printf("run: %v", err)
		os.Exit(1)
	}
}

func run(stdin io.Reader, stdout io.Writer) error {
	s := bufio.NewScanner(stdin)
	for s.Scan() {
		// Read the keyboad input.
		line := s.Text()
		cmd := strings.TrimSpace(line)
		if cmd == "exit" {
			break
		}

		_, err := fmt.Fprintln(stdout, line)
		if err != nil {
			return err
		}
	}

	err := s.Err()
	if err != nil {
		return err
	}

	return nil
}
