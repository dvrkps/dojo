package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const filePath = "testdata/data.txt"

func main() {

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	ps := scanConcurrently(f)
	fmt.Println(ps)
}

func scanString(r io.Reader) Persons {

	s := bufio.NewScanner(r)

	var all Persons

	for s.Scan() {
		p, err := newPersonString(s.Text())
		if err != nil {
			continue
		}
		all = append(all, p)
	}
	if s.Err() != nil {
		fmt.Println(s.Err())
		return nil
	}
	return all
}

func scanBytes(r io.Reader) Persons {

	s := bufio.NewScanner(r)

	var all Persons

	for s.Scan() {
		p, err := newPersonBytes(s.Bytes())
		if err != nil {
			continue
		}
		all = append(all, p)
	}
	if s.Err() != nil {
		fmt.Println(s.Err())
		return nil
	}
	return all
}

// Person holds person's data.
type Person struct {
	ID   int
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%4d %-10s %4d\n", p.ID, p.Name, p.Age)
}

// Persons holds all persons.
type Persons []Person

func (ps Persons) String() string {
	all := ""
	for _, p := range ps {
		all += fmt.Sprint(p)
	}
	return all
}

func newPersonString(in string) (Person, error) {

	fields := strings.Split(in, ",")
	if len(fields) != 3 {
		return Person{}, errors.New("invalid row")
	}

	id, err := strconv.ParseInt(fields[0], 10, 64)
	if err != nil {
		return Person{}, errors.New("invalid id")
	}

	name := fields[1]

	age, err := strconv.ParseInt(fields[2], 10, 64)
	if err != nil {
		return Person{}, errors.New("invalid age")
	}

	p := Person{
		ID:   int(id),
		Name: name,
		Age:  int(age),
	}

	return p, nil
}

func newPersonBytes(in []byte) (Person, error) {

	fields := bytes.Split(in, []byte(","))
	if len(fields) != 3 {
		return Person{}, errors.New("invalid row")
	}

	id, err := strconv.ParseInt(string(fields[0]), 10, 64)
	if err != nil {
		return Person{}, errors.New("invalid id")
	}

	name := fields[1]

	age, err := strconv.ParseInt(string(fields[2]), 10, 64)
	if err != nil {
		return Person{}, errors.New("invalid age")
	}

	p := Person{
		ID:   int(id),
		Name: string(name),
		Age:  int(age),
	}

	return p, nil
}
