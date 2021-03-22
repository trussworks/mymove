// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Order order
//
// swagger:model Order
type Order struct {

	// customer
	Customer *Customer `json:"customer,omitempty"`

	// customer ID
	// Format: uuid
	CustomerID strfmt.UUID `json:"customerID,omitempty"`

	// destination duty station
	DestinationDutyStation *DutyStation `json:"destinationDutyStation,omitempty"`

	// e tag
	// Read Only: true
	ETag string `json:"eTag,omitempty"`

	// entitlement
	Entitlement *Entitlements `json:"entitlement,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// lines of accounting
	// Required: true
	LinesOfAccounting *string `json:"linesOfAccounting"`

	// order number
	// Required: true
	OrderNumber *string `json:"orderNumber"`

	// origin duty station
	OriginDutyStation *DutyStation `json:"originDutyStation,omitempty"`

	// rank
	// Required: true
	Rank *string `json:"rank"`

	// report by date
	// Format: date
	ReportByDate strfmt.Date `json:"reportByDate,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinesOfAccounting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportByDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Order) validateCustomer(formats strfmt.Registry) error {

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

func (m *Order) validateCustomerID(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerID) { // not required
		return nil
	}

	if err := validate.FormatOf("customerID", "body", "uuid", m.CustomerID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateDestinationDutyStation(formats strfmt.Registry) error {

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

func (m *Order) validateEntitlement(formats strfmt.Registry) error {

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

func (m *Order) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateLinesOfAccounting(formats strfmt.Registry) error {

	if err := validate.Required("linesOfAccounting", "body", m.LinesOfAccounting); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateOrderNumber(formats strfmt.Registry) error {

	if err := validate.Required("orderNumber", "body", m.OrderNumber); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateOriginDutyStation(formats strfmt.Registry) error {

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

func (m *Order) validateRank(formats strfmt.Registry) error {

	if err := validate.Required("rank", "body", m.Rank); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateReportByDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportByDate) { // not required
		return nil
	}

	if err := validate.FormatOf("reportByDate", "body", "date", m.ReportByDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}