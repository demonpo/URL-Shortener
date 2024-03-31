package mocks

import (
	"github.com/stretchr/testify/mock"
	"goHexBoilerplate/src/modules/user/domain/contracts/entities"
	"goHexBoilerplate/src/modules/user/domain/contracts/repositories"
)

// UserRepositoryMock is a mock implementation of UserRepository interface
type UserRepositoryMock struct {
	mock.Mock
}

// GetById provides a mock function for GetById method
func (m *UserRepositoryMock) GetById(id int) (*entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

// GetByProperties provides a mock function for GetByProperties method
func (m *UserRepositoryMock) GetByProperties(params repositories.GetByPropertiesParams) ([]entities.User, error) {
	args := m.Called(params)
	return args.Get(0).([]entities.User), args.Error(1)
}

// Create provides a mock function for Create method
func (m *UserRepositoryMock) Create(params repositories.Create) (*entities.User, error) {
	args := m.Called(params)
	return args.Get(0).(*entities.User), args.Error(1)
}
