package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 {
		log.Println("no arguments")
		return
	}
	command := os.Args[1]

	const max = 10
	defer over(max)

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	var i int
loop:
	for {
		select {
		case <-s:
			log.Printf("%s: signal: %v", command, s)
			break loop
		case <-time.After(1 * time.Second):
			i++
			fmt.Printf("%s: timeout: %v\n", command, i)
		}
	}

	fmt.Println("end main")
}

func over(max int) {
	fmt.Println("start defer")
	for i := 0; i <= max; i++ {
		log.Printf("defer: %v/%v", i, max)
		time.Sleep(1e6)
	}
	fmt.Println("end defer")
}
