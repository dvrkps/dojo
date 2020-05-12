package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func openFile(path string, date time.Time) (*Data, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return parseFile(f, date)
}

func filePath(user string) (string, error) {
	d, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	p := fmt.Sprintf("%s/pills/%s/pills.txt", d, user)

	if _, err := os.Stat(p); os.IsNotExist(err) {
		return "", err
	}

	return p, nil
}

func parseFile(r io.Reader, date time.Time) (*Data, error) {
	s := bufio.NewScanner(r)

	var d Data

	var err error

	for s.Scan() {
		line := bytes.TrimSpace(s.Bytes())
		if bytes.HasPrefix(line, []byte{'/', '/'}) {
			continue
		}

		d, err = d.Add(line, date)
		if err != nil {
			return nil, err
		}
	}

	err = s.Err()
	if err != nil {
		return nil, err
	}

	d = sortData(d)

	return &d, nil
}
