package spotifyapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func requestAccessToken() {
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

func GetArtistInfo() {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://api.spotify.com/v1/artists/"+"4Z8W4fKeB5YxbusRsdQVPb",
		nil,
	)
	req.Header.Set("Authorization", "Bearer BQBV-z5avnz0EhXwdNNPhGX9dax95uV3EMF8o187qXFCxJFzo5QfG02xSsOAjYD1RtIAvZG-OJxCPxSPozrbTFitGIIqFtOTFd24CpKWV6EzdtACGms")
	if err != nil {
		log.Fatal("failed to build request:", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("failed to send request:", err)
	}
	defer res.Body.Close()

	resBytes, _ := io.ReadAll(res.Body)
	fmt.Print(string(resBytes))
}
