package myconfig

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/coreos/etcd/clientv3"
)

func setup(cli *Client) {

	teardown(cli)

	all := map[string]string{
		"/com/test/global/words":  "This is sentence.",
		"/com/test/global/port":   "1234",
		"/com/test/global/istest": "true",
		"/com/test/app/server":    "127.0.0.1",
		"/com/test/app/devil":     "666",
		"/com/test/app/secret":    "true",
	}

	for k, v := range all {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		_, err := cli.etcdClient.Put(ctx, k, v)
		cancel()
		if err != nil {
			log.Printf("setup: %q:%q err: %v", k, v, err)
		}
	}
}

func teardown(cli *Client) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := cli.etcdClient.Delete(ctx, "/com/", clientv3.WithPrefix())
	cancel()
	if err != nil {
		log.Printf("teardown: %v", err)
	}
}

var (
	testEndpoints = []string{":2379"}
)

func testClient(t *testing.T) (*Client, func()) {
	c, err := New(
		Config{
			Endpoints: testEndpoints,
			Env:       "test",
			Service:   "app",
		},
	)
	if err != nil {
		t.Errorf("testConfig: %v", err)
		return nil, func() {}
	}

	return c, func() { Close(c) }
}

func Test(t *testing.T) {
	c, close := testClient(t)
	defer close()

	setup(c)
	//prefix := "/com/test/global"
	for k, v := range c.storage.all() {
		fmt.Printf("%20v: %40v \n", k, v)
	}

}
