package repositories

import (
	"errors"
	"goHexBoilerplate/src/db"
	"goHexBoilerplate/src/modules/user/domain/contracts/entities"
	"goHexBoilerplate/src/modules/user/domain/contracts/repositories"
	entitiesInfra "goHexBoilerplate/src/modules/user/infra/entities"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *db.DB
}

func NewPostgresUserRepository(db *db.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (userRepository *PostgresUserRepository) Create(params repositories.Create) (*entities.User, error) {
	newUser := entitiesInfra.User{
		Name:  params.Name,
		Email: params.Email,
	}
	if err := userRepository.db.DB.Model(&entitiesInfra.User{}).Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &entities.User{
		Id:        int(newUser.ID),
		Name:      newUser.Name,
		Email:     newUser.Email,
		UpdatedAt: newUser.UpdatedAt,
		CreatedAt: newUser.CreatedAt,
	}, nil
}

func (userRepository *PostgresUserRepository) GetById(id int) (*entities.User, error) {
	foundUser := entitiesInfra.User{}
	if err := userRepository.db.DB.Model(&entitiesInfra.User{}).First(&foundUser, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &entities.User{
		Id:        int(foundUser.ID),
		Name:      foundUser.Name,
		Email:     foundUser.Email,
		UpdatedAt: foundUser.UpdatedAt,
		CreatedAt: foundUser.CreatedAt,
	}, nil
}

func (userRepository *PostgresUserRepository) GetByProperties(params repositories.GetByPropertiesParams) ([]entities.User, error) {
	return []entities.User{entities.User{Id: 1, Name: "Daniel", Email: "soul.daniel@hotmail.com"}, entities.User{Id: 2, Name: "Xavier", Email: "xavicoGarcia12@gmail.com"}}, nil
}
