package main

import "errors"

type UserController struct {
	userModel UserModelInterface
}

func NewUserController(userModel UserModelInterface) *UserController {
	return &UserController{
		userModel: userModel,
	}
}

func (c *UserController) RegisterUser(body RegisterUserBody) (*User, error) {
	username := body.Username
	if username == "" {
		return nil, errors.New("username is required")
	}

	_, exists := c.userModel.GetUserByUsername(username)
	if exists {
		return nil, errors.New("user already exists")
	}

	return c.userModel.CreateUser(username)
}
