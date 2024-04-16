package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	dbToken := os.Getenv("DB_TOKEN")

	connURL := fmt.Sprintf("%s?authToken=%s", dbURL, dbToken)
	db, err := sql.Open("libsql", connURL)
	if err != nil {
		log.Fatal("failed to open SQL conn", err)
	}
	err = db.Ping()

	if err != nil {
		log.Fatal("failed to ping db", err)
	}
	fmt.Println("vim-go")
}
