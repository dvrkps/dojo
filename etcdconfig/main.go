package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dvrkps/dojo/etcdconfig/myconfig"
)

func main() {
	var (
		endpoints = []string{":2379"}
	)

	cfg, err := myconfig.New(
		myconfig.Config{
			Endpoints: endpoints,
			Env:       "test",
			Service:   "app",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	global(cfg)

	defer myconfig.Close(cfg)

}

func global(c *myconfig.Client) {
	var (
		b bool
		i int
		s string
	)

	for {
		time.Sleep(1e9)
		nb, _ := c.Bool("/global/istest")
		ni, _ := c.Int("/global/port")
		ns, _ := c.String("/global/words")
		if b != nb || i != ni || s != ns {
			b = nb
			i = ni
			s = ns
			fmt.Printf("global: istest: %v port: %v words: %q \n", b, i, s)
		}

	}
}
