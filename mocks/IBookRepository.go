// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IBookRepository is an autogenerated mock type for the IBookRepository type
type IBookRepository struct {
	mock.Mock
}

type mockConstructorTestingTNewIBookRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIBookRepository creates a new instance of IBookRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIBookRepository(t mockConstructorTestingTNewIBookRepository) *IBookRepository {
	mock := &IBookRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}