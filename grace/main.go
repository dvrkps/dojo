package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()

	err := ctx.Err()
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}

		println(err.Error())
		stop()
		os.Exit(1)
	}

}
