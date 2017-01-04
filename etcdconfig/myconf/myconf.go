package myconf

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// Config holds configuration data.
type Config struct {
	Endpoints      []string
	DialTimeout    time.Duration
	RequestTimeout time.Duration
}

// Client is configuration client.
type Client struct {
	etcdClient *clientv3.Client
	timeout    time.Duration
}

// New creates configuration client.
func New(cfg Config) (*Client, error) {

	ec, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: cfg.DialTimeout,
	})

	if err != nil {
		return nil, err
	}

	c := &Client{
		etcdClient: ec,
		timeout:    cfg.RequestTimeout,
	}

	return c, nil

}

func (c *Client) Value(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	resp, err := c.etcdClient.Get(ctx, key)
	if err != nil {
		return err
	}
	defer cancel()

	if len(resp.Kvs) < 1 {
		return errors.New("not exists")
	}

	value = resp.Kvs[0].Value
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
