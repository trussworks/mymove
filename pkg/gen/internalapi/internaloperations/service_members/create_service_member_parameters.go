// Code generated by go-swagger; DO NOT EDIT.

package service_members

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"
)

// NewCreateServiceMemberParams creates a new CreateServiceMemberParams object
// no default values defined in spec.
func NewCreateServiceMemberParams() CreateServiceMemberParams {

	return CreateServiceMemberParams{}
}

// CreateServiceMemberParams contains all the bound params for the create service member operation
// typically these are obtained from a http.Request
//
// swagger:parameters createServiceMember
type CreateServiceMemberParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	CreateServiceMemberPayload *internalmessages.CreateServiceMemberPayload
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateServiceMemberParams() beforehand.
func (o *CreateServiceMemberParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body internalmessages.CreateServiceMemberPayload
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("createServiceMemberPayload", "body"))
			} else {
				res = append(res, errors.NewParseError("createServiceMemberPayload", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.CreateServiceMemberPayload = &body
			}
		}
	} else {
		res = append(res, errors.Required("createServiceMemberPayload", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
