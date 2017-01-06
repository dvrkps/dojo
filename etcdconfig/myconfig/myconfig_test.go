package myconfig

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"
)

func proba(v interface{}) error {
	//data := "5"
	vof := reflect.ValueOf(v)
	fmt.Println(vof)
	return nil
}

func TestProba(t *testing.T) {
	var v int
	err := proba(&v)
	println(err)
	println(v)
}

func setup(t *testing.T) (*Client, func()) {
	var (
		endpoints      = []string{":2379"}
		dialTimeout    = 5 * time.Second
		requestTimeout = 1 * time.Second
	)

	cfg, err := New(
		Config{
			Endpoints:      endpoints,
			DialTimeout:    dialTimeout,
			RequestTimeout: requestTimeout})
	if err != nil {
		t.Errorf("setup: new: %v", err)
		return nil, func() {}
	}

	return cfg, func() { Close(cfg) }
}

func Test(t *testing.T) {
	cfg, teardown := setup(t)
	defer teardown()

	key := "foo"

	var got string
	if err := cfg.Value(key, &got); err != nil {
		log.Print("value:", err)
	}

	want := "bar"
	if got != want {
		t.Errorf("Value(%v, ...) = %v, want %v", key, got, want)
	}
	fmt.Printf("%s: %s", key, got)

}
