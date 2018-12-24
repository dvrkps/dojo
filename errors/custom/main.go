package main

import (
	"log"

	"github.com/dvrkps/dojo/errors/custom/packa"
)

func main() {
	log.SetFlags(0)

	// temporary error
	err := packa.New(-1, 0)
	if packa.IsTemporary(err) {
		println("temporary")
	}
	log.Printf("packa.New(-1,0) = %v", err)
	// basic error
	err = packa.New(1, 2)
	if packa.IsTemporary(err) {
		println("temporary")
	}
	log.Printf("packa.New(1,2) = %v", err)
}
