// Code generated by go-swagger; DO NOT EDIT.

package payment_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// NewUpdatePaymentServiceItemStatusParams creates a new UpdatePaymentServiceItemStatusParams object
// no default values defined in spec.
func NewUpdatePaymentServiceItemStatusParams() UpdatePaymentServiceItemStatusParams {

	return UpdatePaymentServiceItemStatusParams{}
}

// UpdatePaymentServiceItemStatusParams contains all the bound params for the update payment service item status operation
// typically these are obtained from a http.Request
//
// swagger:parameters updatePaymentServiceItemStatus
type UpdatePaymentServiceItemStatusParams struct {

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
	Body *ghcmessages.PaymentServiceItem
	/*ID of move to use
	  Required: true
	  In: path
	*/
	MoveTaskOrderID string
	/*ID of payment service item to use
	  Required: true
	  In: path
	*/
	PaymentServiceItemID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdatePaymentServiceItemStatusParams() beforehand.
func (o *UpdatePaymentServiceItemStatusParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindIfMatch(r.Header[http.CanonicalHeaderKey("If-Match")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body ghcmessages.PaymentServiceItem
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
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
		res = append(res, errors.Required("body", "body", ""))
	}
	rMoveTaskOrderID, rhkMoveTaskOrderID, _ := route.Params.GetOK("moveTaskOrderID")
	if err := o.bindMoveTaskOrderID(rMoveTaskOrderID, rhkMoveTaskOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	rPaymentServiceItemID, rhkPaymentServiceItemID, _ := route.Params.GetOK("paymentServiceItemID")
	if err := o.bindPaymentServiceItemID(rPaymentServiceItemID, rhkPaymentServiceItemID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIfMatch binds and validates parameter IfMatch from header.
func (o *UpdatePaymentServiceItemStatusParams) bindIfMatch(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("If-Match", "header", rawData)
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
func (o *UpdatePaymentServiceItemStatusParams) bindMoveTaskOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.MoveTaskOrderID = raw

	return nil
}

// bindPaymentServiceItemID binds and validates parameter PaymentServiceItemID from path.
func (o *UpdatePaymentServiceItemStatusParams) bindPaymentServiceItemID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PaymentServiceItemID = raw

	return nil
}
