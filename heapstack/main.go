package main

const newID = 42

func main() {
	cc := newClientCopy(newID)
	setID(&cc)

	cp := newClientPointer(newID)
	setID(cp)
}

func setID(c *client) {
	c.setID(newID + 1)
}
