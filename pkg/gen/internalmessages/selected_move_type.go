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

// SelectedMoveType Selected Move Type
//
// swagger:model SelectedMoveType
type SelectedMoveType string

const (

	// SelectedMoveTypeHHG captures enum value "HHG"
	SelectedMoveTypeHHG SelectedMoveType = "HHG"

	// SelectedMoveTypePPM captures enum value "PPM"
	SelectedMoveTypePPM SelectedMoveType = "PPM"

	// SelectedMoveTypeUB captures enum value "UB"
	SelectedMoveTypeUB SelectedMoveType = "UB"

	// SelectedMoveTypePOV captures enum value "POV"
	SelectedMoveTypePOV SelectedMoveType = "POV"

	// SelectedMoveTypeHHGINTONTSDOMESTIC captures enum value "HHG_INTO_NTS_DOMESTIC"
	SelectedMoveTypeHHGINTONTSDOMESTIC SelectedMoveType = "HHG_INTO_NTS_DOMESTIC"

	// SelectedMoveTypeHHGOUTOFNTSDOMESTIC captures enum value "HHG_OUTOF_NTS_DOMESTIC"
	SelectedMoveTypeHHGOUTOFNTSDOMESTIC SelectedMoveType = "HHG_OUTOF_NTS_DOMESTIC"

	// SelectedMoveTypeHHGPPM captures enum value "HHG_PPM"
	SelectedMoveTypeHHGPPM SelectedMoveType = "HHG_PPM"
)

// for schema
var selectedMoveTypeEnum []interface{}

func init() {
	var res []SelectedMoveType
	if err := json.Unmarshal([]byte(`["HHG","PPM","UB","POV","HHG_INTO_NTS_DOMESTIC","HHG_OUTOF_NTS_DOMESTIC","HHG_PPM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		selectedMoveTypeEnum = append(selectedMoveTypeEnum, v)
	}
}

func (m SelectedMoveType) validateSelectedMoveTypeEnum(path, location string, value SelectedMoveType) error {
	if err := validate.EnumCase(path, location, value, selectedMoveTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this selected move type
func (m SelectedMoveType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSelectedMoveTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
