// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Order order
//
// swagger:model Order
type Order struct {

	// customer
	Customer *Customer `json:"customer,omitempty"`

	// ID of the Customer this Order belongs to.
	//
	// If creating a MoveTaskOrder. either an existing customerID should be provided or the nested customer object should be populated for creation.
	//
	// Format: uuid
	CustomerID *strfmt.UUID `json:"customerID,omitempty"`

	// destination duty station
	DestinationDutyStation *DutyStation `json:"destinationDutyStation,omitempty"`

	// ID of the destination duty station.
	//
	// If creating a MoveTaskOrder, this should match an existing duty station.
	//
	// Required: true
	// Format: uuid
	DestinationDutyStationID *strfmt.UUID `json:"destinationDutyStationID"`

	// Uniquely identifies the state of the Order object (but not the nested objects)
	//
	// It will change everytime the object is updated. Client should store the value.
	// Updates to this Order will require that this eTag be passed in with the If-Match header.
	//
	// Read Only: true
	ETag string `json:"eTag,omitempty"`

	// entitlement
	Entitlement *Entitlement `json:"entitlement,omitempty"`

	// ID of the Order object.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The date the orders were issued.
	// Required: true
	// Format: date
	IssueDate *strfmt.Date `json:"issueDate"`

	// ID of the military orders associated with this move.
	// Required: true
	OrderNumber *string `json:"orderNumber"`

	// orders type
	// Required: true
	OrdersType OrdersType `json:"ordersType"`

	// origin duty station
	OriginDutyStation *DutyStation `json:"originDutyStation,omitempty"`

	// ID of the origin duty station.
	//
	// If creating a MoveTaskOrder, this should match an existing duty station.
	//
	// Required: true
	// Format: uuid
	OriginDutyStationID *strfmt.UUID `json:"originDutyStationID"`

	// rank
	// Required: true
	Rank Rank `json:"rank"`

	// Date that the service member must report to the new DutyStation by.
	// Required: true
	// Format: date
	ReportByDate *strfmt.Date `json:"reportByDate"`

	// status
	// Required: true
	Status OrdersStatus `json:"status"`

	// TAC
	// Required: true
	Tac *string `json:"tac"`

	// uploaded orders
	UploadedOrders *Document `json:"uploadedOrders,omitempty"`

	// ID of the uploaded document.
	// Required: true
	// Format: uuid
	UploadedOrdersID *strfmt.UUID `json:"uploadedOrdersID"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationDutyStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrdersType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginDutyStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportByDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTac(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadedOrders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadedOrdersID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Order) validateCustomer(formats strfmt.Registry) error {

	if swag.IsZero(m.Customer) { // not required
		return nil
	}

	if m.Customer != nil {
		if err := m.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *Order) validateCustomerID(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerID) { // not required
		return nil
	}

	if err := validate.FormatOf("customerID", "body", "uuid", m.CustomerID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateDestinationDutyStation(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationDutyStation) { // not required
		return nil
	}

	if m.DestinationDutyStation != nil {
		if err := m.DestinationDutyStation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationDutyStation")
			}
			return err
		}
	}

	return nil
}

func (m *Order) validateDestinationDutyStationID(formats strfmt.Registry) error {

	if err := validate.Required("destinationDutyStationID", "body", m.DestinationDutyStationID); err != nil {
		return err
	}

	if err := validate.FormatOf("destinationDutyStationID", "body", "uuid", m.DestinationDutyStationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateEntitlement(formats strfmt.Registry) error {

	if swag.IsZero(m.Entitlement) { // not required
		return nil
	}

	if m.Entitlement != nil {
		if err := m.Entitlement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entitlement")
			}
			return err
		}
	}

	return nil
}

func (m *Order) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateIssueDate(formats strfmt.Registry) error {

	if err := validate.Required("issueDate", "body", m.IssueDate); err != nil {
		return err
	}

	if err := validate.FormatOf("issueDate", "body", "date", m.IssueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateOrderNumber(formats strfmt.Registry) error {

	if err := validate.Required("orderNumber", "body", m.OrderNumber); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateOrdersType(formats strfmt.Registry) error {

	if err := m.OrdersType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ordersType")
		}
		return err
	}

	return nil
}

func (m *Order) validateOriginDutyStation(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginDutyStation) { // not required
		return nil
	}

	if m.OriginDutyStation != nil {
		if err := m.OriginDutyStation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originDutyStation")
			}
			return err
		}
	}

	return nil
}

func (m *Order) validateOriginDutyStationID(formats strfmt.Registry) error {

	if err := validate.Required("originDutyStationID", "body", m.OriginDutyStationID); err != nil {
		return err
	}

	if err := validate.FormatOf("originDutyStationID", "body", "uuid", m.OriginDutyStationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateRank(formats strfmt.Registry) error {

	if err := m.Rank.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rank")
		}
		return err
	}

	return nil
}

func (m *Order) validateReportByDate(formats strfmt.Registry) error {

	if err := validate.Required("reportByDate", "body", m.ReportByDate); err != nil {
		return err
	}

	if err := validate.FormatOf("reportByDate", "body", "date", m.ReportByDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateStatus(formats strfmt.Registry) error {

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *Order) validateTac(formats strfmt.Registry) error {

	if err := validate.Required("tac", "body", m.Tac); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateUploadedOrders(formats strfmt.Registry) error {

	if swag.IsZero(m.UploadedOrders) { // not required
		return nil
	}

	if m.UploadedOrders != nil {
		if err := m.UploadedOrders.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadedOrders")
			}
			return err
		}
	}

	return nil
}

func (m *Order) validateUploadedOrdersID(formats strfmt.Registry) error {

	if err := validate.Required("uploadedOrdersID", "body", m.UploadedOrdersID); err != nil {
		return err
	}

	if err := validate.FormatOf("uploadedOrdersID", "body", "uuid", m.UploadedOrdersID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}