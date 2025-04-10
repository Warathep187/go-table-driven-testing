package controllers

import (
	"errors"
	"strings"

	"practical-example/models"
	"practical-example/types"
)

type UserController struct {
	userModel        models.UserModelInterface
	badUsernameModel models.BadUsernameModelInterface
}

func NewUserController(userModel models.UserModelInterface, badUsernameModel models.BadUsernameModelInterface) *UserController {
	return &UserController{
		userModel:        userModel,
		badUsernameModel: badUsernameModel,
	}
}

func (c *UserController) RegisterUser(body types.RegisterUserBody) (*models.User, error) {
	username := body.Username
	if username == "" {
		return nil, errors.New("username is required")
	}

	if strings.Contains(username, " ") {
		return nil, errors.New("username cannot contain spaces")
	}

	_, exists := c.badUsernameModel.GetBadUsernameByUsername(username)
	if exists {
		return nil, errors.New("username is not allowed")
	}

	_, exists = c.userModel.GetUserByUsername(username)
	if exists {
		return nil, errors.New("user already exists")
	}

	user, err := c.userModel.CreateUser(username)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	return user, nil
}
