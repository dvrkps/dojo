package main

import (
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
)

func newConsulClient(addr string) (*api.Client, error) {
	config := api.DefaultConfig()
	config.Address = addr
	return api.NewClient(config)
}

func consulClusterAddress() string {
	const (
		host = "127.0.0.1"
		port = "8500"
	)
	return net.JoinHostPort(host, port)
}

const defaultKeyPrefix = "app"

func consulSetupKV(c *api.Client) (func() error, error) {

	fail := func(err error) (func() error, error) {
		empty := func() error { return nil }
		return empty, err
	}

	clear := func() error {
		_, err := c.KV().DeleteTree(defaultKeyPrefix, nil)
		return err
	}

	err := clear()
	if err != nil {
		err = fmt.Errorf("init clear fail: %v", err)
		fail(err)
	}

	pairs := []struct {
		k string
		v string
	}{
		{"k1", "v1"},
		{"k2", "v2"},
	}
	for _, p := range pairs {
		kv := &api.KVPair{
			Key:   fmt.Sprintf("%s/%s", defaultKeyPrefix, p.k),
			Value: []byte(p.v),
		}
		_, err := c.KV().Put(kv, nil)
		if err != nil {
			err = fmt.Errorf("put for key %v fail: %v",
				kv.Key,
				err)
			return fail(err)
		}
	}
	return clear, nil
}
