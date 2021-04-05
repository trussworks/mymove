// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	services "github.com/transcom/mymove/pkg/services"

	time "time"

	unit "github.com/transcom/mymove/pkg/unit"
)

// DomesticOriginAdditionalDaysSITPricer is an autogenerated mock type for the DomesticOriginAdditionalDaysSITPricer type
type DomesticOriginAdditionalDaysSITPricer struct {
	mock.Mock
}

// Price provides a mock function with given fields: contractCode, requestedPickupDate, weight, serviceArea, numberOfDaysInSIT
func (_m *DomesticOriginAdditionalDaysSITPricer) Price(contractCode string, requestedPickupDate time.Time, weight unit.Pound, serviceArea string, numberOfDaysInSIT int) (unit.Cents, services.PricingDisplayParams, error) {
	ret := _m.Called(contractCode, requestedPickupDate, weight, serviceArea, numberOfDaysInSIT)

	var r0 unit.Cents
	if rf, ok := ret.Get(0).(func(string, time.Time, unit.Pound, string, int) unit.Cents); ok {
		r0 = rf(contractCode, requestedPickupDate, weight, serviceArea, numberOfDaysInSIT)
	} else {
		r0 = ret.Get(0).(unit.Cents)
	}

	var r1 services.PricingDisplayParams
	if rf, ok := ret.Get(1).(func(string, time.Time, unit.Pound, string, int) services.PricingDisplayParams); ok {
		r1 = rf(contractCode, requestedPickupDate, weight, serviceArea, numberOfDaysInSIT)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(services.PricingDisplayParams)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, time.Time, unit.Pound, string, int) error); ok {
		r2 = rf(contractCode, requestedPickupDate, weight, serviceArea, numberOfDaysInSIT)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PriceUsingParams provides a mock function with given fields: params
func (_m *DomesticOriginAdditionalDaysSITPricer) PriceUsingParams(params models.PaymentServiceItemParams) (unit.Cents, services.PricingDisplayParams, error) {
	ret := _m.Called(params)

	var r0 unit.Cents
	if rf, ok := ret.Get(0).(func(models.PaymentServiceItemParams) unit.Cents); ok {
		r0 = rf(params)
	} else {
		r0 = ret.Get(0).(unit.Cents)
	}

	var r1 services.PricingDisplayParams
	if rf, ok := ret.Get(1).(func(models.PaymentServiceItemParams) services.PricingDisplayParams); ok {
		r1 = rf(params)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(services.PricingDisplayParams)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(models.PaymentServiceItemParams) error); ok {
		r2 = rf(params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
