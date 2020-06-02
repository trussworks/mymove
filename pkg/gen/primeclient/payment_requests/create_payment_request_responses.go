// Code generated by go-swagger; DO NOT EDIT.

package payment_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// CreatePaymentRequestReader is a Reader for the CreatePaymentRequest structure.
type CreatePaymentRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePaymentRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePaymentRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreatePaymentRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePaymentRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreatePaymentRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreatePaymentRequestUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreatePaymentRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentRequestCreated creates a CreatePaymentRequestCreated with default headers values
func NewCreatePaymentRequestCreated() *CreatePaymentRequestCreated {
	return &CreatePaymentRequestCreated{}
}

/*CreatePaymentRequestCreated handles this case with default header values.

Successfully created a paymentRequest object.
*/
type CreatePaymentRequestCreated struct {
	Payload *primemessages.PaymentRequest
}

func (o *CreatePaymentRequestCreated) Error() string {
	return fmt.Sprintf("[POST /payment-requests][%d] createPaymentRequestCreated  %+v", 201, o.Payload)
}

func (o *CreatePaymentRequestCreated) GetPayload() *primemessages.PaymentRequest {
	return o.Payload
}

func (o *CreatePaymentRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.PaymentRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentRequestBadRequest creates a CreatePaymentRequestBadRequest with default headers values
func NewCreatePaymentRequestBadRequest() *CreatePaymentRequestBadRequest {
	return &CreatePaymentRequestBadRequest{}
}

/*CreatePaymentRequestBadRequest handles this case with default header values.

Request payload is invalid.
*/
type CreatePaymentRequestBadRequest struct {
	Payload *primemessages.ClientError
}

func (o *CreatePaymentRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /payment-requests][%d] createPaymentRequestBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePaymentRequestBadRequest) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreatePaymentRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentRequestUnauthorized creates a CreatePaymentRequestUnauthorized with default headers values
func NewCreatePaymentRequestUnauthorized() *CreatePaymentRequestUnauthorized {
	return &CreatePaymentRequestUnauthorized{}
}

/*CreatePaymentRequestUnauthorized handles this case with default header values.

The request was denied.
*/
type CreatePaymentRequestUnauthorized struct {
	Payload *primemessages.ClientError
}

func (o *CreatePaymentRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payment-requests][%d] createPaymentRequestUnauthorized  %+v", 401, o.Payload)
}

func (o *CreatePaymentRequestUnauthorized) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreatePaymentRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentRequestForbidden creates a CreatePaymentRequestForbidden with default headers values
func NewCreatePaymentRequestForbidden() *CreatePaymentRequestForbidden {
	return &CreatePaymentRequestForbidden{}
}

/*CreatePaymentRequestForbidden handles this case with default header values.

The request was denied.
*/
type CreatePaymentRequestForbidden struct {
	Payload *primemessages.ClientError
}

func (o *CreatePaymentRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /payment-requests][%d] createPaymentRequestForbidden  %+v", 403, o.Payload)
}

func (o *CreatePaymentRequestForbidden) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreatePaymentRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentRequestNotFound creates a CreatePaymentRequestNotFound with default headers values
func NewCreatePaymentRequestNotFound() *CreatePaymentRequestNotFound {
	return &CreatePaymentRequestNotFound{}
}

/*CreatePaymentRequestNotFound handles this case with default header values.

The requested resource wasn't found.
*/
type CreatePaymentRequestNotFound struct {
	Payload *primemessages.ClientError
}

func (o *CreatePaymentRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /payment-requests][%d] createPaymentRequestNotFound  %+v", 404, o.Payload)
}

func (o *CreatePaymentRequestNotFound) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreatePaymentRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentRequestUnprocessableEntity creates a CreatePaymentRequestUnprocessableEntity with default headers values
func NewCreatePaymentRequestUnprocessableEntity() *CreatePaymentRequestUnprocessableEntity {
	return &CreatePaymentRequestUnprocessableEntity{}
}

/*CreatePaymentRequestUnprocessableEntity handles this case with default header values.

The payload was unprocessable.
*/
type CreatePaymentRequestUnprocessableEntity struct {
	Payload *primemessages.ValidationError
}

func (o *CreatePaymentRequestUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /payment-requests][%d] createPaymentRequestUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreatePaymentRequestUnprocessableEntity) GetPayload() *primemessages.ValidationError {
	return o.Payload
}

func (o *CreatePaymentRequestUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentRequestInternalServerError creates a CreatePaymentRequestInternalServerError with default headers values
func NewCreatePaymentRequestInternalServerError() *CreatePaymentRequestInternalServerError {
	return &CreatePaymentRequestInternalServerError{}
}

/*CreatePaymentRequestInternalServerError handles this case with default header values.

A server error occurred.
*/
type CreatePaymentRequestInternalServerError struct {
	Payload *primemessages.Error
}

func (o *CreatePaymentRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /payment-requests][%d] createPaymentRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *CreatePaymentRequestInternalServerError) GetPayload() *primemessages.Error {
	return o.Payload
}

func (o *CreatePaymentRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
