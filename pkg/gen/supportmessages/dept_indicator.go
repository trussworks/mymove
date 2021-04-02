// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DeptIndicator Dept. indicator
//
// swagger:model DeptIndicator
type DeptIndicator string

const (

	// DeptIndicatorNAVYANDMARINES captures enum value "NAVY_AND_MARINES"
	DeptIndicatorNAVYANDMARINES DeptIndicator = "NAVY_AND_MARINES"

	// DeptIndicatorARMY captures enum value "ARMY"
	DeptIndicatorARMY DeptIndicator = "ARMY"

	// DeptIndicatorAIRFORCE captures enum value "AIR_FORCE"
	DeptIndicatorAIRFORCE DeptIndicator = "AIR_FORCE"

	// DeptIndicatorCOASTGUARD captures enum value "COAST_GUARD"
	DeptIndicatorCOASTGUARD DeptIndicator = "COAST_GUARD"
)

// for schema
var deptIndicatorEnum []interface{}

func init() {
	var res []DeptIndicator
	if err := json.Unmarshal([]byte(`["NAVY_AND_MARINES","ARMY","AIR_FORCE","COAST_GUARD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deptIndicatorEnum = append(deptIndicatorEnum, v)
	}
}

func (m DeptIndicator) validateDeptIndicatorEnum(path, location string, value DeptIndicator) error {
	if err := validate.EnumCase(path, location, value, deptIndicatorEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this dept indicator
func (m DeptIndicator) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDeptIndicatorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
