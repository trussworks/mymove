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

// MakeMoveTaskOrderAvailableReader is a Reader for the MakeMoveTaskOrderAvailable structure.
type MakeMoveTaskOrderAvailableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MakeMoveTaskOrderAvailableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMakeMoveTaskOrderAvailableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMakeMoveTaskOrderAvailableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMakeMoveTaskOrderAvailableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMakeMoveTaskOrderAvailableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMakeMoveTaskOrderAvailableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewMakeMoveTaskOrderAvailablePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewMakeMoveTaskOrderAvailableUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMakeMoveTaskOrderAvailableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMakeMoveTaskOrderAvailableOK creates a MakeMoveTaskOrderAvailableOK with default headers values
func NewMakeMoveTaskOrderAvailableOK() *MakeMoveTaskOrderAvailableOK {
	return &MakeMoveTaskOrderAvailableOK{}
}

/*MakeMoveTaskOrderAvailableOK handles this case with default header values.

Successfully updated move task order status.
*/
type MakeMoveTaskOrderAvailableOK struct {
	Payload *supportmessages.MoveTaskOrder
}

func (o *MakeMoveTaskOrderAvailableOK) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/available-to-prime][%d] makeMoveTaskOrderAvailableOK  %+v", 200, o.Payload)
}

func (o *MakeMoveTaskOrderAvailableOK) GetPayload() *supportmessages.MoveTaskOrder {
	return o.Payload
}

func (o *MakeMoveTaskOrderAvailableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.MoveTaskOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMakeMoveTaskOrderAvailableBadRequest creates a MakeMoveTaskOrderAvailableBadRequest with default headers values
func NewMakeMoveTaskOrderAvailableBadRequest() *MakeMoveTaskOrderAvailableBadRequest {
	return &MakeMoveTaskOrderAvailableBadRequest{}
}

/*MakeMoveTaskOrderAvailableBadRequest handles this case with default header values.

The request payload is invalid.
*/
type MakeMoveTaskOrderAvailableBadRequest struct {
	Payload *supportmessages.ClientError
}

func (o *MakeMoveTaskOrderAvailableBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/available-to-prime][%d] makeMoveTaskOrderAvailableBadRequest  %+v", 400, o.Payload)
}

func (o *MakeMoveTaskOrderAvailableBadRequest) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *MakeMoveTaskOrderAvailableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMakeMoveTaskOrderAvailableUnauthorized creates a MakeMoveTaskOrderAvailableUnauthorized with default headers values
func NewMakeMoveTaskOrderAvailableUnauthorized() *MakeMoveTaskOrderAvailableUnauthorized {
	return &MakeMoveTaskOrderAvailableUnauthorized{}
}

/*MakeMoveTaskOrderAvailableUnauthorized handles this case with default header values.

The request was unauthorized.
*/
type MakeMoveTaskOrderAvailableUnauthorized struct {
	Payload interface{}
}

func (o *MakeMoveTaskOrderAvailableUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/available-to-prime][%d] makeMoveTaskOrderAvailableUnauthorized  %+v", 401, o.Payload)
}

func (o *MakeMoveTaskOrderAvailableUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *MakeMoveTaskOrderAvailableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMakeMoveTaskOrderAvailableForbidden creates a MakeMoveTaskOrderAvailableForbidden with default headers values
func NewMakeMoveTaskOrderAvailableForbidden() *MakeMoveTaskOrderAvailableForbidden {
	return &MakeMoveTaskOrderAvailableForbidden{}
}

/*MakeMoveTaskOrderAvailableForbidden handles this case with default header values.

The client doesn't have permissions to perform the request.
*/
type MakeMoveTaskOrderAvailableForbidden struct {
	Payload interface{}
}

func (o *MakeMoveTaskOrderAvailableForbidden) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/available-to-prime][%d] makeMoveTaskOrderAvailableForbidden  %+v", 403, o.Payload)
}

func (o *MakeMoveTaskOrderAvailableForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *MakeMoveTaskOrderAvailableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMakeMoveTaskOrderAvailableNotFound creates a MakeMoveTaskOrderAvailableNotFound with default headers values
func NewMakeMoveTaskOrderAvailableNotFound() *MakeMoveTaskOrderAvailableNotFound {
	return &MakeMoveTaskOrderAvailableNotFound{}
}

/*MakeMoveTaskOrderAvailableNotFound handles this case with default header values.

The requested resource wasn't found.
*/
type MakeMoveTaskOrderAvailableNotFound struct {
	Payload *supportmessages.ClientError
}

func (o *MakeMoveTaskOrderAvailableNotFound) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/available-to-prime][%d] makeMoveTaskOrderAvailableNotFound  %+v", 404, o.Payload)
}

func (o *MakeMoveTaskOrderAvailableNotFound) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *MakeMoveTaskOrderAvailableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMakeMoveTaskOrderAvailablePreconditionFailed creates a MakeMoveTaskOrderAvailablePreconditionFailed with default headers values
func NewMakeMoveTaskOrderAvailablePreconditionFailed() *MakeMoveTaskOrderAvailablePreconditionFailed {
	return &MakeMoveTaskOrderAvailablePreconditionFailed{}
}

/*MakeMoveTaskOrderAvailablePreconditionFailed handles this case with default header values.

Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.
*/
type MakeMoveTaskOrderAvailablePreconditionFailed struct {
	Payload *supportmessages.ClientError
}

func (o *MakeMoveTaskOrderAvailablePreconditionFailed) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/available-to-prime][%d] makeMoveTaskOrderAvailablePreconditionFailed  %+v", 412, o.Payload)
}

func (o *MakeMoveTaskOrderAvailablePreconditionFailed) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *MakeMoveTaskOrderAvailablePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMakeMoveTaskOrderAvailableUnprocessableEntity creates a MakeMoveTaskOrderAvailableUnprocessableEntity with default headers values
func NewMakeMoveTaskOrderAvailableUnprocessableEntity() *MakeMoveTaskOrderAvailableUnprocessableEntity {
	return &MakeMoveTaskOrderAvailableUnprocessableEntity{}
}

/*MakeMoveTaskOrderAvailableUnprocessableEntity handles this case with default header values.

The payload was unprocessable.
*/
type MakeMoveTaskOrderAvailableUnprocessableEntity struct {
	Payload *supportmessages.ValidationError
}

func (o *MakeMoveTaskOrderAvailableUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/available-to-prime][%d] makeMoveTaskOrderAvailableUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *MakeMoveTaskOrderAvailableUnprocessableEntity) GetPayload() *supportmessages.ValidationError {
	return o.Payload
}

func (o *MakeMoveTaskOrderAvailableUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMakeMoveTaskOrderAvailableInternalServerError creates a MakeMoveTaskOrderAvailableInternalServerError with default headers values
func NewMakeMoveTaskOrderAvailableInternalServerError() *MakeMoveTaskOrderAvailableInternalServerError {
	return &MakeMoveTaskOrderAvailableInternalServerError{}
}

/*MakeMoveTaskOrderAvailableInternalServerError handles this case with default header values.

A server error occurred.
*/
type MakeMoveTaskOrderAvailableInternalServerError struct {
	Payload *supportmessages.Error
}

func (o *MakeMoveTaskOrderAvailableInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/available-to-prime][%d] makeMoveTaskOrderAvailableInternalServerError  %+v", 500, o.Payload)
}

func (o *MakeMoveTaskOrderAvailableInternalServerError) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *MakeMoveTaskOrderAvailableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
