// Code generated by go-swagger; DO NOT EDIT.

package addresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
)

// ShowAddressOKCode is the HTTP code returned for type ShowAddressOK
const ShowAddressOKCode int = 200

/*ShowAddressOK the requested address

swagger:response showAddressOK
*/
type ShowAddressOK struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.Address `json:"body,omitempty"`
}

// NewShowAddressOK creates ShowAddressOK with default headers values
func NewShowAddressOK() *ShowAddressOK {

	return &ShowAddressOK{}
}

// WithPayload adds the payload to the show address o k response
func (o *ShowAddressOK) WithPayload(payload *internalmessages.Address) *ShowAddressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show address o k response
func (o *ShowAddressOK) SetPayload(payload *internalmessages.Address) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowAddressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowAddressBadRequestCode is the HTTP code returned for type ShowAddressBadRequest
const ShowAddressBadRequestCode int = 400

/*ShowAddressBadRequest invalid request

swagger:response showAddressBadRequest
*/
type ShowAddressBadRequest struct {
}

// NewShowAddressBadRequest creates ShowAddressBadRequest with default headers values
func NewShowAddressBadRequest() *ShowAddressBadRequest {

	return &ShowAddressBadRequest{}
}

// WriteResponse to the client
func (o *ShowAddressBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ShowAddressForbiddenCode is the HTTP code returned for type ShowAddressForbidden
const ShowAddressForbiddenCode int = 403

/*ShowAddressForbidden not authorized

swagger:response showAddressForbidden
*/
type ShowAddressForbidden struct {
}

// NewShowAddressForbidden creates ShowAddressForbidden with default headers values
func NewShowAddressForbidden() *ShowAddressForbidden {

	return &ShowAddressForbidden{}
}

// WriteResponse to the client
func (o *ShowAddressForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ShowAddressNotFoundCode is the HTTP code returned for type ShowAddressNotFound
const ShowAddressNotFoundCode int = 404

/*ShowAddressNotFound not found

swagger:response showAddressNotFound
*/
type ShowAddressNotFound struct {
}

// NewShowAddressNotFound creates ShowAddressNotFound with default headers values
func NewShowAddressNotFound() *ShowAddressNotFound {

	return &ShowAddressNotFound{}
}

// WriteResponse to the client
func (o *ShowAddressNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ShowAddressInternalServerErrorCode is the HTTP code returned for type ShowAddressInternalServerError
const ShowAddressInternalServerErrorCode int = 500

/*ShowAddressInternalServerError server error

swagger:response showAddressInternalServerError
*/
type ShowAddressInternalServerError struct {
}

// NewShowAddressInternalServerError creates ShowAddressInternalServerError with default headers values
func NewShowAddressInternalServerError() *ShowAddressInternalServerError {

	return &ShowAddressInternalServerError{}
}

// WriteResponse to the client
func (o *ShowAddressInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
