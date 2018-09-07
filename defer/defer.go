package main

import (
	"fmt"
	"time"
)

func main() {
	println("start")
	now := time.Now()
	c := client{msg: "message"}

	defer trackTime(now, c.message())

	time.Sleep(50 * time.Millisecond)

	now = time.Now()

	c.msg = "new one"

	proba()

	println("end")
}

func proba() {
	time.Sleep(1e9)
}
func trackTime(start time.Time, msg string) {
	end := time.Since(start)
	fmt.Printf("tt: %v: %q\n", end, msg)
}

type client struct {
	msg string
}

func (c *client) message() string {
	println("client message run")
	return c.msg
}
