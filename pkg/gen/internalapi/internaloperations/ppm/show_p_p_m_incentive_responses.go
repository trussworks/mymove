// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"
)

// ShowPPMIncentiveOKCode is the HTTP code returned for type ShowPPMIncentiveOK
const ShowPPMIncentiveOKCode int = 200

/*ShowPPMIncentiveOK Made calculation of PPM incentive

swagger:response showPPMIncentiveOK
*/
type ShowPPMIncentiveOK struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.PPMIncentive `json:"body,omitempty"`
}

// NewShowPPMIncentiveOK creates ShowPPMIncentiveOK with default headers values
func NewShowPPMIncentiveOK() *ShowPPMIncentiveOK {

	return &ShowPPMIncentiveOK{}
}

// WithPayload adds the payload to the show p p m incentive o k response
func (o *ShowPPMIncentiveOK) WithPayload(payload *internalmessages.PPMIncentive) *ShowPPMIncentiveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show p p m incentive o k response
func (o *ShowPPMIncentiveOK) SetPayload(payload *internalmessages.PPMIncentive) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowPPMIncentiveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowPPMIncentiveBadRequestCode is the HTTP code returned for type ShowPPMIncentiveBadRequest
const ShowPPMIncentiveBadRequestCode int = 400

/*ShowPPMIncentiveBadRequest invalid request

swagger:response showPPMIncentiveBadRequest
*/
type ShowPPMIncentiveBadRequest struct {
}

// NewShowPPMIncentiveBadRequest creates ShowPPMIncentiveBadRequest with default headers values
func NewShowPPMIncentiveBadRequest() *ShowPPMIncentiveBadRequest {

	return &ShowPPMIncentiveBadRequest{}
}

// WriteResponse to the client
func (o *ShowPPMIncentiveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ShowPPMIncentiveUnauthorizedCode is the HTTP code returned for type ShowPPMIncentiveUnauthorized
const ShowPPMIncentiveUnauthorizedCode int = 401

/*ShowPPMIncentiveUnauthorized request requires user authentication

swagger:response showPPMIncentiveUnauthorized
*/
type ShowPPMIncentiveUnauthorized struct {
}

// NewShowPPMIncentiveUnauthorized creates ShowPPMIncentiveUnauthorized with default headers values
func NewShowPPMIncentiveUnauthorized() *ShowPPMIncentiveUnauthorized {

	return &ShowPPMIncentiveUnauthorized{}
}

// WriteResponse to the client
func (o *ShowPPMIncentiveUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ShowPPMIncentiveForbiddenCode is the HTTP code returned for type ShowPPMIncentiveForbidden
const ShowPPMIncentiveForbiddenCode int = 403

/*ShowPPMIncentiveForbidden user is not authorized

swagger:response showPPMIncentiveForbidden
*/
type ShowPPMIncentiveForbidden struct {
}

// NewShowPPMIncentiveForbidden creates ShowPPMIncentiveForbidden with default headers values
func NewShowPPMIncentiveForbidden() *ShowPPMIncentiveForbidden {

	return &ShowPPMIncentiveForbidden{}
}

// WriteResponse to the client
func (o *ShowPPMIncentiveForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ShowPPMIncentiveConflictCode is the HTTP code returned for type ShowPPMIncentiveConflict
const ShowPPMIncentiveConflictCode int = 409

/*ShowPPMIncentiveConflict distance is less than 50 miles (no short haul moves)

swagger:response showPPMIncentiveConflict
*/
type ShowPPMIncentiveConflict struct {
}

// NewShowPPMIncentiveConflict creates ShowPPMIncentiveConflict with default headers values
func NewShowPPMIncentiveConflict() *ShowPPMIncentiveConflict {

	return &ShowPPMIncentiveConflict{}
}

// WriteResponse to the client
func (o *ShowPPMIncentiveConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// ShowPPMIncentiveInternalServerErrorCode is the HTTP code returned for type ShowPPMIncentiveInternalServerError
const ShowPPMIncentiveInternalServerErrorCode int = 500

/*ShowPPMIncentiveInternalServerError internal server error

swagger:response showPPMIncentiveInternalServerError
*/
type ShowPPMIncentiveInternalServerError struct {
}

// NewShowPPMIncentiveInternalServerError creates ShowPPMIncentiveInternalServerError with default headers values
func NewShowPPMIncentiveInternalServerError() *ShowPPMIncentiveInternalServerError {

	return &ShowPPMIncentiveInternalServerError{}
}

// WriteResponse to the client
func (o *ShowPPMIncentiveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
