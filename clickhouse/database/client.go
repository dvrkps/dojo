package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (c *Client) CreateIfNotExists() error {
	if c.db == nil {
		return errors.New("nil db")
	}

	const createDatabase = "CREATE DATABASE IF NOT EXISTS dojodb"

	_, err := c.db.Exec(createDatabase)

	return err
}
