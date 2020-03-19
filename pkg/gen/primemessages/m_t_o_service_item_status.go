// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOServiceItemStatus m t o service item status
// swagger:model MTOServiceItemStatus
type MTOServiceItemStatus struct {

	// status
	// Enum: [APPROVED SUBMITTED REJECTED]
	Status string `json:"status,omitempty"`
}

// Validate validates this m t o service item status
func (m *MTOServiceItemStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mTOServiceItemStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APPROVED","SUBMITTED","REJECTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOServiceItemStatusTypeStatusPropEnum = append(mTOServiceItemStatusTypeStatusPropEnum, v)
	}
}

const (

	// MTOServiceItemStatusStatusAPPROVED captures enum value "APPROVED"
	MTOServiceItemStatusStatusAPPROVED string = "APPROVED"

	// MTOServiceItemStatusStatusSUBMITTED captures enum value "SUBMITTED"
	MTOServiceItemStatusStatusSUBMITTED string = "SUBMITTED"

	// MTOServiceItemStatusStatusREJECTED captures enum value "REJECTED"
	MTOServiceItemStatusStatusREJECTED string = "REJECTED"
)

// prop value enum
func (m *MTOServiceItemStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, mTOServiceItemStatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MTOServiceItemStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOServiceItemStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOServiceItemStatus) UnmarshalBinary(b []byte) error {
	var res MTOServiceItemStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}