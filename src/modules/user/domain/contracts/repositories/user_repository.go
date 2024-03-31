package repositories

import "goHexBoilerplate/src/modules/user/domain/contracts/entities"

type GetByPropertiesParams struct {
	id    int
	name  string
	email string
}

type Create struct {
	Name  string
	Email string
}

type UserRepository interface {
	GetById(id int) (*entities.User, error)
	GetByProperties(params GetByPropertiesParams) ([]entities.User, error)
	Create(params Create) (*entities.User, error)
}
