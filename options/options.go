package main

import (
	"fmt"

	"github.com/dvrkps/dojo/options/server"
)

func main() {
	srv := server.New(
		server.Port(80),
		server.Host("example.com"),
		server.Port(82))
	fmt.Printf("%+v\n", srv)
}
