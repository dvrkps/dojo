package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/dvrkps/dojo/fileload"
)

func main() {
	d, err := run()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	fmt.Print(&d)
}

func run() (fileload.Data, error) {
	empty := fileload.Data{}

	f, err := os.Open(fileload.Path)
	if err != nil {
		return empty, fmt.Errorf("open: %v", err)
	}
	defer f.Close()

	d := fileload.NewData()

	ch := make(chan []byte)

	var wg sync.WaitGroup

	const steps = 2

	wg.Add(steps)

	go func() {
		defer wg.Done()

		for b := range ch {
			err := d.Add(b)
			if err != nil {
				fmt.Printf("add: %v\n", err)
			}
		}
	}()

	err = parse(&wg, f, ch)
	if err != nil {
		return empty, fmt.Errorf("parse: %v", err)
	}

	wg.Wait()

	return d, nil
}

func parse(wg *sync.WaitGroup, r io.Reader, ch chan []byte) error {
	defer wg.Done()

	s := bufio.NewScanner(r)

	for s.Scan() {
		ch <- s.Bytes()
	}

	close(ch)

	return s.Err()
}
