package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	godotenv.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1", handleWelcome)

	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Musiqa server, let's go!")
	fmt.Println("Listening on port 8080")
	s.ListenAndServe()
}

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Musiqa API\n"))
}

func initDB(dbURL, dbToken string) (*sql.DB, error) {
	connURL := fmt.Sprintf("%s?authToken=%s", dbURL, dbToken)
	db, err := sql.Open("libsql", connURL)
	if err != nil {
		return db, fmt.Errorf("failed to open sql conn: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("", err)
		return db, fmt.Errorf("failed to ping db: %v", err)
	}

	return db, err
}
