// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// NewUpdateMTOServiceItemStatusParams creates a new UpdateMTOServiceItemStatusParams object
// no default values defined in spec.
func NewUpdateMTOServiceItemStatusParams() UpdateMTOServiceItemStatusParams {

	return UpdateMTOServiceItemStatusParams{}
}

// UpdateMTOServiceItemStatusParams contains all the bound params for the update m t o service item status operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateMTOServiceItemStatus
type UpdateMTOServiceItemStatusParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	IfMatch string
	/*
	  Required: true
	  In: body
	*/
	Body *ghcmessages.MTOServiceItem
	/*ID of move order to use
	  Required: true
	  In: path
	*/
	MoveTaskOrderID string
	/*ID of line item to use
	  Required: true
	  In: path
	*/
	MtoServiceItemID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateMTOServiceItemStatusParams() beforehand.
func (o *UpdateMTOServiceItemStatusParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindIfMatch(r.Header[http.CanonicalHeaderKey("If-Match")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body ghcmessages.MTOServiceItem
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	rMoveTaskOrderID, rhkMoveTaskOrderID, _ := route.Params.GetOK("moveTaskOrderID")
	if err := o.bindMoveTaskOrderID(rMoveTaskOrderID, rhkMoveTaskOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	rMtoServiceItemID, rhkMtoServiceItemID, _ := route.Params.GetOK("mtoServiceItemID")
	if err := o.bindMtoServiceItemID(rMtoServiceItemID, rhkMtoServiceItemID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIfMatch binds and validates parameter IfMatch from header.
func (o *UpdateMTOServiceItemStatusParams) bindIfMatch(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("If-Match", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("If-Match", "header", raw); err != nil {
		return err
	}

	o.IfMatch = raw

	return nil
}

// bindMoveTaskOrderID binds and validates parameter MoveTaskOrderID from path.
func (o *UpdateMTOServiceItemStatusParams) bindMoveTaskOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.MoveTaskOrderID = raw

	return nil
}

// bindMtoServiceItemID binds and validates parameter MtoServiceItemID from path.
func (o *UpdateMTOServiceItemStatusParams) bindMtoServiceItemID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.MtoServiceItemID = raw

	return nil
}
