package main

import "strconv"

type pointerID struct {
	id int
}

func newPointerID(id int) *pointerID {
	return &pointerID{id: id}
}

func (pid *pointerID) setID(id int) {
	pid.id = id
}

func (pid *pointerID) string() string {
	return strconv.Itoa(pid.id)
}
