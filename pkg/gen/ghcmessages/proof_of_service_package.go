// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProofOfServicePackage proof of service package
//
// swagger:model ProofOfServicePackage
type ProofOfServicePackage struct {

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// uploads
	Uploads []*Upload `json:"uploads"`
}

// Validate validates this proof of service package
func (m *ProofOfServicePackage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploads(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProofOfServicePackage) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProofOfServicePackage) validateUploads(formats strfmt.Registry) error {

	if swag.IsZero(m.Uploads) { // not required
		return nil
	}

	for i := 0; i < len(m.Uploads); i++ {
		if swag.IsZero(m.Uploads[i]) { // not required
			continue
		}

		if m.Uploads[i] != nil {
			if err := m.Uploads[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("uploads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProofOfServicePackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProofOfServicePackage) UnmarshalBinary(b []byte) error {
	var res ProofOfServicePackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
