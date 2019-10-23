// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// UpdateMoveTaskOrderActualWeightOKCode is the HTTP code returned for type UpdateMoveTaskOrderActualWeightOK
const UpdateMoveTaskOrderActualWeightOKCode int = 200

/*UpdateMoveTaskOrderActualWeightOK Successfully retrieved move task order

swagger:response updateMoveTaskOrderActualWeightOK
*/
type UpdateMoveTaskOrderActualWeightOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.MoveTaskOrder `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderActualWeightOK creates UpdateMoveTaskOrderActualWeightOK with default headers values
func NewUpdateMoveTaskOrderActualWeightOK() *UpdateMoveTaskOrderActualWeightOK {

	return &UpdateMoveTaskOrderActualWeightOK{}
}

// WithPayload adds the payload to the update move task order actual weight o k response
func (o *UpdateMoveTaskOrderActualWeightOK) WithPayload(payload *ghcmessages.MoveTaskOrder) *UpdateMoveTaskOrderActualWeightOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order actual weight o k response
func (o *UpdateMoveTaskOrderActualWeightOK) SetPayload(payload *ghcmessages.MoveTaskOrder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderActualWeightOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMoveTaskOrderActualWeightBadRequestCode is the HTTP code returned for type UpdateMoveTaskOrderActualWeightBadRequest
const UpdateMoveTaskOrderActualWeightBadRequestCode int = 400

/*UpdateMoveTaskOrderActualWeightBadRequest The request payload is invalid

swagger:response updateMoveTaskOrderActualWeightBadRequest
*/
type UpdateMoveTaskOrderActualWeightBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderActualWeightBadRequest creates UpdateMoveTaskOrderActualWeightBadRequest with default headers values
func NewUpdateMoveTaskOrderActualWeightBadRequest() *UpdateMoveTaskOrderActualWeightBadRequest {

	return &UpdateMoveTaskOrderActualWeightBadRequest{}
}

// WithPayload adds the payload to the update move task order actual weight bad request response
func (o *UpdateMoveTaskOrderActualWeightBadRequest) WithPayload(payload interface{}) *UpdateMoveTaskOrderActualWeightBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order actual weight bad request response
func (o *UpdateMoveTaskOrderActualWeightBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderActualWeightBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderActualWeightUnauthorizedCode is the HTTP code returned for type UpdateMoveTaskOrderActualWeightUnauthorized
const UpdateMoveTaskOrderActualWeightUnauthorizedCode int = 401

/*UpdateMoveTaskOrderActualWeightUnauthorized The request was denied

swagger:response updateMoveTaskOrderActualWeightUnauthorized
*/
type UpdateMoveTaskOrderActualWeightUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderActualWeightUnauthorized creates UpdateMoveTaskOrderActualWeightUnauthorized with default headers values
func NewUpdateMoveTaskOrderActualWeightUnauthorized() *UpdateMoveTaskOrderActualWeightUnauthorized {

	return &UpdateMoveTaskOrderActualWeightUnauthorized{}
}

// WithPayload adds the payload to the update move task order actual weight unauthorized response
func (o *UpdateMoveTaskOrderActualWeightUnauthorized) WithPayload(payload interface{}) *UpdateMoveTaskOrderActualWeightUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order actual weight unauthorized response
func (o *UpdateMoveTaskOrderActualWeightUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderActualWeightUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderActualWeightForbiddenCode is the HTTP code returned for type UpdateMoveTaskOrderActualWeightForbidden
const UpdateMoveTaskOrderActualWeightForbiddenCode int = 403

/*UpdateMoveTaskOrderActualWeightForbidden The request was denied

swagger:response updateMoveTaskOrderActualWeightForbidden
*/
type UpdateMoveTaskOrderActualWeightForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderActualWeightForbidden creates UpdateMoveTaskOrderActualWeightForbidden with default headers values
func NewUpdateMoveTaskOrderActualWeightForbidden() *UpdateMoveTaskOrderActualWeightForbidden {

	return &UpdateMoveTaskOrderActualWeightForbidden{}
}

// WithPayload adds the payload to the update move task order actual weight forbidden response
func (o *UpdateMoveTaskOrderActualWeightForbidden) WithPayload(payload interface{}) *UpdateMoveTaskOrderActualWeightForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order actual weight forbidden response
func (o *UpdateMoveTaskOrderActualWeightForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderActualWeightForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderActualWeightNotFoundCode is the HTTP code returned for type UpdateMoveTaskOrderActualWeightNotFound
const UpdateMoveTaskOrderActualWeightNotFoundCode int = 404

/*UpdateMoveTaskOrderActualWeightNotFound The requested resource wasn't found

swagger:response updateMoveTaskOrderActualWeightNotFound
*/
type UpdateMoveTaskOrderActualWeightNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderActualWeightNotFound creates UpdateMoveTaskOrderActualWeightNotFound with default headers values
func NewUpdateMoveTaskOrderActualWeightNotFound() *UpdateMoveTaskOrderActualWeightNotFound {

	return &UpdateMoveTaskOrderActualWeightNotFound{}
}

// WithPayload adds the payload to the update move task order actual weight not found response
func (o *UpdateMoveTaskOrderActualWeightNotFound) WithPayload(payload interface{}) *UpdateMoveTaskOrderActualWeightNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order actual weight not found response
func (o *UpdateMoveTaskOrderActualWeightNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderActualWeightNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderActualWeightInternalServerErrorCode is the HTTP code returned for type UpdateMoveTaskOrderActualWeightInternalServerError
const UpdateMoveTaskOrderActualWeightInternalServerErrorCode int = 500

/*UpdateMoveTaskOrderActualWeightInternalServerError A server error occurred

swagger:response updateMoveTaskOrderActualWeightInternalServerError
*/
type UpdateMoveTaskOrderActualWeightInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderActualWeightInternalServerError creates UpdateMoveTaskOrderActualWeightInternalServerError with default headers values
func NewUpdateMoveTaskOrderActualWeightInternalServerError() *UpdateMoveTaskOrderActualWeightInternalServerError {

	return &UpdateMoveTaskOrderActualWeightInternalServerError{}
}

// WithPayload adds the payload to the update move task order actual weight internal server error response
func (o *UpdateMoveTaskOrderActualWeightInternalServerError) WithPayload(payload interface{}) *UpdateMoveTaskOrderActualWeightInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order actual weight internal server error response
func (o *UpdateMoveTaskOrderActualWeightInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderActualWeightInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
