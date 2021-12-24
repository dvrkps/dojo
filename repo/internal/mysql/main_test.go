package mysql

import (
	"context"
	"testing"
)

func TestAll(t *testing.T) {
	ctx := context.Background()

	const dsn = "root:example@tcp(127.0.0.1:4406)/repodb"
	db, err := ConnectDB(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}

	createTableIfNotExists(t, ctx, db)

	defer func() {
		dropTable(t, ctx, db)
		err = db.db.Close()
		if err != nil {
			t.Fatalf("close: %v", err)
		}
	}()
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
