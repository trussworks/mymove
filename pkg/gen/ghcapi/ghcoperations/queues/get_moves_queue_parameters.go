// Code generated by go-swagger; DO NOT EDIT.

package queues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMovesQueueParams creates a new GetMovesQueueParams object
// no default values defined in spec.
func NewGetMovesQueueParams() GetMovesQueueParams {

	return GetMovesQueueParams{}
}

// GetMovesQueueParams contains all the bound params for the get moves queue operation
// typically these are obtained from a http.Request
//
// swagger:parameters getMovesQueue
type GetMovesQueueParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Branch *string
	/*Filtering for the status.
	  Unique: true
	  In: query
	*/
	Status []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetMovesQueueParams() beforehand.
func (o *GetMovesQueueParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qBranch, qhkBranch, _ := qs.GetOK("branch")
	if err := o.bindBranch(qBranch, qhkBranch, route.Formats); err != nil {
		res = append(res, err)
	}

	qStatus, qhkStatus, _ := qs.GetOK("status")
	if err := o.bindStatus(qStatus, qhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBranch binds and validates parameter Branch from query.
func (o *GetMovesQueueParams) bindBranch(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Branch = &raw

	return nil
}

// bindStatus binds and validates array parameter Status from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetMovesQueueParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {

	var qvStatus string
	if len(rawData) > 0 {
		qvStatus = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	statusIC := swag.SplitByFormat(qvStatus, "")
	if len(statusIC) == 0 {
		return nil
	}

	var statusIR []string
	for i, statusIV := range statusIC {
		statusI := statusIV

		if err := validate.Enum(fmt.Sprintf("%s.%v", "status", i), "query", statusI, []interface{}{"New move", "Approvals requested", "Move approved"}); err != nil {
			return err
		}

		statusIR = append(statusIR, statusI)
	}

	o.Status = statusIR
	if err := o.validateStatus(formats); err != nil {
		return err
	}

	return nil
}

// validateStatus carries on validations for parameter Status
func (o *GetMovesQueueParams) validateStatus(formats strfmt.Registry) error {

	// uniqueItems: true
	if err := validate.UniqueItems("status", "query", o.Status); err != nil {
		return err
	}

	return nil
}
