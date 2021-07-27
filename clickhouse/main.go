package main

import (
	"context"
	"log"
	"time"

	"github.com/dvrkps/dojo/clickhouse/database"
	"github.com/dvrkps/dojo/clickhouse/real"
)

func main() {
	const dsn = "tcp://127.0.0.1:9000?" +
		// "debug=true&" +
		// "database=dojodb&" +
		"password=dojopassword"

	rc, err := real.NewClient(dsn)
	if err != nil {
		log.Printf("real client new: %v", err)
		return
	}

	c, err := database.NewClient(rc)
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

	n := time.Now().UTC()

	r := database.Row{
		UID:   n.Format("20060102150405"),
		Title: n.Format(time.RFC3339),
		Date:  n,
	}

	err = c.InsertRow(ctx, r)
	if err != nil {
		log.Printf("insert row: %v", err)
		return
	}

	println("done.")
}
