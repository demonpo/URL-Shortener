package entities

import "time"

type Click struct {
	Id          string    `json:"id"`
	ShortenerId string    `json:"shortenerId"`
	UserIp      *string   `json:"userIp"`
	ReferrerUrl *string   `json:"referrerUrl"`
	UserAgent   *string   `json:"userAgent"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
