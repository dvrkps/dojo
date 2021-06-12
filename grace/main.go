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

	numbers := make(chan int)

	go func() {
		var i int
		for {
			select {
			case <-ctx.Done():
				break
			default:
				numbers <- i
				i++
			}
		}
		close(numbers)
		println("generator: ", i)
	}()

	go func() {
		last := 0
		for n := range numbers {
			if n > 0 && n != last+1 {
				println(n, last)
			}
			last = n
		}
		println("producer: ", last)
	}()

	<-ctx.Done()

	err := ctx.Err()
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}
