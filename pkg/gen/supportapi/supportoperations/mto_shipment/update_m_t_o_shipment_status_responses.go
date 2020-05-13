// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	supportmessages "github.com/transcom/mymove/pkg/gen/supportmessages"
)

// UpdateMTOShipmentStatusOKCode is the HTTP code returned for type UpdateMTOShipmentStatusOK
const UpdateMTOShipmentStatusOKCode int = 200

/*UpdateMTOShipmentStatusOK Successfully updated the shipment's status.

swagger:response updateMTOShipmentStatusOK
*/
type UpdateMTOShipmentStatusOK struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.MTOShipment `json:"body,omitempty"`
}

// NewUpdateMTOShipmentStatusOK creates UpdateMTOShipmentStatusOK with default headers values
func NewUpdateMTOShipmentStatusOK() *UpdateMTOShipmentStatusOK {

	return &UpdateMTOShipmentStatusOK{}
}

// WithPayload adds the payload to the update m t o shipment status o k response
func (o *UpdateMTOShipmentStatusOK) WithPayload(payload *supportmessages.MTOShipment) *UpdateMTOShipmentStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment status o k response
func (o *UpdateMTOShipmentStatusOK) SetPayload(payload *supportmessages.MTOShipment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentStatusBadRequestCode is the HTTP code returned for type UpdateMTOShipmentStatusBadRequest
const UpdateMTOShipmentStatusBadRequestCode int = 400

/*UpdateMTOShipmentStatusBadRequest The parameters were invalid.

swagger:response updateMTOShipmentStatusBadRequest
*/
type UpdateMTOShipmentStatusBadRequest struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewUpdateMTOShipmentStatusBadRequest creates UpdateMTOShipmentStatusBadRequest with default headers values
func NewUpdateMTOShipmentStatusBadRequest() *UpdateMTOShipmentStatusBadRequest {

	return &UpdateMTOShipmentStatusBadRequest{}
}

// WithPayload adds the payload to the update m t o shipment status bad request response
func (o *UpdateMTOShipmentStatusBadRequest) WithPayload(payload *supportmessages.ClientError) *UpdateMTOShipmentStatusBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment status bad request response
func (o *UpdateMTOShipmentStatusBadRequest) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentStatusUnauthorizedCode is the HTTP code returned for type UpdateMTOShipmentStatusUnauthorized
const UpdateMTOShipmentStatusUnauthorizedCode int = 401

/*UpdateMTOShipmentStatusUnauthorized The request was unauthorized.

swagger:response updateMTOShipmentStatusUnauthorized
*/
type UpdateMTOShipmentStatusUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOShipmentStatusUnauthorized creates UpdateMTOShipmentStatusUnauthorized with default headers values
func NewUpdateMTOShipmentStatusUnauthorized() *UpdateMTOShipmentStatusUnauthorized {

	return &UpdateMTOShipmentStatusUnauthorized{}
}

// WithPayload adds the payload to the update m t o shipment status unauthorized response
func (o *UpdateMTOShipmentStatusUnauthorized) WithPayload(payload interface{}) *UpdateMTOShipmentStatusUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment status unauthorized response
func (o *UpdateMTOShipmentStatusUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOShipmentStatusForbiddenCode is the HTTP code returned for type UpdateMTOShipmentStatusForbidden
const UpdateMTOShipmentStatusForbiddenCode int = 403

/*UpdateMTOShipmentStatusForbidden The client doesn't have permissions to perform the request.

swagger:response updateMTOShipmentStatusForbidden
*/
type UpdateMTOShipmentStatusForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOShipmentStatusForbidden creates UpdateMTOShipmentStatusForbidden with default headers values
func NewUpdateMTOShipmentStatusForbidden() *UpdateMTOShipmentStatusForbidden {

	return &UpdateMTOShipmentStatusForbidden{}
}

// WithPayload adds the payload to the update m t o shipment status forbidden response
func (o *UpdateMTOShipmentStatusForbidden) WithPayload(payload interface{}) *UpdateMTOShipmentStatusForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment status forbidden response
func (o *UpdateMTOShipmentStatusForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentStatusForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOShipmentStatusNotFoundCode is the HTTP code returned for type UpdateMTOShipmentStatusNotFound
const UpdateMTOShipmentStatusNotFoundCode int = 404

/*UpdateMTOShipmentStatusNotFound The requested resource wasn't found.

swagger:response updateMTOShipmentStatusNotFound
*/
type UpdateMTOShipmentStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewUpdateMTOShipmentStatusNotFound creates UpdateMTOShipmentStatusNotFound with default headers values
func NewUpdateMTOShipmentStatusNotFound() *UpdateMTOShipmentStatusNotFound {

	return &UpdateMTOShipmentStatusNotFound{}
}

// WithPayload adds the payload to the update m t o shipment status not found response
func (o *UpdateMTOShipmentStatusNotFound) WithPayload(payload *supportmessages.ClientError) *UpdateMTOShipmentStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment status not found response
func (o *UpdateMTOShipmentStatusNotFound) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentStatusConflictCode is the HTTP code returned for type UpdateMTOShipmentStatusConflict
const UpdateMTOShipmentStatusConflictCode int = 409

/*UpdateMTOShipmentStatusConflict Conflict error due to trying to change the status of shipment that is not currently "SUBMITTED".

swagger:response updateMTOShipmentStatusConflict
*/
type UpdateMTOShipmentStatusConflict struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOShipmentStatusConflict creates UpdateMTOShipmentStatusConflict with default headers values
func NewUpdateMTOShipmentStatusConflict() *UpdateMTOShipmentStatusConflict {

	return &UpdateMTOShipmentStatusConflict{}
}

// WithPayload adds the payload to the update m t o shipment status conflict response
func (o *UpdateMTOShipmentStatusConflict) WithPayload(payload interface{}) *UpdateMTOShipmentStatusConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment status conflict response
func (o *UpdateMTOShipmentStatusConflict) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentStatusConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOShipmentStatusPreconditionFailedCode is the HTTP code returned for type UpdateMTOShipmentStatusPreconditionFailed
const UpdateMTOShipmentStatusPreconditionFailedCode int = 412

/*UpdateMTOShipmentStatusPreconditionFailed Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.

swagger:response updateMTOShipmentStatusPreconditionFailed
*/
type UpdateMTOShipmentStatusPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewUpdateMTOShipmentStatusPreconditionFailed creates UpdateMTOShipmentStatusPreconditionFailed with default headers values
func NewUpdateMTOShipmentStatusPreconditionFailed() *UpdateMTOShipmentStatusPreconditionFailed {

	return &UpdateMTOShipmentStatusPreconditionFailed{}
}

// WithPayload adds the payload to the update m t o shipment status precondition failed response
func (o *UpdateMTOShipmentStatusPreconditionFailed) WithPayload(payload *supportmessages.ClientError) *UpdateMTOShipmentStatusPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment status precondition failed response
func (o *UpdateMTOShipmentStatusPreconditionFailed) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentStatusPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentStatusUnprocessableEntityCode is the HTTP code returned for type UpdateMTOShipmentStatusUnprocessableEntity
const UpdateMTOShipmentStatusUnprocessableEntityCode int = 422

/*UpdateMTOShipmentStatusUnprocessableEntity The payload was unprocessable.

swagger:response updateMTOShipmentStatusUnprocessableEntity
*/
type UpdateMTOShipmentStatusUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ValidationError `json:"body,omitempty"`
}

// NewUpdateMTOShipmentStatusUnprocessableEntity creates UpdateMTOShipmentStatusUnprocessableEntity with default headers values
func NewUpdateMTOShipmentStatusUnprocessableEntity() *UpdateMTOShipmentStatusUnprocessableEntity {

	return &UpdateMTOShipmentStatusUnprocessableEntity{}
}

// WithPayload adds the payload to the update m t o shipment status unprocessable entity response
func (o *UpdateMTOShipmentStatusUnprocessableEntity) WithPayload(payload *supportmessages.ValidationError) *UpdateMTOShipmentStatusUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment status unprocessable entity response
func (o *UpdateMTOShipmentStatusUnprocessableEntity) SetPayload(payload *supportmessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentStatusUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentStatusInternalServerErrorCode is the HTTP code returned for type UpdateMTOShipmentStatusInternalServerError
const UpdateMTOShipmentStatusInternalServerErrorCode int = 500

/*UpdateMTOShipmentStatusInternalServerError A server error occurred.

swagger:response updateMTOShipmentStatusInternalServerError
*/
type UpdateMTOShipmentStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.Error `json:"body,omitempty"`
}

// NewUpdateMTOShipmentStatusInternalServerError creates UpdateMTOShipmentStatusInternalServerError with default headers values
func NewUpdateMTOShipmentStatusInternalServerError() *UpdateMTOShipmentStatusInternalServerError {

	return &UpdateMTOShipmentStatusInternalServerError{}
}

// WithPayload adds the payload to the update m t o shipment status internal server error response
func (o *UpdateMTOShipmentStatusInternalServerError) WithPayload(payload *supportmessages.Error) *UpdateMTOShipmentStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment status internal server error response
func (o *UpdateMTOShipmentStatusInternalServerError) SetPayload(payload *supportmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
