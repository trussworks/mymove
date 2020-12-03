// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"
)

// OfficeUserFetcherPop is an autogenerated mock type for the OfficeUserFetcherPop type
type OfficeUserFetcherPop struct {
	mock.Mock
}

// FetchOfficeUserByID provides a mock function with given fields: id
func (_m *OfficeUserFetcherPop) FetchOfficeUserByID(id uuid.UUID) (models.OfficeUser, error) {
	ret := _m.Called(id)

	var r0 models.OfficeUser
	if rf, ok := ret.Get(0).(func(uuid.UUID) models.OfficeUser); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.OfficeUser)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
