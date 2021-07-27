package fake

import (
	"context"
	"errors"

	"github.com/dvrkps/dojo/clickhouse/real"
)

type Client struct {
	PingFunc              func(context.Context) error
	CloseFunc             func() error
	CreateIfNotExistsFunc func(context.Context) error
	InsertRowFunc         func(context.Context, real.Row) error
}

func (c *Client) Close() error {
	if c.CloseFunc == nil {
		return errors.New("nil close func")
	}
	return c.CloseFunc()
}

func (c *Client) Ping(ctx context.Context) error {
	if c.PingFunc == nil {
		return errors.New("nil ping func")
	}
	return c.PingFunc(ctx)
}

func (c *Client) CreateIfNotExists(ctx context.Context) error {
	if c.CreateIfNotExistsFunc == nil {
		return errors.New("nil createIfNotExists func")
	}
	return c.CreateIfNotExistsFunc(ctx)
}

func (c *Client) InsertRow(ctx context.Context, r real.Row) error {
	if c.CreateIfNotExistsFunc == nil {
		return errors.New("nil createIfNotExists func")
	}
	return c.InsertRowFunc(ctx, r)
}
