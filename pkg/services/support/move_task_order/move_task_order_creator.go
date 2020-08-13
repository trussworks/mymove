package supportmovetaskorder

import (
	"time"

	"github.com/transcom/mymove/pkg/services/support"

	"github.com/go-openapi/swag"
	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/gen/supportmessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
	"github.com/transcom/mymove/pkg/services/office_user/customer"
	"github.com/transcom/mymove/pkg/unit"
)

type moveTaskOrderCreator struct {
	db *pop.Connection
}

// InternalCreateMoveTaskOrder creates a move task order for the supportapi (internal use only, not used in production)
func (f moveTaskOrderCreator) InternalCreateMoveTaskOrder(payload supportmessages.MoveTaskOrder, logger handlers.Logger) (*models.Move, error) {
	var moveTaskOrder *models.Move
	var refID string
	if payload.MoveOrder == nil {
		return nil, services.NewQueryError("MoveTaskOrder", nil, "MoveOrder is necessary")
	}

	transactionError := f.db.Transaction(func(tx *pop.Connection) error {
		// Create or get customer
		customer, err := createOrGetCustomer(tx, customer.NewCustomerFetcher(tx), payload.MoveOrder.CustomerID.String(), payload.MoveOrder.Customer, logger)
		if err != nil {
			return err
		}

		// Create move order and entitlement
		moveOrder, err := createMoveOrder(tx, customer, payload.MoveOrder, logger)
		if err != nil {
			return err
		}

		moveTaskOrder = MoveTaskOrderModel(&payload)
		if *moveTaskOrder.ReferenceID == "" {
			refID, err = models.GenerateReferenceID(tx)
			moveTaskOrder.ReferenceID = &refID
		}
		if err != nil {
			return err
		}
		if moveTaskOrder.Locator == "" {
			moveTaskOrder.Locator = models.GenerateLocator()
		}
		if moveTaskOrder.Status == "" {
			moveTaskOrder.Status = models.MoveStatusDRAFT
		}
		moveTaskOrder.Show = swag.Bool(true)
		moveTaskOrder.Orders = *moveOrder
		moveTaskOrder.OrdersID = moveOrder.ID

		verrs, err := tx.ValidateAndCreate(moveTaskOrder)

		if verrs.Count() > 0 {
			logger.Error("supportapi.createMoveTaskOrderSupport error", zap.Error(verrs))
			return services.NewInvalidInputError(uuid.Nil, nil, verrs, "")
		} else if err != nil {
			logger.Error("supportapi.createMoveTaskOrderSupport error", zap.Error(err))
			return services.NewQueryError("MoveTaskOrder", err, "Unable to create MoveTaskOrder.")
		}
		return nil
	})

	if transactionError != nil {
		return nil, transactionError
	}

	return moveTaskOrder, nil
}

// NewInternalMoveTaskOrderCreator creates a new struct with the service dependencies
func NewInternalMoveTaskOrderCreator(db *pop.Connection) support.InternalMoveTaskOrderCreator {
	return &moveTaskOrderCreator{db}
}

// createMoveOrder creates a basic move order - this is a support function do not use in production
func createMoveOrder(tx *pop.Connection, customer *models.ServiceMember, moveOrderPayload *supportmessages.MoveOrder, logger handlers.Logger) (*models.Order, error) {
	if moveOrderPayload == nil {
		returnErr := services.NewInvalidInputError(uuid.Nil, nil, nil, "MoveOrder definition is required to create MoveTaskOrder")
		return nil, returnErr
	}

	// Move order model will also contain the entitlement
	moveOrder := MoveOrderModel(moveOrderPayload)

	// Add customer to mO
	moveOrder.ServiceMember = *customer
	moveOrder.ServiceMemberID = customer.ID

	// Creates the moveOrder and the entitlement at the same time
	verrs, err := tx.Eager().ValidateAndCreate(moveOrder)
	if verrs.Count() > 0 {
		logger.Error("supportapi.createMoveOrder error", zap.Error(verrs))
		return nil, services.NewInvalidInputError(uuid.Nil, nil, verrs, "")
	} else if err != nil {
		logger.Error("supportapi.createMoveOrder error", zap.Error(err))
		e := services.NewQueryError("MoveOrder", err, "Unable to create MoveOrder.")
		return nil, e
	}
	return moveOrder, nil
}

