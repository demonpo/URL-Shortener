package services

import (
	"goHexBoilerplate/src/modules/shortener/domain/contracts/entities"
	"goHexBoilerplate/src/modules/shortener/domain/contracts/repositories"
)

type ShortenerService struct {
	shortenerRepository repositories.ShortenerRepository
}

type CreateInput struct {
	Url string
}

type RedirectInput struct {
	shortenerId string
}

func NewShortenerService(shortenerRepository repositories.ShortenerRepository) *ShortenerService {
	return &ShortenerService{shortenerRepository: shortenerRepository}
}

func (shortenerService *ShortenerService) GetById(id string) (*entities.Shortener, error) {
	shortenerRepository := shortenerService.shortenerRepository
	return shortenerRepository.GetById(id)
}

func (shortenerService *ShortenerService) Create(input CreateInput) (*entities.Shortener, error) {
	shortenerRepository := shortenerService.shortenerRepository
	return shortenerRepository.Create(repositories.Create{Url: input.Url})
}
