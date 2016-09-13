package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const keyServiceID = "SERVICE_ID"

// EnvServiceID returns value from
// environment variable.
func EnvServiceID() string {
	return os.Getenv(keyServiceID)
}

const (
	minServiceID = 1
	maxServiceID = 99
)

// ServiceID returns valid ServiceID.
func ServiceID(sid string) (int, error) {

	if sid == "" {
		return 0, errors.New("empty")
	}

	id, err := strconv.Atoi(sid)
	if err != nil {
		return 0, errors.New("not a number")
	}

	if id < minServiceID || id > maxServiceID {
		return 0, fmt.Errorf("out of range (%d - %d)",
			minServiceID, maxServiceID)
	}

	return id, nil
}
