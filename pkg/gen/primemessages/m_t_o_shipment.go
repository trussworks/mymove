// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOShipment m t o shipment
//
// swagger:model MTOShipment
type MTOShipment struct {

	// actual pickup date
	// Format: date
	ActualPickupDate strfmt.Date `json:"actualPickupDate,omitempty"`

	// agents
	Agents MTOAgents `json:"agents,omitempty"`

	// date when the shipment was given the status "APPROVED"
	// Read Only: true
	// Format: date
	ApprovedDate strfmt.Date `json:"approvedDate,omitempty"`

	// created at
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// customer remarks
	// Read Only: true
	CustomerRemarks *string `json:"customerRemarks,omitempty"`

	// destination address
	DestinationAddress *Address `json:"destinationAddress,omitempty"`

	// e tag
	// Read Only: true
	ETag string `json:"eTag,omitempty"`

	// first available delivery date
	// Format: date
	FirstAvailableDeliveryDate strfmt.Date `json:"firstAvailableDeliveryDate,omitempty"`

	// id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// move task order ID
	// Read Only: true
	// Format: uuid
	MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

	mtoServiceItemsField []MTOServiceItem

	// pickup address
	PickupAddress *Address `json:"pickupAddress,omitempty"`

	// Email or id of a contact person for this update.
	PointOfContact string `json:"pointOfContact,omitempty"`

	// prime actual weight
	PrimeActualWeight int64 `json:"primeActualWeight,omitempty"`

	// prime estimated weight
	PrimeEstimatedWeight int64 `json:"primeEstimatedWeight,omitempty"`

	// prime estimated weight recorded date
	// Read Only: true
	// Format: date
	PrimeEstimatedWeightRecordedDate strfmt.Date `json:"primeEstimatedWeightRecordedDate,omitempty"`

	// rejection reason
	// Read Only: true
	RejectionReason *string `json:"rejectionReason,omitempty"`

	// requested pickup date
	// Read Only: true
	// Format: date
	RequestedPickupDate strfmt.Date `json:"requestedPickupDate,omitempty"`

	// required delivery date
	// Read Only: true
	// Format: date
	RequiredDeliveryDate strfmt.Date `json:"requiredDeliveryDate,omitempty"`

	// scheduled pickup date
	// Format: date
	ScheduledPickupDate strfmt.Date `json:"scheduledPickupDate,omitempty"`

	// secondary delivery address
	SecondaryDeliveryAddress *Address `json:"secondaryDeliveryAddress,omitempty"`

	// secondary pickup address
	SecondaryPickupAddress *Address `json:"secondaryPickupAddress,omitempty"`

	// shipment type
	ShipmentType MTOShipmentType `json:"shipmentType,omitempty"`

	// status
	// Read Only: true
	// Enum: [APPROVED SUBMITTED REJECTED CANCELLATION_REQUESTED]
	Status string `json:"status,omitempty"`

	// updated at
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// MtoServiceItems gets the mto service items of this base type
func (m *MTOShipment) MtoServiceItems() []MTOServiceItem {
	return m.mtoServiceItemsField
}

// SetMtoServiceItems sets the mto service items of this base type
func (m *MTOShipment) SetMtoServiceItems(val []MTOServiceItem) {
	m.mtoServiceItemsField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *MTOShipment) UnmarshalJSON(raw []byte) error {
	var data struct {
		ActualPickupDate strfmt.Date `json:"actualPickupDate,omitempty"`

		Agents MTOAgents `json:"agents,omitempty"`

		ApprovedDate strfmt.Date `json:"approvedDate,omitempty"`

		CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

		CustomerRemarks *string `json:"customerRemarks,omitempty"`

		DestinationAddress *Address `json:"destinationAddress,omitempty"`

		ETag string `json:"eTag,omitempty"`

		FirstAvailableDeliveryDate strfmt.Date `json:"firstAvailableDeliveryDate,omitempty"`

		ID strfmt.UUID `json:"id,omitempty"`

		MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

		MtoServiceItems json.RawMessage `json:"mtoServiceItems"`

		PickupAddress *Address `json:"pickupAddress,omitempty"`

		PointOfContact string `json:"pointOfContact,omitempty"`

		PrimeActualWeight int64 `json:"primeActualWeight,omitempty"`

		PrimeEstimatedWeight int64 `json:"primeEstimatedWeight,omitempty"`

		PrimeEstimatedWeightRecordedDate strfmt.Date `json:"primeEstimatedWeightRecordedDate,omitempty"`

		RejectionReason *string `json:"rejectionReason,omitempty"`

		RequestedPickupDate strfmt.Date `json:"requestedPickupDate,omitempty"`

		RequiredDeliveryDate strfmt.Date `json:"requiredDeliveryDate,omitempty"`

		ScheduledPickupDate strfmt.Date `json:"scheduledPickupDate,omitempty"`

		SecondaryDeliveryAddress *Address `json:"secondaryDeliveryAddress,omitempty"`

		SecondaryPickupAddress *Address `json:"secondaryPickupAddress,omitempty"`

		ShipmentType MTOShipmentType `json:"shipmentType,omitempty"`

		Status string `json:"status,omitempty"`

		UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
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

	var result MTOShipment

	// actualPickupDate
	result.ActualPickupDate = data.ActualPickupDate

	// agents
	result.Agents = data.Agents

	// approvedDate
	result.ApprovedDate = data.ApprovedDate

	// createdAt
	result.CreatedAt = data.CreatedAt

	// customerRemarks
	result.CustomerRemarks = data.CustomerRemarks

	// destinationAddress
	result.DestinationAddress = data.DestinationAddress

	// eTag
	result.ETag = data.ETag

	// firstAvailableDeliveryDate
	result.FirstAvailableDeliveryDate = data.FirstAvailableDeliveryDate

	// id
	result.ID = data.ID

	// moveTaskOrderID
	result.MoveTaskOrderID = data.MoveTaskOrderID

	// mtoServiceItems
	result.mtoServiceItemsField = propMtoServiceItems

	// pickupAddress
	result.PickupAddress = data.PickupAddress

	// pointOfContact
	result.PointOfContact = data.PointOfContact

	// primeActualWeight
	result.PrimeActualWeight = data.PrimeActualWeight

	// primeEstimatedWeight
	result.PrimeEstimatedWeight = data.PrimeEstimatedWeight

	// primeEstimatedWeightRecordedDate
	result.PrimeEstimatedWeightRecordedDate = data.PrimeEstimatedWeightRecordedDate

	// rejectionReason
	result.RejectionReason = data.RejectionReason

	// requestedPickupDate
	result.RequestedPickupDate = data.RequestedPickupDate

	// requiredDeliveryDate
	result.RequiredDeliveryDate = data.RequiredDeliveryDate

	// scheduledPickupDate
	result.ScheduledPickupDate = data.ScheduledPickupDate

	// secondaryDeliveryAddress
	result.SecondaryDeliveryAddress = data.SecondaryDeliveryAddress

	// secondaryPickupAddress
	result.SecondaryPickupAddress = data.SecondaryPickupAddress

	// shipmentType
	result.ShipmentType = data.ShipmentType

	// status
	result.Status = data.Status

	// updatedAt
	result.UpdatedAt = data.UpdatedAt

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m MTOShipment) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		ActualPickupDate strfmt.Date `json:"actualPickupDate,omitempty"`

		Agents MTOAgents `json:"agents,omitempty"`

		ApprovedDate strfmt.Date `json:"approvedDate,omitempty"`

		CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

		CustomerRemarks *string `json:"customerRemarks,omitempty"`

		DestinationAddress *Address `json:"destinationAddress,omitempty"`

		ETag string `json:"eTag,omitempty"`

		FirstAvailableDeliveryDate strfmt.Date `json:"firstAvailableDeliveryDate,omitempty"`

		ID strfmt.UUID `json:"id,omitempty"`

		MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

		PickupAddress *Address `json:"pickupAddress,omitempty"`

		PointOfContact string `json:"pointOfContact,omitempty"`

		PrimeActualWeight int64 `json:"primeActualWeight,omitempty"`

		PrimeEstimatedWeight int64 `json:"primeEstimatedWeight,omitempty"`

		PrimeEstimatedWeightRecordedDate strfmt.Date `json:"primeEstimatedWeightRecordedDate,omitempty"`

		RejectionReason *string `json:"rejectionReason,omitempty"`

		RequestedPickupDate strfmt.Date `json:"requestedPickupDate,omitempty"`

		RequiredDeliveryDate strfmt.Date `json:"requiredDeliveryDate,omitempty"`

		ScheduledPickupDate strfmt.Date `json:"scheduledPickupDate,omitempty"`

		SecondaryDeliveryAddress *Address `json:"secondaryDeliveryAddress,omitempty"`

		SecondaryPickupAddress *Address `json:"secondaryPickupAddress,omitempty"`

		ShipmentType MTOShipmentType `json:"shipmentType,omitempty"`

		Status string `json:"status,omitempty"`

		UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
	}{

		ActualPickupDate: m.ActualPickupDate,

		Agents: m.Agents,

		ApprovedDate: m.ApprovedDate,

		CreatedAt: m.CreatedAt,

		CustomerRemarks: m.CustomerRemarks,

		DestinationAddress: m.DestinationAddress,

		ETag: m.ETag,

		FirstAvailableDeliveryDate: m.FirstAvailableDeliveryDate,

		ID: m.ID,

		MoveTaskOrderID: m.MoveTaskOrderID,

		PickupAddress: m.PickupAddress,

		PointOfContact: m.PointOfContact,

		PrimeActualWeight: m.PrimeActualWeight,

		PrimeEstimatedWeight: m.PrimeEstimatedWeight,

		PrimeEstimatedWeightRecordedDate: m.PrimeEstimatedWeightRecordedDate,

		RejectionReason: m.RejectionReason,

		RequestedPickupDate: m.RequestedPickupDate,

		RequiredDeliveryDate: m.RequiredDeliveryDate,

		ScheduledPickupDate: m.ScheduledPickupDate,

		SecondaryDeliveryAddress: m.SecondaryDeliveryAddress,

		SecondaryPickupAddress: m.SecondaryPickupAddress,

		ShipmentType: m.ShipmentType,

		Status: m.Status,

		UpdatedAt: m.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		MtoServiceItems []MTOServiceItem `json:"mtoServiceItems"`
	}{

		MtoServiceItems: m.mtoServiceItemsField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this m t o shipment
func (m *MTOShipment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActualPickupDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApprovedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstAvailableDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

	if err := m.validatePrimeEstimatedWeightRecordedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedPickupDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledPickupDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryDeliveryAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryPickupAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentType(formats); err != nil {
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

func (m *MTOShipment) validateActualPickupDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ActualPickupDate) { // not required
		return nil
	}

	if err := validate.FormatOf("actualPickupDate", "body", "date", m.ActualPickupDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateAgents(formats strfmt.Registry) error {

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

func (m *MTOShipment) validateApprovedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ApprovedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("approvedDate", "body", "date", m.ApprovedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateDestinationAddress(formats strfmt.Registry) error {

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

func (m *MTOShipment) validateFirstAvailableDeliveryDate(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstAvailableDeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("firstAvailableDeliveryDate", "body", "date", m.FirstAvailableDeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateMtoServiceItems(formats strfmt.Registry) error {

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

func (m *MTOShipment) validatePickupAddress(formats strfmt.Registry) error {

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

func (m *MTOShipment) validatePrimeEstimatedWeightRecordedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimeEstimatedWeightRecordedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("primeEstimatedWeightRecordedDate", "body", "date", m.PrimeEstimatedWeightRecordedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateRequestedPickupDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedPickupDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedPickupDate", "body", "date", m.RequestedPickupDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateRequiredDeliveryDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RequiredDeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requiredDeliveryDate", "body", "date", m.RequiredDeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateScheduledPickupDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledPickupDate) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduledPickupDate", "body", "date", m.ScheduledPickupDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateSecondaryDeliveryAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryDeliveryAddress) { // not required
		return nil
	}

	if m.SecondaryDeliveryAddress != nil {
		if err := m.SecondaryDeliveryAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryDeliveryAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MTOShipment) validateSecondaryPickupAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryPickupAddress) { // not required
		return nil
	}

	if m.SecondaryPickupAddress != nil {
		if err := m.SecondaryPickupAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryPickupAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MTOShipment) validateShipmentType(formats strfmt.Registry) error {

	if swag.IsZero(m.ShipmentType) { // not required
		return nil
	}

	if err := m.ShipmentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shipmentType")
		}
		return err
	}

	return nil
}

var mTOShipmentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APPROVED","SUBMITTED","REJECTED","CANCELLATION_REQUESTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOShipmentTypeStatusPropEnum = append(mTOShipmentTypeStatusPropEnum, v)
	}
}

const (

	// MTOShipmentStatusAPPROVED captures enum value "APPROVED"
	MTOShipmentStatusAPPROVED string = "APPROVED"

	// MTOShipmentStatusSUBMITTED captures enum value "SUBMITTED"
	MTOShipmentStatusSUBMITTED string = "SUBMITTED"

	// MTOShipmentStatusREJECTED captures enum value "REJECTED"
	MTOShipmentStatusREJECTED string = "REJECTED"

	// MTOShipmentStatusCANCELLATIONREQUESTED captures enum value "CANCELLATION_REQUESTED"
	MTOShipmentStatusCANCELLATIONREQUESTED string = "CANCELLATION_REQUESTED"
)

// prop value enum
func (m *MTOShipment) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mTOShipmentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MTOShipment) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOShipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOShipment) UnmarshalBinary(b []byte) error {
	var res MTOShipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
