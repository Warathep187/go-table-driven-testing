package tests

import (
	"errors"
	"testing"

	"practical-example/controllers"
	"practical-example/models"
	mocks "practical-example/tests/mocks/models"
	"practical-example/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestSimpleUserController(t *testing.T) {
	suite.Run(t, new(SimpleTestUserControllerSuite))
}

type SimpleTestUserControllerSuite struct {
	suite.Suite
	userController       *controllers.UserController
	mockUserModel        *mocks.MockUserModel
	mockBadUsernameModel *mocks.MockBadUsernameModel
}

func (s *SimpleTestUserControllerSuite) SetupTest() {
	s.mockUserModel = new(mocks.MockUserModel)
	s.mockBadUsernameModel = new(mocks.MockBadUsernameModel)
	s.userController = controllers.NewUserController(s.mockUserModel, s.mockBadUsernameModel)
}

func (s *SimpleTestUserControllerSuite) TestRegisterUserFailed_UsernameRequired() {
	body := types.RegisterUserBody{Username: ""}

	user, err := s.userController.RegisterUser(body)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), err, errors.New("username is required"))
}

func (s *SimpleTestUserControllerSuite) TestRegisterUserFailed_UsernameContainsSpaces() {
	body := types.RegisterUserBody{Username: "user name"}

	user, err := s.userController.RegisterUser(body)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), err, errors.New("username cannot contain spaces"))
}

func (s *SimpleTestUserControllerSuite) TestRegisterUserFailed_UsernameNotAllowed() {
	body := types.RegisterUserBody{Username: "admin"}

	s.mockBadUsernameModel.On("GetBadUsernameByUsername", body.Username).Return(&models.BadUsername{Username: "admin"}, true)

	user, err := s.userController.RegisterUser(body)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), err, errors.New("username is not allowed"))

	s.mockBadUsernameModel.AssertExpectations(s.T())
}

func (s *SimpleTestUserControllerSuite) TestRegisterUserFailed_UsernameAlreadyExists() {
	body := types.RegisterUserBody{Username: "existinguser"}

	s.mockBadUsernameModel.On("GetBadUsernameByUsername", body.Username).Return(nil, false)
	s.mockUserModel.On("GetUserByUsername", body.Username).Return(&models.User{Username: "existinguser"}, true)

	user, err := s.userController.RegisterUser(body)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), err, errors.New("user already exists"))

	s.mockBadUsernameModel.AssertExpectations(s.T())
	s.mockUserModel.AssertExpectations(s.T())
}

func (s *SimpleTestUserControllerSuite) TestRegisterUserFailed_FailedToCreateUser() {
	body := types.RegisterUserBody{Username: "newuser"}

	s.mockBadUsernameModel.On("GetBadUsernameByUsername", body.Username).Return(nil, false)
	s.mockUserModel.On("GetUserByUsername", body.Username).Return(nil, false)
	s.mockUserModel.On("CreateUser", body.Username).Return(nil, errors.New("failed to create user"))

	user, err := s.userController.RegisterUser(body)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), err, errors.New("failed to create user"))

	s.mockBadUsernameModel.AssertExpectations(s.T())
	s.mockUserModel.AssertExpectations(s.T())
}

func (s *SimpleTestUserControllerSuite) TestRegisterUserSuccess() {
	body := types.RegisterUserBody{Username: "newuser"}

	s.mockBadUsernameModel.On("GetBadUsernameByUsername", body.Username).Return(nil, false)
	s.mockUserModel.On("GetUserByUsername", body.Username).Return(nil, false)
	s.mockUserModel.On("CreateUser", body.Username).Return(&models.User{Username: "newuser"}, nil)

	user, err := s.userController.RegisterUser(body)

	assert.Equal(s.T(), user, &models.User{Username: "newuser"})
	assert.Equal(s.T(), err, nil)

	s.mockBadUsernameModel.AssertExpectations(s.T())
	s.mockUserModel.AssertExpectations(s.T())
}
