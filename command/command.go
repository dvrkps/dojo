package main

import (
	"errors"
	"log"
)

func newCommand(args []string) (*Command, error) {
	var argsError = errors.New(`use "one", "two" or "three".`)

	if len(args) < 2 {
		return nil, argsError
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
		err = argsError
	}

	if err != nil {
		return nil, err
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
