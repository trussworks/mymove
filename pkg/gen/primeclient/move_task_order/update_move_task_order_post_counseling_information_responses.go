// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// UpdateMoveTaskOrderPostCounselingInformationReader is a Reader for the UpdateMoveTaskOrderPostCounselingInformation structure.
type UpdateMoveTaskOrderPostCounselingInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMoveTaskOrderPostCounselingInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMoveTaskOrderPostCounselingInformationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateMoveTaskOrderPostCounselingInformationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMoveTaskOrderPostCounselingInformationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMoveTaskOrderPostCounselingInformationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateMoveTaskOrderPostCounselingInformationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMoveTaskOrderPostCounselingInformationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateMoveTaskOrderPostCounselingInformationOK creates a UpdateMoveTaskOrderPostCounselingInformationOK with default headers values
func NewUpdateMoveTaskOrderPostCounselingInformationOK() *UpdateMoveTaskOrderPostCounselingInformationOK {
	return &UpdateMoveTaskOrderPostCounselingInformationOK{}
}

/*UpdateMoveTaskOrderPostCounselingInformationOK handles this case with default header values.

Successfully updated move task order post counseling information
*/
type UpdateMoveTaskOrderPostCounselingInformationOK struct {
	Payload *primemessages.MoveTaskOrder
}

func (o *UpdateMoveTaskOrderPostCounselingInformationOK) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMoveTaskOrderPostCounselingInformationOK  %+v", 200, o.Payload)
}

func (o *UpdateMoveTaskOrderPostCounselingInformationOK) GetPayload() *primemessages.MoveTaskOrder {
	return o.Payload
}

func (o *UpdateMoveTaskOrderPostCounselingInformationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.MoveTaskOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderPostCounselingInformationUnauthorized creates a UpdateMoveTaskOrderPostCounselingInformationUnauthorized with default headers values
func NewUpdateMoveTaskOrderPostCounselingInformationUnauthorized() *UpdateMoveTaskOrderPostCounselingInformationUnauthorized {
	return &UpdateMoveTaskOrderPostCounselingInformationUnauthorized{}
}

/*UpdateMoveTaskOrderPostCounselingInformationUnauthorized handles this case with default header values.

The request was denied
*/
type UpdateMoveTaskOrderPostCounselingInformationUnauthorized struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderPostCounselingInformationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMoveTaskOrderPostCounselingInformationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateMoveTaskOrderPostCounselingInformationUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderPostCounselingInformationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderPostCounselingInformationForbidden creates a UpdateMoveTaskOrderPostCounselingInformationForbidden with default headers values
func NewUpdateMoveTaskOrderPostCounselingInformationForbidden() *UpdateMoveTaskOrderPostCounselingInformationForbidden {
	return &UpdateMoveTaskOrderPostCounselingInformationForbidden{}
}

/*UpdateMoveTaskOrderPostCounselingInformationForbidden handles this case with default header values.

The request was denied
*/
type UpdateMoveTaskOrderPostCounselingInformationForbidden struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderPostCounselingInformationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMoveTaskOrderPostCounselingInformationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateMoveTaskOrderPostCounselingInformationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderPostCounselingInformationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderPostCounselingInformationNotFound creates a UpdateMoveTaskOrderPostCounselingInformationNotFound with default headers values
func NewUpdateMoveTaskOrderPostCounselingInformationNotFound() *UpdateMoveTaskOrderPostCounselingInformationNotFound {
	return &UpdateMoveTaskOrderPostCounselingInformationNotFound{}
}

/*UpdateMoveTaskOrderPostCounselingInformationNotFound handles this case with default header values.

The requested resource wasn't found
*/
type UpdateMoveTaskOrderPostCounselingInformationNotFound struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderPostCounselingInformationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMoveTaskOrderPostCounselingInformationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateMoveTaskOrderPostCounselingInformationNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderPostCounselingInformationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderPostCounselingInformationUnprocessableEntity creates a UpdateMoveTaskOrderPostCounselingInformationUnprocessableEntity with default headers values
func NewUpdateMoveTaskOrderPostCounselingInformationUnprocessableEntity() *UpdateMoveTaskOrderPostCounselingInformationUnprocessableEntity {
	return &UpdateMoveTaskOrderPostCounselingInformationUnprocessableEntity{}
}

/*UpdateMoveTaskOrderPostCounselingInformationUnprocessableEntity handles this case with default header values.

The request payload is invalid
*/
type UpdateMoveTaskOrderPostCounselingInformationUnprocessableEntity struct {
	Payload *primemessages.ValidationError
}

func (o *UpdateMoveTaskOrderPostCounselingInformationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMoveTaskOrderPostCounselingInformationUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateMoveTaskOrderPostCounselingInformationUnprocessableEntity) GetPayload() *primemessages.ValidationError {
	return o.Payload
}

func (o *UpdateMoveTaskOrderPostCounselingInformationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderPostCounselingInformationInternalServerError creates a UpdateMoveTaskOrderPostCounselingInformationInternalServerError with default headers values
func NewUpdateMoveTaskOrderPostCounselingInformationInternalServerError() *UpdateMoveTaskOrderPostCounselingInformationInternalServerError {
	return &UpdateMoveTaskOrderPostCounselingInformationInternalServerError{}
}

/*UpdateMoveTaskOrderPostCounselingInformationInternalServerError handles this case with default header values.

A server error occurred
*/
type UpdateMoveTaskOrderPostCounselingInformationInternalServerError struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderPostCounselingInformationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMoveTaskOrderPostCounselingInformationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateMoveTaskOrderPostCounselingInformationInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderPostCounselingInformationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateMoveTaskOrderPostCounselingInformationBody update move task order post counseling information body
swagger:model UpdateMoveTaskOrderPostCounselingInformationBody
*/
type UpdateMoveTaskOrderPostCounselingInformationBody struct {

	// ppm
	Ppm *primemessages.PrimePPM `json:"ppm,omitempty"`
}

// Validate validates this update move task order post counseling information body
func (o *UpdateMoveTaskOrderPostCounselingInformationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePpm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateMoveTaskOrderPostCounselingInformationBody) validatePpm(formats strfmt.Registry) error {

	if swag.IsZero(o.Ppm) { // not required
		return nil
	}

	if o.Ppm != nil {
		if err := o.Ppm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "ppm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateMoveTaskOrderPostCounselingInformationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateMoveTaskOrderPostCounselingInformationBody) UnmarshalBinary(b []byte) error {
	var res UpdateMoveTaskOrderPostCounselingInformationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
