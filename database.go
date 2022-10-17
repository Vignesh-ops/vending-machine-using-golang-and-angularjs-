package main

import (
	"database/sql"
	"fmt"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "vicky"
	dbname   = "postgres"
)

func connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port = %d user= %s password =%s dbname = %s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	fmt.Println("The database is connected and error is ", err)
	return db
}
