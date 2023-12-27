package database

import (
	"spotifyapis/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/music?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	db.AutoMigrate(&models.Track{})
}

func InsertTrack(track *models.Track) error {
	newTrack := models.Track{
		ISRC:        track.ISRC,
		Title:       track.Title,
		ArtistNames: track.ArtistNames,
		ImageURI:    track.ImageURI,
		Popularity:  track.Popularity,
	}

	result := DB.Create(&newTrack)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetTrackByISRC(isrc string) (*models.Track, error) {
	var track models.Track
	result := DB.Where("isrc = ?", isrc).First(&track)
	if result.Error != nil {
		return nil, result.Error
	}
	return &track, nil
}

func GetTracksByArtist(artist string) ([]models.Track, error) {
	var tracks []models.Track
	result := DB.Where("artist_names = ?", artist).Find(&tracks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tracks, nil
}
