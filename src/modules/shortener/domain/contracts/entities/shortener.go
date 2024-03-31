package entities

import "time"

type Shortener struct {
	Id        string    `json:"id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
