// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// UpdateMTOShipmentAddressReader is a Reader for the UpdateMTOShipmentAddress structure.
type UpdateMTOShipmentAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMTOShipmentAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMTOShipmentAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMTOShipmentAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateMTOShipmentAddressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMTOShipmentAddressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMTOShipmentAddressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateMTOShipmentAddressConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewUpdateMTOShipmentAddressPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateMTOShipmentAddressUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMTOShipmentAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateMTOShipmentAddressOK creates a UpdateMTOShipmentAddressOK with default headers values
func NewUpdateMTOShipmentAddressOK() *UpdateMTOShipmentAddressOK {
	return &UpdateMTOShipmentAddressOK{}
}

/*UpdateMTOShipmentAddressOK handles this case with default header values.

Successfully updated the address.
*/
type UpdateMTOShipmentAddressOK struct {
	Payload *primemessages.Address
}

func (o *UpdateMTOShipmentAddressOK) Error() string {
	return fmt.Sprintf("[PUT /mto-shipments/{mtoShipmentID}/addresses/{addressID}][%d] updateMTOShipmentAddressOK  %+v", 200, o.Payload)
}

func (o *UpdateMTOShipmentAddressOK) GetPayload() *primemessages.Address {
	return o.Payload
}

func (o *UpdateMTOShipmentAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.Address)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentAddressBadRequest creates a UpdateMTOShipmentAddressBadRequest with default headers values
func NewUpdateMTOShipmentAddressBadRequest() *UpdateMTOShipmentAddressBadRequest {
	return &UpdateMTOShipmentAddressBadRequest{}
}

/*UpdateMTOShipmentAddressBadRequest handles this case with default header values.

The request payload is invalid.
*/
type UpdateMTOShipmentAddressBadRequest struct {
	Payload *primemessages.ClientError
}

