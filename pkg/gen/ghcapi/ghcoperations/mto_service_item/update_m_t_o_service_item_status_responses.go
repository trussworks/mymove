// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// UpdateMTOServiceItemStatusOKCode is the HTTP code returned for type UpdateMTOServiceItemStatusOK
const UpdateMTOServiceItemStatusOKCode int = 200

/*UpdateMTOServiceItemStatusOK Successfully updated status for a line item for a move task order by ID

swagger:response updateMTOServiceItemStatusOK
*/
type UpdateMTOServiceItemStatusOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.MTOServiceItem `json:"body,omitempty"`
}

// NewUpdateMTOServiceItemStatusOK creates UpdateMTOServiceItemStatusOK with default headers values
func NewUpdateMTOServiceItemStatusOK() *UpdateMTOServiceItemStatusOK {

	return &UpdateMTOServiceItemStatusOK{}
}

// WithPayload adds the payload to the update m t o service item status o k response
func (o *UpdateMTOServiceItemStatusOK) WithPayload(payload *ghcmessages.MTOServiceItem) *UpdateMTOServiceItemStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o service item status o k response
func (o *UpdateMTOServiceItemStatusOK) SetPayload(payload *ghcmessages.MTOServiceItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOServiceItemStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOServiceItemStatusBadRequestCode is the HTTP code returned for type UpdateMTOServiceItemStatusBadRequest
const UpdateMTOServiceItemStatusBadRequestCode int = 400

/*UpdateMTOServiceItemStatusBadRequest The request payload is invalid

swagger:response updateMTOServiceItemStatusBadRequest
*/
type UpdateMTOServiceItemStatusBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOServiceItemStatusBadRequest creates UpdateMTOServiceItemStatusBadRequest with default headers values
func NewUpdateMTOServiceItemStatusBadRequest() *UpdateMTOServiceItemStatusBadRequest {

	return &UpdateMTOServiceItemStatusBadRequest{}
}

// WithPayload adds the payload to the update m t o service item status bad request response
func (o *UpdateMTOServiceItemStatusBadRequest) WithPayload(payload interface{}) *UpdateMTOServiceItemStatusBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o service item status bad request response
func (o *UpdateMTOServiceItemStatusBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOServiceItemStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOServiceItemStatusUnauthorizedCode is the HTTP code returned for type UpdateMTOServiceItemStatusUnauthorized
const UpdateMTOServiceItemStatusUnauthorizedCode int = 401

/*UpdateMTOServiceItemStatusUnauthorized The request was denied

swagger:response updateMTOServiceItemStatusUnauthorized
*/
type UpdateMTOServiceItemStatusUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOServiceItemStatusUnauthorized creates UpdateMTOServiceItemStatusUnauthorized with default headers values
func NewUpdateMTOServiceItemStatusUnauthorized() *UpdateMTOServiceItemStatusUnauthorized {

	return &UpdateMTOServiceItemStatusUnauthorized{}
}

// WithPayload adds the payload to the update m t o service item status unauthorized response
func (o *UpdateMTOServiceItemStatusUnauthorized) WithPayload(payload interface{}) *UpdateMTOServiceItemStatusUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o service item status unauthorized response
func (o *UpdateMTOServiceItemStatusUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOServiceItemStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOServiceItemStatusForbiddenCode is the HTTP code returned for type UpdateMTOServiceItemStatusForbidden
const UpdateMTOServiceItemStatusForbiddenCode int = 403

/*UpdateMTOServiceItemStatusForbidden The request was denied

swagger:response updateMTOServiceItemStatusForbidden
*/
type UpdateMTOServiceItemStatusForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOServiceItemStatusForbidden creates UpdateMTOServiceItemStatusForbidden with default headers values
func NewUpdateMTOServiceItemStatusForbidden() *UpdateMTOServiceItemStatusForbidden {

	return &UpdateMTOServiceItemStatusForbidden{}
}

// WithPayload adds the payload to the update m t o service item status forbidden response
func (o *UpdateMTOServiceItemStatusForbidden) WithPayload(payload interface{}) *UpdateMTOServiceItemStatusForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o service item status forbidden response
func (o *UpdateMTOServiceItemStatusForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOServiceItemStatusForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOServiceItemStatusNotFoundCode is the HTTP code returned for type UpdateMTOServiceItemStatusNotFound
const UpdateMTOServiceItemStatusNotFoundCode int = 404

/*UpdateMTOServiceItemStatusNotFound The requested resource wasn't found

swagger:response updateMTOServiceItemStatusNotFound
*/
type UpdateMTOServiceItemStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOServiceItemStatusNotFound creates UpdateMTOServiceItemStatusNotFound with default headers values
func NewUpdateMTOServiceItemStatusNotFound() *UpdateMTOServiceItemStatusNotFound {

	return &UpdateMTOServiceItemStatusNotFound{}
}

// WithPayload adds the payload to the update m t o service item status not found response
func (o *UpdateMTOServiceItemStatusNotFound) WithPayload(payload interface{}) *UpdateMTOServiceItemStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o service item status not found response
func (o *UpdateMTOServiceItemStatusNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOServiceItemStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOServiceItemStatusPreconditionFailedCode is the HTTP code returned for type UpdateMTOServiceItemStatusPreconditionFailed
const UpdateMTOServiceItemStatusPreconditionFailedCode int = 412

/*UpdateMTOServiceItemStatusPreconditionFailed Precondition Failed

swagger:response updateMTOServiceItemStatusPreconditionFailed
*/
type UpdateMTOServiceItemStatusPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOServiceItemStatusPreconditionFailed creates UpdateMTOServiceItemStatusPreconditionFailed with default headers values
func NewUpdateMTOServiceItemStatusPreconditionFailed() *UpdateMTOServiceItemStatusPreconditionFailed {

	return &UpdateMTOServiceItemStatusPreconditionFailed{}
}

// WithPayload adds the payload to the update m t o service item status precondition failed response
func (o *UpdateMTOServiceItemStatusPreconditionFailed) WithPayload(payload interface{}) *UpdateMTOServiceItemStatusPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o service item status precondition failed response
func (o *UpdateMTOServiceItemStatusPreconditionFailed) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOServiceItemStatusPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOServiceItemStatusUnprocessableEntityCode is the HTTP code returned for type UpdateMTOServiceItemStatusUnprocessableEntity
const UpdateMTOServiceItemStatusUnprocessableEntityCode int = 422

/*UpdateMTOServiceItemStatusUnprocessableEntity Validation error

swagger:response updateMTOServiceItemStatusUnprocessableEntity
*/
type UpdateMTOServiceItemStatusUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.ValidationError `json:"body,omitempty"`
}

// NewUpdateMTOServiceItemStatusUnprocessableEntity creates UpdateMTOServiceItemStatusUnprocessableEntity with default headers values
func NewUpdateMTOServiceItemStatusUnprocessableEntity() *UpdateMTOServiceItemStatusUnprocessableEntity {

	return &UpdateMTOServiceItemStatusUnprocessableEntity{}
}

// WithPayload adds the payload to the update m t o service item status unprocessable entity response
func (o *UpdateMTOServiceItemStatusUnprocessableEntity) WithPayload(payload *ghcmessages.ValidationError) *UpdateMTOServiceItemStatusUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o service item status unprocessable entity response
func (o *UpdateMTOServiceItemStatusUnprocessableEntity) SetPayload(payload *ghcmessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOServiceItemStatusUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOServiceItemStatusInternalServerErrorCode is the HTTP code returned for type UpdateMTOServiceItemStatusInternalServerError
const UpdateMTOServiceItemStatusInternalServerErrorCode int = 500

/*UpdateMTOServiceItemStatusInternalServerError A server error occurred

swagger:response updateMTOServiceItemStatusInternalServerError
*/
type UpdateMTOServiceItemStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOServiceItemStatusInternalServerError creates UpdateMTOServiceItemStatusInternalServerError with default headers values
func NewUpdateMTOServiceItemStatusInternalServerError() *UpdateMTOServiceItemStatusInternalServerError {

	return &UpdateMTOServiceItemStatusInternalServerError{}
}

// WithPayload adds the payload to the update m t o service item status internal server error response
func (o *UpdateMTOServiceItemStatusInternalServerError) WithPayload(payload interface{}) *UpdateMTOServiceItemStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o service item status internal server error response
func (o *UpdateMTOServiceItemStatusInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOServiceItemStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
