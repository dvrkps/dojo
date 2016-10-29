package main

import (
	"context"
	"fmt"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		println("defer")
		cancel()
	}()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}

	println("end")
}

func gen(ctx context.Context) <-chan int {

	ch := make(chan int)

	go func() {

		var n int

		for {
			select {
			case <-ctx.Done():
				println("stop")
				return
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}
