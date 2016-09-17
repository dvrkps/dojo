package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	os.Exit(run(&CLI{}, 1))
}

const (
	exitOk = iota
	exitErr
)

func run(c *CLI, in int) int {

	c = NewCLI(c)

	if in < 1 {
		c.log.Printf("%d < 1", in)
		return exitErr
	}

	c.printf("in = %d\n", in)

	return exitOk
}

// CLI holds cli environment.
type CLI struct {
	out io.Writer
	log *log.Logger
}

// NewCLI creates CLI.
func NewCLI(c *CLI) *CLI {

	if c.out == nil {
		c.out = os.Stdout
	}

	if c.log == nil {
		c.log = log.New(os.Stderr, "clog:", 0)
	}

	return c

}

func (c *CLI) fprintf(w io.Writer, format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(w, format, a...)
}

func (c *CLI) printf(format string, v ...interface{}) {
	_, err := c.fprintf(c.out, format, v...)
	if err != nil {
		c.log.Print(err)
	}
}
