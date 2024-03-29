package mysql

import (
	"context"
	"testing"
)

func TestAll(t *testing.T) {
	const dsn = "root:example@tcp(127.0.0.1:4406)/repodb"

	ctx, db := setupAndCleanup(t, dsn)

	_, _ = ctx, db
}

func setupAndCleanup(t *testing.T, dsn string) (context.Context, *DB) {
	t.Helper()

	ctx := context.Background()

	db, err := ConnectDB(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}

	createTableIfNotExists(t, ctx, db)

	t.Cleanup(func() {
		dropTable(t, ctx, db)
		err := db.Close()
		if err != nil {
			t.Fatalf("close: %v", err)
		}
	})

	return ctx, db
}

func createTableIfNotExists(t *testing.T, ctx context.Context, db *DB) {
	t.Helper()
	const query = `CREATE TABLE IF NOT EXISTS repotable (
    				id INT AUTO_INCREMENT PRIMARY KEY,
    				title VARCHAR(255) NOT NULL,
    				description TEXT NOT NULL,
    				priority TINYINT,
    				date_of_birth TIMESTAMP,
    				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			)`

	_, err := db.db.Exec(query)
	if err != nil {
		t.Fatal(err)
	}
}

func dropTable(t *testing.T, ctx context.Context, db *DB) {
	t.Helper()
	const query = `DROP TABLE repotable`

	_, err := db.db.Exec(query)
	if err != nil {
		t.Fatal(err)
	}
}