// createUser creates a user but this is a fake login.gov user
// this is support code only, do not use in a production case
func createUser(tx *pop.Connection, userEmail *string, logger handlers.Logger) (*models.User, error) {
	if userEmail == nil {
		defaultEmail := "generatedMTOuser@example.com"
		userEmail = &defaultEmail
	}
	id := uuid.Must(uuid.NewV4())
	user := models.User{
		LoginGovUUID:  id,
		LoginGovEmail: *userEmail,
		Active:        true,
	}
	verrs, err := tx.ValidateAndCreate(&user)
	if verrs.Count() > 0 {
		logger.Error("supportapi.createUser error", zap.Error(verrs))
		return nil, services.NewInvalidInputError(uuid.Nil, nil, verrs, "")
	} else if err != nil {
		logger.Error("supportapi.createUser error", zap.Error(err))
		e := services.NewQueryError("User", err, "Unable to create User.")
		return nil, e
	}
	return &user, nil
}

// createOrGetCustomer creates a customer or gets one if id was provided
func createOrGetCustomer(tx *pop.Connection, f services.CustomerFetcher, customerIDString string, customerBody *supportmessages.Customer, logger handlers.Logger) (*models.ServiceMember, error) {
	// If customer ID string is provided, we should find this customer
	if customerIDString != "" {
		customerID, err := uuid.FromString(customerIDString)
		// Error on bad customer id string
		if err != nil {
			returnErr := services.NewInvalidInputError(uuid.Nil, err, nil, "Invalid customerID: params CustomerID cannot be converted to a UUID")
			return nil, returnErr
		}
		// Find customer and return
		customer, err := f.FetchCustomer(customerID)
		if err != nil {
			returnErr := services.NewNotFoundError(customerID, "Customer with that ID not found")
			return nil, returnErr
		}
		return customer, nil
	}
	// Else customerIDString is empty and we need to create a customer
	// Since each customer has a unique userid we need to create a user
	if customerBody == nil {
		returnErr := services.NewInvalidInputError(uuid.Nil, nil, nil, "If CustomerID is not provided, customer object is required to create Customer")
		return nil, returnErr
	}
	user, err := createUser(tx, customerBody.Email, logger)
	if err != nil {
		return nil, err
	}

	customer := CustomerModel(customerBody)
	customer.User = *user
	customer.UserID = user.ID

	// Create the new customer in the db
	verrs, err := tx.ValidateAndCreate(customer)
	if verrs.Count() > 0 {
		logger.Error("supportapi.createOrGetCustomer error", zap.Error(verrs))
		return nil, services.NewInvalidInputError(uuid.Nil, nil, verrs, "")
	} else if err != nil {
		logger.Error("supportapi.createOrGetCustomer error", zap.Error(err))
		e := services.NewQueryError("Customer", err, "Unable to create Customer")
		return nil, e
	}
	return customer, nil

}

// CustomerModel converts payload to model - currently does not tackle addresses
func CustomerModel(customer *supportmessages.Customer) *models.ServiceMember {
	if customer == nil {
		return nil
	}
	return &models.ServiceMember{
		ID:            uuid.FromStringOrNil(customer.ID.String()),
		Affiliation:   (*models.ServiceMemberAffiliation)(&customer.Agency),
		Edipi:         &customer.DodID,
		FirstName:     &customer.FirstName,
		LastName:      &customer.LastName,
		PersonalEmail: customer.Email,
		Telephone:     customer.Phone,
	}
}

