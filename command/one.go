package main

import (
	"errors"
	"log"
)

func newOneCommand(c *Command) error {
	log.SetPrefix("one: ")

	if c == nil {
		return errors.New("nil command")
	}

	c.kind = OneCommand

	return nil

}
