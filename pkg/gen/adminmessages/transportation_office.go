// Code generated by go-swagger; DO NOT EDIT.

package adminmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TransportationOffice transportation office
// swagger:model TransportationOffice
type TransportationOffice struct {

	// address
	// Required: true
	Address *Address `json:"address"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// gbloc
	// Pattern: ^[A-Z]{4}$
	Gbloc string `json:"gbloc,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// latitude
	Latitude float32 `json:"latitude,omitempty"`

	// longitude
	Longitude float32 `json:"longitude,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// phone lines
	PhoneLines []string `json:"phoneLines"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`
}

// Validate validates this transportation office
func (m *TransportationOffice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGbloc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransportationOffice) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *TransportationOffice) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransportationOffice) validateGbloc(formats strfmt.Registry) error {

	if swag.IsZero(m.Gbloc) { // not required
		return nil
	}

	if err := validate.Pattern("gbloc", "body", string(m.Gbloc), `^[A-Z]{4}$`); err != nil {
		return err
	}

	return nil
}

func (m *TransportationOffice) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransportationOffice) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *TransportationOffice) validatePhoneLines(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneLines) { // not required
		return nil
	}

	for i := 0; i < len(m.PhoneLines); i++ {

		if err := validate.Pattern("phoneLines"+"."+strconv.Itoa(i), "body", string(m.PhoneLines[i]), `^[2-9]\d{2}-\d{3}-\d{4}$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *TransportationOffice) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransportationOffice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransportationOffice) UnmarshalBinary(b []byte) error {
	var res TransportationOffice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
