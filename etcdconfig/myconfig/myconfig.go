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
	etcdClient *clientv3.Client
	keyPrefix  struct {
		global  string
		service string
	}
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

	if err := c.initKeyPrefixes(cfg.Env, cfg.Service); err != nil {
		return nil, fmt.Errorf("prefixes: %v", err)
	}

	fmt.Printf("%#v\n", c)
	return c, nil

}

const companyKey = "com"

func (c *Client) initKeyPrefixes(env, service string) error {
	if env == "" {
		return errors.New("empty env")
	}
	if service == "" {
		return errors.New("empty service")
	}
	root := fmt.Sprintf("/%s/%s", companyKey, env)
	c.keyPrefix.global = fmt.Sprintf("%s/global", root)
	c.keyPrefix.service = fmt.Sprintf("%s/%s", root, service)
	return nil
}

func (c *Client) get(prefix string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := c.etcdClient.Get(
		ctx,
		prefix,
		clientv3.WithPrefix(),
		clientv3.WithSerializable(),
		clientv3.WithSort(
			clientv3.SortByKey,
			clientv3.SortAscend,
		),
	)
	defer cancel()
	if err != nil {
		return err
	}
	defer cancel()

	if len(resp.Kvs) < 1 {
		return errors.New("not exists")
	}

	m := map[string]string{}

	var k, v string

	for _, ev := range resp.Kvs {
		k = string(ev.Key)
		v = string(ev.Value)

		m[k] = v
		//fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
	//fmt.Println(m)
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
