package dbconfig

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectDB(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	fmt.Println("Connected to DB")
	return db
}
