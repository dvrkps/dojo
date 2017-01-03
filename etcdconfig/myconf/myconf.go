package myconf

import (
	"time"

	"github.com/coreos/etcd/clientv3"
)

// Config holds configuration data.
type Config struct {
	Endpoints   []string
	DialTimeout time.Duration
}

// Client is configuration client.
type Client struct {
	cli *clientv3.Client
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
		cli: ec,
	}

	return c, nil

}
