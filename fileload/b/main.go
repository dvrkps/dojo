package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/dvrkps/dojo/fileload"
)

func main() {
	os.Exit(fileload.Run(parse, fileload.Rows99))
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

	var err error
	for b := range ch {
		*d, err = d.Add(b)
		if err != nil {
			return err
		}
	}

	return nil
}
