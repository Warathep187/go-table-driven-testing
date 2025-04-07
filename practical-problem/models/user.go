package models

import (
	"math/rand"
	"strconv"
)

type User struct {
	ID       string
	Username string
}

//go:generate mockery --name UserModelInterface --structname MockUserModel --filename=mock_user_model.go --output ../tests/mocks/models
type UserModelInterface interface {
	GetUserByUsername(username string) (*User, bool)
	CreateUser(username string) (*User, error)
}

type UserModel struct {
	users map[string]User
}

func NewUserModel() *UserModel {
	return &UserModel{
		users: make(map[string]User),
	}
}

func (u *UserModel) GetUserByUsername(username string) (*User, bool) {
	user, exists := u.users[username]
	return &user, exists
}

func (u *UserModel) CreateUser(username string) (*User, error) {
	user := User{
		ID:       strconv.Itoa(rand.Intn(1_000_000)),
		Username: username,
	}
	u.users[user.Username] = user
	return &user, nil
}
