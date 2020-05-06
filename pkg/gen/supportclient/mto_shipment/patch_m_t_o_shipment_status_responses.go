// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	supportmessages "github.com/transcom/mymove/pkg/gen/supportmessages"
)

// PatchMTOShipmentStatusReader is a Reader for the PatchMTOShipmentStatus structure.
type PatchMTOShipmentStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchMTOShipmentStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchMTOShipmentStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchMTOShipmentStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchMTOShipmentStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchMTOShipmentStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchMTOShipmentStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchMTOShipmentStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewPatchMTOShipmentStatusPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPatchMTOShipmentStatusUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchMTOShipmentStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchMTOShipmentStatusOK creates a PatchMTOShipmentStatusOK with default headers values
func NewPatchMTOShipmentStatusOK() *PatchMTOShipmentStatusOK {
	return &PatchMTOShipmentStatusOK{}
}

/*PatchMTOShipmentStatusOK handles this case with default header values.

Successfully updated the shipment's status.
*/
type PatchMTOShipmentStatusOK struct {
	Payload *supportmessages.MTOShipment
}

func (o *PatchMTOShipmentStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /mto-shipments/{mtoShipmentID}/status][%d] patchMTOShipmentStatusOK  %+v", 200, o.Payload)
}

func (o *PatchMTOShipmentStatusOK) GetPayload() *supportmessages.MTOShipment {
	return o.Payload
}

func (o *PatchMTOShipmentStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.MTOShipment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMTOShipmentStatusBadRequest creates a PatchMTOShipmentStatusBadRequest with default headers values
func NewPatchMTOShipmentStatusBadRequest() *PatchMTOShipmentStatusBadRequest {
	return &PatchMTOShipmentStatusBadRequest{}
}

/*PatchMTOShipmentStatusBadRequest handles this case with default header values.

The parameters were invalid.
*/
type PatchMTOShipmentStatusBadRequest struct {
	Payload *supportmessages.Error
}

func (o *PatchMTOShipmentStatusBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /mto-shipments/{mtoShipmentID}/status][%d] patchMTOShipmentStatusBadRequest  %+v", 400, o.Payload)
}

func (o *PatchMTOShipmentStatusBadRequest) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *PatchMTOShipmentStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMTOShipmentStatusUnauthorized creates a PatchMTOShipmentStatusUnauthorized with default headers values
func NewPatchMTOShipmentStatusUnauthorized() *PatchMTOShipmentStatusUnauthorized {
	return &PatchMTOShipmentStatusUnauthorized{}
}

/*PatchMTOShipmentStatusUnauthorized handles this case with default header values.

The request was unauthorized.
*/
type PatchMTOShipmentStatusUnauthorized struct {
	Payload interface{}
}

func (o *PatchMTOShipmentStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /mto-shipments/{mtoShipmentID}/status][%d] patchMTOShipmentStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchMTOShipmentStatusUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchMTOShipmentStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMTOShipmentStatusForbidden creates a PatchMTOShipmentStatusForbidden with default headers values
func NewPatchMTOShipmentStatusForbidden() *PatchMTOShipmentStatusForbidden {
	return &PatchMTOShipmentStatusForbidden{}
}

/*PatchMTOShipmentStatusForbidden handles this case with default header values.

The client doesn't have permissions to perform the request.
*/
type PatchMTOShipmentStatusForbidden struct {
	Payload interface{}
}

func (o *PatchMTOShipmentStatusForbidden) Error() string {
	return fmt.Sprintf("[PATCH /mto-shipments/{mtoShipmentID}/status][%d] patchMTOShipmentStatusForbidden  %+v", 403, o.Payload)
}

func (o *PatchMTOShipmentStatusForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchMTOShipmentStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMTOShipmentStatusNotFound creates a PatchMTOShipmentStatusNotFound with default headers values
func NewPatchMTOShipmentStatusNotFound() *PatchMTOShipmentStatusNotFound {
	return &PatchMTOShipmentStatusNotFound{}
}

/*PatchMTOShipmentStatusNotFound handles this case with default header values.

The requested resource wasn't found.
*/
type PatchMTOShipmentStatusNotFound struct {
	Payload *supportmessages.Error
}

func (o *PatchMTOShipmentStatusNotFound) Error() string {
	return fmt.Sprintf("[PATCH /mto-shipments/{mtoShipmentID}/status][%d] patchMTOShipmentStatusNotFound  %+v", 404, o.Payload)
}

func (o *PatchMTOShipmentStatusNotFound) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *PatchMTOShipmentStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMTOShipmentStatusConflict creates a PatchMTOShipmentStatusConflict with default headers values
func NewPatchMTOShipmentStatusConflict() *PatchMTOShipmentStatusConflict {
	return &PatchMTOShipmentStatusConflict{}
}

/*PatchMTOShipmentStatusConflict handles this case with default header values.

Conflict error due to trying to change the status of shipment that is not currently "SUBMITTED".
*/
type PatchMTOShipmentStatusConflict struct {
	Payload interface{}
}

func (o *PatchMTOShipmentStatusConflict) Error() string {
	return fmt.Sprintf("[PATCH /mto-shipments/{mtoShipmentID}/status][%d] patchMTOShipmentStatusConflict  %+v", 409, o.Payload)
}

func (o *PatchMTOShipmentStatusConflict) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchMTOShipmentStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMTOShipmentStatusPreconditionFailed creates a PatchMTOShipmentStatusPreconditionFailed with default headers values
func NewPatchMTOShipmentStatusPreconditionFailed() *PatchMTOShipmentStatusPreconditionFailed {
	return &PatchMTOShipmentStatusPreconditionFailed{}
}

/*PatchMTOShipmentStatusPreconditionFailed handles this case with default header values.

Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.
*/
type PatchMTOShipmentStatusPreconditionFailed struct {
	Payload *supportmessages.Error
}

func (o *PatchMTOShipmentStatusPreconditionFailed) Error() string {
	return fmt.Sprintf("[PATCH /mto-shipments/{mtoShipmentID}/status][%d] patchMTOShipmentStatusPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PatchMTOShipmentStatusPreconditionFailed) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *PatchMTOShipmentStatusPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMTOShipmentStatusUnprocessableEntity creates a PatchMTOShipmentStatusUnprocessableEntity with default headers values
func NewPatchMTOShipmentStatusUnprocessableEntity() *PatchMTOShipmentStatusUnprocessableEntity {
	return &PatchMTOShipmentStatusUnprocessableEntity{}
}

/*PatchMTOShipmentStatusUnprocessableEntity handles this case with default header values.

The payload was unprocessable.
*/
type PatchMTOShipmentStatusUnprocessableEntity struct {
	Payload *supportmessages.ValidationError
}

func (o *PatchMTOShipmentStatusUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /mto-shipments/{mtoShipmentID}/status][%d] patchMTOShipmentStatusUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PatchMTOShipmentStatusUnprocessableEntity) GetPayload() *supportmessages.ValidationError {
	return o.Payload
}

func (o *PatchMTOShipmentStatusUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMTOShipmentStatusInternalServerError creates a PatchMTOShipmentStatusInternalServerError with default headers values
func NewPatchMTOShipmentStatusInternalServerError() *PatchMTOShipmentStatusInternalServerError {
	return &PatchMTOShipmentStatusInternalServerError{}
}

/*PatchMTOShipmentStatusInternalServerError handles this case with default header values.

A server error occurred.
*/
type PatchMTOShipmentStatusInternalServerError struct {
	Payload *supportmessages.Error
}

func (o *PatchMTOShipmentStatusInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /mto-shipments/{mtoShipmentID}/status][%d] patchMTOShipmentStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchMTOShipmentStatusInternalServerError) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *PatchMTOShipmentStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
