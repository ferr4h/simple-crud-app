package domain

import "errors"

var (
	GameNotFound = errors.New("Game not found")
)

type Game struct {
	ID              int64   `json:"id"`
	Title           string  `json:"title"`
	Developer       string  `json:"developer"`
	Publisher       string  `json:"publisher"`
	Genre           string  `json:"genre"`
	PublicationDate string  `json:"publication_date"`
	Rating          float32 `json:"rating"`
}

type GameInput struct {
	Title           *string  `json:"title"`
	Developer       *string  `json:"developer"`
	Publisher       *string  `json:"publisher"`
	Genre           *string  `json:"genre"`
	PublicationDate *string  `json:"publication_date"`
	Rating          *float32 `json:"rating"`
}
