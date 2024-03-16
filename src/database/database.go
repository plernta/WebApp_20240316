package database

//usually the name of package is the same as the name of directory

import (
	"database/sql"

	_ "github.com/lib/pq"
	// _ will used this library for sure, don't delete
)

func Connect(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	return db, err
}
