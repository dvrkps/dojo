package database

import (
	"database/sql"
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
