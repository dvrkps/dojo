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
	Env       string
	Service   string
}

// Client is configuration client.
type Client struct {
	etcdClient    *clientv3.Client
	globalPrefix  string
	servicePrefix string
	data          *data
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
		data:       newData(),
	}

	if err := c.setPrefixes(cfg.Env, cfg.Service); err != nil {
		return nil, fmt.Errorf("prefixes: %v", err)
	}

	dm, err := c.get(c.globalPrefix)

	c.data.update(dm)
	return c, nil

}

const companyKey = "com"

func (c *Client) setPrefixes(env, service string) error {
	if env == "" {
		return errors.New("empty env")
	}
	if service == "" {
		return errors.New("empty service")
	}
	root := fmt.Sprintf("/%s/%s", companyKey, env)
	c.globalPrefix = fmt.Sprintf("%s/global", root)
	c.servicePrefix = fmt.Sprintf("%s/%s", root, service)
	return nil
}

func (c *Client) get(prefix string) (dataMap, error) {

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := c.etcdClient.Get(
		ctx,
		prefix,
		clientv3.WithPrefix(),
		clientv3.WithSerializable(),
	)

	defer cancel()

	dm := dataMap{}

	if err != nil {
		return dm, err
	}

	if len(resp.Kvs) < 1 {
		return dm, errors.New("not exists")
	}

	var k, v string

	for _, ev := range resp.Kvs {
		k = string(ev.Key)
		v = string(ev.Value)

		dm[k] = v
	}
	return dm, nil
}

// Close closes client.
func Close(c *Client) {

	err := c.etcdClient.Close()
	if err != nil {
		// TODO(dvrkps): add better logging
		log.Print(err)
	}
}
