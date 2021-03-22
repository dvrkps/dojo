package fake

import (
	"context"

	"github.com/dvrkps/dojo/clickhouse/database"
)

type Client struct {
	Ping              func(context.Context) error
	Close             func() error
	CreateIfNotExists func(context.Context) error
	InsertRow         func(context.Context, database.Row) error
}
