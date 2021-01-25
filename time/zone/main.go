package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	localNow := time.Now()
	fmt.Println(localNow)

	utcNow := time.Now().UTC()
	fmt.Println(utcNow)

	newYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Println(err)

		return
	}

	nyNow := utcNow.In(newYork)
	fmt.Println(nyNow)
}
