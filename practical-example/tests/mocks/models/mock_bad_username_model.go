// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "practical-example/models"

	mock "github.com/stretchr/testify/mock"
)

// MockBadUsernameModel is an autogenerated mock type for the BadUsernameModelInterface type
type MockBadUsernameModel struct {
	mock.Mock
}

// GetBadUsernameByUsername provides a mock function with given fields: username
func (_m *MockBadUsernameModel) GetBadUsernameByUsername(username string) (*models.BadUsername, bool) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for GetBadUsernameByUsername")
	}

	var r0 *models.BadUsername
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (*models.BadUsername, bool)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *models.BadUsername); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.BadUsername)
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// NewMockBadUsernameModel creates a new instance of MockBadUsernameModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBadUsernameModel(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBadUsernameModel {
	mock := &MockBadUsernameModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
