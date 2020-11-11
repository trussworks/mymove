package moveorder

import (
	"testing"

	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/services"

	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *MoveOrderServiceSuite) TestMoveOrderFetcher() {
	expectedMoveTaskOrder := testdatagen.MakeDefaultMove(suite.DB())
	expectedMoveOrder := expectedMoveTaskOrder.Orders
	moveOrderFetcher := NewMoveOrderFetcher(suite.DB())

	moveOrder, err := moveOrderFetcher.FetchMoveOrder(expectedMoveOrder.ID)
	suite.FatalNoError(err)

	suite.Equal(expectedMoveOrder.ID, moveOrder.ID)
	suite.Equal(expectedMoveOrder.ServiceMemberID, moveOrder.ServiceMemberID)
	suite.NotNil(moveOrder.NewDutyStation)
	suite.Equal(expectedMoveOrder.NewDutyStationID, moveOrder.NewDutyStation.ID)
	suite.Equal(expectedMoveOrder.NewDutyStation.AddressID, moveOrder.NewDutyStation.AddressID)
	suite.Equal(expectedMoveOrder.NewDutyStation.Address.StreetAddress1, moveOrder.NewDutyStation.Address.StreetAddress1)
	suite.NotNil(moveOrder.Entitlement)
	suite.Equal(*expectedMoveOrder.EntitlementID, moveOrder.Entitlement.ID)
	suite.Equal(expectedMoveOrder.OriginDutyStation.ID, moveOrder.OriginDutyStation.ID)
	suite.Equal(expectedMoveOrder.OriginDutyStation.AddressID, moveOrder.OriginDutyStation.AddressID)
	suite.Equal(expectedMoveOrder.OriginDutyStation.Address.StreetAddress1, moveOrder.OriginDutyStation.Address.StreetAddress1)
	suite.NotZero(moveOrder.OriginDutyStation)
}

func (suite *MoveOrderServiceSuite) TestMoveOrderFetcherWithEmptyFields() {
	// When move_orders and orders were consolidated, we moved the OriginDutyStation
	// field that used to only exist on the move_orders table into the orders table.
	// This means that existing orders in production won't have any values in the
	// OriginDutyStation column. To mimic that and to surface any issues, we didn't
	// update the testdatagen MakeOrder function so that new orders would have
	// an empty OriginDutyStation. During local testing in the office app, we
	// noticed an exception due to trying to load empty OriginDutyStations.
	// This was not caught by any tests, so we're adding one now.
	expectedOrder := testdatagen.MakeDefaultOrder(suite.DB())

	expectedOrder.Entitlement = nil
	expectedOrder.EntitlementID = nil
	expectedOrder.Grade = nil
	expectedOrder.OriginDutyStation = nil
	expectedOrder.OriginDutyStationID = nil
	suite.MustSave(&expectedOrder)

	testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Order: expectedOrder,
	})
	moveOrderFetcher := NewMoveOrderFetcher(suite.DB())
	moveOrder, err := moveOrderFetcher.FetchMoveOrder(expectedOrder.ID)

	suite.FatalNoError(err)
	suite.Nil(moveOrder.Entitlement)
	suite.Nil(moveOrder.OriginDutyStation)
	suite.Nil(moveOrder.Grade)
}

func (suite *MoveOrderServiceSuite) TestListMoveOrders() {
	// Create a Move without a shipment to test that only Orders with shipments
	// are displayed to the TOO
	testdatagen.MakeDefaultMove(suite.DB())

	expectedMoveTaskOrder := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{Move: models.Move{Status: models.MoveStatusSUBMITTED}})

	// Only orders with shipments are returned, so we need to add a shipment
	// to the move we just created
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: expectedMoveTaskOrder,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())

	expectedMoveOrder := expectedMoveTaskOrder.Orders
	moveOrderFetcher := NewMoveOrderFetcher(suite.DB())

	suite.T().Run("returns move orders", func(t *testing.T) {
		moveOrders, _, err := moveOrderFetcher.ListMoveOrders(officeUser.ID, &services.ListMoveOrderParams{PerPage: swag.Int64(1), Page: swag.Int64(1)})

		suite.FatalNoError(err)
		suite.Len(moveOrders, 1)

		moveOrder := moveOrders[0]

		suite.NotNil(moveOrder.ServiceMember)
		suite.Equal(expectedMoveOrder.ServiceMember.FirstName, moveOrder.ServiceMember.FirstName)
		suite.Equal(expectedMoveOrder.ServiceMember.LastName, moveOrder.ServiceMember.LastName)
		suite.Equal(expectedMoveOrder.ID, moveOrder.ID)
		suite.Equal(expectedMoveOrder.ServiceMemberID, moveOrder.ServiceMemberID)
		suite.NotNil(moveOrder.NewDutyStation)
		suite.Equal(expectedMoveOrder.NewDutyStationID, moveOrder.NewDutyStation.ID)
		suite.NotNil(moveOrder.Entitlement)
		suite.Equal(*expectedMoveOrder.EntitlementID, moveOrder.Entitlement.ID)
		suite.Equal(expectedMoveOrder.OriginDutyStation.ID, moveOrder.OriginDutyStation.ID)
		suite.NotNil(moveOrder.OriginDutyStation)
		suite.Equal(expectedMoveOrder.OriginDutyStation.AddressID, moveOrder.OriginDutyStation.AddressID)
		suite.Equal(expectedMoveOrder.OriginDutyStation.Address.StreetAddress1, moveOrder.OriginDutyStation.Address.StreetAddress1)
	})

	suite.T().Run("returns move orders filtered by GBLOC", func(t *testing.T) {
		testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusSubmitted,
			},
			TransportationOffice: models.TransportationOffice{
				Gbloc: "AGFM",
			},
			Move: models.Move{
				Status: models.MoveStatusSUBMITTED,
			},
		})

		moveOrders, _, err := moveOrderFetcher.ListMoveOrders(officeUser.ID, &services.ListMoveOrderParams{PerPage: swag.Int64(1), Page: swag.Int64(1)})

		suite.FatalNoError(err)
		suite.Equal(1, len(moveOrders))
	})

	suite.T().Run("returns orders filtered by an arbitrary query", func(t *testing.T) {
		army := "ARMY"
		params := services.ListMoveOrderParams{Branch: &army, PerPage: swag.Int64(1), Page: swag.Int64(1)}
		testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusSubmitted,
			},
			Order: models.Order{
				DepartmentIndicator: &army,
			},
			Move: models.Move{
				Status: models.MoveStatusSUBMITTED,
			},
		})

		moveOrders, _, err := moveOrderFetcher.ListMoveOrders(officeUser.ID, &params)

		suite.FatalNoError(err)
		suite.Equal(1, len(moveOrders))
	})
}

