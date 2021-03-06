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

// UpdateMoveOrderPayload update move order payload
//
// swagger:model UpdateMoveOrderPayload
type UpdateMoveOrderPayload struct {

	// the branch that the service member belongs to
	// Required: true
	Agency Branch `json:"agency"`

	// unit is in lbs
	// Minimum: 1
	AuthorizedWeight *int64 `json:"authorizedWeight,omitempty"`

	// department indicator
	DepartmentIndicator *DeptIndicator `json:"departmentIndicator,omitempty"`

	// dependents authorized
	DependentsAuthorized *bool `json:"dependentsAuthorized,omitempty"`

	// grade
	Grade *Grade `json:"grade,omitempty"`

	// Orders date
	//
	// The date and time that these orders were cut.
	// Required: true
	// Format: date
	IssueDate *strfmt.Date `json:"issueDate"`

	// new duty station Id
	// Required: true
	// Format: uuid
	NewDutyStationID *strfmt.UUID `json:"newDutyStationId"`

	// Orders Number
	OrdersNumber *string `json:"ordersNumber,omitempty"`

	// orders type
	// Required: true
	OrdersType OrdersType `json:"ordersType"`

	// orders type detail
	OrdersTypeDetail *OrdersTypeDetail `json:"ordersTypeDetail,omitempty"`

	// origin duty station Id
	// Required: true
	// Format: uuid
	OriginDutyStationID *strfmt.UUID `json:"originDutyStationId"`

	// Report-by date
	//
	// Report By Date
	// Required: true
	// Format: date
	ReportByDate *strfmt.Date `json:"reportByDate"`

	// SAC
	Sac *string `json:"sac,omitempty"`

	// TAC
	Tac *string `json:"tac,omitempty"`
}

// Validate validates this update move order payload
func (m *UpdateMoveOrderPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorizedWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartmentIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrade(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewDutyStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrdersType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrdersTypeDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginDutyStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportByDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateMoveOrderPayload) validateAgency(formats strfmt.Registry) error {

	if err := m.Agency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("agency")
		}
		return err
	}

	return nil
}

func (m *UpdateMoveOrderPayload) validateAuthorizedWeight(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthorizedWeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("authorizedWeight", "body", int64(*m.AuthorizedWeight), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMoveOrderPayload) validateDepartmentIndicator(formats strfmt.Registry) error {

	if swag.IsZero(m.DepartmentIndicator) { // not required
		return nil
	}

	if m.DepartmentIndicator != nil {
		if err := m.DepartmentIndicator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("departmentIndicator")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateMoveOrderPayload) validateGrade(formats strfmt.Registry) error {

	if swag.IsZero(m.Grade) { // not required
		return nil
	}

	if m.Grade != nil {
		if err := m.Grade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grade")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateMoveOrderPayload) validateIssueDate(formats strfmt.Registry) error {

	if err := validate.Required("issueDate", "body", m.IssueDate); err != nil {
		return err
	}

	if err := validate.FormatOf("issueDate", "body", "date", m.IssueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMoveOrderPayload) validateNewDutyStationID(formats strfmt.Registry) error {

	if err := validate.Required("newDutyStationId", "body", m.NewDutyStationID); err != nil {
		return err
	}

	if err := validate.FormatOf("newDutyStationId", "body", "uuid", m.NewDutyStationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMoveOrderPayload) validateOrdersType(formats strfmt.Registry) error {

	if err := m.OrdersType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ordersType")
		}
		return err
	}

	return nil
}

func (m *UpdateMoveOrderPayload) validateOrdersTypeDetail(formats strfmt.Registry) error {

	if swag.IsZero(m.OrdersTypeDetail) { // not required
		return nil
	}

	if m.OrdersTypeDetail != nil {
		if err := m.OrdersTypeDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ordersTypeDetail")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateMoveOrderPayload) validateOriginDutyStationID(formats strfmt.Registry) error {

	if err := validate.Required("originDutyStationId", "body", m.OriginDutyStationID); err != nil {
		return err
	}

	if err := validate.FormatOf("originDutyStationId", "body", "uuid", m.OriginDutyStationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMoveOrderPayload) validateReportByDate(formats strfmt.Registry) error {

	if err := validate.Required("reportByDate", "body", m.ReportByDate); err != nil {
		return err
	}

	if err := validate.FormatOf("reportByDate", "body", "date", m.ReportByDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateMoveOrderPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateMoveOrderPayload) UnmarshalBinary(b []byte) error {
	var res UpdateMoveOrderPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
