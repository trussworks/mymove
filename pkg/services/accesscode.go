package services

import (
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
)

// AccessCodeValidator is the service object interface for ValidateAccessCode
//go:generate mockery -name AccessCodeValidator
type AccessCodeValidator interface {
	ValidateAccessCode(code string, moveType models.SelectedMoveType) (*models.AccessCode, bool, error)
}

// AccessCodeClaimer is the service object interface for ValidateAccessCode
//go:generate mockery -name AccessCodeClaimer
type AccessCodeClaimer interface {
	ClaimAccessCode(code string, serviceMemberID uuid.UUID) (*models.AccessCode, *validate.Errors, error)
}
