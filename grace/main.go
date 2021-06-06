package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()

	fmt.Println("the end:", ctx.Err())
}

type Number struct {
	Value int
}

func runCounter(ctx context.Context, out chan<- Number) error {
	x := 0

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("stoped at %v, err: %v", x, ctx.Err())
		default:
		}

		x++

		out <- Number{Value: x}

		time.Sleep(1e9)
	}

	return nil
}
