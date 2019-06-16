package main

import "strconv"

type client struct {
	id int
}

func newClientCopy(id int) client {
	return client{id: id}
}

func newClientPointer(id int) *client {
	c := client{id: id}
	return &c
}

func (c *client) setID(id int) {
	c.id = id
}

func (c *client) string() string {
	return strconv.Itoa(c.id)
}
