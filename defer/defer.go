package main

import (
	"log"
	"time"
)

func main() {
	defer trackTime("start msg")("end msg")
	time.Sleep(1 * time.Millisecond)
}

func trackTime(msg string) func(msg string) {
	now := time.Now()
	log.Print("start:", msg)
	return func(msg string) {
		since := time.Since(now)
		log.Printf("end: %s %s", msg, since)
	}
}
