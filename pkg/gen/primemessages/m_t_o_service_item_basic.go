// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOServiceItemBasic Describes a basic service item subtype of a MTOServiceItem.
// swagger:model MTOServiceItemBasic
type MTOServiceItemBasic struct {
	eTagField string

	idField strfmt.UUID

	moveTaskOrderIdField strfmt.UUID

	mtoShipmentIdField strfmt.UUID

	reServiceIdField strfmt.UUID

	reServiceNameField string

	rejectionReasonField *string

	statusField MTOServiceItemStatus

	// re service code
	// Required: true
	ReServiceCode ReServiceCode `json:"reServiceCode"`
}

// ETag gets the e tag of this subtype
func (m *MTOServiceItemBasic) ETag() string {
	return m.eTagField
}

// SetETag sets the e tag of this subtype
func (m *MTOServiceItemBasic) SetETag(val string) {
	m.eTagField = val
}

// ID gets the id of this subtype
func (m *MTOServiceItemBasic) ID() strfmt.UUID {
	return m.idField
}

// SetID sets the id of this subtype
func (m *MTOServiceItemBasic) SetID(val strfmt.UUID) {
	m.idField = val
}

// ModelType gets the model type of this subtype
func (m *MTOServiceItemBasic) ModelType() MTOServiceItemModelType {
	return "MTOServiceItemBasic"
}

// SetModelType sets the model type of this subtype
func (m *MTOServiceItemBasic) SetModelType(val MTOServiceItemModelType) {

}

// MoveTaskOrderID gets the move task order ID of this subtype
func (m *MTOServiceItemBasic) MoveTaskOrderID() strfmt.UUID {
	return m.moveTaskOrderIdField
}

// SetMoveTaskOrderID sets the move task order ID of this subtype
func (m *MTOServiceItemBasic) SetMoveTaskOrderID(val strfmt.UUID) {
	m.moveTaskOrderIdField = val
}

// MtoShipmentID gets the mto shipment ID of this subtype
func (m *MTOServiceItemBasic) MtoShipmentID() strfmt.UUID {
	return m.mtoShipmentIdField
}

// SetMtoShipmentID sets the mto shipment ID of this subtype
func (m *MTOServiceItemBasic) SetMtoShipmentID(val strfmt.UUID) {
	m.mtoShipmentIdField = val
}

// ReServiceID gets the re service ID of this subtype
func (m *MTOServiceItemBasic) ReServiceID() strfmt.UUID {
	return m.reServiceIdField
}

// SetReServiceID sets the re service ID of this subtype
func (m *MTOServiceItemBasic) SetReServiceID(val strfmt.UUID) {
	m.reServiceIdField = val
}

// ReServiceName gets the re service name of this subtype
func (m *MTOServiceItemBasic) ReServiceName() string {
	return m.reServiceNameField
}

// SetReServiceName sets the re service name of this subtype
func (m *MTOServiceItemBasic) SetReServiceName(val string) {
	m.reServiceNameField = val
}

// RejectionReason gets the rejection reason of this subtype
func (m *MTOServiceItemBasic) RejectionReason() *string {
	return m.rejectionReasonField
}

// SetRejectionReason sets the rejection reason of this subtype
func (m *MTOServiceItemBasic) SetRejectionReason(val *string) {
	m.rejectionReasonField = val
}

// Status gets the status of this subtype
func (m *MTOServiceItemBasic) Status() MTOServiceItemStatus {
	return m.statusField
}

// SetStatus sets the status of this subtype
func (m *MTOServiceItemBasic) SetStatus(val MTOServiceItemStatus) {
	m.statusField = val
}

// ReServiceCode gets the re service code of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *MTOServiceItemBasic) UnmarshalJSON(raw []byte) error {
	var data struct {

		// re service code
		// Required: true
		ReServiceCode ReServiceCode `json:"reServiceCode"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ETag string `json:"eTag,omitempty"`

		ID strfmt.UUID `json:"id,omitempty"`

		ModelType MTOServiceItemModelType `json:"modelType"`

		MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

		MtoShipmentID strfmt.UUID `json:"mtoShipmentID,omitempty"`

		ReServiceID strfmt.UUID `json:"reServiceID,omitempty"`

		ReServiceName string `json:"reServiceName,omitempty"`

		RejectionReason *string `json:"rejectionReason,omitempty"`

		Status MTOServiceItemStatus `json:"status,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result MTOServiceItemBasic

	result.eTagField = base.ETag

	result.idField = base.ID

	if base.ModelType != result.ModelType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid modelType value: %q", base.ModelType)
	}

	result.moveTaskOrderIdField = base.MoveTaskOrderID

	result.mtoShipmentIdField = base.MtoShipmentID

	result.reServiceIdField = base.ReServiceID

	result.reServiceNameField = base.ReServiceName

	result.rejectionReasonField = base.RejectionReason

	result.statusField = base.Status

	result.ReServiceCode = data.ReServiceCode

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m MTOServiceItemBasic) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// re service code
		// Required: true
		ReServiceCode ReServiceCode `json:"reServiceCode"`
	}{

		ReServiceCode: m.ReServiceCode,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ETag string `json:"eTag,omitempty"`

		ID strfmt.UUID `json:"id,omitempty"`

		ModelType MTOServiceItemModelType `json:"modelType"`

		MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

		MtoShipmentID strfmt.UUID `json:"mtoShipmentID,omitempty"`

		ReServiceID strfmt.UUID `json:"reServiceID,omitempty"`

		ReServiceName string `json:"reServiceName,omitempty"`

		RejectionReason *string `json:"rejectionReason,omitempty"`

		Status MTOServiceItemStatus `json:"status,omitempty"`
	}{

		ETag: m.ETag(),

		ID: m.ID(),

		ModelType: m.ModelType(),

		MoveTaskOrderID: m.MoveTaskOrderID(),

		MtoShipmentID: m.MtoShipmentID(),

		ReServiceID: m.ReServiceID(),

		ReServiceName: m.ReServiceName(),

		RejectionReason: m.RejectionReason(),

		Status: m.Status(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this m t o service item basic
func (m *MTOServiceItemBasic) Validate(formats strfmt.Registry) error {
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

	if err := m.validateReServiceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReServiceCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MTOServiceItemBasic) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemBasic) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrderID()) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemBasic) validateMtoShipmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.MtoShipmentID()) { // not required
		return nil
	}

	if err := validate.FormatOf("mtoShipmentID", "body", "uuid", m.MtoShipmentID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemBasic) validateReServiceID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReServiceID()) { // not required
		return nil
	}

	if err := validate.FormatOf("reServiceID", "body", "uuid", m.ReServiceID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemBasic) validateStatus(formats strfmt.Registry) error {

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

func (m *MTOServiceItemBasic) validateReServiceCode(formats strfmt.Registry) error {

	if err := m.ReServiceCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reServiceCode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOServiceItemBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOServiceItemBasic) UnmarshalBinary(b []byte) error {
	var res MTOServiceItemBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
