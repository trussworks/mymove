// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOServiceItemCustomerContact Customer contact information for a destination SIT service item
//
// swagger:model MTOServiceItemCustomerContact
type MTOServiceItemCustomerContact struct {

	// First available date that Prime can deliver SIT service item.
	// Format: date
	FirstAvailableDeliveryDate strfmt.Date `json:"firstAvailableDeliveryDate,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Time of delivery corresponding to `firstAvailableDeliveryDate`.
	TimeMilitary string `json:"timeMilitary,omitempty"`

	// type
	Type CustomerContactType `json:"type,omitempty"`
}

// Validate validates this m t o service item customer contact
func (m *MTOServiceItemCustomerContact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstAvailableDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MTOServiceItemCustomerContact) validateFirstAvailableDeliveryDate(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstAvailableDeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("firstAvailableDeliveryDate", "body", "date", m.FirstAvailableDeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemCustomerContact) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemCustomerContact) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOServiceItemCustomerContact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOServiceItemCustomerContact) UnmarshalBinary(b []byte) error {
	var res MTOServiceItemCustomerContact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
