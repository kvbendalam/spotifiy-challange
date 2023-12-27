package main

import (
	"net/http"
	"spotifyapis/database"

	"spotifyapis/handlers"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/tracks", handlers.CreateTrack).Methods("POST")
	r.HandleFunc("/tracks/{isrc}", handlers.GetTrackByISRC).Methods("GET")
	r.HandleFunc("/tracks/artist/{artist}", handlers.GetTracksByArtist).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
