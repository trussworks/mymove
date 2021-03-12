// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/supportmessages"
)

// NewReceiveWebhookNotificationParams creates a new ReceiveWebhookNotificationParams object
// with the default values initialized.
func NewReceiveWebhookNotificationParams() *ReceiveWebhookNotificationParams {
	var ()
	return &ReceiveWebhookNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReceiveWebhookNotificationParamsWithTimeout creates a new ReceiveWebhookNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReceiveWebhookNotificationParamsWithTimeout(timeout time.Duration) *ReceiveWebhookNotificationParams {
	var ()
	return &ReceiveWebhookNotificationParams{

		timeout: timeout,
	}
}

// NewReceiveWebhookNotificationParamsWithContext creates a new ReceiveWebhookNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewReceiveWebhookNotificationParamsWithContext(ctx context.Context) *ReceiveWebhookNotificationParams {
	var ()
	return &ReceiveWebhookNotificationParams{

		Context: ctx,
	}
}

// NewReceiveWebhookNotificationParamsWithHTTPClient creates a new ReceiveWebhookNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReceiveWebhookNotificationParamsWithHTTPClient(client *http.Client) *ReceiveWebhookNotificationParams {
	var ()
	return &ReceiveWebhookNotificationParams{
		HTTPClient: client,
	}
}

/*ReceiveWebhookNotificationParams contains all the parameters to send to the API endpoint
for the receive webhook notification operation typically these are written to a http.Request
*/
type ReceiveWebhookNotificationParams struct {

	/*Body
	  The notification sent by webhook-client.

	*/
	Body *supportmessages.WebhookNotification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the receive webhook notification params
func (o *ReceiveWebhookNotificationParams) WithTimeout(timeout time.Duration) *ReceiveWebhookNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the receive webhook notification params
func (o *ReceiveWebhookNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the receive webhook notification params
func (o *ReceiveWebhookNotificationParams) WithContext(ctx context.Context) *ReceiveWebhookNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the receive webhook notification params
func (o *ReceiveWebhookNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the receive webhook notification params
func (o *ReceiveWebhookNotificationParams) WithHTTPClient(client *http.Client) *ReceiveWebhookNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the receive webhook notification params
func (o *ReceiveWebhookNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the receive webhook notification params
func (o *ReceiveWebhookNotificationParams) WithBody(body *supportmessages.WebhookNotification) *ReceiveWebhookNotificationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the receive webhook notification params
func (o *ReceiveWebhookNotificationParams) SetBody(body *supportmessages.WebhookNotification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ReceiveWebhookNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
