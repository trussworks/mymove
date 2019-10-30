// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MoveTaskOrder move task order
// swagger:model moveTaskOrder
type MoveTaskOrder struct {

	// code
	Code string `json:"code,omitempty"`

	// created at
	// Format: date
	CreatedAt strfmt.Date `json:"createdAt,omitempty"`

	// customer Id
	// Format: uuid
	CustomerID strfmt.UUID `json:"customerId,omitempty"`

	// deleted at
	// Format: date
	DeletedAt strfmt.Date `json:"deletedAt,omitempty"`

	// destination address
	DestinationAddress *Address `json:"destinationAddress,omitempty"`

	// destination duty station
	// Format: uuid
	DestinationDutyStation strfmt.UUID `json:"destinationDutyStation,omitempty"`

	// destination p p s o
	// Format: uuid
	DestinationPPSO strfmt.UUID `json:"destinationPPSO,omitempty"`

	// entitlements
	Entitlements *Entitlements `json:"entitlements,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// move date
	// Format: date
	MoveDate strfmt.Date `json:"moveDate,omitempty"`

	// move ID
	// Format: uuid
	MoveID strfmt.UUID `json:"moveID,omitempty"`

	// move task orders type
	// Enum: [NON_TEMPORARY_STORAGE PRIME]
	MoveTaskOrdersType string `json:"moveTaskOrdersType,omitempty"`

	// origin duty station
	// Format: uuid
	OriginDutyStation strfmt.UUID `json:"originDutyStation,omitempty"`

	// origin p p s o
	// Format: uuid
	OriginPPSO strfmt.UUID `json:"originPPSO,omitempty"`

	// pickup address
	PickupAddress *Address `json:"pickupAddress,omitempty"`

	// remarks
	Remarks string `json:"remarks,omitempty"`

	// requested pickup date
	// Format: date
	RequestedPickupDate strfmt.Date `json:"requestedPickupDate,omitempty"`

	// service items
	ServiceItems []*ServiceItem `json:"serviceItems"`

	// status
	// Enum: [DRAFT APPROVED REJECTED SUBMITTED]
	Status string `json:"status,omitempty"`

	// updated at
	// Format: date
	UpdatedAt strfmt.Date `json:"updatedAt,omitempty"`
}

// Validate validates this move task order
func (m *MoveTaskOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationPPSO(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTaskOrdersType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginPPSO(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickupAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedPickupDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *MoveTaskOrder) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateCustomerID(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerID) { // not required
		return nil
	}

	if err := validate.FormatOf("customerId", "body", "uuid", m.CustomerID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateDeletedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.DeletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("deletedAt", "body", "date", m.DeletedAt.String(), formats); err != nil {
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

func (m *MoveTaskOrder) validateDestinationPPSO(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationPPSO) { // not required
		return nil
	}

	if err := validate.FormatOf("destinationPPSO", "body", "uuid", m.DestinationPPSO.String(), formats); err != nil {
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

func (m *MoveTaskOrder) validateMoveDate(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("moveDate", "body", "date", m.MoveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateMoveID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveID", "body", "uuid", m.MoveID.String(), formats); err != nil {
		return err
	}

	return nil
}

var moveTaskOrderTypeMoveTaskOrdersTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NON_TEMPORARY_STORAGE","PRIME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moveTaskOrderTypeMoveTaskOrdersTypePropEnum = append(moveTaskOrderTypeMoveTaskOrdersTypePropEnum, v)
	}
}

const (

	// MoveTaskOrderMoveTaskOrdersTypeNONTEMPORARYSTORAGE captures enum value "NON_TEMPORARY_STORAGE"
	MoveTaskOrderMoveTaskOrdersTypeNONTEMPORARYSTORAGE string = "NON_TEMPORARY_STORAGE"

	// MoveTaskOrderMoveTaskOrdersTypePRIME captures enum value "PRIME"
	MoveTaskOrderMoveTaskOrdersTypePRIME string = "PRIME"
)

// prop value enum
func (m *MoveTaskOrder) validateMoveTaskOrdersTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, moveTaskOrderTypeMoveTaskOrdersTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MoveTaskOrder) validateMoveTaskOrdersType(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrdersType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMoveTaskOrdersTypeEnum("moveTaskOrdersType", "body", m.MoveTaskOrdersType); err != nil {
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

func (m *MoveTaskOrder) validateOriginPPSO(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginPPSO) { // not required
		return nil
	}

	if err := validate.FormatOf("originPPSO", "body", "uuid", m.OriginPPSO.String(), formats); err != nil {
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

func (m *MoveTaskOrder) validateServiceItems(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceItems) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceItems); i++ {
		if swag.IsZero(m.ServiceItems[i]) { // not required
			continue
		}

		if m.ServiceItems[i] != nil {
			if err := m.ServiceItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var moveTaskOrderTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DRAFT","APPROVED","REJECTED","SUBMITTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moveTaskOrderTypeStatusPropEnum = append(moveTaskOrderTypeStatusPropEnum, v)
	}
}

const (

	// MoveTaskOrderStatusDRAFT captures enum value "DRAFT"
	MoveTaskOrderStatusDRAFT string = "DRAFT"

	// MoveTaskOrderStatusAPPROVED captures enum value "APPROVED"
	MoveTaskOrderStatusAPPROVED string = "APPROVED"

	// MoveTaskOrderStatusREJECTED captures enum value "REJECTED"
	MoveTaskOrderStatusREJECTED string = "REJECTED"

	// MoveTaskOrderStatusSUBMITTED captures enum value "SUBMITTED"
	MoveTaskOrderStatusSUBMITTED string = "SUBMITTED"
)

// prop value enum
func (m *MoveTaskOrder) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, moveTaskOrderTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MoveTaskOrder) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date", m.UpdatedAt.String(), formats); err != nil {
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
