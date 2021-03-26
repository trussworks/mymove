package invoice

import (
	"fmt"

	"github.com/gobuffalo/pop/v5"
	"go.uber.org/zap"

	ediResponse997 "github.com/transcom/mymove/pkg/edi/edi997"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
)

type edi997Processor struct {
	db     *pop.Connection
	logger Logger
}

// NewEDI997Processor returns a new EDI997 processor
func NewEDI997Processor(db *pop.Connection,
	logger Logger) services.EDI997Processor {

	return &edi997Processor{
		db:     db,
		logger: logger,
	}
}

//ProcessEDI997 parses an EDI 997 response and updates the payment request status
func (e *edi997Processor) ProcessEDI997(stringEDI997 string) (ediResponse997.EDI, error) {
	errString := ""

	edi997 := ediResponse997.EDI{}
	err := edi997.Parse(stringEDI997)
	if err != nil {
		errString += err.Error()
	}

	// Find the PaymentRequestID that matches the ICN
	icn := edi997.InterchangeControlEnvelope.ISA.InterchangeControlNumber
	var paymentRequest models.PaymentRequest
	err = e.db.Q().
		Join("payment_request_to_interchange_control_numbers", "payment_request_to_interchange_control_numbers.payment_request_id = payment_requests.id").
		Where("payment_request_to_interchange_control_numbers.interchange_control_number = ?", int(icn)).
		First(&paymentRequest)
	if err != nil {
		errString += fmt.Sprintf("unable to find payment request with ID: %s, %d", err.Error(), int(icn)) + "\n"
	}

	err = edi997.Validate()
	if err != nil {
		errString += err.Error()
	}

	if errString != "" {
		e.logger.Error(errString)
		return edi997, fmt.Errorf(errString)
	}

	var transactionError error
	transactionError = e.db.Transaction(func(tx *pop.Connection) error {
		paymentRequest.Status = models.PaymentRequestStatusReceivedByGex
		err = e.db.Update(&paymentRequest)
		if err != nil {
			e.logger.Error("failure updating payment request", zap.Error(err))
			return fmt.Errorf("failure updating payment request status: %w", err)
		}
		return nil
	})

	if transactionError != nil {
		return edi997, transactionError
	}

	return edi997, nil
}
