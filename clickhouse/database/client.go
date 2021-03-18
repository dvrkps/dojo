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

func (c *Client) CreateIfNotExists(ctx context.Context) error {
	if c.db == nil {
		return errors.New("nil db")
	}

	const databaseQuery = "CREATE DATABASE IF NOT EXISTS dojodb"

	_, err := c.db.ExecContext(ctx, databaseQuery)
	if err != nil {
		return fmt.Errorf("create database: %v", err)
	}

	const tableQuery = `CREATE TABLE IF NOT EXISTS dojodb.dojotable
		(
			uid String,
			title String,
			date DateTime
		)
		engine = MergeTree() PARTITION BY toYYYYMM(date) ORDER BY uid SETTINGS index_granularity = 8192;`

	_, err = c.db.ExecContext(ctx, tableQuery)

	return err
}
