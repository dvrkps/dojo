package main

const newID = 42

//go:noinline
func main() {
	c := newClientCopy(newID)
	setID(&c)
}

//go:noinline
func setID(c *client) {
	c.setID(newID + 1)
}

type client struct {
	id int
}

//go:noinline
func newClient(id int) client {
	return client{id: id}
}

//go:noinline
func (c *client) setID(id int) {
	c.id = id
}
