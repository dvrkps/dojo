package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
)

func main() {
	lgr := log.New(os.Stderr, "grace: ", 0)

	err := run()
	if err != nil {
		lgr.Printf("%v", err)
		os.Exit(1)
	}

}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()

	err := ctx.Err()
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}
