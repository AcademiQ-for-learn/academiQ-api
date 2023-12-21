package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// DB is the database connection
var DB *sql.DB

// Connect connects to the database
func Connect() {

	log.Println("Starting database connection")
	connStr := "user=postgres dbname=academiq password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.Begin()

	log.Println("Connected to database")

	DB = db
}
