package main

import (
	"log"

	google "github.com/google/uuid"
	satori "github.com/satori/go.uuid"
)

func main() {
	s, err := satori.NewV4()
	if err != nil {
		log.Println("s:", err)
	}
	println("s: ", s.String())

	g, err := google.NewRandom()
	if err != nil {
		log.Println("g:", err)
	}

	println("g: ", g.String())
}
