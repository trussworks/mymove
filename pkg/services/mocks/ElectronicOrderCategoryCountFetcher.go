// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	services "github.com/transcom/mymove/pkg/services"
)

// ElectronicOrderCategoryCountFetcher is an autogenerated mock type for the ElectronicOrderCategoryCountFetcher type
type ElectronicOrderCategoryCountFetcher struct {
	mock.Mock
}

// FetchElectronicOrderCategoricalCounts provides a mock function with given fields: filters, andFilters
func (_m *ElectronicOrderCategoryCountFetcher) FetchElectronicOrderCategoricalCounts(filters []services.QueryFilter, andFilters *[]services.QueryFilter) (map[interface{}]int, error) {
	ret := _m.Called(filters, andFilters)

	var r0 map[interface{}]int
	if rf, ok := ret.Get(0).(func([]services.QueryFilter, *[]services.QueryFilter) map[interface{}]int); ok {
		r0 = rf(filters, andFilters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[interface{}]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]services.QueryFilter, *[]services.QueryFilter) error); ok {
		r1 = rf(filters, andFilters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
