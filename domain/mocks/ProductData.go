// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "altaproject2/domain"

	mock "github.com/stretchr/testify/mock"
)

// ProductData is an autogenerated mock type for the ProductData type
type ProductData struct {
	mock.Mock
}

// DeleteItemData provides a mock function with given fields:
func (_m *ProductData) DeleteItemData() {
	_m.Called()
}

// GetAllItemData provides a mock function with given fields:
func (_m *ProductData) GetAllItemData() []domain.Product {
	ret := _m.Called()

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func() []domain.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	return r0
}

// GetItemData provides a mock function with given fields: productID
func (_m *ProductData) GetItemData(productID int) (domain.Product, error) {
	ret := _m.Called(productID)

	var r0 domain.Product
	if rf, ok := ret.Get(0).(func(int) domain.Product); ok {
		r0 = rf(productID)
	} else {
		r0 = ret.Get(0).(domain.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostItemData provides a mock function with given fields:
func (_m *ProductData) PostItemData() {
	_m.Called()
}

// UpdateItemData provides a mock function with given fields:
func (_m *ProductData) UpdateItemData() {
	_m.Called()
}

type mockConstructorTestingTNewProductData interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductData creates a new instance of ProductData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductData(t mockConstructorTestingTNewProductData) *ProductData {
	mock := &ProductData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}