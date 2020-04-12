package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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

	c, err := ioutil.ReadFile(fileload.Path)
	if err != nil {
		return empty, fmt.Errorf("read file: %v", err)
	}

	r := bytes.NewReader(c)

	d := fileload.NewData()

	err = parse(r, &d)
	if err != nil {
		return empty, fmt.Errorf("parse: %v", err)
	}

	return d, nil
}

func parse(r io.Reader, d *fileload.Data) error {
	s := bufio.NewScanner(r)

	for s.Scan() {
		err := d.Add(s.Bytes())
		if err != nil {
			return fmt.Errorf("add: %v", err)
		}
	}

	return s.Err()
}
