package main

import (
	"context"
	"errors"
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
	kind int
	Run  func(context.Context, *Command, []string) error
}

const (
	OneCommand = iota
	TwoCommand
	ThreeCommand
)

func (c *Command) Is(kind int) bool {
	return c.kind == kind
}
