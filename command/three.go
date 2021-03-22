package main

import (
	"context"
	"errors"
	"log"
)

func newThreeCommand(c *Command) error {
	log.SetPrefix("three: ")

	if c == nil {
		return errors.New("nil command")
	}

	c.kind = ThreeCommand

	c.Run = runTwoCommand

	return nil
}

func runTwoCommand(ctx context.Context, c *Command, args []string) error {
	log.Printf("run: args: %v", args)
	return nil
}
