package testdatagen

import (
	"github.com/gobuffalo/pop"

	"github.com/transcom/mymove/pkg/models"
)

// MakeProofOfServiceDoc creates a single ProofOfServiceDoc.
func MakeProofOfServiceDoc(db *pop.Connection, assertions Assertions) models.ProofOfServiceDoc {

	var pr models.PaymentRequest
	prID := assertions.ProofOfServiceDoc.PaymentRequestID
	// ID is required because it must be populated for Eager saving to work.
	if isZeroUUID(prID) {
		if !isZeroUUID(assertions.PaymentRequest.ID) {
			pr = assertions.PaymentRequest
		} else {
			pr = MakePaymentRequest(db, assertions)
		}
		prID = pr.ID
	}

	posDoc := models.ProofOfServiceDoc{
		PaymentRequest:   pr,
		PaymentRequestID: prID,
	}

	// Overwrite values with those from assertions
	mergeModels(&posDoc, assertions.Document)

	mustCreate(db, &posDoc)

	return posDoc
}

// MakeDefaultProofOfServiceDocument returns a ProofOfServiceDoc with default values
func MakeDefaultProofOfServiceDocument(db *pop.Connection) models.ProofOfServiceDoc {
	return MakeProofOfServiceDoc(db, Assertions{})
}
