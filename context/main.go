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
		if abort(n) {
			println("abort")
			cancel()
			break
		}
	}

	println("end")
}

func gen(ctx context.Context) <-chan int {

	ch := make(chan int)

	go inc(ctx, ch)

	return ch
}

func inc(ctx context.Context, ch chan int) {

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
}

func abort(i int) bool {
	if i != 5 {
		return false
	}
	return true
}
