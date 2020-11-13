// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewRequestPPMExpenseSummaryParams creates a new RequestPPMExpenseSummaryParams object
// no default values defined in spec.
func NewRequestPPMExpenseSummaryParams() RequestPPMExpenseSummaryParams {

	return RequestPPMExpenseSummaryParams{}
}

// RequestPPMExpenseSummaryParams contains all the bound params for the request p p m expense summary operation
// typically these are obtained from a http.Request
//
// swagger:parameters requestPPMExpenseSummary
type RequestPPMExpenseSummaryParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*UUID of the PPM
	  Required: true
	  In: path
	*/
	PersonallyProcuredMoveID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRequestPPMExpenseSummaryParams() beforehand.
func (o *RequestPPMExpenseSummaryParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rPersonallyProcuredMoveID, rhkPersonallyProcuredMoveID, _ := route.Params.GetOK("personallyProcuredMoveId")
	if err := o.bindPersonallyProcuredMoveID(rPersonallyProcuredMoveID, rhkPersonallyProcuredMoveID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPersonallyProcuredMoveID binds and validates parameter PersonallyProcuredMoveID from path.
func (o *RequestPPMExpenseSummaryParams) bindPersonallyProcuredMoveID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("personallyProcuredMoveId", "path", "strfmt.UUID", raw)
	}
	o.PersonallyProcuredMoveID = *(value.(*strfmt.UUID))

	if err := o.validatePersonallyProcuredMoveID(formats); err != nil {
		return err
	}

	return nil
}

// validatePersonallyProcuredMoveID carries on validations for parameter PersonallyProcuredMoveID
func (o *RequestPPMExpenseSummaryParams) validatePersonallyProcuredMoveID(formats strfmt.Registry) error {

	if err := validate.FormatOf("personallyProcuredMoveId", "path", "uuid", o.PersonallyProcuredMoveID.String(), formats); err != nil {
		return err
	}
	return nil
}
