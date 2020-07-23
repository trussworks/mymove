// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MoveOrder move order
// swagger:model MoveOrder
type MoveOrder struct {

	// customer
	Customer *Customer `json:"customer,omitempty"`

	// ID of the Customer this MoveOrder belongs to.
	//
	// If creating a MoveTaskOrder. either an existing customerID should be provided or the nested customer object should be populated for creation.
	//
	// Format: uuid
	CustomerID strfmt.UUID `json:"customerID,omitempty"`

	// The date the orders were issued.
	// Format: date
	DateIssued strfmt.Date `json:"dateIssued,omitempty"`

	// destination duty station
	DestinationDutyStation *DutyStation `json:"destinationDutyStation,omitempty"`

	// ID of the destination duty station.
	//
	// If creating a MoveTaskOrder, this should match an existing duty station.
	//
	// Format: uuid
	DestinationDutyStationID strfmt.UUID `json:"destinationDutyStationID,omitempty"`

	// Uniquely identifies the state of the MoveOrder object (but not the nested objects)
	//
	// It will change everytime the object is updated. Client should store the value.
	// Updates to this MoveOrder will require that this eTag be passed in with the If-Match header.
	//
	// Read Only: true
	ETag string `json:"eTag,omitempty"`

	// entitlement
	Entitlement *Entitlement `json:"entitlement,omitempty"`

	// ID of the MoveOrder object.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// ID of the military orders associated with this move.
	OrderNumber *string `json:"orderNumber,omitempty"`

	// order type
	OrderType OrderType `json:"orderType,omitempty"`

	// origin duty station
	OriginDutyStation *DutyStation `json:"originDutyStation,omitempty"`

	// ID of the origin duty station.
	//
	// If creating a MoveTaskOrder, this should match an existing duty station.
	//
	// Format: uuid
	OriginDutyStationID strfmt.UUID `json:"originDutyStationID,omitempty"`

	// Rank of the service member, must match specific list of available ranks.
	Rank string `json:"rank,omitempty"`

	// Date that the service member must report to the new DutyStation by.
	// Format: date
	ReportByDate strfmt.Date `json:"reportByDate,omitempty"`

	// status
	Status OrdersStatus `json:"status,omitempty"`

	// uploaded orders
	UploadedOrders *Document `json:"uploadedOrders,omitempty"`

	// ID of the uploaded document.
	// Format: uuid
	UploadedOrdersID strfmt.UUID `json:"uploadedOrdersID,omitempty"`
}

// Validate validates this move order
func (m *MoveOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateIssued(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationDutyStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginDutyStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportByDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadedOrders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadedOrdersID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveOrder) validateCustomer(formats strfmt.Registry) error {

	if swag.IsZero(m.Customer) { // not required
		return nil
	}

	if m.Customer != nil {
		if err := m.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *MoveOrder) validateCustomerID(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerID) { // not required
		return nil
	}

	if err := validate.FormatOf("customerID", "body", "uuid", m.CustomerID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveOrder) validateDateIssued(formats strfmt.Registry) error {

	if swag.IsZero(m.DateIssued) { // not required
		return nil
	}

	if err := validate.FormatOf("dateIssued", "body", "date", m.DateIssued.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveOrder) validateDestinationDutyStation(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationDutyStation) { // not required
		return nil
	}

	if m.DestinationDutyStation != nil {
		if err := m.DestinationDutyStation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationDutyStation")
			}
			return err
		}
	}

	return nil
}

func (m *MoveOrder) validateDestinationDutyStationID(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationDutyStationID) { // not required
		return nil
	}

	if err := validate.FormatOf("destinationDutyStationID", "body", "uuid", m.DestinationDutyStationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveOrder) validateEntitlement(formats strfmt.Registry) error {

	if swag.IsZero(m.Entitlement) { // not required
		return nil
	}

	if m.Entitlement != nil {
		if err := m.Entitlement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entitlement")
			}
			return err
		}
	}

	return nil
}

func (m *MoveOrder) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveOrder) validateOrderType(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderType) { // not required
		return nil
	}

	if err := m.OrderType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orderType")
		}
		return err
	}

	return nil
}

func (m *MoveOrder) validateOriginDutyStation(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginDutyStation) { // not required
		return nil
	}

	if m.OriginDutyStation != nil {
		if err := m.OriginDutyStation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originDutyStation")
			}
			return err
		}
	}

	return nil
}

func (m *MoveOrder) validateOriginDutyStationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginDutyStationID) { // not required
		return nil
	}

	if err := validate.FormatOf("originDutyStationID", "body", "uuid", m.OriginDutyStationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveOrder) validateReportByDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportByDate) { // not required
		return nil
	}

	if err := validate.FormatOf("reportByDate", "body", "date", m.ReportByDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveOrder) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *MoveOrder) validateUploadedOrders(formats strfmt.Registry) error {

	if swag.IsZero(m.UploadedOrders) { // not required
		return nil
	}

	if m.UploadedOrders != nil {
		if err := m.UploadedOrders.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadedOrders")
			}
			return err
		}
	}

	return nil
}

func (m *MoveOrder) validateUploadedOrdersID(formats strfmt.Registry) error {

	if swag.IsZero(m.UploadedOrdersID) { // not required
		return nil
	}

	if err := validate.FormatOf("uploadedOrdersID", "body", "uuid", m.UploadedOrdersID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveOrder) UnmarshalBinary(b []byte) error {
	var res MoveOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
