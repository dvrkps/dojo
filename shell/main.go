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
	reader := bufio.NewReader(stdin)
	for {
		// Read the keyboad input.
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		cmd := strings.TrimSpace(line)
		if cmd == "quit" {
			break
		}

		_, err = fmt.Fprintln(stdout, cmd)
		if err != nil {
			return err
		}

	}
	return nil
}
