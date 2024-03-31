package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name      string
	Email     string    `gorm:"unique"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt = time.Now()
	return nil
}
