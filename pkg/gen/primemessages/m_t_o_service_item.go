// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOServiceItem m t o service item
// swagger:model MTOServiceItem
type MTOServiceItem struct {

	// id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// move task order ID
	// Format: uuid
	MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

	// re service code
	ReServiceCode string `json:"reServiceCode,omitempty"`

	// re service ID
	// Format: uuid
	ReServiceID strfmt.UUID `json:"reServiceID,omitempty"`

	// re service name
	ReServiceName string `json:"reServiceName,omitempty"`
}

// Validate validates this m t o service item
func (m *MTOServiceItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTaskOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReServiceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MTOServiceItem) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateReServiceID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReServiceID) { // not required
		return nil
	}

	if err := validate.FormatOf("reServiceID", "body", "uuid", m.ReServiceID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOServiceItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOServiceItem) UnmarshalBinary(b []byte) error {
	var res MTOServiceItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}