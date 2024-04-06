package repositories

import (
	"goHexBoilerplate/src/db"
	"goHexBoilerplate/src/modules/analytics/domain/contracts/entities"
	"goHexBoilerplate/src/modules/analytics/domain/contracts/repositories"
	entitiesInfra "goHexBoilerplate/src/modules/analytics/infra/entities"
)

type PostgresClickRepository struct {
	db *db.DB
}

func NewPostgresClickRepository(db *db.DB) *PostgresClickRepository {
	return &PostgresClickRepository{
		db: db,
	}
}

func (shortenerRepository *PostgresClickRepository) Create(params repositories.Create) (*entities.Click, error) {
	newShortener := entitiesInfra.Click{
		ShortenerId: params.ShortenerId,
		UserIp:      params.UserIp,
		ReferrerUrl: params.ReferrerUrl,
		UserAgent:   params.UserAgent,
	}
	if err := shortenerRepository.db.DB.Model(&entitiesInfra.Click{}).Create(&newShortener).Error; err != nil {
		return nil, err
	}
	return &entities.Click{
		ShortenerId: newShortener.ShortenerId,
		UserIp:      newShortener.UserIp,
		ReferrerUrl: newShortener.ReferrerUrl,
		UserAgent:   newShortener.UserAgent,
	}, nil
}
