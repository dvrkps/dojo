package main

import (
	"log"

	satori "github.com/satori/go.uuid"
)

func main() {
	s, err := satori.NewV4()
	if err != nil {
		log.Println(err)
	}
	println(s.String())
}
