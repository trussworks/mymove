package supportapi

import (
	"github.com/gofrs/uuid"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/gen/primemessages"

	"github.com/transcom/mymove/pkg/gen/supportmessages"

	"github.com/transcom/mymove/pkg/handlers/supportapi/internal/payloads"

	"github.com/go-openapi/runtime/middleware"

	movetaskorderops "github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/move_task_order"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/services"
)

// UpdateMoveTaskOrderStatusHandlerFunc updates the status of a Move Task Order
type UpdateMoveTaskOrderStatusHandlerFunc struct {
	handlers.HandlerContext
	moveTaskOrderStatusUpdater services.MoveTaskOrderUpdater
}

// Handle updates the status of a MoveTaskOrder
func (h UpdateMoveTaskOrderStatusHandlerFunc) Handle(params movetaskorderops.UpdateMoveTaskOrderStatusParams) middleware.Responder {
	_, logger := h.SessionAndLoggerFromRequest(params.HTTPRequest)
	eTag := params.IfMatch

	moveTaskOrderID := uuid.FromStringOrNil(params.MoveTaskOrderID)

	mto, err := h.moveTaskOrderStatusUpdater.MakeAvailableToPrime(moveTaskOrderID, eTag)

	if err != nil {
		logger.Error("supportapi.MoveTaskOrderHandler error", zap.Error(err))
		switch err.(type) {
		case services.NotFoundError:
			return movetaskorderops.NewUpdateMoveTaskOrderStatusNotFound().WithPayload(&supportmessages.Error{Message: handlers.FmtString(err.Error())})
		case services.InvalidInputError:
			return movetaskorderops.NewUpdateMoveTaskOrderStatusBadRequest().WithPayload(&supportmessages.Error{Message: handlers.FmtString(err.Error())})
		case services.PreconditionFailedError:
			return movetaskorderops.NewUpdateMoveTaskOrderStatusPreconditionFailed().WithPayload(&supportmessages.Error{Message: handlers.FmtString(err.Error())})
		default:
			return movetaskorderops.NewUpdateMoveTaskOrderStatusInternalServerError().WithPayload(&supportmessages.Error{Message: handlers.FmtString(err.Error())})
		}
	}

	moveTaskOrderPayload := payloads.MoveTaskOrder(mto)

	return movetaskorderops.NewUpdateMoveTaskOrderStatusOK().WithPayload(moveTaskOrderPayload)
}

// GetMoveTaskOrderHandlerFunc updates the status of a Move Task Order
type GetMoveTaskOrderHandlerFunc struct {
	handlers.HandlerContext
	moveTaskOrderFetcher services.MoveTaskOrderFetcher
}

// Handle updates the status of a MoveTaskOrder
func (h GetMoveTaskOrderHandlerFunc) Handle(params movetaskorderops.GetMoveTaskOrderParams) middleware.Responder {
	logger := h.LoggerFromRequest(params.HTTPRequest)

	moveTaskOrderID := uuid.FromStringOrNil(params.MoveTaskOrderID)
	mto, err := h.moveTaskOrderFetcher.FetchMoveTaskOrder(moveTaskOrderID)
	if err != nil {
		logger.Error("primeapi.support.GetMoveTaskOrderHandler error", zap.Error(err))
		switch err.(type) {
		case services.NotFoundError:
			return movetaskorderops.NewGetMoveTaskOrderNotFound().WithPayload(&primemessages.Error{Message: handlers.FmtString(err.Error())})
		case services.InvalidInputError:
			return movetaskorderops.NewGetMoveTaskOrderBadRequest().WithPayload(&primemessages.Error{Message: handlers.FmtString(err.Error())})
		default:
			return movetaskorderops.NewGetMoveTaskOrderInternalServerError().WithPayload(&primemessages.Error{Message: handlers.FmtString(err.Error())})
		}
	}
	moveTaskOrderPayload := payloads.MoveTaskOrder(mto)
	return movetaskorderops.NewGetMoveTaskOrderOK().WithPayload(moveTaskOrderPayload)
}
