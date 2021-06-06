package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()

	err := ctx.Err()
	if err != nil {
		fmt.Println("the end:", ctx.Err())
	}

}
