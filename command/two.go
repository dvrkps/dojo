package main

import (
	"context"
	"errors"
	"log"
)

func newTwoCommand(c *Command) error {
	log.SetPrefix("two: ")

	if c == nil {
		return errors.New("nil command")
	}

	c.kind = TwoCommand

	c.Run = runThreeCommand

	return nil
}

func runThreeCommand(ctx context.Context, c *Command, args []string) error {
	log.Printf("run: args: %v", args)
	return nil
}
