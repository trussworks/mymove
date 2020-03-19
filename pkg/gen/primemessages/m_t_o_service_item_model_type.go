// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// MTOServiceItemModelType Describes all model sub-types for a MTOServiceItem model
// swagger:model MTOServiceItemModelType
type MTOServiceItemModelType string

const (

	// MTOServiceItemModelTypeMTOServiceItemBasic captures enum value "MTOServiceItemBasic"
	MTOServiceItemModelTypeMTOServiceItemBasic MTOServiceItemModelType = "MTOServiceItemBasic"

	// MTOServiceItemModelTypeMTOServiceItemDOFSIT captures enum value "MTOServiceItemDOFSIT"
	MTOServiceItemModelTypeMTOServiceItemDOFSIT MTOServiceItemModelType = "MTOServiceItemDOFSIT"
)

// for schema
var mTOServiceItemModelTypeEnum []interface{}

func init() {
	var res []MTOServiceItemModelType
	if err := json.Unmarshal([]byte(`["MTOServiceItemBasic","MTOServiceItemDOFSIT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOServiceItemModelTypeEnum = append(mTOServiceItemModelTypeEnum, v)
	}
}

func (m MTOServiceItemModelType) validateMTOServiceItemModelTypeEnum(path, location string, value MTOServiceItemModelType) error {
	if err := validate.Enum(path, location, value, mTOServiceItemModelTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this m t o service item model type
func (m MTOServiceItemModelType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMTOServiceItemModelTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}