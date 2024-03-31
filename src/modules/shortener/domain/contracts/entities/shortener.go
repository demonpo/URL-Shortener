package entities

import "time"

type Shortener struct {
	Id        string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
