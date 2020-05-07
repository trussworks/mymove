// Code generated by go-swagger; DO NOT EDIT.

package move

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMoveParams creates a new GetMoveParams object
// no default values defined in spec.
func NewGetMoveParams() GetMoveParams {

	return GetMoveParams{}
}

// GetMoveParams contains all the bound params for the get move operation
// typically these are obtained from a http.Request
//
// swagger:parameters getMove
type GetMoveParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Code used to identify a move in the system
	  Required: true
	  In: path
	*/
	Locator string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetMoveParams() beforehand.
func (o *GetMoveParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rLocator, rhkLocator, _ := route.Params.GetOK("locator")
	if err := o.bindLocator(rLocator, rhkLocator, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLocator binds and validates parameter Locator from path.
func (o *GetMoveParams) bindLocator(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Locator = raw

	return nil
}
