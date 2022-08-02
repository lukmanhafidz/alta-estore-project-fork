// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "altaproject2/domain"

	mock "github.com/stretchr/testify/mock"
)

// ProductUseCase is an autogenerated mock type for the ProductUseCase type
type ProductUseCase struct {
	mock.Mock
}

// DeleteItemAdmin provides a mock function with given fields:
func (_m *ProductUseCase) DeleteItemAdmin() {
	_m.Called()
}

// GetAllItems provides a mock function with given fields:
func (_m *ProductUseCase) GetAllItems() ([]domain.Product, int) {
	ret := _m.Called()

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func() []domain.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// GetItemUser provides a mock function with given fields: id
func (_m *ProductUseCase) GetItemUser(id int) (domain.Product, error) {
	ret := _m.Called(id)

	var r0 domain.Product
	if rf, ok := ret.Get(0).(func(int) domain.Product); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostItemAdmin provides a mock function with given fields:
func (_m *ProductUseCase) PostItemAdmin() {
	_m.Called()
}

// UpdateItemAdmin provides a mock function with given fields:
func (_m *ProductUseCase) UpdateItemAdmin() {
	_m.Called()
}

type mockConstructorTestingTNewProductUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductUseCase creates a new instance of ProductUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductUseCase(t mockConstructorTestingTNewProductUseCase) *ProductUseCase {
	mock := &ProductUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
