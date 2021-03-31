// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	services "github.com/transcom/mymove/pkg/services"

	time "time"

	unit "github.com/transcom/mymove/pkg/unit"
)

// DomesticLinehaulPricer is an autogenerated mock type for the DomesticLinehaulPricer type
type DomesticLinehaulPricer struct {
	mock.Mock
}

// Price provides a mock function with given fields: contractCode, requestedPickupDate, distance, weight, serviceArea
func (_m *DomesticLinehaulPricer) Price(contractCode string, requestedPickupDate time.Time, distance unit.Miles, weight unit.Pound, serviceArea string) (unit.Cents, services.PricingDisplayParams, error) {
	ret := _m.Called(contractCode, requestedPickupDate, distance, weight, serviceArea)

	var r0 unit.Cents
	if rf, ok := ret.Get(0).(func(string, time.Time, unit.Miles, unit.Pound, string) unit.Cents); ok {
		r0 = rf(contractCode, requestedPickupDate, distance, weight, serviceArea)
	} else {
		r0 = ret.Get(0).(unit.Cents)
	}

	var r1 services.PricingDisplayParams
	if rf, ok := ret.Get(1).(func(string, time.Time, unit.Miles, unit.Pound, string) services.PricingDisplayParams); ok {
		r1 = rf(contractCode, requestedPickupDate, distance, weight, serviceArea)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(services.PricingDisplayParams)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, time.Time, unit.Miles, unit.Pound, string) error); ok {
		r2 = rf(contractCode, requestedPickupDate, distance, weight, serviceArea)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PriceUsingParams provides a mock function with given fields: params
func (_m *DomesticLinehaulPricer) PriceUsingParams(params models.PaymentServiceItemParams) (unit.Cents, services.PricingDisplayParams, error) {
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
