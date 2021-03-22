package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("command: ")

	log.Println("setup")

	c, err := newCommand(os.Args)

	if err != nil {
		log.Printf("new: %v", err)

		return
	}

	if c.Is(OneCommand) {
		log.Println("k1")
	}

	if c.Is(TwoCommand) {
		log.Println("k2")
	}

	if c.Is(ThreeCommand) {
		log.Println("k3")
	}

	log.Println("done")
}
