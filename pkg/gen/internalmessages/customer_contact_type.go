// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CustomerContactType Describes a customer contact type for a MTOServiceItemDomesticDestSIT.
// swagger:model CustomerContactType
type CustomerContactType string

const (

	// CustomerContactTypeFIRST captures enum value "FIRST"
	CustomerContactTypeFIRST CustomerContactType = "FIRST"

	// CustomerContactTypeSECOND captures enum value "SECOND"
	CustomerContactTypeSECOND CustomerContactType = "SECOND"
)

// for schema
var customerContactTypeEnum []interface{}

func init() {
	var res []CustomerContactType
	if err := json.Unmarshal([]byte(`["FIRST","SECOND"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerContactTypeEnum = append(customerContactTypeEnum, v)
	}
}

func (m CustomerContactType) validateCustomerContactTypeEnum(path, location string, value CustomerContactType) error {
	if err := validate.Enum(path, location, value, customerContactTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this customer contact type
func (m CustomerContactType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCustomerContactTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
