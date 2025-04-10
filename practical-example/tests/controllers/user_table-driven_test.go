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

func TestUserController(t *testing.T) {
	suite.Run(t, new(TableDrivenTestUserControllerSuite))
}

type TableDrivenTestUserControllerSuite struct {
	suite.Suite
	userController       *controllers.UserController
	mockUserModel        *mocks.MockUserModel
	mockBadUsernameModel *mocks.MockBadUsernameModel
}

type UserControllerTestcase struct {
	name             string
	body             types.RegisterUserBody
	mockFunc         func()
	expectedUser     *models.User
	expectedError    error
	expectedMockFunc func()
}

func (s *TableDrivenTestUserControllerSuite) SetupTest() {
	s.mockUserModel = new(mocks.MockUserModel)
	s.mockBadUsernameModel = new(mocks.MockBadUsernameModel)
	s.userController = controllers.NewUserController(s.mockUserModel, s.mockBadUsernameModel)
}

func (s *TableDrivenTestUserControllerSuite) TestRegisterUser() {
	testcases := []UserControllerTestcase{
		{
			name:             "UsernameRequired",
			body:             types.RegisterUserBody{Username: ""},
			mockFunc:         func() {},
			expectedUser:     nil,
			expectedError:    errors.New("username is required"),
			expectedMockFunc: func() {},
		},
		{
			name:             "UsernameContainsSpaces",
			body:             types.RegisterUserBody{Username: " user "},
			mockFunc:         func() {},
			expectedUser:     nil,
			expectedError:    errors.New("username cannot contain spaces"),
			expectedMockFunc: func() {},
		},
		{
			name: "UsernameNotAllowed",
			body: types.RegisterUserBody{Username: "admin"},
			mockFunc: func() {
				s.mockBadUsernameModel.On("GetBadUsernameByUsername", "admin").Return(&models.BadUsername{Username: "admin"}, true)
			},
			expectedUser:  nil,
			expectedError: errors.New("username is not allowed"),
			expectedMockFunc: func() {
				s.mockBadUsernameModel.AssertExpectations(s.T())
			},
		},
		{
			name: "UsernameAlreadyExists",
			body: types.RegisterUserBody{Username: "existinguser"},
			mockFunc: func() {
				s.mockBadUsernameModel.On("GetBadUsernameByUsername", "existinguser").Return(nil, false)
				s.mockUserModel.On("GetUserByUsername", "existinguser").Return(&models.User{Username: "existinguser"}, true)
			},
			expectedUser:  nil,
			expectedError: errors.New("user already exists"),
			expectedMockFunc: func() {
				s.mockBadUsernameModel.AssertExpectations(s.T())
				s.mockUserModel.AssertExpectations(s.T())
			},
		},
		{
			name: "FailedToCreateUser",
			body: types.RegisterUserBody{Username: "newuser"},
			mockFunc: func() {
				s.mockBadUsernameModel.On("GetBadUsernameByUsername", "newuser").Return(nil, false)
				s.mockUserModel.On("GetUserByUsername", "newuser").Return(nil, false)
				s.mockUserModel.On("CreateUser", "newuser").Return(nil, errors.New("failed to create user"))
			},
			expectedUser:  nil,
			expectedError: errors.New("failed to create user"),
			expectedMockFunc: func() {
				s.mockBadUsernameModel.AssertExpectations(s.T())
				s.mockUserModel.AssertExpectations(s.T())
			},
		},
		{
			name: "UsernameSuccess",
			body: types.RegisterUserBody{Username: "newuser"},
			mockFunc: func() {
				s.mockBadUsernameModel.On("GetBadUsernameByUsername", "newuser").Return(nil, false)
				s.mockUserModel.On("GetUserByUsername", "newuser").Return(nil, false)
				s.mockUserModel.On("CreateUser", "newuser").Return(&models.User{Username: "newuser"}, nil)
			},
			expectedUser:  &models.User{Username: "newuser"},
			expectedError: nil,
			expectedMockFunc: func() {
				s.mockBadUsernameModel.AssertExpectations(s.T())
				s.mockUserModel.AssertExpectations(s.T())
			},
		},
	}

	for _, testcase := range testcases {
		s.Run(testcase.name, func() {
			s.SetupTest()
			testcase.mockFunc()

			user, err := s.userController.RegisterUser(testcase.body)

			assert.Equal(s.T(), user, testcase.expectedUser)
			assert.Equal(s.T(), err, testcase.expectedError)
		})
	}
}
