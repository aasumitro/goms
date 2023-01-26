// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/aasumitro/goms/internal/book/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IBookService is an autogenerated mock type for the IBookService type
type IBookService struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx, args
func (_m *IBookService) All(ctx context.Context, args ...string) ([]*entity.Book, error) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*entity.Book
	if rf, ok := ret.Get(0).(func(context.Context, ...string) []*entity.Book); ok {
		r0 = rf(ctx, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...string) error); ok {
		r1 = rf(ctx, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Erase provides a mock function with given fields: ctx, arg
func (_m *IBookService) Erase(ctx context.Context, arg *entity.Book) error {
	ret := _m.Called(ctx, arg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Book) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, arg
func (_m *IBookService) Find(ctx context.Context, arg *entity.Book) (*entity.Book, error) {
	ret := _m.Called(ctx, arg)

	var r0 *entity.Book
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Book) *entity.Book); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Book) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, arg
func (_m *IBookService) Patch(ctx context.Context, arg *entity.Book) error {
	ret := _m.Called(ctx, arg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Book) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Record provides a mock function with given fields: ctx, args
func (_m *IBookService) Record(ctx context.Context, args ...*entity.Book) error {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*entity.Book) error); ok {
		r0 = rf(ctx, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIBookService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIBookService creates a new instance of IBookService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIBookService(t mockConstructorTestingTNewIBookService) *IBookService {
	mock := &IBookService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
