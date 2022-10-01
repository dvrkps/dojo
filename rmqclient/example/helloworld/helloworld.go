package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup
	wg.Add(2)

	const delay = time.Second

	go func() {
		defer wg.Done()
		err := mainReceive(ctx)
		if err != nil {
			log.Printf("receive: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := mainSend(ctx, delay)
		if err != nil {
			log.Printf("send: %v", err)
		}
	}()

	wg.Wait()

	err := ctx.Err()
	if !errors.Is(err, context.Canceled) {
		log.Printf("ctx: %v", err)
	}

	log.Println("the end.")
}
