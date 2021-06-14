package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(2)

	numbers := make(chan int)

	go runGenerator(ctx, &wg, numbers)

	go runProducer(ctx, &wg, numbers)

	wg.Wait()

	err := ctx.Err()
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}

func runGenerator(ctx context.Context, wg *sync.WaitGroup, numbers chan<- int) {
	defer wg.Done()

	var i int

	var done bool
	for {
		if done {
			break
		}

		select {
		case <-ctx.Done():
			done = true
		default:
			i++
			numbers <- i
		}
	}

	close(numbers)
	println("generator: ", i)
}

func runProducer(ctx context.Context, wg *sync.WaitGroup, numbers <-chan int) {
	defer wg.Done()

	last := 0
	for n := range numbers {
		last = n
	}

	println("producer: ", last)
}
