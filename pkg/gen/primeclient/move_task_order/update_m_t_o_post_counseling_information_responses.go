// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// UpdateMTOPostCounselingInformationReader is a Reader for the UpdateMTOPostCounselingInformation structure.
type UpdateMTOPostCounselingInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMTOPostCounselingInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMTOPostCounselingInformationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateMTOPostCounselingInformationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMTOPostCounselingInformationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMTOPostCounselingInformationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewUpdateMTOPostCounselingInformationPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateMTOPostCounselingInformationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMTOPostCounselingInformationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateMTOPostCounselingInformationOK creates a UpdateMTOPostCounselingInformationOK with default headers values
func NewUpdateMTOPostCounselingInformationOK() *UpdateMTOPostCounselingInformationOK {
	return &UpdateMTOPostCounselingInformationOK{}
}

/*UpdateMTOPostCounselingInformationOK handles this case with default header values.

Successfully updated move task order post counseling information
*/
type UpdateMTOPostCounselingInformationOK struct {
	Payload *primemessages.MoveTaskOrder
}

func (o *UpdateMTOPostCounselingInformationOK) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationOK  %+v", 200, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationOK) GetPayload() *primemessages.MoveTaskOrder {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.MoveTaskOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationUnauthorized creates a UpdateMTOPostCounselingInformationUnauthorized with default headers values
func NewUpdateMTOPostCounselingInformationUnauthorized() *UpdateMTOPostCounselingInformationUnauthorized {
	return &UpdateMTOPostCounselingInformationUnauthorized{}
}

/*UpdateMTOPostCounselingInformationUnauthorized handles this case with default header values.

The request was denied
*/
type UpdateMTOPostCounselingInformationUnauthorized struct {
	Payload interface{}
}

func (o *UpdateMTOPostCounselingInformationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationForbidden creates a UpdateMTOPostCounselingInformationForbidden with default headers values
func NewUpdateMTOPostCounselingInformationForbidden() *UpdateMTOPostCounselingInformationForbidden {
	return &UpdateMTOPostCounselingInformationForbidden{}
}

/*UpdateMTOPostCounselingInformationForbidden handles this case with default header values.

The request was denied
*/
type UpdateMTOPostCounselingInformationForbidden struct {
	Payload interface{}
}

func (o *UpdateMTOPostCounselingInformationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationNotFound creates a UpdateMTOPostCounselingInformationNotFound with default headers values
func NewUpdateMTOPostCounselingInformationNotFound() *UpdateMTOPostCounselingInformationNotFound {
	return &UpdateMTOPostCounselingInformationNotFound{}
}

/*UpdateMTOPostCounselingInformationNotFound handles this case with default header values.

The requested resource wasn't found
*/
type UpdateMTOPostCounselingInformationNotFound struct {
	Payload interface{}
}

func (o *UpdateMTOPostCounselingInformationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationPreconditionFailed creates a UpdateMTOPostCounselingInformationPreconditionFailed with default headers values
func NewUpdateMTOPostCounselingInformationPreconditionFailed() *UpdateMTOPostCounselingInformationPreconditionFailed {
	return &UpdateMTOPostCounselingInformationPreconditionFailed{}
}

/*UpdateMTOPostCounselingInformationPreconditionFailed handles this case with default header values.

precondition failed
*/
type UpdateMTOPostCounselingInformationPreconditionFailed struct {
	Payload interface{}
}

func (o *UpdateMTOPostCounselingInformationPreconditionFailed) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationPreconditionFailed) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationUnprocessableEntity creates a UpdateMTOPostCounselingInformationUnprocessableEntity with default headers values
func NewUpdateMTOPostCounselingInformationUnprocessableEntity() *UpdateMTOPostCounselingInformationUnprocessableEntity {
	return &UpdateMTOPostCounselingInformationUnprocessableEntity{}
}

/*UpdateMTOPostCounselingInformationUnprocessableEntity handles this case with default header values.

The request payload is invalid
*/
type UpdateMTOPostCounselingInformationUnprocessableEntity struct {
	Payload *primemessages.ValidationError
}

func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) GetPayload() *primemessages.ValidationError {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationInternalServerError creates a UpdateMTOPostCounselingInformationInternalServerError with default headers values
func NewUpdateMTOPostCounselingInformationInternalServerError() *UpdateMTOPostCounselingInformationInternalServerError {
	return &UpdateMTOPostCounselingInformationInternalServerError{}
}

/*UpdateMTOPostCounselingInformationInternalServerError handles this case with default header values.

A server error occurred
*/
type UpdateMTOPostCounselingInformationInternalServerError struct {
	Payload interface{}
}

func (o *UpdateMTOPostCounselingInformationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateMTOPostCounselingInformationBody update m t o post counseling information body
swagger:model UpdateMTOPostCounselingInformationBody
*/
type UpdateMTOPostCounselingInformationBody struct {

	// ppm estimated weight
	PpmEstimatedWeight int64 `json:"ppm_estimated_weight,omitempty"`

	// ppm type
	// Enum: [FULL PARTIAL]
	PpmType string `json:"ppm_type,omitempty"`
}

// Validate validates this update m t o post counseling information body
func (o *UpdateMTOPostCounselingInformationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePpmType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateMTOPostCounselingInformationBodyTypePpmTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FULL","PARTIAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateMTOPostCounselingInformationBodyTypePpmTypePropEnum = append(updateMTOPostCounselingInformationBodyTypePpmTypePropEnum, v)
	}
}

const (

	// UpdateMTOPostCounselingInformationBodyPpmTypeFULL captures enum value "FULL"
	UpdateMTOPostCounselingInformationBodyPpmTypeFULL string = "FULL"

	// UpdateMTOPostCounselingInformationBodyPpmTypePARTIAL captures enum value "PARTIAL"
	UpdateMTOPostCounselingInformationBodyPpmTypePARTIAL string = "PARTIAL"
)

// prop value enum
func (o *UpdateMTOPostCounselingInformationBody) validatePpmTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateMTOPostCounselingInformationBodyTypePpmTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *UpdateMTOPostCounselingInformationBody) validatePpmType(formats strfmt.Registry) error {

	if swag.IsZero(o.PpmType) { // not required
		return nil
	}

	// value enum
	if err := o.validatePpmTypeEnum("body"+"."+"ppm_type", "body", o.PpmType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateMTOPostCounselingInformationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateMTOPostCounselingInformationBody) UnmarshalBinary(b []byte) error {
	var res UpdateMTOPostCounselingInformationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
