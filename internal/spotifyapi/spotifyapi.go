package spotifyapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Service struct {
	spotifyClientID     string
	spotifyClientSecret string
	accessToken         string
}

func NewService(spotifyClientID, spotifyClientSecret string) *Service {
	return &Service{
		spotifyClientID:     spotifyClientID,
		spotifyClientSecret: spotifyClientSecret,
	}
}

func (s *Service) requestAccessToken() {
	body := bytes.NewBufferString(
		fmt.Sprintf(
			"grant_type=client_credentials&client_id=%s&client_secret=%s",
			s.spotifyClientID,
			s.spotifyClientSecret,
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

func (s *Service) GetArtistInfo(artistID string) {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://api.spotify.com/v1/artists/"+artistID,
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+s.accessToken)
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
