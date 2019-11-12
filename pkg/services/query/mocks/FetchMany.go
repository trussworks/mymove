// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	services "github.com/transcom/mymove/pkg/services"
)

// FetchMany is an autogenerated mock type for the FetchMany type
type FetchMany struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *FetchMany) Execute() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithAssociations provides a mock function with given fields: associations
func (_m *FetchMany) WithAssociations(associations services.QueryAssociations) *query.fetchMany {
	ret := _m.Called(associations)

	var r0 *query.fetchMany
	if rf, ok := ret.Get(0).(func(services.QueryAssociations) *query.fetchMany); ok {
		r0 = rf(associations)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*query.fetchMany)
		}
	}

	return r0
}

// WithFilters provides a mock function with given fields: filters
func (_m *FetchMany) WithFilters(filters []services.QueryFilter) *query.fetchMany {
	ret := _m.Called(filters)

	var r0 *query.fetchMany
	if rf, ok := ret.Get(0).(func([]services.QueryFilter) *query.fetchMany); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*query.fetchMany)
		}
	}

	return r0
}

// WithModel provides a mock function with given fields: model
func (_m *FetchMany) WithModel(model interface{}) *query.fetchMany {
	ret := _m.Called(model)

	var r0 *query.fetchMany
	if rf, ok := ret.Get(0).(func(interface{}) *query.fetchMany); ok {
		r0 = rf(model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*query.fetchMany)
		}
	}

	return r0
}

// WithPagination provides a mock function with given fields: pagination
func (_m *FetchMany) WithPagination(pagination services.Pagination) *query.fetchMany {
	ret := _m.Called(pagination)

	var r0 *query.fetchMany
	if rf, ok := ret.Get(0).(func(services.Pagination) *query.fetchMany); ok {
		r0 = rf(pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*query.fetchMany)
		}
	}

	return r0
}
