package main

import "strconv"

type valueID struct {
	id int
}

func newValueID(id int) valueID {
	return valueID{id: id}
}

func (vid *valueID) setID(id int) {
	vid.id = id
}

func (vid *valueID) string() string {
	return strconv.Itoa(vid.id)
}
