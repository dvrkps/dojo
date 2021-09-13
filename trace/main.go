package main

import (
	"context"
	"errors"
	"log"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

const (
	fileName = "trace.out"
	duration = 1 * time.Millisecond
)

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
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	numbers := make(chan int)

	go procesor(&wg, numbers)

	go generator(ctx, &wg, numbers)

	wg.Wait()

	err := ctx.Err()
	if errors.Is(err, context.DeadlineExceeded) {
		return nil
	}

	return err
}

func procesor(wg *sync.WaitGroup, numbers <-chan int) {
	defer wg.Done()

	var last int
	for n := range numbers {
		last = n
	}

	println("procesor:\t", last)
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
	println("generator:\t", i)
}

func deferClose(lgr *log.Logger, f func() error) {
	err := f()
	if err != nil {
		lgr.Printf("close: %v", err)
	}
}
