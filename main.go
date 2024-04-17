package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	godotenv.Load()

	getToken()

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

func getToken() {
	body := bytes.NewBufferString(
		fmt.Sprintf(
			"grant_type=client_credentials&client_id=%s&client_secret=%s",
			os.Getenv("SPOTIFY_CLIENT_ID"),
			os.Getenv("SPOTIFY_CLIENT_SECRET"),
		),
	)
	req, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", body)
	if err != nil {
		log.Fatal("failed to build request:", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("failed to send request:", err)
	}
	defer res.Body.Close()

	type spotifyAccessToken struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	var token spotifyAccessToken
	resBytes, err := io.ReadAll(res.Body)

	fmt.Println(res.Status)
	fmt.Println(string(resBytes))

	if err != nil {
		log.Fatal("failed to read response body:", err)
	}

	json.Unmarshal(resBytes, &token)
	fmt.Printf("%+v\n", token)
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
