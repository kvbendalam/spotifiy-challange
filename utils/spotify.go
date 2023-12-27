package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
)

type SpotifyAPIResponse struct {
	Tracks struct {
		Items []SpotifyTrack `json:"items"`
	} `json:"tracks"`
}

type SpotifyTrack struct {
	Album struct {
		Artists []struct {
			Name string `json:"name"`
		} `json:"artists"`
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
	} `json:"album"`
	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
}

func getSpotifyToken(clientID, clientSecret string) (string, error) {
	authURL := "https://accounts.spotify.com/api/token"

	authHeader := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", authURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Basic "+authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var tokenResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return "", err
	}

	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("failed to retrieve access token from Spotify")
	}

	return accessToken, nil
}

func SpotifyAPI(isrc string) (*SpotifyAPIResponse, error) {
	clientID := "c7aac042e9704c869f601bdd75206b5b"
	clientSecret := "9ab2e06a7f174c74b4ee228aa7983df6"

	token, err := getSpotifyToken(clientID, clientSecret)
	if err != nil {
		return nil, err
	}

	apiURL := fmt.Sprintf("https://api.spotify.com/v1/search?type=track&q=isrc:%s", isrc)

	response, err := makeRequest("GET", apiURL, map[string]string{"Authorization": "Bearer " + token}, nil)
	if err != nil {
		return nil, err
	}

	var spotifyResponse SpotifyAPIResponse
	err = json.Unmarshal(response, &spotifyResponse)
	if err != nil {
		return nil, err
	}

	return &spotifyResponse, nil
}

func makeRequest(method, url string, headers map[string]string, body []byte) ([]byte, error) {
	client := resty.New()

	request := client.R().
		SetHeaders(headers).
		SetBody(body)

	resp, err := request.Execute(method, url)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
