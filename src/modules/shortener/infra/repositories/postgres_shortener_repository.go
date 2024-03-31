package repositories

import (
	"errors"
	"fmt"
	"goHexBoilerplate/src/db"
	"goHexBoilerplate/src/modules/shortener/domain/contracts/entities"
	"goHexBoilerplate/src/modules/shortener/domain/contracts/repositories"
	entitiesInfra "goHexBoilerplate/src/modules/shortener/infra/entities"
	"gorm.io/gorm"
)

type PostgresShortenerRepository struct {
	db *db.DB
}

func NewPostgresShortenerRepository(db *db.DB) *PostgresShortenerRepository {
	return &PostgresShortenerRepository{
		db: db,
	}
}

func (shortenerRepository *PostgresShortenerRepository) Create(params repositories.Create) (*entities.Shortener, error) {
	newShortener := entitiesInfra.Shortener{
		Url: params.Url,
	}
	if err := shortenerRepository.db.DB.Model(&entitiesInfra.Shortener{}).Create(&newShortener).Error; err != nil {
		return nil, err
	}
	return &entities.Shortener{
		Id:        newShortener.Id,
		Url:       newShortener.Url,
		UpdatedAt: newShortener.UpdatedAt,
		CreatedAt: newShortener.CreatedAt,
	}, nil
}

func (shortenerRepository *PostgresShortenerRepository) GetById(id string) (*entities.Shortener, error) {
	fmt.Print("PostgresShortenerRepository GetById\n")
	foundShortener := entitiesInfra.Shortener{Id: id}
	if err := shortenerRepository.db.DB.Model(&entitiesInfra.Shortener{}).First(&foundShortener).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("shortener not found")
		}
		return nil, err
	}
	return &entities.Shortener{
		Id:        foundShortener.Id,
		Url:       foundShortener.Url,
		UpdatedAt: foundShortener.UpdatedAt,
		CreatedAt: foundShortener.CreatedAt,
	}, nil
}
