package main

import "time"

func main() {
	const (
		cap = 4
		max = 6
	)
	ch := make(chan int, cap)
	go worker(ch)
	for i := 0; i < max; i++ {
		ch <- i
		println("add", i)
	}
	close(ch)
	println("end of sending")
	time.Sleep(3e9)
}

func worker(in chan int) {
	for {
		select {
		case ch, ok := <-in:
			if !ok {
				println("chan close")
				return
			}
			time.Sleep(1e6)
			println(ch, len(in), "/", cap(in))
		default:
			println("default")
		}
	}
}
