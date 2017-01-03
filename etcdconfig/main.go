package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	endpoints := []string{
		":2379",
	}
	dialTimeout := 5 * time.Second
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	/*
		_, err = cli.Put(context.TODO(), "foo", "bar")
		if err != nil {
			log.Fatal(err)
		}
	*/

	requestTimeout := 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, "fo", clientv3.WithPrefix())
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n\n", resp)
	for _, ev := range resp.Kvs {
		fmt.Printf("%#v : %#v \n", ev.Key, ev.Value)
	}
}
