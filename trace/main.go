package main

import (
	"log"
	"os"
	"runtime/trace"
)

const fileName = "trace.out"

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

	in := make(chan int)

	go procesor(in)

	generator(in)
}

func procesor(ch <-chan int) {
	var sum int
	var last int
	for out := range ch {
		last = out
		sum = sum + out
	}
	println(last, ": ", sum)
}

func generator(ch chan<- int) {
	var i int
	for {
		ch <- i
		i++
		if i == 12 {
			return
		}
	}

}

func deferClose(lgr *log.Logger, f func() error) {
	err := f()
	if err != nil {
		lgr.Printf("close: %v", err)
	}
}
