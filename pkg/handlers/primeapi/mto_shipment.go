package primeapi

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/models"

	"github.com/transcom/mymove/pkg/gen/primemessages"

	mtoshipmentops "github.com/transcom/mymove/pkg/gen/primeapi/primeoperations/mto_shipment"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/handlers/primeapi/internal/payloads"
	"github.com/transcom/mymove/pkg/services"
)

// CreateMTOShipmentHandler is the handler to create MTO shipments
type CreateMTOShipmentHandler struct {
	handlers.HandlerContext
	mtoShipmentCreator services.MTOShipmentCreator
}

// Handle creates the mto shipment
func (h CreateMTOShipmentHandler) Handle(params mtoshipmentops.CreateMTOShipmentParams) middleware.Responder {

	logger := h.LoggerFromRequest(params.HTTPRequest)

	payload := params.Body
	if payload == nil {
		logger.Error("Invalid mto shipment: params Body is nil")
		return mtoshipmentops.NewCreateMTOShipmentBadRequest().WithPayload(payloads.ClientError(handlers.BadRequestErrMessage,
			"The MTO Shipment request body cannot be empty.", h.GetTraceID()))
	}

	mtoShipment := payloads.MTOShipmentModelFromCreate(payload)

	mtoServiceItemsList, verrs := payloads.MTOServiceItemList(payload)
	if verrs != nil && verrs.HasAny() {
		logger.Error("Error validating mto service item list: ", zap.Error(verrs))

		return mtoshipmentops.NewCreateMTOShipmentUnprocessableEntity().WithPayload(payloads.ValidationError(
			"The MTO service item list is invalid.", h.GetTraceID(), nil))
	}

	//mtoShipment.MTOServiceItems = *mtoServiceItemsList
	mtoShipment, err := h.mtoShipmentCreator.CreateMTOShipment(mtoShipment, mtoServiceItemsList)
	if err != nil {
		switch err.(type) {
		case services.NotFoundError:
			logger.Error("move task order not found", zap.Error(err))
			return mtoshipmentops.NewCreateMTOShipmentNotFound().WithPayload(payloads.ClientError(handlers.NotFoundMessage, err.Error(), h.GetTraceID()))
		case services.InvalidInputError:
			logger.Error("invalid input for creating mto shipment", zap.Error(err))
			return mtoshipmentops.NewCreateMTOShipmentBadRequest().WithPayload(payloads.ClientError(handlers.BadRequestErrMessage, err.Error(), h.GetTraceID()))
		default:
			logger.Error("Error creating mto shipment: ", zap.Error(err))
			return mtoshipmentops.NewCreateMTOShipmentInternalServerError().WithPayload(payloads.InternalServerError(nil, h.GetTraceID()))
		}
	}

	returnPayload := payloads.MTOShipment(mtoShipment)
	return mtoshipmentops.NewCreateMTOShipmentOK().WithPayload(returnPayload)
}

// UpdateMTOShipmentModel checks that only the fields that can be updated were passed into the payload,
// then grabs the model
func UpdateMTOShipmentModel(mtoShipmentID strfmt.UUID, payload *primemessages.MTOShipment) (*models.MTOShipment, *validate.Errors) {
	fieldsInError := validate.NewErrors()
	// this is the value that gets generated by uuid.FromStringOrNil("") for tests:
	nilUUID := strfmt.UUID(uuid.Nil.String())

	if payload.ID != "" && payload.ID != nilUUID && payload.ID != mtoShipmentID {
		fieldsInError.Add("id", "value does not agree with mtoShipmentID in path - omit from body or correct")
	}
	payload.ID = mtoShipmentID // set the ID from the path into the body for use w/ the model

	if payload.MoveTaskOrderID != "" && payload.MoveTaskOrderID != nilUUID {
		fieldsInError.Add("moveTaskOrderID", "cannot be updated")
	}
	createdAt := time.Time(payload.CreatedAt)
	if !createdAt.IsZero() {
		fieldsInError.Add("createdAt", "cannot be updated")
	}
	updatedAt := time.Time(payload.UpdatedAt)
	if !updatedAt.IsZero() {
		fieldsInError.Add("updatedAt", "cannot be manually modified - updated automatically")
	}
	primeEstimatedWeightRecordedDate := time.Time(payload.PrimeEstimatedWeightRecordedDate)
	if !primeEstimatedWeightRecordedDate.IsZero() {
		fieldsInError.Add("primeEstimatedWeightRecordedDate", "cannot be manually modified - updated automatically")
	}
	requiredDeliveryDate := time.Time(payload.RequiredDeliveryDate)
	if !requiredDeliveryDate.IsZero() {
		fieldsInError.Add("requiredDeliveryDate", "cannot be manually modified - updated automatically")
	}
	approvedDate := time.Time(payload.ApprovedDate)
	if !approvedDate.IsZero() {
		fieldsInError.Add("approvedDate", "cannot be manually modified - updated automatically with status change")
	}
	if payload.Status != "" {
		fieldsInError.Add("status", "cannot be updated")
	}
	if payload.RejectionReason != nil {
		fieldsInError.Add("rejectionReason", "cannot be updated")
	}
	if payload.CustomerRemarks != nil {
		fieldsInError.Add("customerRemarks", "cannot be updated")
	}

	if payload.Agents != nil {
		var mtoShipmentIDErr = false
		var createdAtErr = false
		var updatedAtErr = false

		for _, agent := range payload.Agents {
			if agent.MtoShipmentID != "" && agent.MtoShipmentID != nilUUID && agent.MtoShipmentID != payload.ID &&
				!mtoShipmentIDErr {
				fieldsInError.Add("mtoShipmentID", "cannot be updated for agents")
				mtoShipmentIDErr = true
			}
			createdAt := time.Time(agent.CreatedAt)
			if !createdAt.IsZero() {
				fieldsInError.Add("createdAt", "cannot be updated for agents")
				createdAtErr = true
			}
			updatedAt := time.Time(agent.UpdatedAt)
			if !updatedAt.IsZero() {
				fieldsInError.Add("updatedAt", "cannot be manually modified for agents")
				updatedAtErr = true
			}

			if mtoShipmentIDErr && createdAtErr && updatedAtErr {
				break // we've found all the errors we're gonna find here
			}
		}
	}

	if fieldsInError.HasAny() {
		return nil, fieldsInError
	}

	mtoShipment := payloads.MTOShipmentModel(payload)

	return mtoShipment, nil
}

