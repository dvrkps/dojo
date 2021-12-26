package main

import (
	"context"
	"log"

	"github.com/dvrkps/dojo/repo/internal/mysql"
)

func main() {
	const repoDSN = "root:example@tcp(127.0.0.1:4406)/repodb"

	db, err := mysql.ConnectDB(context.Background(), repoDSN)
	if err != nil {
		log.Printf("connect: %v", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			log.Printf("close: %v", err)
		}
	}()

	println("done")
}
