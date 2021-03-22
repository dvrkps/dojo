package main

import (
	"context"
	"errors"
	"log"
)

func newOneCommand(c *Command) error {
	log.SetPrefix("one: ")

	if c == nil {
		return errors.New("nil command")
	}

	c.kind = OneCommand

	c.Run = runOneCommand

	return nil

}

func runOneCommand(ctx context.Context, c *Command, args []string) error {
	log.Printf("run: args: %v", args)
	return nil
}
