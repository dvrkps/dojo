package main

import (
	"log"
	"os"
	"runtime/trace"
)

func main() {
	lgr := log.New(os.Stderr, "", 0)

	f, err := os.Create("trace.out")
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
}
