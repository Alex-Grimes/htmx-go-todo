package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Setup() {

	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=todo password=todo123 dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println("Could not connect to postgres database", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Could not ping to postgres database", err)
	}
}
