package main

import (
	"errors"
	"testing"
)

type MockUserModel struct {
	GetUserByUsernameImpl func(username string) (*User, bool)
	CreateUserImpl        func(username string) (*User, error)
}

func (m *MockUserModel) GetUserByUsername(username string) (*User, bool) {
	return m.GetUserByUsernameImpl(username)
}

func (m *MockUserModel) CreateUser(username string) (*User, error) {
	return m.CreateUserImpl(username)
}

func TestRegisterUser_FailureEmptyUsername(t *testing.T) {
	username := ""

	mockUserModel := &MockUserModel{}
	userController := NewUserController(mockUserModel)

	body := RegisterUserBody{Username: username}

	user, err := userController.RegisterUser(body)
	if err == nil {
		t.Errorf("Expected error for empty username, got nil")
	}
	if user != nil {
		t.Errorf("Expected nil user for empty username, got %v", user)
	}
}

func TestRegisterUser_FailureUserExists(t *testing.T) {
	username := "existing_user"

	mockUserModel := &MockUserModel{
		GetUserByUsernameImpl: func(username string) (*User, bool) {
			return &User{Username: username}, true
		},
	}
	userController := NewUserController(mockUserModel)

	body := RegisterUserBody{Username: username}

	user, err := userController.RegisterUser(body)
	if err == nil {
		t.Errorf("Expected error for existing user, got nil")
	}
	if user != nil {
		t.Errorf("Expected nil user for existing user, got %v", user)
	}
}

func TestRegisterUser_FailOnCreateUser(t *testing.T) {
	username := "new_user"

	mockUserModel := &MockUserModel{
		GetUserByUsernameImpl: func(username string) (*User, bool) {
			return &User{}, true
		},
		CreateUserImpl: func(username string) (*User, error) {
			return nil, errors.New("failed to create user")
		},
	}
	userController := NewUserController(mockUserModel)

	body := RegisterUserBody{Username: username}

	user, err := userController.RegisterUser(body)
	if err == nil {
		t.Errorf("Expected error for failed to create user, got nil")
	}
	if user != nil {
		t.Errorf("Expected nil user for failed to create user, got %v", user)
	}
}

func TestRegisterUser_Success(t *testing.T) {
	username := "new_user"

	mockUserModel := &MockUserModel{
		GetUserByUsernameImpl: func(username string) (*User, bool) {
			return nil, false
		},
		CreateUserImpl: func(username string) (*User, error) {
			return &User{ID: "1", Username: username}, nil
		},
	}
	userController := NewUserController(mockUserModel)

	body := RegisterUserBody{Username: username}

	user, err := userController.RegisterUser(body)
	if err != nil {
		t.Errorf("Expected no error for successful user creation, got %v", err)
	}
	if user == nil {
		t.Errorf("Expected user for successful user creation, got nil")
	}
	if user.Username != username {
		t.Errorf("Expected user with username '%s', got %v", username, user.Username)
	}
}
