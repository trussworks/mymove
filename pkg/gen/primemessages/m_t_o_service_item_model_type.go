// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// MTOServiceItemModelType Describes all model sub-types for a MTOServiceItem model.
//
// Using this list, choose the correct modelType in the dropdown, corresponding to the service item type.
//   * DOFSIT, DOASIT - MTOServiceItemOriginSIT
//   * DDFSIT, DDASIT - MTOServiceItemDestSIT
//   * DOSHUT, DDSHUT - MTOServiceItemShuttle
//   * DCRT, DCRTSA, DUCRT - MTOServiceItemDomesticCrating
//
// The documentation will then update with the supported fields.
//
//
// swagger:model MTOServiceItemModelType
type MTOServiceItemModelType string

const (

	// MTOServiceItemModelTypeMTOServiceItemBasic captures enum value "MTOServiceItemBasic"
	MTOServiceItemModelTypeMTOServiceItemBasic MTOServiceItemModelType = "MTOServiceItemBasic"

	// MTOServiceItemModelTypeMTOServiceItemOriginSIT captures enum value "MTOServiceItemOriginSIT"
	MTOServiceItemModelTypeMTOServiceItemOriginSIT MTOServiceItemModelType = "MTOServiceItemOriginSIT"

	// MTOServiceItemModelTypeMTOServiceItemDestSIT captures enum value "MTOServiceItemDestSIT"
	MTOServiceItemModelTypeMTOServiceItemDestSIT MTOServiceItemModelType = "MTOServiceItemDestSIT"

	// MTOServiceItemModelTypeMTOServiceItemShuttle captures enum value "MTOServiceItemShuttle"
	MTOServiceItemModelTypeMTOServiceItemShuttle MTOServiceItemModelType = "MTOServiceItemShuttle"

	// MTOServiceItemModelTypeMTOServiceItemDomesticCrating captures enum value "MTOServiceItemDomesticCrating"
	MTOServiceItemModelTypeMTOServiceItemDomesticCrating MTOServiceItemModelType = "MTOServiceItemDomesticCrating"

	// MTOServiceItemModelTypeMTOServiceItemSITDeparture captures enum value "MTOServiceItemSITDeparture"
	MTOServiceItemModelTypeMTOServiceItemSITDeparture MTOServiceItemModelType = "MTOServiceItemSITDeparture"
)

// for schema
var mTOServiceItemModelTypeEnum []interface{}

func init() {
	var res []MTOServiceItemModelType
	if err := json.Unmarshal([]byte(`["MTOServiceItemBasic","MTOServiceItemOriginSIT","MTOServiceItemDestSIT","MTOServiceItemShuttle","MTOServiceItemDomesticCrating","MTOServiceItemSITDeparture"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOServiceItemModelTypeEnum = append(mTOServiceItemModelTypeEnum, v)
	}
}

func (m MTOServiceItemModelType) validateMTOServiceItemModelTypeEnum(path, location string, value MTOServiceItemModelType) error {
	if err := validate.EnumCase(path, location, value, mTOServiceItemModelTypeEnum, true); err != nil {
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
