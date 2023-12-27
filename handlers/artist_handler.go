package handlers

import (
	"encoding/json"
	"net/http"
	"spotifyapis/database"

	"github.com/gorilla/mux"
)

func GetTracksByArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	artist := vars["artist"]

	tracks, err := database.GetTracksByArtist(artist)
	if err != nil {
		http.Error(w, "Error retrieving tracks from the database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tracks)
}
