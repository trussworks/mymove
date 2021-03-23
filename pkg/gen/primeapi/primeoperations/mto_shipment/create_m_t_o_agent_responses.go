// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/primemessages"
)

// CreateMTOAgentOKCode is the HTTP code returned for type CreateMTOAgentOK
const CreateMTOAgentOKCode int = 200

/*CreateMTOAgentOK Successfully added the agent.

swagger:response createMTOAgentOK
*/
type CreateMTOAgentOK struct {

	/*
	  In: Body
	*/
	Payload *primemessages.MTOAgent `json:"body,omitempty"`
}

// NewCreateMTOAgentOK creates CreateMTOAgentOK with default headers values
func NewCreateMTOAgentOK() *CreateMTOAgentOK {

	return &CreateMTOAgentOK{}
}

// WithPayload adds the payload to the create m t o agent o k response
func (o *CreateMTOAgentOK) WithPayload(payload *primemessages.MTOAgent) *CreateMTOAgentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o agent o k response
func (o *CreateMTOAgentOK) SetPayload(payload *primemessages.MTOAgent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOAgentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMTOAgentBadRequestCode is the HTTP code returned for type CreateMTOAgentBadRequest
const CreateMTOAgentBadRequestCode int = 400

/*CreateMTOAgentBadRequest The request payload is invalid.

swagger:response createMTOAgentBadRequest
*/
type CreateMTOAgentBadRequest struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewCreateMTOAgentBadRequest creates CreateMTOAgentBadRequest with default headers values
func NewCreateMTOAgentBadRequest() *CreateMTOAgentBadRequest {

	return &CreateMTOAgentBadRequest{}
}

// WithPayload adds the payload to the create m t o agent bad request response
func (o *CreateMTOAgentBadRequest) WithPayload(payload *primemessages.ClientError) *CreateMTOAgentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o agent bad request response
func (o *CreateMTOAgentBadRequest) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOAgentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMTOAgentUnauthorizedCode is the HTTP code returned for type CreateMTOAgentUnauthorized
const CreateMTOAgentUnauthorizedCode int = 401

/*CreateMTOAgentUnauthorized The request was denied.

swagger:response createMTOAgentUnauthorized
*/
type CreateMTOAgentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewCreateMTOAgentUnauthorized creates CreateMTOAgentUnauthorized with default headers values
func NewCreateMTOAgentUnauthorized() *CreateMTOAgentUnauthorized {

	return &CreateMTOAgentUnauthorized{}
}

// WithPayload adds the payload to the create m t o agent unauthorized response
func (o *CreateMTOAgentUnauthorized) WithPayload(payload *primemessages.ClientError) *CreateMTOAgentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o agent unauthorized response
func (o *CreateMTOAgentUnauthorized) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOAgentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMTOAgentForbiddenCode is the HTTP code returned for type CreateMTOAgentForbidden
const CreateMTOAgentForbiddenCode int = 403

/*CreateMTOAgentForbidden The request was denied.

swagger:response createMTOAgentForbidden
*/
type CreateMTOAgentForbidden struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewCreateMTOAgentForbidden creates CreateMTOAgentForbidden with default headers values
func NewCreateMTOAgentForbidden() *CreateMTOAgentForbidden {

	return &CreateMTOAgentForbidden{}
}

// WithPayload adds the payload to the create m t o agent forbidden response
func (o *CreateMTOAgentForbidden) WithPayload(payload *primemessages.ClientError) *CreateMTOAgentForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o agent forbidden response
func (o *CreateMTOAgentForbidden) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOAgentForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMTOAgentNotFoundCode is the HTTP code returned for type CreateMTOAgentNotFound
const CreateMTOAgentNotFoundCode int = 404

/*CreateMTOAgentNotFound The requested resource wasn't found.

swagger:response createMTOAgentNotFound
*/
type CreateMTOAgentNotFound struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewCreateMTOAgentNotFound creates CreateMTOAgentNotFound with default headers values
func NewCreateMTOAgentNotFound() *CreateMTOAgentNotFound {

	return &CreateMTOAgentNotFound{}
}

// WithPayload adds the payload to the create m t o agent not found response
func (o *CreateMTOAgentNotFound) WithPayload(payload *primemessages.ClientError) *CreateMTOAgentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o agent not found response
func (o *CreateMTOAgentNotFound) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOAgentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMTOAgentUnprocessableEntityCode is the HTTP code returned for type CreateMTOAgentUnprocessableEntity
const CreateMTOAgentUnprocessableEntityCode int = 422

/*CreateMTOAgentUnprocessableEntity The payload was unprocessable.

swagger:response createMTOAgentUnprocessableEntity
*/
type CreateMTOAgentUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ValidationError `json:"body,omitempty"`
}

// NewCreateMTOAgentUnprocessableEntity creates CreateMTOAgentUnprocessableEntity with default headers values
func NewCreateMTOAgentUnprocessableEntity() *CreateMTOAgentUnprocessableEntity {

	return &CreateMTOAgentUnprocessableEntity{}
}

// WithPayload adds the payload to the create m t o agent unprocessable entity response
func (o *CreateMTOAgentUnprocessableEntity) WithPayload(payload *primemessages.ValidationError) *CreateMTOAgentUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o agent unprocessable entity response
func (o *CreateMTOAgentUnprocessableEntity) SetPayload(payload *primemessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOAgentUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMTOAgentInternalServerErrorCode is the HTTP code returned for type CreateMTOAgentInternalServerError
const CreateMTOAgentInternalServerErrorCode int = 500

/*CreateMTOAgentInternalServerError A server error occurred.

swagger:response createMTOAgentInternalServerError
*/
type CreateMTOAgentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *primemessages.Error `json:"body,omitempty"`
}

// NewCreateMTOAgentInternalServerError creates CreateMTOAgentInternalServerError with default headers values
func NewCreateMTOAgentInternalServerError() *CreateMTOAgentInternalServerError {

	return &CreateMTOAgentInternalServerError{}
}

// WithPayload adds the payload to the create m t o agent internal server error response
func (o *CreateMTOAgentInternalServerError) WithPayload(payload *primemessages.Error) *CreateMTOAgentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o agent internal server error response
func (o *CreateMTOAgentInternalServerError) SetPayload(payload *primemessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOAgentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
