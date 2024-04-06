package entities

import (
	"gorm.io/gorm"
	"time"
)

type Click struct {
	gorm.Model
	ShortenerId string
	UserIp      *string
	ReferrerUrl *string
	UserAgent   *string
}

func (click *Click) BeforeCreate(tx *gorm.DB) (err error) {
	click.CreatedAt = time.Now()
	click.UpdatedAt = time.Now()
	return nil
}

func (click *Click) BeforeUpdate(tx *gorm.DB) (err error) {
	click.UpdatedAt = time.Now()
	return nil
}
