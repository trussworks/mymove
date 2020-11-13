// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// DeleteMoveTaskOrderOKCode is the HTTP code returned for type DeleteMoveTaskOrderOK
const DeleteMoveTaskOrderOKCode int = 200

/*DeleteMoveTaskOrderOK Successfully deleted move task order

swagger:response deleteMoveTaskOrderOK
*/
type DeleteMoveTaskOrderOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.MoveTaskOrder `json:"body,omitempty"`
}

// NewDeleteMoveTaskOrderOK creates DeleteMoveTaskOrderOK with default headers values
func NewDeleteMoveTaskOrderOK() *DeleteMoveTaskOrderOK {

	return &DeleteMoveTaskOrderOK{}
}

// WithPayload adds the payload to the delete move task order o k response
func (o *DeleteMoveTaskOrderOK) WithPayload(payload *ghcmessages.MoveTaskOrder) *DeleteMoveTaskOrderOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete move task order o k response
func (o *DeleteMoveTaskOrderOK) SetPayload(payload *ghcmessages.MoveTaskOrder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMoveTaskOrderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteMoveTaskOrderBadRequestCode is the HTTP code returned for type DeleteMoveTaskOrderBadRequest
const DeleteMoveTaskOrderBadRequestCode int = 400

/*DeleteMoveTaskOrderBadRequest The request payload is invalid

swagger:response deleteMoveTaskOrderBadRequest
*/
type DeleteMoveTaskOrderBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteMoveTaskOrderBadRequest creates DeleteMoveTaskOrderBadRequest with default headers values
func NewDeleteMoveTaskOrderBadRequest() *DeleteMoveTaskOrderBadRequest {

	return &DeleteMoveTaskOrderBadRequest{}
}

// WithPayload adds the payload to the delete move task order bad request response
func (o *DeleteMoveTaskOrderBadRequest) WithPayload(payload interface{}) *DeleteMoveTaskOrderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete move task order bad request response
func (o *DeleteMoveTaskOrderBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMoveTaskOrderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteMoveTaskOrderUnauthorizedCode is the HTTP code returned for type DeleteMoveTaskOrderUnauthorized
const DeleteMoveTaskOrderUnauthorizedCode int = 401

/*DeleteMoveTaskOrderUnauthorized The request was denied

swagger:response deleteMoveTaskOrderUnauthorized
*/
type DeleteMoveTaskOrderUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteMoveTaskOrderUnauthorized creates DeleteMoveTaskOrderUnauthorized with default headers values
func NewDeleteMoveTaskOrderUnauthorized() *DeleteMoveTaskOrderUnauthorized {

	return &DeleteMoveTaskOrderUnauthorized{}
}

// WithPayload adds the payload to the delete move task order unauthorized response
func (o *DeleteMoveTaskOrderUnauthorized) WithPayload(payload interface{}) *DeleteMoveTaskOrderUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete move task order unauthorized response
func (o *DeleteMoveTaskOrderUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMoveTaskOrderUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteMoveTaskOrderForbiddenCode is the HTTP code returned for type DeleteMoveTaskOrderForbidden
const DeleteMoveTaskOrderForbiddenCode int = 403

/*DeleteMoveTaskOrderForbidden The request was denied

swagger:response deleteMoveTaskOrderForbidden
*/
type DeleteMoveTaskOrderForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteMoveTaskOrderForbidden creates DeleteMoveTaskOrderForbidden with default headers values
func NewDeleteMoveTaskOrderForbidden() *DeleteMoveTaskOrderForbidden {

	return &DeleteMoveTaskOrderForbidden{}
}

// WithPayload adds the payload to the delete move task order forbidden response
func (o *DeleteMoveTaskOrderForbidden) WithPayload(payload interface{}) *DeleteMoveTaskOrderForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete move task order forbidden response
func (o *DeleteMoveTaskOrderForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMoveTaskOrderForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteMoveTaskOrderNotFoundCode is the HTTP code returned for type DeleteMoveTaskOrderNotFound
const DeleteMoveTaskOrderNotFoundCode int = 404

/*DeleteMoveTaskOrderNotFound The requested resource wasn't found

swagger:response deleteMoveTaskOrderNotFound
*/
type DeleteMoveTaskOrderNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteMoveTaskOrderNotFound creates DeleteMoveTaskOrderNotFound with default headers values
func NewDeleteMoveTaskOrderNotFound() *DeleteMoveTaskOrderNotFound {

	return &DeleteMoveTaskOrderNotFound{}
}

// WithPayload adds the payload to the delete move task order not found response
func (o *DeleteMoveTaskOrderNotFound) WithPayload(payload interface{}) *DeleteMoveTaskOrderNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete move task order not found response
func (o *DeleteMoveTaskOrderNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMoveTaskOrderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteMoveTaskOrderInternalServerErrorCode is the HTTP code returned for type DeleteMoveTaskOrderInternalServerError
const DeleteMoveTaskOrderInternalServerErrorCode int = 500

/*DeleteMoveTaskOrderInternalServerError A server error occurred

swagger:response deleteMoveTaskOrderInternalServerError
*/
type DeleteMoveTaskOrderInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteMoveTaskOrderInternalServerError creates DeleteMoveTaskOrderInternalServerError with default headers values
func NewDeleteMoveTaskOrderInternalServerError() *DeleteMoveTaskOrderInternalServerError {

	return &DeleteMoveTaskOrderInternalServerError{}
}

// WithPayload adds the payload to the delete move task order internal server error response
func (o *DeleteMoveTaskOrderInternalServerError) WithPayload(payload interface{}) *DeleteMoveTaskOrderInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete move task order internal server error response
func (o *DeleteMoveTaskOrderInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMoveTaskOrderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
