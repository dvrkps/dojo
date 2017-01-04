package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dvrkps/dojo/etcdconfig/myconf"
)

func main() {
	var (
		endpoints      = []string{":2379"}
		dialTimeout    = 5 * time.Second
		requestTimeout = 1 * time.Second
	)

	cfg, err := myconf.New(
		myconf.Config{
			Endpoints:      endpoints,
			DialTimeout:    dialTimeout,
			RequestTimeout: requestTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer myconf.Close(cfg)

	key := "foo"

	var val string
	if err := cfg.Value(key, &val); err != nil {
		log.Print("value:", err)
	}
	fmt.Printf("%s: %s", key, val)
}
