package myconfig

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func setup(cli *Client) {

	all := map[string]string{
		"/com/test/global/words":  "This is sentence.",
		"/com/test/global/port":   "1234",
		"/com/test/global/istest": "test",
	}

	for k, v := range all {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		_, err := cli.etcdClient.Put(ctx, k, v)
		cancel()
		if err != nil {
			log.Printf("setupTestData: %q:%q err: %v", k, v, err)
		}
	}
}

var (
	endpoints      = []string{":2379"}
	dialTimeout    = 5 * time.Second
	requestTimeout = 1 * time.Second
)

func testClient(t *testing.T) (*Client, func()) {
	c, err := New(
		Config{
			Endpoints:      endpoints,
			DialTimeout:    dialTimeout,
			RequestTimeout: requestTimeout})
	if err != nil {
		t.Errorf("testConfig: %v", err)
		return nil, func() {}
	}

	return c, func() { Close(c) }
}

func Test(t *testing.T) {
	c, close := testClient(t)
	defer close()

	key := "foo"

	var got string
	if err := c.Value(key, &got); err != nil {
		log.Print("value:", err)
	}

	want := "bar"
	if got != want {
		t.Errorf("Value(%v, ...) = %v, want %v", key, got, want)
	}
	fmt.Printf("%s: %s", key, got)

}
