package handlers

import (
	"encoding/json"
	"net/http"

	"spotifyapis/models"
	"spotifyapis/utils"

	"spotifyapis/database"

	"github.com/gorilla/mux"
)

func CreateTrack(w http.ResponseWriter, r *http.Request) {
	isrc := "USVT10300001"

	spotifyResponse, err := utils.SpotifyAPI(isrc)
	if err != nil {
		http.Error(w, "Error fetching data from Spotify API", http.StatusInternalServerError)
		return
	}

	firstItem := spotifyResponse.Tracks.Items[0]

	track := models.Track{
		ISRC:        isrc,
		Title:       firstItem.Name,
		ArtistNames: firstItem.Album.Artists[0].Name,
		ImageURI:    firstItem.Album.Images[0].URL,
		Popularity:  firstItem.Popularity,
	}

	err = database.InsertTrack(&track)
	if err != nil {
		http.Error(w, "Error inserting data into the database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spotifyResponse)

	w.Write([]byte("Track created successfully"))
}

func GetTrackByISRC(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isrc := vars["isrc"]

	track, err := database.GetTrackByISRC(isrc)
	if err != nil {
		http.Error(w, "Error retrieving track from the database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(track)
}
