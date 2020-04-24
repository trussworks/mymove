package primeapi

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/transcom/mymove/pkg/services"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/handlers/primeapi/internal/payloads"
	"github.com/transcom/mymove/pkg/models"

	movetaskorderops "github.com/transcom/mymove/pkg/gen/primeapi/primeoperations/move_task_order"
	"github.com/transcom/mymove/pkg/gen/primemessages"
	"github.com/transcom/mymove/pkg/handlers"
)

// FetchMTOUpdatesHandler lists move task orders with the option to filter since a particular date
type FetchMTOUpdatesHandler struct {
	handlers.HandlerContext
}

// Handle fetches all move task orders with the option to filter since a particular date
func (h FetchMTOUpdatesHandler) Handle(params movetaskorderops.FetchMTOUpdatesParams) middleware.Responder {
	logger := h.LoggerFromRequest(params.HTTPRequest)

	var mtos models.MoveTaskOrders

	query := h.DB().Where("is_available_to_prime = ?", true).Eager(
		"PaymentRequests",
		"MTOServiceItems",
		"MTOServiceItems.ReService",
		"MTOServiceItems.Dimensions",
		"MTOServiceItems.CustomerContacts",
		"MTOShipments",
		"MTOShipments.DestinationAddress",
		"MTOShipments.PickupAddress",
		"MTOShipments.SecondaryDeliveryAddress",
		"MTOShipments.SecondaryPickupAddress",
		"MTOShipments.MTOAgents",
		"MoveOrder",
		"MoveOrder.Customer",
		"MoveOrder.Entitlement")
	if params.Since != nil {
		since := time.Unix(*params.Since, 0)
		query = query.Where("updated_at > ?", since)
	}

	err := query.All(&mtos)

	if err != nil {
		logger.Error("Unable to fetch records:", zap.Error(err))
		return movetaskorderops.NewFetchMTOUpdatesInternalServerError()
	}

	payload := payloads.MoveTaskOrders(&mtos)

	return movetaskorderops.NewFetchMTOUpdatesOK().WithPayload(payload)
}

// UpdateMTOPostCounselingInformationHandler updates the move task order with post-counseling information
type UpdateMTOPostCounselingInformationHandler struct {
	handlers.HandlerContext
	services.Fetcher
	services.MoveTaskOrderUpdater
}

// Handle updates to move task order post-counseling
func (h UpdateMTOPostCounselingInformationHandler) Handle(params movetaskorderops.UpdateMTOPostCounselingInformationParams) middleware.Responder {
	logger := h.LoggerFromRequest(params.HTTPRequest)
	mtoID := uuid.FromStringOrNil(params.MoveTaskOrderID)
	eTag := params.IfMatch
	logger.Info("primeapi.UpdateMTOShipmentHandler info", zap.String("pointOfContact", params.Body.PointOfContact))

	mto, err := h.MoveTaskOrderUpdater.UpdatePostCounselingInfo(mtoID, params.Body, eTag)
	if err != nil {
		logger.Error("primeapi.UpdateMTOPostCounselingInformation error", zap.Error(err))
		switch err.(type) {
		case services.NotFoundError:
			return movetaskorderops.NewUpdateMTOPostCounselingInformationNotFound()
		case services.PreconditionFailedError:
			return movetaskorderops.NewUpdateMTOPostCounselingInformationPreconditionFailed().WithPayload(&primemessages.Error{Message: handlers.FmtString(err.Error())})
		case services.InvalidInputError:
			return movetaskorderops.NewUpdateMTOPostCounselingInformationUnprocessableEntity()
		default:
			return movetaskorderops.NewUpdateMTOPostCounselingInformationInternalServerError()
		}
	}
	mtoPayload := payloads.MoveTaskOrder(mto)
	return movetaskorderops.NewUpdateMTOPostCounselingInformationOK().WithPayload(mtoPayload)
}

// CreateMoveTaskOrderHandler creates a move task order
type CreateMoveTaskOrderHandler struct {
	handlers.HandlerContext
	services.CustomerFetcher
	services.MoveTaskOrderCreator
}

// Handle updates to move task order post-counseling
func (h CreateMoveTaskOrderHandler) Handle(params movetaskorderops.CreateMoveTaskOrderParams) middleware.Responder {
	logger := h.LoggerFromRequest(params.HTTPRequest)
	payload := params.Body

	// If customer id was provided, check if customer exists
	customerIDString := payload.MoveOrder.CustomerID.String()
	customerID, err := uuid.FromString(customerIDString)
	if err != nil {
		logger.Error("Invalid customer: params CustomerID cannot be converted to a UUID",
			zap.String("CustomerID", customerIDString), zap.Error(err))
		return movetaskorderops.NewCreateMoveTaskOrderBadRequest()
	}

	customer, err := h.FetchCustomer(customerID)
	if err != nil {
		logger.Error("Customer fetch error", zap.Error(err))
		switch err {
		case sql.ErrNoRows:
			return movetaskorderops.NewCreateMoveTaskOrderBadRequest()
		default:
			return movetaskorderops.NewCreateMoveTaskOrderInternalServerError()
		}
	}

	fmt.Println("\n\n >>", *customer.FirstName, *customer.LastName)

	fmt.Println("\n\n --")
	return movetaskorderops.NewCreateMoveTaskOrderCreated()

}
