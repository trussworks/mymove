// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/supportmessages"
)

// ReceiveWebhookNotificationOKCode is the HTTP code returned for type ReceiveWebhookNotificationOK
const ReceiveWebhookNotificationOKCode int = 200

/*ReceiveWebhookNotificationOK Received notification

swagger:response receiveWebhookNotificationOK
*/
type ReceiveWebhookNotificationOK struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.WebhookNotification `json:"body,omitempty"`
}

// NewReceiveWebhookNotificationOK creates ReceiveWebhookNotificationOK with default headers values
func NewReceiveWebhookNotificationOK() *ReceiveWebhookNotificationOK {

	return &ReceiveWebhookNotificationOK{}
}

// WithPayload adds the payload to the receive webhook notification o k response
func (o *ReceiveWebhookNotificationOK) WithPayload(payload *supportmessages.WebhookNotification) *ReceiveWebhookNotificationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the receive webhook notification o k response
func (o *ReceiveWebhookNotificationOK) SetPayload(payload *supportmessages.WebhookNotification) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReceiveWebhookNotificationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReceiveWebhookNotificationBadRequestCode is the HTTP code returned for type ReceiveWebhookNotificationBadRequest
const ReceiveWebhookNotificationBadRequestCode int = 400

/*ReceiveWebhookNotificationBadRequest The request payload is invalid.

swagger:response receiveWebhookNotificationBadRequest
*/
type ReceiveWebhookNotificationBadRequest struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewReceiveWebhookNotificationBadRequest creates ReceiveWebhookNotificationBadRequest with default headers values
func NewReceiveWebhookNotificationBadRequest() *ReceiveWebhookNotificationBadRequest {

	return &ReceiveWebhookNotificationBadRequest{}
}

// WithPayload adds the payload to the receive webhook notification bad request response
func (o *ReceiveWebhookNotificationBadRequest) WithPayload(payload *supportmessages.ClientError) *ReceiveWebhookNotificationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the receive webhook notification bad request response
func (o *ReceiveWebhookNotificationBadRequest) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReceiveWebhookNotificationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReceiveWebhookNotificationUnauthorizedCode is the HTTP code returned for type ReceiveWebhookNotificationUnauthorized
const ReceiveWebhookNotificationUnauthorizedCode int = 401

/*ReceiveWebhookNotificationUnauthorized The request was denied.

swagger:response receiveWebhookNotificationUnauthorized
*/
type ReceiveWebhookNotificationUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewReceiveWebhookNotificationUnauthorized creates ReceiveWebhookNotificationUnauthorized with default headers values
func NewReceiveWebhookNotificationUnauthorized() *ReceiveWebhookNotificationUnauthorized {

	return &ReceiveWebhookNotificationUnauthorized{}
}

// WithPayload adds the payload to the receive webhook notification unauthorized response
func (o *ReceiveWebhookNotificationUnauthorized) WithPayload(payload *supportmessages.ClientError) *ReceiveWebhookNotificationUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the receive webhook notification unauthorized response
func (o *ReceiveWebhookNotificationUnauthorized) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReceiveWebhookNotificationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReceiveWebhookNotificationForbiddenCode is the HTTP code returned for type ReceiveWebhookNotificationForbidden
const ReceiveWebhookNotificationForbiddenCode int = 403

/*ReceiveWebhookNotificationForbidden The request was denied.

swagger:response receiveWebhookNotificationForbidden
*/
type ReceiveWebhookNotificationForbidden struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewReceiveWebhookNotificationForbidden creates ReceiveWebhookNotificationForbidden with default headers values
func NewReceiveWebhookNotificationForbidden() *ReceiveWebhookNotificationForbidden {

	return &ReceiveWebhookNotificationForbidden{}
}

// WithPayload adds the payload to the receive webhook notification forbidden response
func (o *ReceiveWebhookNotificationForbidden) WithPayload(payload *supportmessages.ClientError) *ReceiveWebhookNotificationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the receive webhook notification forbidden response
func (o *ReceiveWebhookNotificationForbidden) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReceiveWebhookNotificationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReceiveWebhookNotificationInternalServerErrorCode is the HTTP code returned for type ReceiveWebhookNotificationInternalServerError
const ReceiveWebhookNotificationInternalServerErrorCode int = 500

/*ReceiveWebhookNotificationInternalServerError A server error occurred.

swagger:response receiveWebhookNotificationInternalServerError
*/
type ReceiveWebhookNotificationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.Error `json:"body,omitempty"`
}

// NewReceiveWebhookNotificationInternalServerError creates ReceiveWebhookNotificationInternalServerError with default headers values
func NewReceiveWebhookNotificationInternalServerError() *ReceiveWebhookNotificationInternalServerError {

	return &ReceiveWebhookNotificationInternalServerError{}
}

// WithPayload adds the payload to the receive webhook notification internal server error response
func (o *ReceiveWebhookNotificationInternalServerError) WithPayload(payload *supportmessages.Error) *ReceiveWebhookNotificationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the receive webhook notification internal server error response
func (o *ReceiveWebhookNotificationInternalServerError) SetPayload(payload *supportmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReceiveWebhookNotificationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
