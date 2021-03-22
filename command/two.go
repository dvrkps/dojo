package main

import (
	"errors"
	"log"
)

func newTwoCommand(c *Command) error {
	log.SetPrefix("two: ")

	if c == nil {
		return errors.New("nil command")
	}

	c.kind = TwoCommand

	return nil
}
