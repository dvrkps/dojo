package main

import (
	"log"
	"os"
	"runtime/trace"
)

const fileName = "trace.out"

var sum int

func main() {
	lgr := log.New(os.Stderr, "", 0)

	f, err := os.Create(fileName)
	if err != nil {
		lgr.Printf("create: %v", err)
	}
	defer func() {
		errClose := f.Close()
		if errClose != nil {
			lgr.Printf("close: %v", errClose)
		}
	}()

	err = trace.Start(f)
	if err != nil {
		lgr.Printf("start: %v", err)
	}
	defer trace.Stop()

	in := make(chan int)

	go result(sum, in)

	generator(in)

}

func result(sum int, ch <-chan int) {
	for out := range ch {
		sum = sum + out
	}
	println(sum)
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
