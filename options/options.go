package main

import (
	"fmt"

	"github.com/dvrkps/dojo/options/server"
)

func main() {
	srv, err := server.New(
		server.Port(80),
		server.Host("example.com"),
		server.Port(82))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", srv)
}
