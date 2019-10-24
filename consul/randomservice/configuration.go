package randomservice

import (
	"errors"

	"github.com/dvrkps/consuldojo/guard"
	"github.com/dvrkps/consuldojo/logger"
)

// Configuration holds service configuration.
type Configuration struct {
	Guard      *guard.Guard
	Log        *logger.Logger
	NoWorkers  int
	MaxNumber  int
	OutputSize int
}

// Err reports configuration error.
func (c *Configuration) Err() error {
	if c.Guard == nil {
		return errors.New("nil guard")
	}
	if c.Log == nil {
		return errors.New("nil logger")
	}
	if c.NoWorkers < 1 {
		return errors.New("workers < 1")
	}
	if c.MaxNumber < 1 {
		return errors.New("MaxNumber < 1")
	}
	if c.OutputSize < 1 {
		return errors.New("OutputSize < 1")
	}
	return nil
}
