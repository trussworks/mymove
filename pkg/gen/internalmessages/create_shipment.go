// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateShipment create shipment
// swagger:model CreateShipment
type CreateShipment struct {

	// agents
	Agents MTOAgents `json:"agents,omitempty"`

	// customer remarks
	CustomerRemarks *string `json:"customerRemarks,omitempty"`

	// destination address
	DestinationAddress *Address `json:"destinationAddress,omitempty"`

	// move task order ID
	// Required: true
	// Format: uuid
	MoveTaskOrderID *strfmt.UUID `json:"moveTaskOrderID"`

	// pickup address
	// Required: true
	PickupAddress *Address `json:"pickupAddress"`

	// requested delivery date
	// Format: date
	RequestedDeliveryDate strfmt.Date `json:"requestedDeliveryDate,omitempty"`

	// requested pickup date
	// Format: date
	RequestedPickupDate strfmt.Date `json:"requestedPickupDate,omitempty"`

	// shipment type
	// Required: true
	ShipmentType MTOShipmentType `json:"shipmentType"`
}

// Validate validates this create shipment
func (m *CreateShipment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTaskOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickupAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedPickupDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateShipment) validateAgents(formats strfmt.Registry) error {

	if swag.IsZero(m.Agents) { // not required
		return nil
	}

	if err := m.Agents.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("agents")
		}
		return err
	}

	return nil
}

func (m *CreateShipment) validateDestinationAddress(formats strfmt.Registry) error {

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

func (m *CreateShipment) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if err := validate.Required("moveTaskOrderID", "body", m.MoveTaskOrderID); err != nil {
		return err
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateShipment) validatePickupAddress(formats strfmt.Registry) error {

	if err := validate.Required("pickupAddress", "body", m.PickupAddress); err != nil {
		return err
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

func (m *CreateShipment) validateRequestedDeliveryDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedDeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedDeliveryDate", "body", "date", m.RequestedDeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateShipment) validateRequestedPickupDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedPickupDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedPickupDate", "body", "date", m.RequestedPickupDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateShipment) validateShipmentType(formats strfmt.Registry) error {

	if err := m.ShipmentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shipmentType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateShipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateShipment) UnmarshalBinary(b []byte) error {
	var res CreateShipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
