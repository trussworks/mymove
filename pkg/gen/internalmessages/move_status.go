// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// MoveStatus Move status
//
// swagger:model MoveStatus
type MoveStatus string

const (

	// MoveStatusDRAFT captures enum value "DRAFT"
	MoveStatusDRAFT MoveStatus = "DRAFT"

	// MoveStatusSUBMITTED captures enum value "SUBMITTED"
	MoveStatusSUBMITTED MoveStatus = "SUBMITTED"

	// MoveStatusAPPROVED captures enum value "APPROVED"
	MoveStatusAPPROVED MoveStatus = "APPROVED"

	// MoveStatusCANCELED captures enum value "CANCELED"
	MoveStatusCANCELED MoveStatus = "CANCELED"

	// MoveStatusNEEDSSERVICECOUNSELING captures enum value "NEEDS SERVICE COUNSELING"
	MoveStatusNEEDSSERVICECOUNSELING MoveStatus = "NEEDS SERVICE COUNSELING"
)

// for schema
var moveStatusEnum []interface{}

func init() {
	var res []MoveStatus
	if err := json.Unmarshal([]byte(`["DRAFT","SUBMITTED","APPROVED","CANCELED","NEEDS SERVICE COUNSELING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moveStatusEnum = append(moveStatusEnum, v)
	}
}

func (m MoveStatus) validateMoveStatusEnum(path, location string, value MoveStatus) error {
	if err := validate.EnumCase(path, location, value, moveStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this move status
func (m MoveStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMoveStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
