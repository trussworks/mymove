// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	supportmessages "github.com/transcom/mymove/pkg/gen/supportmessages"
)

// CreateMoveTaskOrderReader is a Reader for the CreateMoveTaskOrder structure.
type CreateMoveTaskOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMoveTaskOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMoveTaskOrderCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateMoveTaskOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateMoveTaskOrderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateMoveTaskOrderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateMoveTaskOrderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateMoveTaskOrderUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateMoveTaskOrderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateMoveTaskOrderCreated creates a CreateMoveTaskOrderCreated with default headers values
func NewCreateMoveTaskOrderCreated() *CreateMoveTaskOrderCreated {
	return &CreateMoveTaskOrderCreated{}
}

/*CreateMoveTaskOrderCreated handles this case with default header values.

Successfully created MoveTaskOrder object.
*/
type CreateMoveTaskOrderCreated struct {
	Payload *supportmessages.MoveTaskOrder
}

func (o *CreateMoveTaskOrderCreated) Error() string {
	return fmt.Sprintf("[POST /move-task-orders][%d] createMoveTaskOrderCreated  %+v", 201, o.Payload)
}

func (o *CreateMoveTaskOrderCreated) GetPayload() *supportmessages.MoveTaskOrder {
	return o.Payload
}

func (o *CreateMoveTaskOrderCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.MoveTaskOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMoveTaskOrderBadRequest creates a CreateMoveTaskOrderBadRequest with default headers values
func NewCreateMoveTaskOrderBadRequest() *CreateMoveTaskOrderBadRequest {
	return &CreateMoveTaskOrderBadRequest{}
}

/*CreateMoveTaskOrderBadRequest handles this case with default header values.

The parameters were invalid.
*/
type CreateMoveTaskOrderBadRequest struct {
	Payload *supportmessages.Error
}

func (o *CreateMoveTaskOrderBadRequest) Error() string {
	return fmt.Sprintf("[POST /move-task-orders][%d] createMoveTaskOrderBadRequest  %+v", 400, o.Payload)
}

func (o *CreateMoveTaskOrderBadRequest) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *CreateMoveTaskOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMoveTaskOrderUnauthorized creates a CreateMoveTaskOrderUnauthorized with default headers values
func NewCreateMoveTaskOrderUnauthorized() *CreateMoveTaskOrderUnauthorized {
	return &CreateMoveTaskOrderUnauthorized{}
}

/*CreateMoveTaskOrderUnauthorized handles this case with default header values.

The request was unauthorized.
*/
type CreateMoveTaskOrderUnauthorized struct {
	Payload interface{}
}

func (o *CreateMoveTaskOrderUnauthorized) Error() string {
	return fmt.Sprintf("[POST /move-task-orders][%d] createMoveTaskOrderUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateMoveTaskOrderUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateMoveTaskOrderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMoveTaskOrderForbidden creates a CreateMoveTaskOrderForbidden with default headers values
func NewCreateMoveTaskOrderForbidden() *CreateMoveTaskOrderForbidden {
	return &CreateMoveTaskOrderForbidden{}
}

/*CreateMoveTaskOrderForbidden handles this case with default header values.

The client doesn't have permissions to perform the request.
*/
type CreateMoveTaskOrderForbidden struct {
	Payload interface{}
}

func (o *CreateMoveTaskOrderForbidden) Error() string {
	return fmt.Sprintf("[POST /move-task-orders][%d] createMoveTaskOrderForbidden  %+v", 403, o.Payload)
}

func (o *CreateMoveTaskOrderForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateMoveTaskOrderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMoveTaskOrderNotFound creates a CreateMoveTaskOrderNotFound with default headers values
func NewCreateMoveTaskOrderNotFound() *CreateMoveTaskOrderNotFound {
	return &CreateMoveTaskOrderNotFound{}
}

/*CreateMoveTaskOrderNotFound handles this case with default header values.

The requested resource wasn't found.
*/
type CreateMoveTaskOrderNotFound struct {
	Payload *supportmessages.Error
}

func (o *CreateMoveTaskOrderNotFound) Error() string {
	return fmt.Sprintf("[POST /move-task-orders][%d] createMoveTaskOrderNotFound  %+v", 404, o.Payload)
}

func (o *CreateMoveTaskOrderNotFound) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *CreateMoveTaskOrderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMoveTaskOrderUnprocessableEntity creates a CreateMoveTaskOrderUnprocessableEntity with default headers values
func NewCreateMoveTaskOrderUnprocessableEntity() *CreateMoveTaskOrderUnprocessableEntity {
	return &CreateMoveTaskOrderUnprocessableEntity{}
}

/*CreateMoveTaskOrderUnprocessableEntity handles this case with default header values.

The payload was unprocessable.
*/
type CreateMoveTaskOrderUnprocessableEntity struct {
	Payload *supportmessages.ValidationError
}

func (o *CreateMoveTaskOrderUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /move-task-orders][%d] createMoveTaskOrderUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateMoveTaskOrderUnprocessableEntity) GetPayload() *supportmessages.ValidationError {
	return o.Payload
}

func (o *CreateMoveTaskOrderUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMoveTaskOrderInternalServerError creates a CreateMoveTaskOrderInternalServerError with default headers values
func NewCreateMoveTaskOrderInternalServerError() *CreateMoveTaskOrderInternalServerError {
	return &CreateMoveTaskOrderInternalServerError{}
}

/*CreateMoveTaskOrderInternalServerError handles this case with default header values.

A server error occurred.
*/
type CreateMoveTaskOrderInternalServerError struct {
	Payload *supportmessages.Error
}

func (o *CreateMoveTaskOrderInternalServerError) Error() string {
	return fmt.Sprintf("[POST /move-task-orders][%d] createMoveTaskOrderInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateMoveTaskOrderInternalServerError) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *CreateMoveTaskOrderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
