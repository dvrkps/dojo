package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ch := make(chan Number)

	go func() {
		err := runCounter(ctx, ch)
		if err != nil {
			log.Printf("counter: %v", err)
		}
	}()

	go func() {
		for v := range ch {
			println(v.Value)
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		stop()
	}
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
