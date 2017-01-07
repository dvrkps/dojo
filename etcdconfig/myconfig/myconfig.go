package myconfig

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

const (
	dialTimeout    = 5 * time.Second
	requestTimeout = 1 * time.Second
)

// Config holds configuration data.
type Config struct {
	Endpoints []string
}

// Client is configuration client.
type Client struct {
	etcdClient *clientv3.Client
}

// New creates configuration client.
func New(cfg Config) (*Client, error) {

	ec, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: dialTimeout,
	})

	if err != nil {
		return nil, err
	}

	c := &Client{
		etcdClient: ec,
	}

	return c, nil

}

func (c *Client) Value(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := c.etcdClient.Get(ctx, key)
	if err != nil {
		return err
	}
	defer cancel()

	if len(resp.Kvs) < 1 {
		return errors.New("not exists")
	}

	v := resp.Kvs[0].Value

	fmt.Printf("%T", v)
	return nil
}

// Close closes client.
func Close(c *Client) {

	err := c.etcdClient.Close()
	if err != nil {
		// TODO(dvrkps): add better logging
		log.Print(err)
	}
}
