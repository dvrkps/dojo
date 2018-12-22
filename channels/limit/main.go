package main

import "time"

func main() {
	const (
		cap = 4
		max = 8
	)
	ch := make(chan int, cap)
	go worker(ch)
	for i := 0; i < max; i++ {
		select {
		case ch <- i:
			println("send", i)
		default:
			println("block", i)
			ch <- i
		}
	}
	close(ch)
	println("end of sending")
	time.Sleep(3e9)
}

func worker(ch chan int) {
	for v := range ch {
		time.Sleep(1e6)
		println("receive", v)
	}
}
