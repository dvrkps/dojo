package main

import (
	"context"
	"log"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/dvrkps/dojo/clickhouse/database"
)

func main() {
	const dsn = "tcp://127.0.0.1:9000?" +
		// "debug=true&" +
		// "database=dojodb&" +
		"password=dojopassword"

	c, err := database.NewClient(dsn)

	if err != nil {
		log.Printf("client new: %v", err)
		return
	}

	defer func() {
		err := c.Close()
		if err != nil {
			log.Printf("client close: %v", err)
			return
		}
	}()

	const pingTimeout = 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	err = c.Ping(ctx)
	if err != nil {
		log.Printf("ping: %v", err)
		return
	}

	err = c.CreateIfNotExists(ctx)
	if err != nil {
		log.Printf("create if not exists: %v", err)
		return
	}

	println("done.")
}
