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

// CreateWebhookNotificationReader is a Reader for the CreateWebhookNotification structure.
type CreateWebhookNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWebhookNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateWebhookNotificationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateWebhookNotificationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateWebhookNotificationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateWebhookNotificationCreated creates a CreateWebhookNotificationCreated with default headers values
func NewCreateWebhookNotificationCreated() *CreateWebhookNotificationCreated {
	return &CreateWebhookNotificationCreated{}
}

/*CreateWebhookNotificationCreated handles this case with default header values.

Successful creation
*/
type CreateWebhookNotificationCreated struct {
	Payload *supportmessages.WebhookNotification
}

func (o *CreateWebhookNotificationCreated) Error() string {
	return fmt.Sprintf("[POST /webhook-notifications][%d] createWebhookNotificationCreated  %+v", 201, o.Payload)
}

func (o *CreateWebhookNotificationCreated) GetPayload() *supportmessages.WebhookNotification {
	return o.Payload
}

func (o *CreateWebhookNotificationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.WebhookNotification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWebhookNotificationUnprocessableEntity creates a CreateWebhookNotificationUnprocessableEntity with default headers values
func NewCreateWebhookNotificationUnprocessableEntity() *CreateWebhookNotificationUnprocessableEntity {
	return &CreateWebhookNotificationUnprocessableEntity{}
}

/*CreateWebhookNotificationUnprocessableEntity handles this case with default header values.

The payload was unprocessable.
*/
type CreateWebhookNotificationUnprocessableEntity struct {
	Payload *supportmessages.ValidationError
}

func (o *CreateWebhookNotificationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /webhook-notifications][%d] createWebhookNotificationUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateWebhookNotificationUnprocessableEntity) GetPayload() *supportmessages.ValidationError {
	return o.Payload
}

func (o *CreateWebhookNotificationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWebhookNotificationInternalServerError creates a CreateWebhookNotificationInternalServerError with default headers values
func NewCreateWebhookNotificationInternalServerError() *CreateWebhookNotificationInternalServerError {
	return &CreateWebhookNotificationInternalServerError{}
}

/*CreateWebhookNotificationInternalServerError handles this case with default header values.

A server error occurred.
*/
type CreateWebhookNotificationInternalServerError struct {
	Payload *supportmessages.Error
}

func (o *CreateWebhookNotificationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /webhook-notifications][%d] createWebhookNotificationInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateWebhookNotificationInternalServerError) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *CreateWebhookNotificationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}