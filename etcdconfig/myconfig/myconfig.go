package myconfig

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path"
	"strings"
	"sync"
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
	storage       *data
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
		storage:    newData(),
	}

	if err := c.setPrefixes(cfg.Env, cfg.Service); err != nil {
		return nil, fmt.Errorf("prefixes: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go c.updateStorage(&wg, c.globalPrefix)

	go c.updateStorage(&wg, c.servicePrefix)

	wg.Wait()

	return c, nil

}

func (c *Client) String(key string) (string, bool) {
	v, ok := c.storage.get(key)
	return v, ok
}

func (c *Client) updateStorage(wg *sync.WaitGroup, prefix string) {
	defer wg.Done()
	dm, err := c.get(prefix)
	if err != nil {
		log.Printf("update client: %s: %v", prefix, err)
	}
	c.storage.update(dm)
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
		if strings.HasPrefix(k, c.globalPrefix) {
			k = "/global" + strings.TrimPrefix(k, c.globalPrefix)
		}
		k = strings.TrimPrefix(k, c.servicePrefix)

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

func checkKey(key string) error {
	if key == "" || key == "/" {
		return errors.New("empty key")
	}
	if !path.IsAbs(key) {
		return errors.New("relative key")
	}
	return nil
}
