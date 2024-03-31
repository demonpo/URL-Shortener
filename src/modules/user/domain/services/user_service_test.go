package services

import (
	"goHexBoilerplate/src/modules/user/domain/contracts/entities"
	"goHexBoilerplate/src/modules/user/test/mocks"
	"testing"
)

func TestUserService_GetUserById(t *testing.T) {
	// Create an instance of the mock
	userRepositoryMock := &mocks.UserRepositoryMock{}

	// Set up expected behavior for GetById method
	expectedUser := &entities.User{Id: 1, Name: "John Doe", Email: "some@email.com"}
	userRepositoryMock.On("GetById", 1).Return(expectedUser, nil)

	// Pass the mock to your UserService or whatever service you are testing
	userService := NewUserService(userRepositoryMock)

	// Call the method you want to test
	user, err := userService.GetById(1)

	// Assert the results
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if user.Id != expectedUser.Id || user.Name != expectedUser.Name || user.Email != expectedUser.Email {
		t.Errorf("Expected user %+v, got %+v", expectedUser, user)
	}

	// Assert that the GetById method was called with the expected parameters
	userRepositoryMock.AssertCalled(t, "GetById", 1)
}
