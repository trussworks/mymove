// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentRequest payment request
//
// swagger:model PaymentRequest
type PaymentRequest struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// e tag
	ETag string `json:"eTag,omitempty"`

	// id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// is final
	IsFinal *bool `json:"isFinal,omitempty"`

	// move task order
	MoveTaskOrder *Move `json:"moveTaskOrder,omitempty"`

	// move task order ID
	// Format: uuid
	MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

	// payment request number
	// Read Only: true
	PaymentRequestNumber string `json:"paymentRequestNumber,omitempty"`

	// proof of service docs
	ProofOfServiceDocs ProofOfServiceDocs `json:"proofOfServiceDocs,omitempty"`

	// rejection reason
	RejectionReason *string `json:"rejectionReason,omitempty"`

	// reviewed at
	// Format: date-time
	ReviewedAt *strfmt.DateTime `json:"reviewedAt,omitempty"`

	// service items
	ServiceItems PaymentServiceItems `json:"serviceItems,omitempty"`

	// status
	Status PaymentRequestStatus `json:"status,omitempty"`
}

// Validate validates this payment request
func (m *PaymentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTaskOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTaskOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProofOfServiceDocs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReviewedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRequest) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentRequest) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentRequest) validateMoveTaskOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrder) { // not required
		return nil
	}

	if m.MoveTaskOrder != nil {
		if err := m.MoveTaskOrder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moveTaskOrder")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRequest) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentRequest) validateProofOfServiceDocs(formats strfmt.Registry) error {

	if swag.IsZero(m.ProofOfServiceDocs) { // not required
		return nil
	}

	if err := m.ProofOfServiceDocs.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("proofOfServiceDocs")
		}
		return err
	}

	return nil
}

func (m *PaymentRequest) validateReviewedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ReviewedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("reviewedAt", "body", "date-time", m.ReviewedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentRequest) validateServiceItems(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceItems) { // not required
		return nil
	}

	if err := m.ServiceItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceItems")
		}
		return err
	}

	return nil
}

func (m *PaymentRequest) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRequest) UnmarshalBinary(b []byte) error {
	var res PaymentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
