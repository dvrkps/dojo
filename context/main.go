package main

import "context"

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		println("defer")
		cancel()
	}()

	v := value{
		ctx:    ctx,
		cancel: cancel,
	}

	for n := range gen(ctx) {
		v.n = n

		println(n)

		abort(v)

		select {
		case <-ctx.Done():
			println("main done")
			return
		default:
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
			println("inc done")
			return
		case ch <- n:
			n++
		}
	}
}

type value struct {
	n      int
	ctx    context.Context
	cancel context.CancelFunc
}

func abort(v value) {

	if v.n == 5 {
		println("abort", v.n)
		v.cancel()
	}
}
