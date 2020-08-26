package primeapi

import (
	"fmt"
	"time"

	mtoshipment "github.com/transcom/mymove/pkg/services/mto_shipment"

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
	mtoShipmentCreator     services.MTOShipmentCreator
	mtoAvailabilityChecker services.MoveTaskOrderChecker
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

	for _, mtoServiceItem := range params.Body.MtoServiceItems() {
		// restrict creation to a list
		if _, ok := AllowedServiceItemMap[mtoServiceItem.ModelType()]; !ok {
			// throw error if modelType() not on the list
			mapKeys := GetMapKeys(AllowedServiceItemMap)
			detailErr := fmt.Sprintf("MTOServiceItem modelType() not allowed: %s ", mtoServiceItem.ModelType())
			verrs := validate.NewErrors()
			verrs.Add("modelType", fmt.Sprintf("allowed modelType() %v", mapKeys))

			logger.Error("primeapi.CreateMTOShipmentHandler error", zap.Error(verrs))
			return mtoshipmentops.NewCreateMTOShipmentUnprocessableEntity().WithPayload(payloads.ValidationError(
				detailErr, h.GetTraceID(), verrs))
		}
	}

	mtoShipment := payloads.MTOShipmentModelFromCreate(payload)
	mtoShipment.Status = models.MTOShipmentStatusSubmitted
	mtoServiceItemsList, verrs := payloads.MTOServiceItemModelListFromCreate(payload)

	if verrs != nil && verrs.HasAny() {
		logger.Error("Error validating mto service item list: ", zap.Error(verrs))

		return mtoshipmentops.NewCreateMTOShipmentUnprocessableEntity().WithPayload(payloads.ValidationError(
			"The MTO service item list is invalid.", h.GetTraceID(), nil))
	}

	moveTaskOrderID := uuid.FromStringOrNil(payload.MoveTaskOrderID.String())
	mtoAvailableToPrime, err := h.mtoAvailabilityChecker.MTOAvailableToPrime(moveTaskOrderID)

	if mtoAvailableToPrime {
		mtoShipment, err = h.mtoShipmentCreator.CreateMTOShipment(mtoShipment, mtoServiceItemsList)
	} else if err == nil {
		logger.Error("primeapi.CreateMTOShipmentHandler error - MTO is not available to Prime")
		return mtoshipmentops.NewCreateMTOShipmentNotFound().WithPayload(payloads.ClientError(
			handlers.NotFoundMessage, fmt.Sprintf("id: %s not found for moveTaskOrder", moveTaskOrderID), h.GetTraceID()))
	}

	// Could be the error from MTOAvailableToPrime or CreateMTOShipment:
	if err != nil {
		logger.Error("primeapi.CreateMTOShipmentHandler error", zap.Error(err))
		switch e := err.(type) {
		case services.NotFoundError:
			return mtoshipmentops.NewCreateMTOShipmentNotFound().WithPayload(payloads.ClientError(handlers.NotFoundMessage, err.Error(), h.GetTraceID()))
		case services.InvalidInputError:
			return mtoshipmentops.NewCreateMTOShipmentUnprocessableEntity().WithPayload(payloads.ValidationError(err.Error(), h.GetTraceID(), e.ValidationErrors))
		case services.QueryError:
			if e.Unwrap() != nil {
				// If you can unwrap, log the internal error (usually a pq error) for better debugging
				logger.Error("primeapi.CreateMTOShipmentHandler query error", zap.Error(e.Unwrap()))
			}
			return mtoshipmentops.NewCreateMTOShipmentInternalServerError().WithPayload(payloads.InternalServerError(nil, h.GetTraceID()))
		default:
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

	mtoAvailableToPrime, err := h.mtoShipmentUpdater.MTOShipmentsMTOAvailableToPrime(mtoShipmentID)
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

		var dbShipment models.MTOShipment
		err := h.DB().Eager("PickupAddress", "DestinationAddress").Find(&dbShipment, params.MtoShipmentID)
		if err != nil {
			return mtoshipmentops.NewUpdateMTOShipmentNotFound().WithPayload(payloads.ClientError(handlers.NotFoundMessage, err.Error(), h.GetTraceID()))
		}

		mtoShipment, validationErrs := h.checkPrimeValidationsOnModel(mtoShipment, &dbShipment)
		if validationErrs != nil && validationErrs.HasAny() {
			logger.Error("primeapi.UpdateMTOShipmentHandler error - extra fields in request", zap.Error(validationErrs))

			errPayload := payloads.ValidationError("Fields that cannot be updated found in input",
				h.GetTraceID(), validationErrs)

			return mtoshipmentops.NewUpdateMTOShipmentUnprocessableEntity().WithPayload(errPayload)
		}

		eTag := params.IfMatch
		logger.Info("primeapi.UpdateMTOShipmentHandler info", zap.String("pointOfContact", params.Body.PointOfContact))
		mtoShipment, err = h.mtoShipmentUpdater.UpdateMTOShipment(mtoShipment, eTag)
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

func (h UpdateMTOShipmentHandler) checkPrimeValidationsOnModel(mtoShipment *models.MTOShipment, dbShipment *models.MTOShipment) (*models.MTOShipment, *validate.Errors) {
	verrs := validate.NewErrors()
	if mtoShipment.RequestedPickupDate != nil {
		requestedPickupDate := mtoShipment.RequestedPickupDate
		// Prime cannot edit the customer's requestedPickupDate
		// if requestedPickupDate isn't valid then return InvalidInputError
		if !requestedPickupDate.Equal(*dbShipment.RequestedPickupDate) {
			verrs.Add("requestedPickupDate", "must match what customer has requested")
		}
		mtoShipment.RequestedPickupDate = requestedPickupDate
	}

	if mtoShipment.PrimeEstimatedWeight != nil {
		if dbShipment.PrimeEstimatedWeight != nil {
			verrs.Add("primeEstimatedWeight", "cannot be updated after initial estimation")
		}
		now := time.Now()
		if mtoShipment.ApprovedDate != nil {
			err := validatePrimeEstimatedWeightRecordedDate(now, *dbShipment.ScheduledPickupDate, *dbShipment.ApprovedDate)
			if err != nil {
				verrs.Add("primeEstimatedWeight", "the time period for updating the estimated weight for a shipment has expired, please contact the TOO directly to request updates to this shipment’s estimated weight")
				verrs.Add("primeEstimatedWeight", err.Error())
			}
		}
		mtoShipment.PrimeEstimatedWeightRecordedDate = &now
	}

	// Updated based on existing fields that may have been updated:
	if mtoShipment.ScheduledPickupDate != nil && mtoShipment.PrimeEstimatedWeight != nil {
		pickupAddress := dbShipment.PickupAddress
		if mtoShipment.PickupAddress != nil {
			pickupAddress = mtoShipment.PickupAddress
		}
		destinationAddress := dbShipment.DestinationAddress
		if mtoShipment.DestinationAddress != nil {
			destinationAddress = mtoShipment.DestinationAddress
		}
		requiredDeliveryDate, err := mtoshipment.CalculateRequiredDeliveryDate(h.Planner(), h.DB(), *pickupAddress,
			*destinationAddress, *mtoShipment.ScheduledPickupDate, mtoShipment.PrimeEstimatedWeight.Int())
		if err != nil {
			verrs.Add("requiredDeliveryDate", err.Error())
		}
		mtoShipment.RequiredDeliveryDate = requiredDeliveryDate
	}

	if len(mtoShipment.MTOAgents) > 0 {
		if len(dbShipment.MTOAgents) < len(mtoShipment.MTOAgents) {
			verrs.Add("agents", "cannot add MTO agents to a shipment")
		}
	}
	return mtoShipment, verrs
}

func validatePrimeEstimatedWeightRecordedDate(estimatedWeightRecordedDate time.Time, scheduledPickupDate time.Time, approvedDate time.Time) error {
	approvedDaysFromScheduled := scheduledPickupDate.Sub(approvedDate).Hours() / 24
	daysFromScheduled := scheduledPickupDate.Sub(estimatedWeightRecordedDate).Hours() / 24
	if approvedDaysFromScheduled >= 10 && daysFromScheduled >= 10 {
		return nil
	}

	if (approvedDaysFromScheduled >= 3 && approvedDaysFromScheduled <= 9) && daysFromScheduled >= 3 {
		return nil
	}

	if approvedDaysFromScheduled < 3 && daysFromScheduled >= 1 {
		return nil
	}

	return services.InvalidInputError{}
}
