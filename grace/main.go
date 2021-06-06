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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()

	err := ctx.Err()
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}

		lgr.Printf("%v", err)

		stop()

		os.Exit(1)
	}

}
