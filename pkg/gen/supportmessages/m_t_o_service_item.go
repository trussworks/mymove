// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOServiceItem MTOServiceItem describes a base type of a service item. Polymorphic type. Both Move Task Orders and MTO Shipments will have MTO Service Items.
//
// swagger:discriminator MTOServiceItem modelType
type MTOServiceItem interface {
	runtime.Validatable

	// ETag identifier required to update this object
	// Read Only: true
	ETag() string
	SetETag(string)

	// ID of the service item
	// Format: uuid
	ID() strfmt.UUID
	SetID(strfmt.UUID)

	// model type
	// Required: true
	ModelType() MTOServiceItemModelType
	SetModelType(MTOServiceItemModelType)

	// ID of the associated moveTaskOrder
	// Required: true
	// Format: uuid
	MoveTaskOrderID() *strfmt.UUID
	SetMoveTaskOrderID(*strfmt.UUID)

	// ID of the associated mtoShipment
	// Format: uuid
	MtoShipmentID() strfmt.UUID
	SetMtoShipmentID(strfmt.UUID)

	// Full descriptive name of the service
	// Read Only: true
	ReServiceName() string
	SetReServiceName(string)

	// Reason the service item was rejected by the TOO
	RejectionReason() *string
	SetRejectionReason(*string)

	// status
	Status() MTOServiceItemStatus
	SetStatus(MTOServiceItemStatus)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type mTOServiceItem struct {
	eTagField string

	idField strfmt.UUID

	modelTypeField MTOServiceItemModelType

	moveTaskOrderIdField *strfmt.UUID

	mtoShipmentIdField strfmt.UUID

	reServiceNameField string

	rejectionReasonField *string

	statusField MTOServiceItemStatus
}

// ETag gets the e tag of this polymorphic type
func (m *mTOServiceItem) ETag() string {
	return m.eTagField
}

// SetETag sets the e tag of this polymorphic type
func (m *mTOServiceItem) SetETag(val string) {
	m.eTagField = val
}

// ID gets the id of this polymorphic type
func (m *mTOServiceItem) ID() strfmt.UUID {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *mTOServiceItem) SetID(val strfmt.UUID) {
	m.idField = val
}

// ModelType gets the model type of this polymorphic type
func (m *mTOServiceItem) ModelType() MTOServiceItemModelType {
	return "MTOServiceItem"
}

// SetModelType sets the model type of this polymorphic type
func (m *mTOServiceItem) SetModelType(val MTOServiceItemModelType) {
}

// MoveTaskOrderID gets the move task order ID of this polymorphic type
func (m *mTOServiceItem) MoveTaskOrderID() *strfmt.UUID {
	return m.moveTaskOrderIdField
}

// SetMoveTaskOrderID sets the move task order ID of this polymorphic type
func (m *mTOServiceItem) SetMoveTaskOrderID(val *strfmt.UUID) {
	m.moveTaskOrderIdField = val
}

// MtoShipmentID gets the mto shipment ID of this polymorphic type
func (m *mTOServiceItem) MtoShipmentID() strfmt.UUID {
	return m.mtoShipmentIdField
}

// SetMtoShipmentID sets the mto shipment ID of this polymorphic type
func (m *mTOServiceItem) SetMtoShipmentID(val strfmt.UUID) {
	m.mtoShipmentIdField = val
}

// ReServiceName gets the re service name of this polymorphic type
func (m *mTOServiceItem) ReServiceName() string {
	return m.reServiceNameField
}

// SetReServiceName sets the re service name of this polymorphic type
func (m *mTOServiceItem) SetReServiceName(val string) {
	m.reServiceNameField = val
}

// RejectionReason gets the rejection reason of this polymorphic type
func (m *mTOServiceItem) RejectionReason() *string {
	return m.rejectionReasonField
}

// SetRejectionReason sets the rejection reason of this polymorphic type
func (m *mTOServiceItem) SetRejectionReason(val *string) {
	m.rejectionReasonField = val
}

// Status gets the status of this polymorphic type
func (m *mTOServiceItem) Status() MTOServiceItemStatus {
	return m.statusField
}

// SetStatus sets the status of this polymorphic type
func (m *mTOServiceItem) SetStatus(val MTOServiceItemStatus) {
	m.statusField = val
}

// UnmarshalMTOServiceItemSlice unmarshals polymorphic slices of MTOServiceItem
func UnmarshalMTOServiceItemSlice(reader io.Reader, consumer runtime.Consumer) ([]MTOServiceItem, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []MTOServiceItem
	for _, element := range elements {
		obj, err := unmarshalMTOServiceItem(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalMTOServiceItem unmarshals polymorphic MTOServiceItem
func UnmarshalMTOServiceItem(reader io.Reader, consumer runtime.Consumer) (MTOServiceItem, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalMTOServiceItem(data, consumer)
}

func unmarshalMTOServiceItem(data []byte, consumer runtime.Consumer) (MTOServiceItem, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the modelType property.
	var getType struct {
		ModelType string `json:"modelType"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("modelType", "body", getType.ModelType); err != nil {
		return nil, err
	}

	// The value of modelType is used to determine which type to create and unmarshal the data into
	switch getType.ModelType {
	case "MTOServiceItem":
		var result mTOServiceItem
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "MTOServiceItemBasic":
		var result MTOServiceItemBasic
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "MTOServiceItemDestSIT":
		var result MTOServiceItemDestSIT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "MTOServiceItemDomesticCrating":
		var result MTOServiceItemDomesticCrating
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "MTOServiceItemOriginSIT":
		var result MTOServiceItemOriginSIT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "MTOServiceItemShuttle":
		var result MTOServiceItemShuttle
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid modelType value: %q", getType.ModelType)
}

// Validate validates this m t o service item
func (m *mTOServiceItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTaskOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtoShipmentID(formats); err != nil {
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

func (m *mTOServiceItem) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *mTOServiceItem) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if err := validate.Required("moveTaskOrderID", "body", m.MoveTaskOrderID()); err != nil {
		return err
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *mTOServiceItem) validateMtoShipmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.MtoShipmentID()) { // not required
		return nil
	}

	if err := validate.FormatOf("mtoShipmentID", "body", "uuid", m.MtoShipmentID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *mTOServiceItem) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status()) { // not required
		return nil
	}

	if err := m.Status().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}
