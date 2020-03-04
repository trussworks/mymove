// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// UpdateMTOPostCounselingInformationOKCode is the HTTP code returned for type UpdateMTOPostCounselingInformationOK
const UpdateMTOPostCounselingInformationOKCode int = 200

/*UpdateMTOPostCounselingInformationOK Successfully updated move task order post counseling information

swagger:response updateMTOPostCounselingInformationOK
*/
type UpdateMTOPostCounselingInformationOK struct {

	/*
	  In: Body
	*/
	Payload *primemessages.MoveTaskOrderWithEtag `json:"body,omitempty"`
}

// NewUpdateMTOPostCounselingInformationOK creates UpdateMTOPostCounselingInformationOK with default headers values
func NewUpdateMTOPostCounselingInformationOK() *UpdateMTOPostCounselingInformationOK {

	return &UpdateMTOPostCounselingInformationOK{}
}

// WithPayload adds the payload to the update m t o post counseling information o k response
func (o *UpdateMTOPostCounselingInformationOK) WithPayload(payload *primemessages.MoveTaskOrderWithEtag) *UpdateMTOPostCounselingInformationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o post counseling information o k response
func (o *UpdateMTOPostCounselingInformationOK) SetPayload(payload *primemessages.MoveTaskOrderWithEtag) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOPostCounselingInformationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOPostCounselingInformationUnauthorizedCode is the HTTP code returned for type UpdateMTOPostCounselingInformationUnauthorized
const UpdateMTOPostCounselingInformationUnauthorizedCode int = 401

/*UpdateMTOPostCounselingInformationUnauthorized The request was denied

swagger:response updateMTOPostCounselingInformationUnauthorized
*/
type UpdateMTOPostCounselingInformationUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOPostCounselingInformationUnauthorized creates UpdateMTOPostCounselingInformationUnauthorized with default headers values
func NewUpdateMTOPostCounselingInformationUnauthorized() *UpdateMTOPostCounselingInformationUnauthorized {

	return &UpdateMTOPostCounselingInformationUnauthorized{}
}

// WithPayload adds the payload to the update m t o post counseling information unauthorized response
func (o *UpdateMTOPostCounselingInformationUnauthorized) WithPayload(payload interface{}) *UpdateMTOPostCounselingInformationUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o post counseling information unauthorized response
func (o *UpdateMTOPostCounselingInformationUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOPostCounselingInformationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOPostCounselingInformationForbiddenCode is the HTTP code returned for type UpdateMTOPostCounselingInformationForbidden
const UpdateMTOPostCounselingInformationForbiddenCode int = 403

/*UpdateMTOPostCounselingInformationForbidden The request was denied

swagger:response updateMTOPostCounselingInformationForbidden
*/
type UpdateMTOPostCounselingInformationForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOPostCounselingInformationForbidden creates UpdateMTOPostCounselingInformationForbidden with default headers values
func NewUpdateMTOPostCounselingInformationForbidden() *UpdateMTOPostCounselingInformationForbidden {

	return &UpdateMTOPostCounselingInformationForbidden{}
}

// WithPayload adds the payload to the update m t o post counseling information forbidden response
func (o *UpdateMTOPostCounselingInformationForbidden) WithPayload(payload interface{}) *UpdateMTOPostCounselingInformationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o post counseling information forbidden response
func (o *UpdateMTOPostCounselingInformationForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOPostCounselingInformationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOPostCounselingInformationNotFoundCode is the HTTP code returned for type UpdateMTOPostCounselingInformationNotFound
const UpdateMTOPostCounselingInformationNotFoundCode int = 404

/*UpdateMTOPostCounselingInformationNotFound The requested resource wasn't found

swagger:response updateMTOPostCounselingInformationNotFound
*/
type UpdateMTOPostCounselingInformationNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOPostCounselingInformationNotFound creates UpdateMTOPostCounselingInformationNotFound with default headers values
func NewUpdateMTOPostCounselingInformationNotFound() *UpdateMTOPostCounselingInformationNotFound {

	return &UpdateMTOPostCounselingInformationNotFound{}
}

// WithPayload adds the payload to the update m t o post counseling information not found response
func (o *UpdateMTOPostCounselingInformationNotFound) WithPayload(payload interface{}) *UpdateMTOPostCounselingInformationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o post counseling information not found response
func (o *UpdateMTOPostCounselingInformationNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOPostCounselingInformationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOPostCounselingInformationPreconditionFailedCode is the HTTP code returned for type UpdateMTOPostCounselingInformationPreconditionFailed
const UpdateMTOPostCounselingInformationPreconditionFailedCode int = 412

/*UpdateMTOPostCounselingInformationPreconditionFailed precondition failed

swagger:response updateMTOPostCounselingInformationPreconditionFailed
*/
type UpdateMTOPostCounselingInformationPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOPostCounselingInformationPreconditionFailed creates UpdateMTOPostCounselingInformationPreconditionFailed with default headers values
func NewUpdateMTOPostCounselingInformationPreconditionFailed() *UpdateMTOPostCounselingInformationPreconditionFailed {

	return &UpdateMTOPostCounselingInformationPreconditionFailed{}
}

// WithPayload adds the payload to the update m t o post counseling information precondition failed response
func (o *UpdateMTOPostCounselingInformationPreconditionFailed) WithPayload(payload interface{}) *UpdateMTOPostCounselingInformationPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o post counseling information precondition failed response
func (o *UpdateMTOPostCounselingInformationPreconditionFailed) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOPostCounselingInformationPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOPostCounselingInformationUnprocessableEntityCode is the HTTP code returned for type UpdateMTOPostCounselingInformationUnprocessableEntity
const UpdateMTOPostCounselingInformationUnprocessableEntityCode int = 422

/*UpdateMTOPostCounselingInformationUnprocessableEntity The request payload is invalid

swagger:response updateMTOPostCounselingInformationUnprocessableEntity
*/
type UpdateMTOPostCounselingInformationUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ValidationError `json:"body,omitempty"`
}

// NewUpdateMTOPostCounselingInformationUnprocessableEntity creates UpdateMTOPostCounselingInformationUnprocessableEntity with default headers values
func NewUpdateMTOPostCounselingInformationUnprocessableEntity() *UpdateMTOPostCounselingInformationUnprocessableEntity {

	return &UpdateMTOPostCounselingInformationUnprocessableEntity{}
}

// WithPayload adds the payload to the update m t o post counseling information unprocessable entity response
func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) WithPayload(payload *primemessages.ValidationError) *UpdateMTOPostCounselingInformationUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o post counseling information unprocessable entity response
func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) SetPayload(payload *primemessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOPostCounselingInformationInternalServerErrorCode is the HTTP code returned for type UpdateMTOPostCounselingInformationInternalServerError
const UpdateMTOPostCounselingInformationInternalServerErrorCode int = 500

/*UpdateMTOPostCounselingInformationInternalServerError A server error occurred

swagger:response updateMTOPostCounselingInformationInternalServerError
*/
type UpdateMTOPostCounselingInformationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOPostCounselingInformationInternalServerError creates UpdateMTOPostCounselingInformationInternalServerError with default headers values
func NewUpdateMTOPostCounselingInformationInternalServerError() *UpdateMTOPostCounselingInformationInternalServerError {

	return &UpdateMTOPostCounselingInformationInternalServerError{}
}

// WithPayload adds the payload to the update m t o post counseling information internal server error response
func (o *UpdateMTOPostCounselingInformationInternalServerError) WithPayload(payload interface{}) *UpdateMTOPostCounselingInformationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o post counseling information internal server error response
func (o *UpdateMTOPostCounselingInformationInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOPostCounselingInformationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}