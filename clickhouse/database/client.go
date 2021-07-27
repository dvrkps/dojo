package database

import (
	"context"
	"errors"

	"github.com/dvrkps/dojo/clickhouse/real"
)

type RealClient interface {
	Close() error
	Ping(ctx context.Context) error
	CreateIfNotExists(ctx context.Context) error
	InsertRow(ctx context.Context, r real.Row) error
}

var nilRealClientError = errors.New("nil real client")

type Client struct {
	realClient RealClient
}

func NewClient(rc RealClient) (*Client, error) {
	if rc == nil {
		return nil, nilRealClientError
	}
	return &Client{realClient: rc}, nil
}

func (c *Client) Close() error {
	if c.realClient == nil {
		return nilRealClientError
	}

	return c.realClient.Close()
}

func (c *Client) Ping(ctx context.Context) error {
	if c.realClient == nil {
		return nilRealClientError
	}

	return c.realClient.Ping(ctx)
}

func (c *Client) InsertRow(ctx context.Context, r real.Row) error {
	if c.realClient == nil {
		return nilRealClientError
	}

	return c.realClient.InsertRow(ctx, r)
}

func (c *Client) CreateIfNotExists(ctx context.Context) error {
	if c.realClient == nil {
		return nilRealClientError
	}

	return c.realClient.CreateIfNotExists(ctx)
}
