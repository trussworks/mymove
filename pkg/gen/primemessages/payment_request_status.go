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

// PaymentRequestStatus Payment Request Status
//
// swagger:model PaymentRequestStatus
type PaymentRequestStatus string

const (

	// PaymentRequestStatusPENDING captures enum value "PENDING"
	PaymentRequestStatusPENDING PaymentRequestStatus = "PENDING"

	// PaymentRequestStatusREVIEWED captures enum value "REVIEWED"
	PaymentRequestStatusREVIEWED PaymentRequestStatus = "REVIEWED"

	// PaymentRequestStatusREVIEWEDANDALLSERVICEITEMSREJECTED captures enum value "REVIEWED_AND_ALL_SERVICE_ITEMS_REJECTED"
	PaymentRequestStatusREVIEWEDANDALLSERVICEITEMSREJECTED PaymentRequestStatus = "REVIEWED_AND_ALL_SERVICE_ITEMS_REJECTED"

	// PaymentRequestStatusSENTTOGEX captures enum value "SENT_TO_GEX"
	PaymentRequestStatusSENTTOGEX PaymentRequestStatus = "SENT_TO_GEX"

	// PaymentRequestStatusRECEIVEDBYGEX captures enum value "RECEIVED_BY_GEX"
	PaymentRequestStatusRECEIVEDBYGEX PaymentRequestStatus = "RECEIVED_BY_GEX"

	// PaymentRequestStatusPAID captures enum value "PAID"
	PaymentRequestStatusPAID PaymentRequestStatus = "PAID"
)

// for schema
var paymentRequestStatusEnum []interface{}

func init() {
	var res []PaymentRequestStatus
	if err := json.Unmarshal([]byte(`["PENDING","REVIEWED","REVIEWED_AND_ALL_SERVICE_ITEMS_REJECTED","SENT_TO_GEX","RECEIVED_BY_GEX","PAID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentRequestStatusEnum = append(paymentRequestStatusEnum, v)
	}
}

func (m PaymentRequestStatus) validatePaymentRequestStatusEnum(path, location string, value PaymentRequestStatus) error {
	if err := validate.EnumCase(path, location, value, paymentRequestStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment request status
func (m PaymentRequestStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentRequestStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
