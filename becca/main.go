package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Usage holds command usage description.
const Usage = `Usage
	becca [options]... [values]...

Options:
	-s	source directory
	-d	destination directory
`

// Version is command version.
const Version = "0.2.5"

func main() {
	timeNow := time.Now()

	// flags
	src := flag.String("s", "", "source directory")
	dest := flag.String("d", "", "destination directory")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, Usage)
	}

	// output
	fmt.Printf("becca %s\n\n", Version)
	flag.Parse()

	// args
	if err := validArgs(*src, *dest); err != nil {
		fmt.Println(fmt.Errorf("args: %s", err))
		flag.Usage()
	}

	// tar -cvpzf backup.tar.gz /your/folder
	beccaFile := *dest + "/" + newFilename(timeNow)
	if err := runCmd("tar", "-cvpzf", beccaFile, *src); err != nil {
		fmt.Printf("cmd: %v\n", err)
	}

	fmt.Printf("elapsed time: %v\n", time.Since(timeNow))
}

// newFilename returns filename
// from current timestamp.
func newFilename(now time.Time) string {
	return now.Format("becca-20060102-150405.tar.gz")
}

// runCmd run command.
func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	//cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// validArgs validate source and
// destination flag values.
func validArgs(s, d string) error {
	// source
	if err := validPath(s); err != nil {
		return fmt.Errorf("source %s: %s", s, err)
	}
	// dest
	if err := validPath(d); err != nil {
		return fmt.Errorf("destination %s: %s", d, err)
	}
	// src and dest must not be same
	// dest must not be source subdirectory
	if strings.HasPrefix(d, s) {
		return errors.New("same or subdirectory error")
	}
	return nil
}

// validPath validate path.
// Path is valid if is absolute
// and if exists.
func validPath(path string) error {
	p, err := os.Stat(path)
	if err != nil {
		return errors.New("not exists")
	}
	if !p.IsDir() {
		return errors.New("not directory")
	}
	if !strings.HasPrefix(path, "/") {
		return errors.New("not absolute path")
	}
	return nil
}
