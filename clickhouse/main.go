package main

import (
	"database/sql"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
)

func main() {
	const dsn = "tcp://127.0.0.1:9000?" +
		// "debug=true&" +
		// "database=dojodb&" +
		"password=dojopassword"

	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		log.Printf("open: %v", err)
		return
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Printf("close: %v", err)
			return
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Printf("ping: %v", err)
		return
	}

	println("done.")

}
