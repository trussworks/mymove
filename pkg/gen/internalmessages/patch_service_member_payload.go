// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchServiceMemberPayload patch service member payload
//
// swagger:model PatchServiceMemberPayload
type PatchServiceMemberPayload struct {

	// affiliation
	Affiliation *Affiliation `json:"affiliation,omitempty"`

	// Backup Mailing Address
	BackupMailingAddress *Address `json:"backup_mailing_address,omitempty"`

	// current station id
	// Format: uuid
	CurrentStationID *strfmt.UUID `json:"current_station_id,omitempty"`

	// DoD ID number
	// Max Length: 10
	// Min Length: 10
	// Pattern: ^\d{10}$
	Edipi *string `json:"edipi,omitempty"`

	// Email
	EmailIsPreferred *bool `json:"email_is_preferred,omitempty"`

	// First name
	FirstName *string `json:"first_name,omitempty"`

	// Last name
	LastName *string `json:"last_name,omitempty"`

	// Middle name
	MiddleName *string `json:"middle_name,omitempty"`

	// move id
	// Format: uuid
	MoveID strfmt.UUID `json:"move_id,omitempty"`

	// Personal Email
	// Pattern: ^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
	PersonalEmail *string `json:"personal_email,omitempty"`

	// Phone
	PhoneIsPreferred *bool `json:"phone_is_preferred,omitempty"`

	// rank
	Rank *ServiceMemberRank `json:"rank,omitempty"`

	// Residential Address
	ResidentialAddress *Address `json:"residential_address,omitempty"`

	// Alternate Phone
	// Pattern: ^[2-9]\d{2}-\d{3}-\d{4}$
	SecondaryTelephone *string `json:"secondary_telephone,omitempty"`

	// Suffix
	Suffix *string `json:"suffix,omitempty"`

	// Best Contact Phone
	// Pattern: ^[2-9]\d{2}-\d{3}-\d{4}$
	Telephone *string `json:"telephone,omitempty"`

	// user id
	// Format: uuid
	UserID strfmt.UUID `json:"user_id,omitempty"`
}

// Validate validates this patch service member payload
func (m *PatchServiceMemberPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAffiliation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupMailingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdipi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonalEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResidentialAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryTelephone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelephone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchServiceMemberPayload) validateAffiliation(formats strfmt.Registry) error {

	if swag.IsZero(m.Affiliation) { // not required
		return nil
	}

	if m.Affiliation != nil {
		if err := m.Affiliation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("affiliation")
			}
			return err
		}
	}

	return nil
}

func (m *PatchServiceMemberPayload) validateBackupMailingAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.BackupMailingAddress) { // not required
		return nil
	}

	if m.BackupMailingAddress != nil {
		if err := m.BackupMailingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_mailing_address")
			}
			return err
		}
	}

	return nil
}

func (m *PatchServiceMemberPayload) validateCurrentStationID(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentStationID) { // not required
		return nil
	}

	if err := validate.FormatOf("current_station_id", "body", "uuid", m.CurrentStationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PatchServiceMemberPayload) validateEdipi(formats strfmt.Registry) error {

	if swag.IsZero(m.Edipi) { // not required
		return nil
	}

	if err := validate.MinLength("edipi", "body", string(*m.Edipi), 10); err != nil {
		return err
	}

	if err := validate.MaxLength("edipi", "body", string(*m.Edipi), 10); err != nil {
		return err
	}

	if err := validate.Pattern("edipi", "body", string(*m.Edipi), `^\d{10}$`); err != nil {
		return err
	}

	return nil
}

func (m *PatchServiceMemberPayload) validateMoveID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveID) { // not required
		return nil
	}

	if err := validate.FormatOf("move_id", "body", "uuid", m.MoveID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PatchServiceMemberPayload) validatePersonalEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.PersonalEmail) { // not required
		return nil
	}

	if err := validate.Pattern("personal_email", "body", string(*m.PersonalEmail), `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`); err != nil {
		return err
	}

	return nil
}

func (m *PatchServiceMemberPayload) validateRank(formats strfmt.Registry) error {

	if swag.IsZero(m.Rank) { // not required
		return nil
	}

	if m.Rank != nil {
		if err := m.Rank.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rank")
			}
			return err
		}
	}

	return nil
}

func (m *PatchServiceMemberPayload) validateResidentialAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.ResidentialAddress) { // not required
		return nil
	}

	if m.ResidentialAddress != nil {
		if err := m.ResidentialAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("residential_address")
			}
			return err
		}
	}

	return nil
}

func (m *PatchServiceMemberPayload) validateSecondaryTelephone(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryTelephone) { // not required
		return nil
	}

	if err := validate.Pattern("secondary_telephone", "body", string(*m.SecondaryTelephone), `^[2-9]\d{2}-\d{3}-\d{4}$`); err != nil {
		return err
	}

	return nil
}

func (m *PatchServiceMemberPayload) validateTelephone(formats strfmt.Registry) error {

	if swag.IsZero(m.Telephone) { // not required
		return nil
	}

	if err := validate.Pattern("telephone", "body", string(*m.Telephone), `^[2-9]\d{2}-\d{3}-\d{4}$`); err != nil {
		return err
	}

	return nil
}

func (m *PatchServiceMemberPayload) validateUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.UserID) { // not required
		return nil
	}

	if err := validate.FormatOf("user_id", "body", "uuid", m.UserID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchServiceMemberPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchServiceMemberPayload) UnmarshalBinary(b []byte) error {
	var res PatchServiceMemberPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
