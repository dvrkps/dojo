package database

import (
	"testing"

	"github.com/dvrkps/dojo/clickhouse/real/fake"
)

func TestProba(t *testing.T) {
	fc := fake.Client{}
	c, err := NewClient(&fc)
	if err != nil {
		t.Fatal("new client", err)
	}

	err = c.Close()
	if err != nil {
		t.Fatal("close", err)
	}
}
