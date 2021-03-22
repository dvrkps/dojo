package main

import (
	"errors"
	"log"
)

func newThreeCommand(c *Command) error {
	log.SetPrefix("three: ")

	if c == nil {
		return errors.New("nil command")
	}

	c.kind = ThreeCommand

	return nil
}
