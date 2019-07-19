package main

import (
	"sync"
)

func main() {
	const (
		max       = 4
		noWorkers = 10
	)
	ch := make(chan int, 1)

	wg := startWorkers(noWorkers, ch)
	for i := 0; i < max; i++ {
		select {
		case ch <- i:
			println("send", i)
		default:
			println("*** block", i)
			ch <- i
		}
	}
	close(ch)
	println("sending done")
	wg.Wait()
	println("the end")
}

func startWorkers(max int, ch chan int) *sync.WaitGroup {
	initWG := &sync.WaitGroup{}
	initWG.Add(max)
	wg := &sync.WaitGroup{}
	wg.Add(max)
	for i := 0; i < max; i++ {
		go worker(initWG, wg, ch)
	}
	initWG.Wait()
	println("workers init")
	return wg
}

func worker(initWG *sync.WaitGroup, wg *sync.WaitGroup, ch chan int) {
	initWG.Done()
	for v := range ch {
		println("receive", v)
	}
	wg.Done()
	println("end worker")
}
