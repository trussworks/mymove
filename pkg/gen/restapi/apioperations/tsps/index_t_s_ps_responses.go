// Code generated by go-swagger; DO NOT EDIT.

package tsps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	apimessages "github.com/transcom/mymove/pkg/gen/apimessages"
)

// IndexTSPsOKCode is the HTTP code returned for type IndexTSPsOK
const IndexTSPsOKCode int = 200

/*IndexTSPsOK list of TSPs

swagger:response indexTSPsOK
*/
type IndexTSPsOK struct {

	/*
	  In: Body
	*/
	Payload []*apimessages.TSP `json:"body,omitempty"`
}

// NewIndexTSPsOK creates IndexTSPsOK with default headers values
func NewIndexTSPsOK() *IndexTSPsOK {

	return &IndexTSPsOK{}
}

// WithPayload adds the payload to the index t s ps o k response
func (o *IndexTSPsOK) WithPayload(payload []*apimessages.TSP) *IndexTSPsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the index t s ps o k response
func (o *IndexTSPsOK) SetPayload(payload []*apimessages.TSP) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IndexTSPsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*apimessages.TSP, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IndexTSPsBadRequestCode is the HTTP code returned for type IndexTSPsBadRequest
const IndexTSPsBadRequestCode int = 400

/*IndexTSPsBadRequest invalid request

swagger:response indexTSPsBadRequest
*/
type IndexTSPsBadRequest struct {
}

// NewIndexTSPsBadRequest creates IndexTSPsBadRequest with default headers values
func NewIndexTSPsBadRequest() *IndexTSPsBadRequest {

	return &IndexTSPsBadRequest{}
}

// WriteResponse to the client
func (o *IndexTSPsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// IndexTSPsUnauthorizedCode is the HTTP code returned for type IndexTSPsUnauthorized
const IndexTSPsUnauthorizedCode int = 401

/*IndexTSPsUnauthorized must be authenticated to access this endpoint

swagger:response indexTSPsUnauthorized
*/
type IndexTSPsUnauthorized struct {
}

// NewIndexTSPsUnauthorized creates IndexTSPsUnauthorized with default headers values
func NewIndexTSPsUnauthorized() *IndexTSPsUnauthorized {

	return &IndexTSPsUnauthorized{}
}

// WriteResponse to the client
func (o *IndexTSPsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// IndexTSPsInternalServerErrorCode is the HTTP code returned for type IndexTSPsInternalServerError
const IndexTSPsInternalServerErrorCode int = 500

/*IndexTSPsInternalServerError server error

swagger:response indexTSPsInternalServerError
*/
type IndexTSPsInternalServerError struct {
}

// NewIndexTSPsInternalServerError creates IndexTSPsInternalServerError with default headers values
func NewIndexTSPsInternalServerError() *IndexTSPsInternalServerError {

	return &IndexTSPsInternalServerError{}
}

// WriteResponse to the client
func (o *IndexTSPsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
