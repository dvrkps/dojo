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
	// allStorage(t, c)

}

func allStorage(t *testing.T, c *Client) {
	for k, v := range c.storage.all() {
		fmt.Printf("%20s: %-s \n", k, v)
	}

}

func TestString(t *testing.T) {
	c, close := testClient(t)
	defer close()

	key := "/global/words"
	got, err := c.String(key)
	want := "This is sentence."
	if got != want || err != nil {
		t.Errorf(
			"String(%q) = %q, %v; want %q, <nil>",
			key, got, err, want)
	}
}

func TestInt(t *testing.T) {
	c, close := testClient(t)
	defer close()

	key := "/devil"
	got, err := c.Int(key)
	want := 666
	if got != want || err != nil {
		t.Errorf(
			"Int(%q) = %v, %v; want %v, <nil>",
			key, got, err, want)
	}
}

func TestBool(t *testing.T) {
	c, close := testClient(t)
	defer close()

	key := "/secret"
	got, err := c.Bool(key)
	want := true
	if got != want || err != nil {
		t.Errorf(
			"Bool(%q) = %v, %v; want %v, <nil>",
			key, got, err, want)
	}
}
