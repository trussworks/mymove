// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// UpdateMTOShipmentAddressOKCode is the HTTP code returned for type UpdateMTOShipmentAddressOK
const UpdateMTOShipmentAddressOKCode int = 200

/*UpdateMTOShipmentAddressOK Successfully updated the address.

swagger:response updateMTOShipmentAddressOK
*/
type UpdateMTOShipmentAddressOK struct {

	/*
	  In: Body
	*/
	Payload *primemessages.Address `json:"body,omitempty"`
}

// NewUpdateMTOShipmentAddressOK creates UpdateMTOShipmentAddressOK with default headers values
func NewUpdateMTOShipmentAddressOK() *UpdateMTOShipmentAddressOK {

	return &UpdateMTOShipmentAddressOK{}
}

// WithPayload adds the payload to the update m t o shipment address o k response
func (o *UpdateMTOShipmentAddressOK) WithPayload(payload *primemessages.Address) *UpdateMTOShipmentAddressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment address o k response
func (o *UpdateMTOShipmentAddressOK) SetPayload(payload *primemessages.Address) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentAddressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentAddressBadRequestCode is the HTTP code returned for type UpdateMTOShipmentAddressBadRequest
const UpdateMTOShipmentAddressBadRequestCode int = 400

/*UpdateMTOShipmentAddressBadRequest The request payload is invalid.

swagger:response updateMTOShipmentAddressBadRequest
*/
type UpdateMTOShipmentAddressBadRequest struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewUpdateMTOShipmentAddressBadRequest creates UpdateMTOShipmentAddressBadRequest with default headers values
func NewUpdateMTOShipmentAddressBadRequest() *UpdateMTOShipmentAddressBadRequest {

	return &UpdateMTOShipmentAddressBadRequest{}
}

// WithPayload adds the payload to the update m t o shipment address bad request response
func (o *UpdateMTOShipmentAddressBadRequest) WithPayload(payload *primemessages.ClientError) *UpdateMTOShipmentAddressBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment address bad request response
func (o *UpdateMTOShipmentAddressBadRequest) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentAddressBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentAddressUnauthorizedCode is the HTTP code returned for type UpdateMTOShipmentAddressUnauthorized
const UpdateMTOShipmentAddressUnauthorizedCode int = 401

/*UpdateMTOShipmentAddressUnauthorized The request was denied.

swagger:response updateMTOShipmentAddressUnauthorized
*/
type UpdateMTOShipmentAddressUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewUpdateMTOShipmentAddressUnauthorized creates UpdateMTOShipmentAddressUnauthorized with default headers values
func NewUpdateMTOShipmentAddressUnauthorized() *UpdateMTOShipmentAddressUnauthorized {

	return &UpdateMTOShipmentAddressUnauthorized{}
}

// WithPayload adds the payload to the update m t o shipment address unauthorized response
func (o *UpdateMTOShipmentAddressUnauthorized) WithPayload(payload *primemessages.ClientError) *UpdateMTOShipmentAddressUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment address unauthorized response
func (o *UpdateMTOShipmentAddressUnauthorized) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentAddressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentAddressForbiddenCode is the HTTP code returned for type UpdateMTOShipmentAddressForbidden
const UpdateMTOShipmentAddressForbiddenCode int = 403

/*UpdateMTOShipmentAddressForbidden The request was denied.

swagger:response updateMTOShipmentAddressForbidden
*/
type UpdateMTOShipmentAddressForbidden struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewUpdateMTOShipmentAddressForbidden creates UpdateMTOShipmentAddressForbidden with default headers values
func NewUpdateMTOShipmentAddressForbidden() *UpdateMTOShipmentAddressForbidden {

	return &UpdateMTOShipmentAddressForbidden{}
}

// WithPayload adds the payload to the update m t o shipment address forbidden response
func (o *UpdateMTOShipmentAddressForbidden) WithPayload(payload *primemessages.ClientError) *UpdateMTOShipmentAddressForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment address forbidden response
func (o *UpdateMTOShipmentAddressForbidden) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentAddressForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentAddressNotFoundCode is the HTTP code returned for type UpdateMTOShipmentAddressNotFound
const UpdateMTOShipmentAddressNotFoundCode int = 404

