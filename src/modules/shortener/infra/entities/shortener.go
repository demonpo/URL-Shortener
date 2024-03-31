package entities

import (
	"crypto/rand"
	"gorm.io/gorm"
	"math/big"
	"time"
)

type Shortener struct {
	gorm.Model
	Id  string `gorm:"primarykey"`
	Url string
}

func (shortener *Shortener) BeforeCreate(tx *gorm.DB) (err error) {
	shortener.Id, err = GenerateRandomString(6)
	if err != nil {
		return err
	}
	shortener.CreatedAt = time.Now()
	shortener.UpdatedAt = time.Now()
	return nil
}

func (shortener *Shortener) BeforeUpdate(tx *gorm.DB) (err error) {
	shortener.UpdatedAt = time.Now()
	return nil
}

func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result += string(charset[num.Int64()])
	}
	return result, nil
}
