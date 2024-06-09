package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func GetDB() *sql.DB {
	connection := "user=postgres dbname=go_alura_loja password=123456 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
