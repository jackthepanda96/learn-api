// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	users "19api/features/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Insert provides a mock function with given fields: newUser
func (_m *Repository) Insert(newUser users.User) (users.User, error) {
	ret := _m.Called(newUser)

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(users.User) (users.User, error)); ok {
		return rf(newUser)
	}
	if rf, ok := ret.Get(0).(func(users.User) users.User); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(users.User) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: hp
func (_m *Repository) Login(hp string) (users.User, error) {
	ret := _m.Called(hp)

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (users.User, error)); ok {
		return rf(hp)
	}
	if rf, ok := ret.Get(0).(func(string) users.User); ok {
		r0 = rf(hp)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}