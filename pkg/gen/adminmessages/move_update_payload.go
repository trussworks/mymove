// Code generated by go-swagger; DO NOT EDIT.

package adminmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MoveUpdatePayload move update payload
//
// swagger:model MoveUpdatePayload
type MoveUpdatePayload struct {

	// show
	// Required: true
	Show *bool `json:"show"`
}

// Validate validates this move update payload
func (m *MoveUpdatePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveUpdatePayload) validateShow(formats strfmt.Registry) error {

	if err := validate.Required("show", "body", m.Show); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveUpdatePayload) UnmarshalBinary(b []byte) error {
	var res MoveUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
