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

	go result(in)

	var i int
	for {
		in <- i
		i++
		if i == 10 {
			return
		}
	}

}

func result(ch <-chan int) {
	for out := range ch {
		println(out)
		_ = out
	}
}