// MoveOrderModel converts payload to model - it does not convert nested
// duty stations but will preserve the ID if provided.
// It will create nested customer and entitlement models
// if those are provided in the payload
func MoveOrderModel(moveOrderPayload *supportmessages.MoveOrder) *models.Order {
	if moveOrderPayload == nil {
		return nil
	}
	model := &models.Order{
		ID:               uuid.FromStringOrNil(moveOrderPayload.ID.String()),
		Grade:            moveOrderPayload.Rank,
		OrdersNumber:     moveOrderPayload.OrderNumber,
		ServiceMember:    *CustomerModel(moveOrderPayload.Customer),
		Entitlement:      EntitlementModel(moveOrderPayload.Entitlement),
		Status:           (models.OrderStatus)(moveOrderPayload.Status),
		IssueDate:        (time.Time)(*moveOrderPayload.IssueDate),
		OrdersType:       (internalmessages.OrdersType)(moveOrderPayload.OrdersType),
		UploadedOrdersID: uuid.FromStringOrNil(moveOrderPayload.UploadedOrdersID.String()),
	}

	customerID := uuid.FromStringOrNil(moveOrderPayload.CustomerID.String())
	model.ServiceMemberID = customerID

	destinationDutyStationID := uuid.FromStringOrNil(moveOrderPayload.DestinationDutyStationID.String())
	model.NewDutyStationID = destinationDutyStationID

	originDutyStationID := uuid.FromStringOrNil(moveOrderPayload.OriginDutyStationID.String())
	model.OriginDutyStationID = &originDutyStationID

	reportByDate := time.Time(*moveOrderPayload.ReportByDate)
	if !reportByDate.IsZero() {
		model.ReportByDate = reportByDate
	}
	return model
}

// EntitlementModel converts the payload to model
func EntitlementModel(entitlementPayload *supportmessages.Entitlement) *models.Entitlement {
	if entitlementPayload == nil {
		return nil
	}

	// proGearWeight and ProGearWeightSpouse currently not handled as
	// they are not in the entitlement record in the db.
	model := &models.Entitlement{
		ID:                    uuid.FromStringOrNil(entitlementPayload.ID.String()),
		DependentsAuthorized:  entitlementPayload.DependentsAuthorized,
		NonTemporaryStorage:   entitlementPayload.NonTemporaryStorage,
		PrivatelyOwnedVehicle: entitlementPayload.PrivatelyOwnedVehicle,
	}

	if entitlementPayload.AuthorizedWeight != nil {
		model.DBAuthorizedWeight = swag.Int(int(*entitlementPayload.AuthorizedWeight))
	}

	totalDependents := int(entitlementPayload.TotalDependents)
	model.TotalDependents = &totalDependents

	storageInTransit := int(entitlementPayload.StorageInTransit)
	model.StorageInTransit = &storageInTransit

	return model
}

// MoveTaskOrderModel return an MTO model constructed from the payload.
// Does not create nested mtoServiceItems, mtoShipments, or paymentRequests
func MoveTaskOrderModel(mtoPayload *supportmessages.MoveTaskOrder) *models.Move {
	if mtoPayload == nil {
		return nil
	}
	ppmEstimatedWeight := unit.Pound(mtoPayload.PpmEstimatedWeight)
	contractorID := uuid.FromStringOrNil(mtoPayload.ContractorID.String())
	model := &models.Move{
		ReferenceID:        &mtoPayload.ReferenceID,
		Locator:            mtoPayload.Locator,
		PPMEstimatedWeight: &ppmEstimatedWeight,
		PPMType:            &mtoPayload.PpmType,
		ContractorID:       &contractorID,
		Status:             (models.MoveStatus)(mtoPayload.Status),
	}

	if mtoPayload.AvailableToPrimeAt != nil {
		availableToPrimeAt := time.Time(*mtoPayload.AvailableToPrimeAt)
		model.AvailableToPrimeAt = &availableToPrimeAt
	}

	if mtoPayload.IsCanceled != nil && *mtoPayload.IsCanceled == true {
		model.Status = models.MoveStatusCANCELED
	}

	return model
}
