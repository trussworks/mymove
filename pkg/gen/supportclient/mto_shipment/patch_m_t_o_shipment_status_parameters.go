// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	supportmessages "github.com/transcom/mymove/pkg/gen/supportmessages"
)

// NewPatchMTOShipmentStatusParams creates a new PatchMTOShipmentStatusParams object
// with the default values initialized.
func NewPatchMTOShipmentStatusParams() *PatchMTOShipmentStatusParams {
	var ()
	return &PatchMTOShipmentStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchMTOShipmentStatusParamsWithTimeout creates a new PatchMTOShipmentStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchMTOShipmentStatusParamsWithTimeout(timeout time.Duration) *PatchMTOShipmentStatusParams {
	var ()
	return &PatchMTOShipmentStatusParams{

		timeout: timeout,
	}
}

// NewPatchMTOShipmentStatusParamsWithContext creates a new PatchMTOShipmentStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchMTOShipmentStatusParamsWithContext(ctx context.Context) *PatchMTOShipmentStatusParams {
	var ()
	return &PatchMTOShipmentStatusParams{

		Context: ctx,
	}
}

// NewPatchMTOShipmentStatusParamsWithHTTPClient creates a new PatchMTOShipmentStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchMTOShipmentStatusParamsWithHTTPClient(client *http.Client) *PatchMTOShipmentStatusParams {
	var ()
	return &PatchMTOShipmentStatusParams{
		HTTPClient: client,
	}
}

/*PatchMTOShipmentStatusParams contains all the parameters to send to the API endpoint
for the patch m t o shipment status operation typically these are written to a http.Request
*/
type PatchMTOShipmentStatusParams struct {

	/*IfMatch*/
	IfMatch string
	/*Body*/
	Body *supportmessages.PatchMTOShipmentStatus
	/*ShipmentID
	  ID of the shipment being updated

	*/
	ShipmentID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch m t o shipment status params
func (o *PatchMTOShipmentStatusParams) WithTimeout(timeout time.Duration) *PatchMTOShipmentStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch m t o shipment status params
func (o *PatchMTOShipmentStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch m t o shipment status params
func (o *PatchMTOShipmentStatusParams) WithContext(ctx context.Context) *PatchMTOShipmentStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch m t o shipment status params
func (o *PatchMTOShipmentStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch m t o shipment status params
func (o *PatchMTOShipmentStatusParams) WithHTTPClient(client *http.Client) *PatchMTOShipmentStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch m t o shipment status params
func (o *PatchMTOShipmentStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the patch m t o shipment status params
func (o *PatchMTOShipmentStatusParams) WithIfMatch(ifMatch string) *PatchMTOShipmentStatusParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the patch m t o shipment status params
func (o *PatchMTOShipmentStatusParams) SetIfMatch(ifMatch string) {
	o.IfMatch = ifMatch
}

// WithBody adds the body to the patch m t o shipment status params
func (o *PatchMTOShipmentStatusParams) WithBody(body *supportmessages.PatchMTOShipmentStatus) *PatchMTOShipmentStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch m t o shipment status params
func (o *PatchMTOShipmentStatusParams) SetBody(body *supportmessages.PatchMTOShipmentStatus) {
	o.Body = body
}

// WithShipmentID adds the shipmentID to the patch m t o shipment status params
func (o *PatchMTOShipmentStatusParams) WithShipmentID(shipmentID strfmt.UUID) *PatchMTOShipmentStatusParams {
	o.SetShipmentID(shipmentID)
	return o
}

// SetShipmentID adds the shipmentId to the patch m t o shipment status params
func (o *PatchMTOShipmentStatusParams) SetShipmentID(shipmentID strfmt.UUID) {
	o.ShipmentID = shipmentID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchMTOShipmentStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param If-Match
	if err := r.SetHeaderParam("If-Match", o.IfMatch); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param shipmentID
	if err := r.SetPathParam("shipmentID", o.ShipmentID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
