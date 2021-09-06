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
	r := bufio.NewReader(stdin)
	for {
		_, err := fmt.Fprintf(stdout, "> ")
		if err != nil {
			return err
		}

		raw, err := r.ReadString('\n')
		if err != nil {
			return err
		}

		cmd := strings.TrimSpace(raw)
		if cmd == "exit" {
			break
		}

		_, err = fmt.Fprintln(stdout, cmd)
		if err != nil {
			return fmt.Errorf("write: %v", err)
		}
	}

	return nil
}
