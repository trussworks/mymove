// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateMTOShipment create m t o shipment
// swagger:model CreateMTOShipment
type CreateMTOShipment struct {

	// agents
	Agents MTOAgents `json:"agents,omitempty"`

	// customer remarks
	CustomerRemarks *string `json:"customerRemarks,omitempty"`

	// destination address
	// Required: true
	DestinationAddress *Address `json:"destinationAddress"`

	// move task order ID
	// Required: true
	// Format: uuid
	MoveTaskOrderID *strfmt.UUID `json:"moveTaskOrderID"`

	mtoServiceItemsField []MTOServiceItem

	// pickup address
	// Required: true
	PickupAddress *Address `json:"pickupAddress"`

	// Email or id of a contact person for this update
	PointOfContact string `json:"pointOfContact,omitempty"`

	// requested pickup date
	// Format: date
	RequestedPickupDate strfmt.Date `json:"requestedPickupDate,omitempty"`

	// shipment type
	// Required: true
	ShipmentType MTOShipmentType `json:"shipmentType"`
}

// MtoServiceItems gets the mto service items of this base type
func (m *CreateMTOShipment) MtoServiceItems() []MTOServiceItem {
	return m.mtoServiceItemsField
}

// SetMtoServiceItems sets the mto service items of this base type
func (m *CreateMTOShipment) SetMtoServiceItems(val []MTOServiceItem) {
	m.mtoServiceItemsField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CreateMTOShipment) UnmarshalJSON(raw []byte) error {
	var data struct {
		Agents MTOAgents `json:"agents,omitempty"`

		CustomerRemarks *string `json:"customerRemarks,omitempty"`

		DestinationAddress *Address `json:"destinationAddress"`

		MoveTaskOrderID *strfmt.UUID `json:"moveTaskOrderID"`

		MtoServiceItems json.RawMessage `json:"mtoServiceItems"`

		PickupAddress *Address `json:"pickupAddress"`

		PointOfContact string `json:"pointOfContact,omitempty"`

		RequestedPickupDate strfmt.Date `json:"requestedPickupDate,omitempty"`

		ShipmentType MTOShipmentType `json:"shipmentType"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propMtoServiceItems []MTOServiceItem
	if string(data.MtoServiceItems) != "null" {
		mtoServiceItems, err := UnmarshalMTOServiceItemSlice(bytes.NewBuffer(data.MtoServiceItems), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propMtoServiceItems = mtoServiceItems
	}

	var result CreateMTOShipment

	// agents
	result.Agents = data.Agents

	// customerRemarks
	result.CustomerRemarks = data.CustomerRemarks

	// destinationAddress
	result.DestinationAddress = data.DestinationAddress

	// moveTaskOrderID
	result.MoveTaskOrderID = data.MoveTaskOrderID

	// mtoServiceItems
	result.mtoServiceItemsField = propMtoServiceItems

	// pickupAddress
	result.PickupAddress = data.PickupAddress

	// pointOfContact
	result.PointOfContact = data.PointOfContact

	// requestedPickupDate
	result.RequestedPickupDate = data.RequestedPickupDate

	// shipmentType
	result.ShipmentType = data.ShipmentType

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CreateMTOShipment) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Agents MTOAgents `json:"agents,omitempty"`

		CustomerRemarks *string `json:"customerRemarks,omitempty"`

		DestinationAddress *Address `json:"destinationAddress"`

		MoveTaskOrderID *strfmt.UUID `json:"moveTaskOrderID"`

		PickupAddress *Address `json:"pickupAddress"`

		PointOfContact string `json:"pointOfContact,omitempty"`

		RequestedPickupDate strfmt.Date `json:"requestedPickupDate,omitempty"`

		ShipmentType MTOShipmentType `json:"shipmentType"`
	}{

		Agents: m.Agents,

		CustomerRemarks: m.CustomerRemarks,

		DestinationAddress: m.DestinationAddress,

		MoveTaskOrderID: m.MoveTaskOrderID,

		PickupAddress: m.PickupAddress,

		PointOfContact: m.PointOfContact,

		RequestedPickupDate: m.RequestedPickupDate,

		ShipmentType: m.ShipmentType,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		MtoServiceItems []MTOServiceItem `json:"mtoServiceItems"`
	}{

		MtoServiceItems: m.mtoServiceItemsField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this create m t o shipment
func (m *CreateMTOShipment) Validate(formats strfmt.Registry) error {
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

	if err := m.validateMtoServiceItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickupAddress(formats); err != nil {
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

func (m *CreateMTOShipment) validateAgents(formats strfmt.Registry) error {

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

func (m *CreateMTOShipment) validateDestinationAddress(formats strfmt.Registry) error {

	if err := validate.Required("destinationAddress", "body", m.DestinationAddress); err != nil {
		return err
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

func (m *CreateMTOShipment) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if err := validate.Required("moveTaskOrderID", "body", m.MoveTaskOrderID); err != nil {
		return err
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateMTOShipment) validateMtoServiceItems(formats strfmt.Registry) error {

	if swag.IsZero(m.MtoServiceItems()) { // not required
		return nil
	}

	for i := 0; i < len(m.MtoServiceItems()); i++ {

		if err := m.mtoServiceItemsField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mtoServiceItems" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *CreateMTOShipment) validatePickupAddress(formats strfmt.Registry) error {

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

func (m *CreateMTOShipment) validateRequestedPickupDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedPickupDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedPickupDate", "body", "date", m.RequestedPickupDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateMTOShipment) validateShipmentType(formats strfmt.Registry) error {

	if err := m.ShipmentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shipmentType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateMTOShipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateMTOShipment) UnmarshalBinary(b []byte) error {
	var res CreateMTOShipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}