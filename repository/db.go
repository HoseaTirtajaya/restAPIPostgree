package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type db struct {
	DB *sql.DB
}

const (
	hostname = "localhost"
	hostPort = 5432
	username = "postgres"
	password = "Circumstances123"
	dbname   = "golanguage"
)

//OpenDB connection
func OpenDB() (*sql.DB, error) {
	connString := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostPort, hostname, username, password, dbname)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	// fmt.Println("You are connected indeed")
	return db, nil
}
