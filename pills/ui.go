package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

// Version is command version.
const Version = "0.4.11"

func filePath(user string) (string, error) {
	d, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	p := fmt.Sprintf("%s/pills/%s/pills.txt", d, user)
	return p, nil
}

// fileScanner converts file content to scanner.
func fileScanner(path string) (*bufio.Scanner, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(bytes.NewReader(c)), nil
}

func main() {
	fmt.Print("pills " + Version + "\n\n")

	path, err := filePath("davor")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	pills, err := PillsOldWay(path, midnight(time.Now()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(pills)
}

// midnight returns date with zeroed time.
func midnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

// PillsOldWay returns loaded pills data.
func PillsOldWay(path string, date time.Time) (Data, error) {
	fs, err := fileScanner(path)
	if err != nil {
		return Data{}, err
	}
	// parse data
	pills := parseFile(fs, date)
	return pills, nil
}

// parseFile returns parsed and sorted pills data.
func parseFile(s *bufio.Scanner, date time.Time) Data {
	var d Data
	for s.Scan() {
		line := bytes.TrimSpace(s.Bytes())
		if bytes.HasPrefix(line, []byte{'/', '/'}) {
			continue
		}
		err := d.Add(line, date)
		if err != nil || s.Err() != nil {
			return d
		}
	}
	sort.Sort(d)
	return d
}
