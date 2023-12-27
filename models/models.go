package models

type Track struct {
	ID          uint   `gorm:"primaryKey"`
	ISRC        string `gorm:"unique"`
	Title       string
	ArtistNames string
	ImageURI    string
	Popularity  int
}
