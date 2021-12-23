package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func NewDB(dsn string) (*DB, error) {
	db := DB{}

	return &db, nil
}
