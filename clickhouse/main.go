package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go"
)

func main() {
	const dsn = "tcp://127.0.0.1:9000?" +
		// "database=dojodb&" +
		"password=dojopassword&" +
		"debug=true"

	connect, err := sql.Open("clickhouse", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}

}
