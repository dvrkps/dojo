package clickhouse

import (
	"database/sql"
)

// Client is Clickhouse client.
type Client struct {
	db *sql.DB
}
