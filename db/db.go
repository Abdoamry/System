package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func Connect() {
	var err error
	connStr := "user=postgres password=admin dbname=myapp sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Ping failed:", err)
	}

	fmt.Println("Connected to PostgreSQL")
}
