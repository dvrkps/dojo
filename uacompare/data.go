package main

import (
	"bufio"
	"fmt"
	"os"
)

// Data holds all user agents.
type Data []string

func loadData(lgr *Logger, path string) (Data, error) {

	r, err := os.Open(path)
	if err != nil {
		return Data{}, err
	}

	defer closeFile(r.Close, lgr, "low raw data close")

	s := bufio.NewScanner(r)

	var (
		d Data
		t string
	)
	for s.Scan() {
		t = s.Text()
		if t == "" {
			continue
		}
		d = append(d, t)
	}
	if err = s.Err(); err != nil {
		return Data{}, fmt.Errorf("scan: %v", s.Err())
	}
	return d, nil
}

func closeFile(fn func() error, lgr *Logger, msg string) {
	err := fn()
	if err != nil {
		lgr.Errorf("%s%v", msg, err)
	}
}

func fixUnknown(in string) string {
	if in == "Unknown" || in == "0.0.0" {
		return ""
	}
	return in
}
