package main

import (
	"bufio"
	"io"
	"log"
)

func genRows(r io.Reader) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		s := bufio.NewScanner(r)
		var row string
		for s.Scan() {
			row = s.Text()
			if row == "" {
				continue
			}
			out <- row
		}
		if err := s.Err(); err != nil {
			log.Print(err)
		}
	}()
	return out
}

func genPersons(rows <-chan string) <-chan Person {
	out := make(chan Person)
	go func() {
		defer close(out)
		for r := range rows {
			p, err := newPersonString(r)
			if err != nil {
				return
			}
			out <- p
		}
	}()
	return out
}

func scanConcurrently(r io.Reader) Persons {
	var all Persons

	rows := genRows(r)
	persons := genPersons(rows)
	for p := range persons {
		all = append(all, p)
	}
	return all
}
