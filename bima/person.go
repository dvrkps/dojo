package main

import (
	"errors"
	"fmt"
	"strings"
)

// Person holds person's height and weight.
type Person struct {
	h   float64
	w   float64
	err error
}

// NewPerson creates person.
func NewPerson(h, w float64) (Person, error) {
	if err := checkValues(h, w); err != nil {
		return Person{}, fmt.Errorf("new person: %v", err)
	}
	return Person{w: w, h: h}, nil
}

func (p Person) Err() error {
	return checkValues(p.h, p.w)
}

func checkValue(kind string, v, min, max float64) error {
	const f = "%s %.2f: not in range from %.2f to %.2f"
	if v < min || v > max {
		return fmt.Errorf(f, kind, v, min, max)
	}
	return nil
}

func checkValues(h, w float64) error {
	const msgPattern = "%s %.2f: not in range from %.2f to %.2f"
	var errMsgs []string
	if !isValueValid(h, minHeight, maxHeight) {
		msg := fmt.Sprintf(msgPattern, "height", h, minHeight, maxHeight)
		errMsgs = append(errMsgs, msg)
	}
	if !isValueValid(w, minWeight, maxWeight) {
		msg := fmt.Sprintf(msgPattern, "weight", w, minWeight, maxWeight)
		errMsgs = append(errMsgs, msg)
	}
	if len(errMsgs) > 0 {
		return errors.New(strings.Join(errMsgs, ", "))
	}
	return nil
}

func isValueValid(v, min, max float64) bool {
	if v < min || v > max {
		return false
	}
	return true
}
