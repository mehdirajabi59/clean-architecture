package datastore

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SetDB() (*sql.DB, error) {
	//Connect to db

	conn, err := sql.Open("mysql", "root:s321d#asdASD@/clean")

	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close()

	err = conn.Ping()

	if err != nil {
		log.Fatal(err)
	}
	return conn, nil
}
