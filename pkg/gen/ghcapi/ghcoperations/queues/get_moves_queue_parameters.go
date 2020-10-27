// Code generated by go-swagger; DO NOT EDIT.

package queues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

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
	/*
	  In: query
	*/
	DestinationDutyStation *string
	/*
	  In: query
	*/
	DodID *string
	/*
	  In: query
	*/
	LastName *string
	/*
	  In: query
	*/
	MoveID *string
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

	qDestinationDutyStation, qhkDestinationDutyStation, _ := qs.GetOK("destinationDutyStation")
	if err := o.bindDestinationDutyStation(qDestinationDutyStation, qhkDestinationDutyStation, route.Formats); err != nil {
		res = append(res, err)
	}

	qDodID, qhkDodID, _ := qs.GetOK("dodID")
	if err := o.bindDodID(qDodID, qhkDodID, route.Formats); err != nil {
		res = append(res, err)
	}

	qLastName, qhkLastName, _ := qs.GetOK("lastName")
	if err := o.bindLastName(qLastName, qhkLastName, route.Formats); err != nil {
		res = append(res, err)
	}

	qMoveID, qhkMoveID, _ := qs.GetOK("moveID")
	if err := o.bindMoveID(qMoveID, qhkMoveID, route.Formats); err != nil {
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

// bindDestinationDutyStation binds and validates parameter DestinationDutyStation from query.
func (o *GetMovesQueueParams) bindDestinationDutyStation(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.DestinationDutyStation = &raw

	return nil
}

// bindDodID binds and validates parameter DodID from query.
func (o *GetMovesQueueParams) bindDodID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.DodID = &raw

	return nil
}

// bindLastName binds and validates parameter LastName from query.
func (o *GetMovesQueueParams) bindLastName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.LastName = &raw

	return nil
}

// bindMoveID binds and validates parameter MoveID from query.
func (o *GetMovesQueueParams) bindMoveID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.MoveID = &raw

	return nil
}
