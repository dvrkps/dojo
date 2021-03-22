package main

import (
	"errors"
	"fmt"
	"log"
)

func newCommand(args []string) (*Command, error) {
	const argsError = `use "build", "run" or "test".`
	if len(args) < 2 {
		return nil, errors.New(argsError)
	}

	var (
		c   Command
		err error
	)
	switch args[1] {
	case "one":
		err = newOneCommand(&c)
	case "two":
		err = newTwoCommand(&c)
	case "three":
		err = newThreeCommand(&c)
	default:
		err = errors.New(argsError)
	}

	if err != nil {
		return nil, fmt.Errorf("command: %v", err)
	}

	return &c, nil
}

type Command struct {
	Kind int
}

const (
	OneCommand = iota
	TwoCommand
	ThreeCommand
)

func newOneCommand(c *Command) error {
	log.SetPrefix("one: ")

	if c == nil {
		return errors.New("nil command")
	}

	c.Kind = OneCommand

	return nil
}

func newTwoCommand(c *Command) error {
	log.SetPrefix("two: ")

	if c == nil {
		return errors.New("nil command")
	}

	c.Kind = TwoCommand

	return nil
}

func newThreeCommand(c *Command) error {
	log.SetPrefix("three: ")

	if c == nil {
		return errors.New("nil command")
	}

	c.Kind = ThreeCommand

	return nil
}
