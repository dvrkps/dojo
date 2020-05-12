package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

// Version is command version.
const Version = "0.6.2"

func main() {
	flagVersion := flag.Bool("version", false, "show version")
	flagEdit := flag.Bool("edit", false, "edit user data")
	flagUser := flag.String("user", "davor", "choose user")

	flag.Parse()

	if *flagVersion {
		fmt.Printf("pills v%s\n", Version)

		return
	}

	log := log.New(os.Stderr, "", 0)

	const (
		exitErr  = 1
		exitUser = 2
	)

	path, err := filePath(*flagUser)
	if err != nil {
		log.Printf("file path: %v", err)
		os.Exit(exitUser)
	}

	if *flagEdit {
		err = startEditor(path)
		if err != nil {
			log.Printf("edit: %v", err)
			os.Exit(exitUser)
		}
	}

	pills, err := openFile(path, time.Now())
	if err != nil {
		log.Printf("pills: %v", err)
		os.Exit(exitErr)
	}

	fmt.Println(pills)
}
