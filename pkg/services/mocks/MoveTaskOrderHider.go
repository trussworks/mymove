// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	services "github.com/transcom/mymove/pkg/services"
)

// MoveTaskOrderHider is an autogenerated mock type for the MoveTaskOrderHider type
type MoveTaskOrderHider struct {
	mock.Mock
}

// Hide provides a mock function with given fields:
func (_m *MoveTaskOrderHider) Hide() (services.HiddenMoves, error) {
	ret := _m.Called()

	var r0 services.HiddenMoves
	if rf, ok := ret.Get(0).(func() services.HiddenMoves); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(services.HiddenMoves)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
