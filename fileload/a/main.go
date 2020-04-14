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

	var err error

	for s.Scan() {
		*d, err = d.Add(s.Bytes())
		if err != nil {
			return fmt.Errorf("add: %v", err)
		}
	}

	return s.Err()
}
