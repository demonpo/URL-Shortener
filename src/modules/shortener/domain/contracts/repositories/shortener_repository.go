package repositories

import "goHexBoilerplate/src/modules/shortener/domain/contracts/entities"

type Create struct {
	Url string
}

type ShortenerRepository interface {
	GetById(id string) (*entities.Shortener, error)
	Create(params Create) (*entities.Shortener, error)
}