func (suite *MoveOrderServiceSuite) TestListMoveOrdersUSMCGbloc() {
	moveOrderFetcher := NewMoveOrderFetcher(suite.DB())

	suite.T().Run("returns USMC order for USMC office user", func(t *testing.T) {
		officeUUID, _ := uuid.NewV4()
		marines := models.AffiliationMARINES
		army := models.AffiliationARMY

		testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusSubmitted,
			},
			TransportationOffice: models.TransportationOffice{
				Gbloc: "USMC",
				ID:    officeUUID,
			},
			Move: models.Move{
				Status: models.MoveStatusSUBMITTED,
			},
			ServiceMember: models.ServiceMember{Affiliation: &marines},
		})

		testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusSubmitted,
			},
			Move: models.Move{
				Status: models.MoveStatusSUBMITTED,
			},
			ServiceMember: models.ServiceMember{Affiliation: &marines},
		})

		testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusSubmitted,
			},
			Move: models.Move{
				Status: models.MoveStatusSUBMITTED,
			},
			ServiceMember: models.ServiceMember{Affiliation: &army},
		})

		officeUserUSMC := testdatagen.MakeOfficeUser(suite.DB(), testdatagen.Assertions{OfficeUser: models.OfficeUser{TransportationOfficeID: officeUUID}})

		params := services.ListMoveOrderParams{PerPage: swag.Int64(20), Page: swag.Int64(1)}
		moveOrders, _, err := moveOrderFetcher.ListMoveOrders(officeUserUSMC.ID, &params)

		suite.FatalNoError(err)
		suite.Equal(2, len(moveOrders))
	})
}

func (suite *MoveOrderServiceSuite) TestListMoveOrdersWithEmptyFields() {
	expectedOrder := testdatagen.MakeDefaultOrder(suite.DB())

	expectedOrder.Entitlement = nil
	expectedOrder.EntitlementID = nil
	expectedOrder.Grade = nil
	expectedOrder.OriginDutyStation = nil
	expectedOrder.OriginDutyStationID = nil
	suite.MustSave(&expectedOrder)

	moveTaskOrder := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Order: expectedOrder,
	})
	// Only orders with shipments are returned, so we need to add a shipment
	// to the move we just created
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: moveTaskOrder,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})
	// Add a second shipment to make sure we only return 1 order even if its
	// move has more than one shipment
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: moveTaskOrder,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	officeUser := testdatagen.MakeOfficeUser(suite.DB(), testdatagen.Assertions{})
	moveOrderFetcher := NewMoveOrderFetcher(suite.DB())
	moveOrders, _, err := moveOrderFetcher.ListMoveOrders(officeUser.ID, &services.ListMoveOrderParams{PerPage: swag.Int64(1), Page: swag.Int64(1)})

	suite.FatalNoError(err)
	suite.Nil(moveOrders)

}

func (suite *MoveOrderServiceSuite) TestListMoveOrdersWithPagination() {
	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())

	for i := 0; i < 2; i++ {
		expectedMoveTaskOrder := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{Move: models.Move{Status: models.MoveStatusSUBMITTED}})

		// Only orders with shipments are returned, so we need to add a shipment
		// to the move we just created
		testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			Move: expectedMoveTaskOrder,
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusSubmitted,
			},
		})
	}

	moveOrderFetcher := NewMoveOrderFetcher(suite.DB())
	params := services.ListMoveOrderParams{Page: swag.Int64(1), PerPage: swag.Int64(1)}
	moveOrders, count, err := moveOrderFetcher.ListMoveOrders(officeUser.ID, &params)

	suite.NoError(err)
	suite.Equal(1, len(moveOrders))
	suite.Equal(2, count)

}
