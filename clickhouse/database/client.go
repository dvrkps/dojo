package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "embed"

	_ "github.com/ClickHouse/clickhouse-go"
)

type Client struct {
	db *sql.DB
}

func NewClient(dsn string) (*Client, error) {
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return nil, fmt.Errorf("open: %v", err)
	}

	return &Client{db: db}, nil
}

func (c *Client) Ping(ctx context.Context) error {
	if c.db == nil {
		return errors.New("nil db")
	}

	return c.db.PingContext(ctx)
}

func (c *Client) Close() error {
	if c.db == nil {
		return errors.New("nil db")
	}

	return c.db.Close()
}

//go:embed init.sql
var initSQL string

func (c *Client) CreateIfNotExists(ctx context.Context) error {
	if c.db == nil {
		return errors.New("nil db")
	}

	_, err := c.db.ExecContext(ctx, initSQL)

	return err
}
