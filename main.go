package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	log.Println("Starting academiQ-api")
	connStr := "user=postgres dbname=academiq password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.Begin()

	log.Println("Connected to database")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// auth

	http.HandleFunc("/auth/login", login)

	http.ListenAndServe(":7301", nil)
}

func login(w http.ResponseWriter, r *http.Request) {

}
