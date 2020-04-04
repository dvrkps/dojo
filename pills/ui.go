package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"time"
)

// Version is command version.
const Version = "0.4.11"

func main() {
	flagVersion := flag.Bool("version", false, "show version")
	flagEdit := flag.Bool("edit", false, "edit user data")
	flagUser := flag.String("user", "davor", "choose user")

	flag.Parse()

	if *flagVersion {
		fmt.Print("pills " + Version + "\n\n")

		return
	}

	log := log.New(os.Stderr, "", 0)

	const (
		exitErr  = 1
		exitUser = 2
	)

	path, err := filePath(*flagUser)
	if err != nil {
		log.Printf("file path: %v", err)
		os.Exit(exitUser)
	}

	if *flagEdit {
		err = startEditor(path)
		if err != nil {
			log.Printf("edit: %v", err)
			os.Exit(exitUser)
		}
	}

	pills, err := PillsOldWay(path, midnight(time.Now()))
	if err != nil {
		log.Printf("pills: %v", err)
		os.Exit(exitErr)
	}

	fmt.Println(pills)
}

func startEditor(path string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return errors.New("empty env editor")
	}

	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("run: %v", err)
	}

	return nil
}

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

// midnight returns date with zeroed time.
func midnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

// PillsOldWay returns loaded pills data.
func PillsOldWay(path string, date time.Time) (*Data, error) {
	fs, err := fileScanner(path)
	if err != nil {
		return nil, err
	}

	return parseFile(fs, date)
}

// parseFile returns parsed and sorted pills data.
func parseFile(s *bufio.Scanner, date time.Time) (*Data, error) {
	var d Data

	for s.Scan() {
		line := bytes.TrimSpace(s.Bytes())
		if bytes.HasPrefix(line, []byte{'/', '/'}) {
			continue
		}

		err := d.Add(line, date)
		if err != nil {
			return nil, err
		}
	}

	err := s.Err()
	if err != nil {
		return nil, err
	}

	sort.Sort(d)

	return &d, nil
}
