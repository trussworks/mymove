// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/supportmessages"
)

// ReceiveWebhookNotificationReader is a Reader for the ReceiveWebhookNotification structure.
type ReceiveWebhookNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReceiveWebhookNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReceiveWebhookNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReceiveWebhookNotificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReceiveWebhookNotificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReceiveWebhookNotificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReceiveWebhookNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReceiveWebhookNotificationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReceiveWebhookNotificationOK creates a ReceiveWebhookNotificationOK with default headers values
func NewReceiveWebhookNotificationOK() *ReceiveWebhookNotificationOK {
	return &ReceiveWebhookNotificationOK{}
}

/*ReceiveWebhookNotificationOK handles this case with default header values.

Successful creation
*/
type ReceiveWebhookNotificationOK struct {
	Payload *supportmessages.WebhookNotification
}

func (o *ReceiveWebhookNotificationOK) Error() string {
	return fmt.Sprintf("[POST /webhook-notify][%d] receiveWebhookNotificationOK  %+v", 200, o.Payload)
}

func (o *ReceiveWebhookNotificationOK) GetPayload() *supportmessages.WebhookNotification {
	return o.Payload
}

func (o *ReceiveWebhookNotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.WebhookNotification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReceiveWebhookNotificationBadRequest creates a ReceiveWebhookNotificationBadRequest with default headers values
func NewReceiveWebhookNotificationBadRequest() *ReceiveWebhookNotificationBadRequest {
	return &ReceiveWebhookNotificationBadRequest{}
}

/*ReceiveWebhookNotificationBadRequest handles this case with default header values.

Bad request
*/
type ReceiveWebhookNotificationBadRequest struct {
}

func (o *ReceiveWebhookNotificationBadRequest) Error() string {
	return fmt.Sprintf("[POST /webhook-notify][%d] receiveWebhookNotificationBadRequest ", 400)
}

func (o *ReceiveWebhookNotificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReceiveWebhookNotificationUnauthorized creates a ReceiveWebhookNotificationUnauthorized with default headers values
func NewReceiveWebhookNotificationUnauthorized() *ReceiveWebhookNotificationUnauthorized {
	return &ReceiveWebhookNotificationUnauthorized{}
}

/*ReceiveWebhookNotificationUnauthorized handles this case with default header values.

must be authenticated to use this endpoint
*/
type ReceiveWebhookNotificationUnauthorized struct {
}

func (o *ReceiveWebhookNotificationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /webhook-notify][%d] receiveWebhookNotificationUnauthorized ", 401)
}

func (o *ReceiveWebhookNotificationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReceiveWebhookNotificationForbidden creates a ReceiveWebhookNotificationForbidden with default headers values
func NewReceiveWebhookNotificationForbidden() *ReceiveWebhookNotificationForbidden {
	return &ReceiveWebhookNotificationForbidden{}
}

/*ReceiveWebhookNotificationForbidden handles this case with default header values.

Forbidden
*/
type ReceiveWebhookNotificationForbidden struct {
}

func (o *ReceiveWebhookNotificationForbidden) Error() string {
	return fmt.Sprintf("[POST /webhook-notify][%d] receiveWebhookNotificationForbidden ", 403)
}

func (o *ReceiveWebhookNotificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReceiveWebhookNotificationNotFound creates a ReceiveWebhookNotificationNotFound with default headers values
func NewReceiveWebhookNotificationNotFound() *ReceiveWebhookNotificationNotFound {
	return &ReceiveWebhookNotificationNotFound{}
}

/*ReceiveWebhookNotificationNotFound handles this case with default header values.

No orders found
*/
type ReceiveWebhookNotificationNotFound struct {
}

func (o *ReceiveWebhookNotificationNotFound) Error() string {
	return fmt.Sprintf("[POST /webhook-notify][%d] receiveWebhookNotificationNotFound ", 404)
}

func (o *ReceiveWebhookNotificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReceiveWebhookNotificationInternalServerError creates a ReceiveWebhookNotificationInternalServerError with default headers values
func NewReceiveWebhookNotificationInternalServerError() *ReceiveWebhookNotificationInternalServerError {
	return &ReceiveWebhookNotificationInternalServerError{}
}

/*ReceiveWebhookNotificationInternalServerError handles this case with default header values.

Server error
*/
type ReceiveWebhookNotificationInternalServerError struct {
}

func (o *ReceiveWebhookNotificationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /webhook-notify][%d] receiveWebhookNotificationInternalServerError ", 500)
}

func (o *ReceiveWebhookNotificationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
