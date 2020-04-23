// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SignedCertificationPayload signed certification payload
// swagger:model SignedCertificationPayload
type SignedCertificationPayload struct {

	// certification text
	// Required: true
	CertificationText *string `json:"certification_text"`

	// certification type
	CertificationType *SignedCertificationType `json:"certification_type,omitempty"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// Date
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// move id
	// Required: true
	// Format: uuid
	MoveID *strfmt.UUID `json:"move_id"`

	// personally procured move id
	// Format: uuid
	PersonallyProcuredMoveID *strfmt.UUID `json:"personally_procured_move_id,omitempty"`

	// Signature
	// Required: true
	Signature *string `json:"signature"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this signed certification payload
func (m *SignedCertificationPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificationText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonallyProcuredMoveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignature(formats); err != nil {
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

func (m *SignedCertificationPayload) validateCertificationText(formats strfmt.Registry) error {

	if err := validate.Required("certification_text", "body", m.CertificationText); err != nil {
		return err
	}

	return nil
}

func (m *SignedCertificationPayload) validateCertificationType(formats strfmt.Registry) error {

	if swag.IsZero(m.CertificationType) { // not required
		return nil
	}

	if m.CertificationType != nil {
		if err := m.CertificationType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certification_type")
			}
			return err
		}
	}

	return nil
}

func (m *SignedCertificationPayload) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SignedCertificationPayload) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SignedCertificationPayload) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SignedCertificationPayload) validateMoveID(formats strfmt.Registry) error {

	if err := validate.Required("move_id", "body", m.MoveID); err != nil {
		return err
	}

	if err := validate.FormatOf("move_id", "body", "uuid", m.MoveID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SignedCertificationPayload) validatePersonallyProcuredMoveID(formats strfmt.Registry) error {

	if swag.IsZero(m.PersonallyProcuredMoveID) { // not required
		return nil
	}

	if err := validate.FormatOf("personally_procured_move_id", "body", "uuid", m.PersonallyProcuredMoveID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SignedCertificationPayload) validateSignature(formats strfmt.Registry) error {

	if err := validate.Required("signature", "body", m.Signature); err != nil {
		return err
	}

	return nil
}

func (m *SignedCertificationPayload) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SignedCertificationPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignedCertificationPayload) UnmarshalBinary(b []byte) error {
	var res SignedCertificationPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
