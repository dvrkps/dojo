package main

import "strconv"

type client struct {
	id int
}

func newClient(id int) client {
	return client{id: id}
}

func (c *client) setID(id int) {
	c.id = id
}

func (c *client) string() string {
	return strconv.Itoa(c.id)
}
