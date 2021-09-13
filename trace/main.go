package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"runtime/trace"
	"sync"
)

const fileName = "trace.out"

func main() {
	lgr := log.New(os.Stderr, "", 0)

	f, err := os.Create(fileName)
	if err != nil {
		lgr.Printf("create: %v", err)
	}
	defer deferClose(lgr, f.Close)

	err = trace.Start(f)
	if err != nil {
		lgr.Printf("start: %v", err)
	}
	defer trace.Stop()

	err = run()
	if err != nil {
		lgr.Printf("run: %v", err)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup
	wg.Add(2)

	numbers := make(chan int)

	go procesor(ctx, &wg, numbers)

	go generator(ctx, &wg, numbers)

	wg.Wait()

	err := ctx.Err()
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}

func procesor(_ context.Context, wg *sync.WaitGroup, numbers <-chan int) {
	defer wg.Done()

	var last int
	for n := range numbers {
		last = n
	}

	println("procesor: ", last)
}

func generator(ctx context.Context, wg *sync.WaitGroup, numbers chan<- int) {
	defer wg.Done()

	var i int

	var done bool
	for !done {
		select {
		case <-ctx.Done():
			done = true
		default:
			i++
			numbers <- i
		}
	}

	close(numbers)
	println("\ngenerator: ", i)
}

func deferClose(lgr *log.Logger, f func() error) {
	err := f()
	if err != nil {
		lgr.Printf("close: %v", err)
	}
}
