// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Importer is an autogenerated mock type for the Importer type
type Importer struct {
	mock.Mock
}

// Import provides a mock function with given fields: ctx, sourcePath
func (_m *Importer) Import(ctx context.Context, sourcePath string) error {
	ret := _m.Called(ctx, sourcePath)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, sourcePath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewImporter interface {
	mock.TestingT
	Cleanup(func())
}

// NewImporter creates a new instance of Importer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImporter(t mockConstructorTestingTNewImporter) *Importer {
	mock := &Importer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
