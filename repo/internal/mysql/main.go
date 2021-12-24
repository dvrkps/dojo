package mysql

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func ConnectDB(ctx context.Context, dsn string) (*DB, error) {
	db, err := sqlx.ConnectContext(ctx, "mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect: %v", err)
	}

	return &DB{db: db}, nil
}
