package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

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

	err = parse(f, &d)
	if err != nil {
		return empty, fmt.Errorf("parse: %v", err)
	}

	return d, nil
}

func parse(r io.Reader, d *fileload.Data) error {
	s := bufio.NewScanner(r)

	ch := make(chan []byte)

	go func() {
		for s.Scan() {
			ch <- s.Bytes()
		}
		close(ch)

		err := s.Err()
		if err != nil {
			fmt.Printf("add: %v\n", err)
		}
	}()

	for b := range ch {
		err := d.Add(b)
		if err != nil {
			return err
		}
	}

	return nil
}
