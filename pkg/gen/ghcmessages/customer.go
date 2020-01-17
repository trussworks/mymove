// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Customer customer
// swagger:model Customer
type Customer struct {

	// Agency customer is affilated with
	Agency string `json:"agency,omitempty"`

	// current address
	CurrentAddress *Address `json:"current_address,omitempty"`

	// destination address
	DestinationAddress *Address `json:"destination_address,omitempty"`

	// dod ID
	DodID string `json:"dodID,omitempty"`

	// email
	// Pattern: ^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
	Email *string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// phone
	// Pattern: ^[2-9]\d{2}-\d{3}-\d{4}$
	Phone *string `json:"phone,omitempty"`

	// user ID
	// Format: uuid
	UserID strfmt.UUID `json:"userID,omitempty"`
}

// Validate validates this customer
func (m *Customer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Customer) validateCurrentAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentAddress) { // not required
		return nil
	}

	if m.CurrentAddress != nil {
		if err := m.CurrentAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current_address")
			}
			return err
		}
	}

	return nil
}

func (m *Customer) validateDestinationAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationAddress) { // not required
		return nil
	}

	if m.DestinationAddress != nil {
		if err := m.DestinationAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination_address")
			}
			return err
		}
	}

	return nil
}

func (m *Customer) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.Pattern("email", "body", string(*m.Email), `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validatePhone(formats strfmt.Registry) error {

	if swag.IsZero(m.Phone) { // not required
		return nil
	}

	if err := validate.Pattern("phone", "body", string(*m.Phone), `^[2-9]\d{2}-\d{3}-\d{4}$`); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.UserID) { // not required
		return nil
	}

	if err := validate.FormatOf("userID", "body", "uuid", m.UserID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Customer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Customer) UnmarshalBinary(b []byte) error {
	var res Customer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
