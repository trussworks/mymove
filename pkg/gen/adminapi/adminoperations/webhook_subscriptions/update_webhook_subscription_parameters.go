// Code generated by go-swagger; DO NOT EDIT.

package webhook_subscriptions

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

	"github.com/transcom/mymove/pkg/gen/adminmessages"
)

// NewUpdateWebhookSubscriptionParams creates a new UpdateWebhookSubscriptionParams object
// no default values defined in spec.
func NewUpdateWebhookSubscriptionParams() UpdateWebhookSubscriptionParams {

	return UpdateWebhookSubscriptionParams{}
}

// UpdateWebhookSubscriptionParams contains all the bound params for the update webhook subscription operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateWebhookSubscription
type UpdateWebhookSubscriptionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error.

	  Required: true
	  In: header
	*/
	IfMatch string
	/*Webhook subscription information
	  Required: true
	  In: body
	*/
	WebhookSubscription *adminmessages.WebhookSubscription
	/*
	  Required: true
	  In: path
	*/
	WebhookSubscriptionID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateWebhookSubscriptionParams() beforehand.
func (o *UpdateWebhookSubscriptionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindIfMatch(r.Header[http.CanonicalHeaderKey("If-Match")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body adminmessages.WebhookSubscription
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("webhookSubscription", "body", ""))
			} else {
				res = append(res, errors.NewParseError("webhookSubscription", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.WebhookSubscription = &body
			}
		}
	} else {
		res = append(res, errors.Required("webhookSubscription", "body", ""))
	}
	rWebhookSubscriptionID, rhkWebhookSubscriptionID, _ := route.Params.GetOK("webhookSubscriptionId")
	if err := o.bindWebhookSubscriptionID(rWebhookSubscriptionID, rhkWebhookSubscriptionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIfMatch binds and validates parameter IfMatch from header.
func (o *UpdateWebhookSubscriptionParams) bindIfMatch(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindWebhookSubscriptionID binds and validates parameter WebhookSubscriptionID from path.
func (o *UpdateWebhookSubscriptionParams) bindWebhookSubscriptionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("webhookSubscriptionId", "path", "strfmt.UUID", raw)
	}
	o.WebhookSubscriptionID = *(value.(*strfmt.UUID))

	if err := o.validateWebhookSubscriptionID(formats); err != nil {
		return err
	}

	return nil
}

// validateWebhookSubscriptionID carries on validations for parameter WebhookSubscriptionID
func (o *UpdateWebhookSubscriptionParams) validateWebhookSubscriptionID(formats strfmt.Registry) error {

	if err := validate.FormatOf("webhookSubscriptionId", "path", "uuid", o.WebhookSubscriptionID.String(), formats); err != nil {
		return err
	}
	return nil
}