// UpdateMTOShipmentHandler is the handler to update MTO shipments
type UpdateMTOShipmentHandler struct {
	handlers.HandlerContext
	mtoShipmentUpdater services.MTOShipmentUpdater
}

// Handle handler that updates a mto shipment
func (h UpdateMTOShipmentHandler) Handle(params mtoshipmentops.UpdateMTOShipmentParams) middleware.Responder {
	logger := h.LoggerFromRequest(params.HTTPRequest)

	mtoShipmentID := uuid.FromStringOrNil(params.MtoShipmentID.String())

	mtoAvailableToPrime, err := h.mtoShipmentUpdater.MTOAvailableToPrime(mtoShipmentID)
	if err != nil {
		logger.Error("primeapi.UpdateMTOShipmentHandler error - MTO is not available to prime", zap.Error(err))
		switch e := err.(type) {
		case services.NotFoundError:
			return mtoshipmentops.NewUpdateMTOShipmentNotFound().WithPayload(payloads.ClientError(handlers.NotFoundMessage, e.Error(), h.GetTraceID()))
		default:
			return mtoshipmentops.NewUpdateMTOShipmentInternalServerError().WithPayload(payloads.InternalServerError(nil, h.GetTraceID()))
		}
	}

	if mtoAvailableToPrime {
		mtoShipment, fieldErrs := UpdateMTOShipmentModel(params.MtoShipmentID, params.Body)
		if fieldErrs != nil {
			logger.Error("primeapi.UpdateMTOShipmentHandler error - extra fields in request", zap.Error(fieldErrs))

			errPayload := payloads.ValidationError("Fields that cannot be updated found in input",
				h.GetTraceID(), fieldErrs)

			return mtoshipmentops.NewUpdateMTOShipmentUnprocessableEntity().WithPayload(errPayload)
		}

		eTag := params.IfMatch
		logger.Info("primeapi.UpdateMTOShipmentHandler info", zap.String("pointOfContact", params.Body.PointOfContact))
		mtoShipment, err := h.mtoShipmentUpdater.UpdateMTOShipment(mtoShipment, eTag)
		if err != nil {
			logger.Error("primeapi.UpdateMTOShipmentHandler error", zap.Error(err))
			switch e := err.(type) {
			case services.NotFoundError:
				return mtoshipmentops.NewUpdateMTOShipmentNotFound().WithPayload(payloads.ClientError(handlers.NotFoundMessage, err.Error(), h.GetTraceID()))
			case services.InvalidInputError:
				payload := payloads.ValidationError(err.Error(), h.GetTraceID(), e.ValidationErrors)
				return mtoshipmentops.NewUpdateMTOShipmentUnprocessableEntity().WithPayload(payload)
			case services.PreconditionFailedError:
				return mtoshipmentops.NewUpdateMTOShipmentPreconditionFailed().WithPayload(payloads.ClientError(handlers.PreconditionErrMessage, err.Error(), h.GetTraceID()))
			default:
				return mtoshipmentops.NewUpdateMTOShipmentInternalServerError().WithPayload(payloads.InternalServerError(nil, h.GetTraceID()))
			}
		}
		mtoShipmentPayload := payloads.MTOShipment(mtoShipment)
		return mtoshipmentops.NewUpdateMTOShipmentOK().WithPayload(mtoShipmentPayload)
	}
	return mtoshipmentops.NewUpdateMTOShipmentInternalServerError().WithPayload(payloads.InternalServerError(nil, h.GetTraceID()))
}
