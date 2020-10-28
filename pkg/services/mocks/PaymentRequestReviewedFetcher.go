// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"
)

// PaymentRequestReviewedFetcher is an autogenerated mock type for the PaymentRequestReviewedFetcher type
type PaymentRequestReviewedFetcher struct {
	mock.Mock
}

// FetchReviewedPaymentRequest provides a mock function with given fields:
func (_m *PaymentRequestReviewedFetcher) FetchReviewedPaymentRequest() (models.PaymentRequests, error) {
	ret := _m.Called()

	var r0 models.PaymentRequests
	if rf, ok := ret.Get(0).(func() models.PaymentRequests); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.PaymentRequests)
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
