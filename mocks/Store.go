// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "seaports-service-assignment/internal/domain/model"

	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Create provides a mock function with given fields: port
func (_m *Store) Create(port model.Seaport) error {
	ret := _m.Called(port)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Seaport) error); ok {
		r0 = rf(port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: port
func (_m *Store) Update(port model.Seaport) error {
	ret := _m.Called(port)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Seaport) error); ok {
		r0 = rf(port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStore(t mockConstructorTestingTNewStore) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
