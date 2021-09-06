package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main() {
	lgr := log.New(os.Stderr, "", 0)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := run(ctx, os.Stdin, os.Stdout)
	if err != nil {
		if !errors.Is(err, context.Canceled) {
			lgr.Printf("run: %v", err)
			os.Exit(1)
		}
	}
}

func run(ctx context.Context, stdin io.Reader, stdout io.Writer) error {
	r := bufio.NewReader(stdin)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
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
				return nil
			}

			_, err = fmt.Fprintln(stdout, cmd)
			if err != nil {
				return fmt.Errorf("write: %v", err)
			}
		}
	}
}
