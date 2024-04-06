package services

import (
	"goHexBoilerplate/src/modules/analytics/domain/contracts/entities"
	"goHexBoilerplate/src/modules/analytics/domain/contracts/repositories"
)

type ClickService struct {
	clickService repositories.ClickRepository
}

type CreateInput struct {
	ShortenerId string
	UserIp      *string
	ReferrerUrl *string
	UserAgent   *string
}

func NewClickService(clickService repositories.ClickRepository) *ClickService {
	return &ClickService{clickService: clickService}
}

func (clickService *ClickService) Create(input CreateInput) (*entities.Click, error) {
	shortenerRepository := clickService.clickService
	return shortenerRepository.Create(repositories.Create{
		UserAgent:   input.UserAgent,
		ReferrerUrl: input.ReferrerUrl,
		UserIp:      input.UserIp,
		ShortenerId: input.ShortenerId,
	})
}
