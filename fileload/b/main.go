package main

import (
	"bytes"
	"fmt"
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

	s := bytes.Split(c, []byte{10})

	d := fileload.NewData()

	err = parse(s, &d)
	if err != nil {
		return empty, fmt.Errorf("parse: %v", err)
	}

	return d, nil
}

func parse(s [][]byte, d *fileload.Data) error {
	for i := range s {
		if len(s[i]) == 0 {
			continue
		}

		err := d.Add(s[i])
		if err != nil {
			return fmt.Errorf("add: %v", err)
		}
	}

	return nil
}