func (o *UpdateMTOShipmentAddressBadRequest) Error() string {
	return fmt.Sprintf("[PUT /mto-shipments/{mtoShipmentID}/addresses/{addressID}][%d] updateMTOShipmentAddressBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateMTOShipmentAddressBadRequest) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *UpdateMTOShipmentAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentAddressUnauthorized creates a UpdateMTOShipmentAddressUnauthorized with default headers values
func NewUpdateMTOShipmentAddressUnauthorized() *UpdateMTOShipmentAddressUnauthorized {
	return &UpdateMTOShipmentAddressUnauthorized{}
}

/*UpdateMTOShipmentAddressUnauthorized handles this case with default header values.

The request was denied.
*/
type UpdateMTOShipmentAddressUnauthorized struct {
	Payload *primemessages.ClientError
}

func (o *UpdateMTOShipmentAddressUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /mto-shipments/{mtoShipmentID}/addresses/{addressID}][%d] updateMTOShipmentAddressUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateMTOShipmentAddressUnauthorized) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *UpdateMTOShipmentAddressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentAddressForbidden creates a UpdateMTOShipmentAddressForbidden with default headers values
func NewUpdateMTOShipmentAddressForbidden() *UpdateMTOShipmentAddressForbidden {
	return &UpdateMTOShipmentAddressForbidden{}
}

/*UpdateMTOShipmentAddressForbidden handles this case with default header values.

The request was denied.
*/
type UpdateMTOShipmentAddressForbidden struct {
	Payload *primemessages.ClientError
}

func (o *UpdateMTOShipmentAddressForbidden) Error() string {
	return fmt.Sprintf("[PUT /mto-shipments/{mtoShipmentID}/addresses/{addressID}][%d] updateMTOShipmentAddressForbidden  %+v", 403, o.Payload)
}

func (o *UpdateMTOShipmentAddressForbidden) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *UpdateMTOShipmentAddressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentAddressNotFound creates a UpdateMTOShipmentAddressNotFound with default headers values
func NewUpdateMTOShipmentAddressNotFound() *UpdateMTOShipmentAddressNotFound {
	return &UpdateMTOShipmentAddressNotFound{}
}

/*UpdateMTOShipmentAddressNotFound handles this case with default header values.

The requested resource wasn't found.
*/
type UpdateMTOShipmentAddressNotFound struct {
	Payload *primemessages.ClientError
}

func (o *UpdateMTOShipmentAddressNotFound) Error() string {
	return fmt.Sprintf("[PUT /mto-shipments/{mtoShipmentID}/addresses/{addressID}][%d] updateMTOShipmentAddressNotFound  %+v", 404, o.Payload)
}

func (o *UpdateMTOShipmentAddressNotFound) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *UpdateMTOShipmentAddressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentAddressConflict creates a UpdateMTOShipmentAddressConflict with default headers values
func NewUpdateMTOShipmentAddressConflict() *UpdateMTOShipmentAddressConflict {
	return &UpdateMTOShipmentAddressConflict{}
}

/*UpdateMTOShipmentAddressConflict handles this case with default header values.

The request could not be processed because of conflict in the current state of the resource.
*/
type UpdateMTOShipmentAddressConflict struct {
	Payload *primemessages.ClientError
}

func (o *UpdateMTOShipmentAddressConflict) Error() string {
	return fmt.Sprintf("[PUT /mto-shipments/{mtoShipmentID}/addresses/{addressID}][%d] updateMTOShipmentAddressConflict  %+v", 409, o.Payload)
}

func (o *UpdateMTOShipmentAddressConflict) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *UpdateMTOShipmentAddressConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentAddressPreconditionFailed creates a UpdateMTOShipmentAddressPreconditionFailed with default headers values
func NewUpdateMTOShipmentAddressPreconditionFailed() *UpdateMTOShipmentAddressPreconditionFailed {
	return &UpdateMTOShipmentAddressPreconditionFailed{}
}

/*UpdateMTOShipmentAddressPreconditionFailed handles this case with default header values.

Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.
*/
type UpdateMTOShipmentAddressPreconditionFailed struct {
	Payload *primemessages.ClientError
}

func (o *UpdateMTOShipmentAddressPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /mto-shipments/{mtoShipmentID}/addresses/{addressID}][%d] updateMTOShipmentAddressPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateMTOShipmentAddressPreconditionFailed) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *UpdateMTOShipmentAddressPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentAddressUnprocessableEntity creates a UpdateMTOShipmentAddressUnprocessableEntity with default headers values
func NewUpdateMTOShipmentAddressUnprocessableEntity() *UpdateMTOShipmentAddressUnprocessableEntity {
	return &UpdateMTOShipmentAddressUnprocessableEntity{}
}

/*UpdateMTOShipmentAddressUnprocessableEntity handles this case with default header values.

The payload was unprocessable.
*/
type UpdateMTOShipmentAddressUnprocessableEntity struct {
	Payload *primemessages.ValidationError
}

func (o *UpdateMTOShipmentAddressUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /mto-shipments/{mtoShipmentID}/addresses/{addressID}][%d] updateMTOShipmentAddressUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateMTOShipmentAddressUnprocessableEntity) GetPayload() *primemessages.ValidationError {
	return o.Payload
}

func (o *UpdateMTOShipmentAddressUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentAddressInternalServerError creates a UpdateMTOShipmentAddressInternalServerError with default headers values
func NewUpdateMTOShipmentAddressInternalServerError() *UpdateMTOShipmentAddressInternalServerError {
	return &UpdateMTOShipmentAddressInternalServerError{}
}

/*UpdateMTOShipmentAddressInternalServerError handles this case with default header values.

A server error occurred.
*/
type UpdateMTOShipmentAddressInternalServerError struct {
	Payload *primemessages.Error
}

func (o *UpdateMTOShipmentAddressInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /mto-shipments/{mtoShipmentID}/addresses/{addressID}][%d] updateMTOShipmentAddressInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateMTOShipmentAddressInternalServerError) GetPayload() *primemessages.Error {
	return o.Payload
}

func (o *UpdateMTOShipmentAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
