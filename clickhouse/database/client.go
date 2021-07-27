package database

import (
	"context"
	"errors"
)

type RealClient interface {
	Close() error
	Ping(ctx context.Context) error
	InsertRow(ctx context.Context, r Row) error
	CreateIfNotExists(ctx context.Context) error
}

type Client struct {
	realClient RealClient
}

func NewClient(rc RealClient) (*Client, error) {
	if rc == nil {
		return nil, errors.New("nil real client")
	}
	return &Client{realClient: rc}, nil
}
