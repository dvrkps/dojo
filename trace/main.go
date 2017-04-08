package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
)

func setupTrace(path string) func() {

	f, err := os.Create(path)
	if err != nil {
		log.Print(err)
		return func() {}
	}

	err = trace.Start(f)
	if err != nil {
		log.Println(err)
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}

	return func() {

		trace.Stop()

		if f != nil {
			err = f.Close()
			if err != nil {
				log.Println(err)
			}
		}
	}

}

func main() {
	stop := setupTrace("trace.out")
	defer stop()

	out := sq(sq(gen(2, 3)))

	for n := range out {
		fmt.Println(n)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
