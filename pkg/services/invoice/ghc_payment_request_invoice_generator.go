package invoice

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"

	ediinvoice "github.com/transcom/mymove/pkg/edi/invoice"
	edisegment "github.com/transcom/mymove/pkg/edi/segment"
)

// GHCPaymentRequestInvoiceGenerator is a service object to turn payment requests into 858s
type GHCPaymentRequestInvoiceGenerator struct {
	DB *pop.Connection
}

const dateFormat = "060102"
const timeFormat = "1504"

// Generate method takes a payment request and returns an Invoice858C
func (g GHCPaymentRequestInvoiceGenerator) Generate(paymentRequest models.PaymentRequest, sendProductionInvoice bool) (ediinvoice.Invoice858C, error) {
	// TODO: probably need to check if the MTO is loaded on the paymentRequest that is passed in, not sure what is more in line with go standards to error out if it's not there or look it up.
	// TODO: seems ReferenceID is a *string but cannot be saved as nil, do we need to validate it's not nil here

	currentTime := time.Now()

	// TODO: interchangeControlNumber, err := icnSequencer.NextVal()
	// if err != nil {
	// 	return Invoice858C{}, errors.Wrap(err, fmt.Sprintf("Failed to get next Interchange Control Number"))
	// }
	var usageIndicator string
	if sendProductionInvoice {
		usageIndicator = "P"
	} else {
		usageIndicator = "T"
	}

	var edi858 ediinvoice.Invoice858C
	edi858.ISA = edisegment.ISA{
		AuthorizationInformationQualifier: "00", // No authorization information
		AuthorizationInformation:          "0084182369",
		SecurityInformationQualifier:      "00", // No security information
		SecurityInformation:               "_   _",
		InterchangeSenderIDQualifier:      "ZZ",
		InterchangeSenderID:               "GOVDPIBS",
		InterchangeReceiverIDQualifier:    "12",
		InterchangeReceiverID:             "8004171844",
		InterchangeDate:                   currentTime.Format(dateFormat),
		InterchangeTime:                   currentTime.Format(timeFormat),
		InterchangeControlStandards:       "U",
		InterchangeControlVersionNumber:   "00401",
		InterchangeControlNumber:          100001272,
		AcknowledgementRequested:          0,
		UsageIndicator:                    usageIndicator, // T for test, P for production
		ComponentElementSeparator:         "|",
	}
	edi858.IEA = edisegment.IEA{
		NumberOfIncludedFunctionalGroups: 1,
		InterchangeControlNumber:         100001272,
	}

	edi858.GS = edisegment.GS{
		FunctionalIdentifierCode: "SI",
		ApplicationSendersCode:   "MYMOVE",
		ApplicationReceiversCode: "8004171844",
		Date:                     currentTime.Format(dateFormat),
		Time:                     currentTime.Format(timeFormat),
		GroupControlNumber:       100001251,
		ResponsibleAgencyCode:    "X",
		Version:                  "004010",
	}

	edi858.GE = edisegment.GE{
		NumberOfTransactionSetsIncluded: 1,
		GroupControlNumber:              100001251,
	}

	bx := edisegment.BX{
		TransactionSetPurposeCode:    "00",
		TransactionMethodTypeCode:    "J",
		ShipmentMethodOfPayment:      "PP",
		ShipmentIdentificationNumber: *paymentRequest.MoveTaskOrder.ReferenceID,
		StandardCarrierAlphaCode:     "TRUS",
		ShipmentQualifier:            "4",
	}
	edi858.Header = append(edi858.Header, &bx)

	paymentRequestNumberSegment := edisegment.N9{
		ReferenceIdentificationQualifier: "CN",
		ReferenceIdentification:          paymentRequest.PaymentRequestNumber,
	}
	edi858.Header = append(edi858.Header, &paymentRequestNumberSegment)

	// Add service member details to header
	serviceMemberSegments, err := g.createServiceMemberDetailSegments(paymentRequest)
	if err != nil {
		return ediinvoice.Invoice858C{}, err
	}
	edi858.Header = append(edi858.Header, serviceMemberSegments...)

	var paymentServiceItems models.PaymentServiceItems
	error := g.DB.Q().
		Eager("MTOServiceItem.ReService").
		Where("payment_request_id = ?", paymentRequest.ID).
		All(&paymentServiceItems)
	if error != nil {
		return ediinvoice.Invoice858C{}, fmt.Errorf("Could not find payment service items: %w", error)
	}

	paymentServiceItemSegments, err := g.generatePaymentServiceItemSegments(paymentServiceItems)
	if err != nil {
		return ediinvoice.Invoice858C{}, fmt.Errorf("Could not generate payment service item segments: %w", err)
	}
	edi858.ServiceItems = append(edi858.ServiceItems, paymentServiceItemSegments...)

	return edi858, nil
}