/*UpdateMTOShipmentAddressNotFound The requested resource wasn't found.

swagger:response updateMTOShipmentAddressNotFound
*/
type UpdateMTOShipmentAddressNotFound struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewUpdateMTOShipmentAddressNotFound creates UpdateMTOShipmentAddressNotFound with default headers values
func NewUpdateMTOShipmentAddressNotFound() *UpdateMTOShipmentAddressNotFound {

	return &UpdateMTOShipmentAddressNotFound{}
}

// WithPayload adds the payload to the update m t o shipment address not found response
func (o *UpdateMTOShipmentAddressNotFound) WithPayload(payload *primemessages.ClientError) *UpdateMTOShipmentAddressNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment address not found response
func (o *UpdateMTOShipmentAddressNotFound) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentAddressNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentAddressPreconditionFailedCode is the HTTP code returned for type UpdateMTOShipmentAddressPreconditionFailed
const UpdateMTOShipmentAddressPreconditionFailedCode int = 412

/*UpdateMTOShipmentAddressPreconditionFailed Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.

swagger:response updateMTOShipmentAddressPreconditionFailed
*/
type UpdateMTOShipmentAddressPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewUpdateMTOShipmentAddressPreconditionFailed creates UpdateMTOShipmentAddressPreconditionFailed with default headers values
func NewUpdateMTOShipmentAddressPreconditionFailed() *UpdateMTOShipmentAddressPreconditionFailed {

	return &UpdateMTOShipmentAddressPreconditionFailed{}
}

// WithPayload adds the payload to the update m t o shipment address precondition failed response
func (o *UpdateMTOShipmentAddressPreconditionFailed) WithPayload(payload *primemessages.ClientError) *UpdateMTOShipmentAddressPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment address precondition failed response
func (o *UpdateMTOShipmentAddressPreconditionFailed) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentAddressPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentAddressUnprocessableEntityCode is the HTTP code returned for type UpdateMTOShipmentAddressUnprocessableEntity
const UpdateMTOShipmentAddressUnprocessableEntityCode int = 422

/*UpdateMTOShipmentAddressUnprocessableEntity The payload was unprocessable.

swagger:response updateMTOShipmentAddressUnprocessableEntity
*/
type UpdateMTOShipmentAddressUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ValidationError `json:"body,omitempty"`
}

// NewUpdateMTOShipmentAddressUnprocessableEntity creates UpdateMTOShipmentAddressUnprocessableEntity with default headers values
func NewUpdateMTOShipmentAddressUnprocessableEntity() *UpdateMTOShipmentAddressUnprocessableEntity {

	return &UpdateMTOShipmentAddressUnprocessableEntity{}
}

// WithPayload adds the payload to the update m t o shipment address unprocessable entity response
func (o *UpdateMTOShipmentAddressUnprocessableEntity) WithPayload(payload *primemessages.ValidationError) *UpdateMTOShipmentAddressUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment address unprocessable entity response
func (o *UpdateMTOShipmentAddressUnprocessableEntity) SetPayload(payload *primemessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentAddressUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentAddressInternalServerErrorCode is the HTTP code returned for type UpdateMTOShipmentAddressInternalServerError
const UpdateMTOShipmentAddressInternalServerErrorCode int = 500

/*UpdateMTOShipmentAddressInternalServerError A server error occurred.

swagger:response updateMTOShipmentAddressInternalServerError
*/
type UpdateMTOShipmentAddressInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *primemessages.Error `json:"body,omitempty"`
}

// NewUpdateMTOShipmentAddressInternalServerError creates UpdateMTOShipmentAddressInternalServerError with default headers values
func NewUpdateMTOShipmentAddressInternalServerError() *UpdateMTOShipmentAddressInternalServerError {

	return &UpdateMTOShipmentAddressInternalServerError{}
}

// WithPayload adds the payload to the update m t o shipment address internal server error response
func (o *UpdateMTOShipmentAddressInternalServerError) WithPayload(payload *primemessages.Error) *UpdateMTOShipmentAddressInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment address internal server error response
func (o *UpdateMTOShipmentAddressInternalServerError) SetPayload(payload *primemessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentAddressInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
