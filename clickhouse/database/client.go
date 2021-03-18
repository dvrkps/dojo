package database

import (
	"database/sql"
)

type Client struct {
	db *sql.DB
}

func NewClient(dsn string) (*Client, error) {
	return nil, nil
}