func (g GHCPaymentRequestInvoiceGenerator) createServiceMemberDetailSegments(paymentRequest models.PaymentRequest) ([]edisegment.Segment, error) {
	serviceMemberDetails := []edisegment.Segment{}

	serviceMember := paymentRequest.MoveTaskOrder.Orders.ServiceMember
	if serviceMember.ID == uuid.Nil {
		return []edisegment.Segment{}, fmt.Errorf("no ServiceMember found for Payment Request ID: %s", paymentRequest.ID)
	}

	// name
	serviceMemberName := edisegment.N9{
		ReferenceIdentificationQualifier: "1W",
		ReferenceIdentification:          serviceMember.ReverseNameLineFormat(),
	}
	serviceMemberDetails = append(serviceMemberDetails, &serviceMemberName)

	// rank
	rank := serviceMember.Rank
	if rank == nil {
		return []edisegment.Segment{}, fmt.Errorf("no rank found for ServiceMember ID: %s Payment Request ID: %s", serviceMember.ID, paymentRequest.ID)
	}
	serviceMemberRank := edisegment.N9{
		ReferenceIdentificationQualifier: "ML",
		ReferenceIdentification:          string(*rank),
	}
	serviceMemberDetails = append(serviceMemberDetails, &serviceMemberRank)

	// branch
	branch := serviceMember.Affiliation
	if branch == nil {
		return []edisegment.Segment{}, fmt.Errorf("no branch found for ServiceMember ID: %s Payment Request ID: %s", serviceMember.ID, paymentRequest.ID)
	}
	serviceMemberBranch := edisegment.N9{
		ReferenceIdentificationQualifier: "3L",
		ReferenceIdentification:          string(*branch),
	}
	serviceMemberDetails = append(serviceMemberDetails, &serviceMemberBranch)

	return serviceMemberDetails, nil
}

func (g GHCPaymentRequestInvoiceGenerator) fetchPaymentServiceItemParam(serviceItemID uuid.UUID, key models.ServiceItemParamName) (models.PaymentServiceItemParam, error) {
	var paymentServiceItemParam models.PaymentServiceItemParam
	err := g.DB.Q().
		Join("service_item_param_keys sk", "payment_service_item_params.service_item_param_key_id = sk.id").
		Where("payment_service_item_id = ?", serviceItemID).
		Where("sk.key = ?", key).
		First(&paymentServiceItemParam)
	if err != nil {
		return models.PaymentServiceItemParam{}, fmt.Errorf("Could not lookup PaymentServiceItemParam key (%s) payment service item id (%s): %w", key, serviceItemID, err)
	}
	return paymentServiceItemParam, nil
}

func (g GHCPaymentRequestInvoiceGenerator) generatePaymentServiceItemSegments(paymentServiceItems models.PaymentServiceItems) ([]edisegment.Segment, error) {
	//Initialize empty collection of segments
	var segments []edisegment.Segment

	// Iterate over payment service items
	for idx, serviceItem := range paymentServiceItems {
		hierarchicalIDNumber := idx + 1
		// Build and put together the segments
		hlSegment := edisegment.HL{
			HierarchicalIDNumber:  strconv.Itoa(hierarchicalIDNumber), // may need to change if sending multiple payment request in a single edi
			HierarchicalLevelCode: "|",
		}

		n9Segment := edisegment.N9{
			ReferenceIdentificationQualifier: "PO",
			// pending creation of shorter identifier for payment service item
			// https://dp3.atlassian.net/browse/MB-3718
			ReferenceIdentification: serviceItem.ID.String(),
		}

		// TODO: add another n9 for SIT
		// TODO: add a L5 segment/definition

		var weight models.PaymentServiceItemParam
		// TODO: update to have a case statement as different service items may or may not have weight
		// and the distance key can differ (zip3 v zip5, and distances for SIT)
		weight, err := g.fetchPaymentServiceItemParam(serviceItem.ID, models.ServiceItemParamNameWeightBilledActual)
		if err != nil {
			return nil, err
		}
		weightFloat, err := strconv.ParseFloat(weight.Value, 64)
		if err != nil {
			return nil, fmt.Errorf("Could not parse weight for PaymentServiceItem %s: %w", serviceItem.ID, err)
		}
		distance, err := g.fetchPaymentServiceItemParam(serviceItem.ID, models.ServiceItemParamNameDistanceZip3)
		if err != nil {
			return nil, err
		}
		distanceFloat, err := strconv.ParseFloat(distance.Value, 64)
		if err != nil {
			return nil, fmt.Errorf("Could not parse Distance Zip3 for PaymentServiceItem %s: %w", serviceItem.ID, err)
		}

		l0Segment := edisegment.L0{
			LadingLineItemNumber:   hierarchicalIDNumber,
			BilledRatedAsQuantity:  distanceFloat,
			BilledRatedAsQualifier: "DM",
			Weight:                 weightFloat,
			WeightQualifier:        "B",
			WeightUnitCode:         "L",
		}

		segments = append(segments, &hlSegment, &n9Segment, &l0Segment)
	}

	return segments, nil
}
