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

// MoveTaskOrder move task order
//
// swagger:model MoveTaskOrder
type MoveTaskOrder struct {

	// available to prime at
	// Format: date-time
	AvailableToPrimeAt *strfmt.DateTime `json:"availableToPrimeAt,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// destination address
	DestinationAddress *Address `json:"destinationAddress,omitempty"`

	// destination duty station
	// Format: uuid
	DestinationDutyStation strfmt.UUID `json:"destinationDutyStation,omitempty"`

	// e tag
	ETag string `json:"eTag,omitempty"`

	// entitlements
	Entitlements *Entitlements `json:"entitlements,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// is canceled
	IsCanceled *bool `json:"isCanceled,omitempty"`

	// locator
	Locator string `json:"locator,omitempty"`

	// move order ID
	// Format: uuid
	MoveOrderID strfmt.UUID `json:"moveOrderID,omitempty"`

	// origin duty station
	// Format: uuid
	OriginDutyStation strfmt.UUID `json:"originDutyStation,omitempty"`

	// pickup address
	PickupAddress *Address `json:"pickupAddress,omitempty"`

	// reference Id
	ReferenceID string `json:"referenceId,omitempty"`

	// requested pickup date
	// Format: date
	RequestedPickupDate strfmt.Date `json:"requestedPickupDate,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this move task order
func (m *MoveTaskOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableToPrimeAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickupAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedPickupDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveTaskOrder) validateAvailableToPrimeAt(formats strfmt.Registry) error {

	if swag.IsZero(m.AvailableToPrimeAt) { // not required
		return nil
	}

	if err := validate.FormatOf("availableToPrimeAt", "body", "date-time", m.AvailableToPrimeAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateDestinationAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationAddress) { // not required
		return nil
	}

	if m.DestinationAddress != nil {
		if err := m.DestinationAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MoveTaskOrder) validateDestinationDutyStation(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationDutyStation) { // not required
		return nil
	}

	if err := validate.FormatOf("destinationDutyStation", "body", "uuid", m.DestinationDutyStation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateEntitlements(formats strfmt.Registry) error {

	if swag.IsZero(m.Entitlements) { // not required
		return nil
	}

	if m.Entitlements != nil {
		if err := m.Entitlements.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entitlements")
			}
			return err
		}
	}

	return nil
}

func (m *MoveTaskOrder) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateMoveOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveOrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveOrderID", "body", "uuid", m.MoveOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateOriginDutyStation(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginDutyStation) { // not required
		return nil
	}

	if err := validate.FormatOf("originDutyStation", "body", "uuid", m.OriginDutyStation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validatePickupAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.PickupAddress) { // not required
		return nil
	}

	if m.PickupAddress != nil {
		if err := m.PickupAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pickupAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MoveTaskOrder) validateRequestedPickupDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedPickupDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedPickupDate", "body", "date", m.RequestedPickupDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveTaskOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveTaskOrder) UnmarshalBinary(b []byte) error {
	var res MoveTaskOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
