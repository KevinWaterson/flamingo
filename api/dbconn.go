package main

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
 )

const (
	DB_TYPE = "postgres"
	DB_NAME = "wwwlive_new"
	DB_USER = "db_username"
	DB_PASS = "db_password"
	DB_PORT = "5432"
)

// open db
func OpenDB() *sql.DB {
	// db, err := sql.Open("postgres", DB_USER+":"+DB_PASS+"@/"+DB_NAME)
	// db, err := sql.Open("postgres", "user=db_username dbname=db_password sslmode=verify-full")
	db, err := sql.Open("postgres", "user=db_username dbname=flamingo sslmode=disable")
	if err != nil {
		log.Print("Error: Connecting\n")
		panic(err)
	}
	return db
}
