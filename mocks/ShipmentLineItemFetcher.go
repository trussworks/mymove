// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import auth "github.com/transcom/mymove/pkg/auth"
import mock "github.com/stretchr/testify/mock"
import models "github.com/transcom/mymove/pkg/models"

import uuid "github.com/gofrs/uuid"

// ShipmentLineItemFetcher is an autogenerated mock type for the ShipmentLineItemFetcher type
type ShipmentLineItemFetcher struct {
	mock.Mock
}

// GetShipmentLineItemsByShipmentID provides a mock function with given fields: shipmentID, session
func (_m *ShipmentLineItemFetcher) GetShipmentLineItemsByShipmentID(shipmentID uuid.UUID, session *auth.Session) ([]models.ShipmentLineItem, error) {
	ret := _m.Called(shipmentID, session)

	var r0 []models.ShipmentLineItem
	if rf, ok := ret.Get(0).(func(uuid.UUID, *auth.Session) []models.ShipmentLineItem); ok {
		r0 = rf(shipmentID, session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ShipmentLineItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, *auth.Session) error); ok {
		r1 = rf(shipmentID, session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
