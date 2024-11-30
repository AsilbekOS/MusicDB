package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

// AccessToken javob strukturasini olish
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// Qo'shiq ma'lumotlarini olish
type SearchResponse struct {
	Tracks struct {
		Items []struct {
			Name  string `json:"name"`
			Album struct {
				ReleaseDate string `json:"release_date"`
			} `json:"album"`
			ExternalURLs struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
		} `json:"items"`
	} `json:"tracks"`
}

// Qo'shiqni qidirish
func SearchSong(songName, token string) (SearchResponse, error) {
	// URL ni to'g'irlash
	searchURL := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track&limit=1", url.QueryEscape(songName))

	// So'rov yuborish
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return SearchResponse{}, err
	}

	// Header'ni qo'shish
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return SearchResponse{}, err
	}
	defer resp.Body.Close()

	// Javobni o'qish
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SearchResponse{}, err
	}

	// Qo'shiq ma'lumotlarini olish
	var searchResp SearchResponse
	if err := json.Unmarshal(body, &searchResp); err != nil {
		return SearchResponse{}, err
	}

	return searchResp, nil
}

// Access Token olish
func GetAccessToken() (string, error) {
	// Client ID va Client Secret
	clientID := os.Getenv("clientID")
	if clientID == "" {
		log.Fatal("clientID is not set in environment variables")
	}

	clientSecret := os.Getenv("clientSecret")
	if clientSecret == "" {
		log.Fatal("clientSecret is not set in environment variables")
	}

	// Client ID va Client Secret ni Base64 bilan kodlash
	auth := fmt.Sprintf("%s:%s", clientID, clientSecret)
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	// Token olish uchun so'rov yuborish
	url := "https://accounts.spotify.com/api/token"
	data := "grant_type=client_credentials"

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return "", err
	}

	// Header'ni qo'shish
	req.Header.Add("Authorization", "Basic "+encodedAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Javobni o'qish
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Access Token ni olish
	var tokenResp AccessTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", err
	}

	return tokenResp.AccessToken, nil
}
