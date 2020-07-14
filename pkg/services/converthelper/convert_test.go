package converthelper_test

import (
	"time"

	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/services/converthelper"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
	"github.com/transcom/mymove/pkg/unit"
)

func (suite *ConvertSuite) TestConvertFromPPMToGHC() {
	move := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{})
	suite.NotNil(move)

	contractor := testdatagen.MakeContractor(suite.DB(), testdatagen.Assertions{
		Contractor: models.Contractor{
			ContractNumber: "HTC111-11-1-1111",
		},
	})
	suite.NotNil(contractor)

	miramar := testdatagen.MakeDutyStation(suite.DB(), testdatagen.Assertions{
		DutyStation: models.DutyStation{
			Name: "USMC Miramar",
		},
	})
	suite.NotNil(miramar)

	sanDiego := testdatagen.MakeDutyStation(suite.DB(), testdatagen.Assertions{
		DutyStation: models.DutyStation{
			Name: "USMC San Diego",
		},
	})
	suite.NotNil(sanDiego)

	moID, conversionErr := converthelper.ConvertFromPPMToGHC(suite.DB(), move.ID)
	suite.FatalNoError(conversionErr)

	var mo models.MoveOrder
	suite.FatalNoError(suite.DB().Eager("Customer", "Entitlement").Find(&mo, moID))

	suite.NotNil(mo.ReportByDate)
	suite.NotNil(mo.DateIssued)
	suite.NotNil(mo.OrderType)
	suite.NotNil(mo.OrderTypeDetail)
	suite.NotNil(mo.Grade)

	suite.NotEqual(uuid.Nil, mo.DestinationDutyStationID)
	suite.Equal(&move.Orders.NewDutyStationID, mo.DestinationDutyStationID)

	suite.NotEqual(uuid.Nil, mo.OriginDutyStationID)
	suite.Equal(move.Orders.ServiceMember.DutyStationID, mo.OriginDutyStationID)

	suite.NotEqual(uuid.Nil, mo.EntitlementID)
	suite.Equal(false, *mo.Entitlement.DependentsAuthorized)
	suite.Equal(7000, *mo.Entitlement.DBAuthorizedWeight)

	customer := mo.Customer
	suite.Equal(*move.Orders.ServiceMember.Edipi, *customer.Edipi)
	suite.NotEqual(uuid.Nil, customer.UserID)
	suite.Equal(move.Orders.ServiceMember.UserID, customer.UserID)

	suite.NotEqual(uuid.Nil, customer.UserID)
	suite.Equal(&customer.ID, mo.CustomerID)

	var mto models.MoveTaskOrder
	suite.FatalNoError(suite.DB().Eager().Where("move_order_id = ?", mo.ID).First(&mto))

	var mtoShipmentHHG models.MTOShipment
	suite.FatalNoError(suite.DB().Eager().Where("move_task_order_id = ? and shipment_type = ?", mto.ID, models.MTOShipmentTypeHHGLongHaulDom).First(&mtoShipmentHHG))

	suite.NotNil(mtoShipmentHHG.ScheduledPickupDate)
	suite.Equal(unit.Pound(4096), *mtoShipmentHHG.PrimeEstimatedWeight)

	expectedNilTime := time.Time{}
	suite.NotEqual(expectedNilTime, *mtoShipmentHHG.ScheduledPickupDate)
	suite.NotNil(*mtoShipmentHHG.PrimeEstimatedWeightRecordedDate)
	suite.NotEqual(expectedNilTime, *mtoShipmentHHG.PrimeEstimatedWeightRecordedDate)

	var mtoShipmentHHGDomShortHaul models.MTOShipment
	suite.FatalNoError(suite.DB().Eager().Where("move_task_order_id = ? and shipment_type = ?", mto.ID, models.MTOShipmentTypeHHGShortHaulDom).First(&mtoShipmentHHGDomShortHaul))

	suite.NotNil(mtoShipmentHHGDomShortHaul.ScheduledPickupDate)
	suite.Equal(unit.Pound(4096), *mtoShipmentHHGDomShortHaul.PrimeEstimatedWeight)
	suite.NotNil(*mtoShipmentHHGDomShortHaul.PrimeEstimatedWeightRecordedDate)
	suite.NotEqual(expectedNilTime, *mtoShipmentHHGDomShortHaul.PrimeEstimatedWeightRecordedDate)

	suite.NotEqual(expectedNilTime, *mtoShipmentHHGDomShortHaul.ScheduledPickupDate)

	suite.Equal(sanDiego.Address.ID, *mtoShipmentHHGDomShortHaul.PickupAddressID)
	suite.Equal(miramar.Address.ID, *mtoShipmentHHGDomShortHaul.DestinationAddressID)
}

func (suite *ConvertSuite) TestConvertProfileOrdersToGHC() {
	move := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{})
	suite.NotNil(move)

	contractor := testdatagen.MakeContractor(suite.DB(), testdatagen.Assertions{
		Contractor: models.Contractor{
			ContractNumber: "HTC111-11-1-1111",
		},
	})
	suite.NotNil(contractor)

	miramar := testdatagen.MakeDutyStation(suite.DB(), testdatagen.Assertions{
		DutyStation: models.DutyStation{
			Name: "USMC Miramar",
		},
	})
	suite.NotNil(miramar)

	sanDiego := testdatagen.MakeDutyStation(suite.DB(), testdatagen.Assertions{
		DutyStation: models.DutyStation{
			Name: "USMC San Diego",
		},
	})
	suite.NotNil(sanDiego)

	moID, conversionErr := converthelper.ConvertProfileOrdersToGHC(suite.DB(), move.ID)
	suite.FatalNoError(conversionErr)

	var mo models.MoveOrder
	suite.FatalNoError(suite.DB().Eager("Customer", "Entitlement").Find(&mo, moID))

	suite.NotNil(mo.ReportByDate)
	suite.NotNil(mo.DateIssued)
	suite.NotNil(mo.OrderType)
	suite.NotNil(mo.OrderTypeDetail)
	suite.NotNil(mo.Grade)

	suite.NotEqual(uuid.Nil, mo.DestinationDutyStationID)
	suite.Equal(&move.Orders.NewDutyStationID, mo.DestinationDutyStationID)

	suite.NotEqual(uuid.Nil, mo.OriginDutyStationID)
	suite.Equal(move.Orders.ServiceMember.DutyStationID, mo.OriginDutyStationID)

	suite.NotEqual(uuid.Nil, mo.EntitlementID)
	suite.Equal(false, *mo.Entitlement.DependentsAuthorized)
	suite.Equal(7000, *mo.Entitlement.DBAuthorizedWeight)

	customer := mo.Customer
	suite.Equal(*move.Orders.ServiceMember.Edipi, *customer.Edipi)
	suite.NotEqual(uuid.Nil, customer.UserID)
	suite.Equal(move.Orders.ServiceMember.UserID, customer.UserID)

	suite.NotEqual(uuid.Nil, customer.UserID)
	suite.Equal(&customer.ID, mo.CustomerID)

	var mto models.MoveTaskOrder
	suite.FatalNoError(suite.DB().Eager().Where("move_order_id = ?", mo.ID).First(&mto))
}

func (suite *ConvertSuite) TestConvertFromPPMToGHCMoveOrdersExist() {
	move := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{})
	suite.NotNil(move)

	sm := move.Orders.ServiceMember
	contractor := testdatagen.MakeContractor(suite.DB(), testdatagen.Assertions{
		Contractor: models.Contractor{
			ContractNumber: "HTC111-11-1-1111",
		},
	})
	suite.NotNil(contractor)

	miramar := testdatagen.MakeDutyStation(suite.DB(), testdatagen.Assertions{
		DutyStation: models.DutyStation{
			Name: "USMC Miramar",
		},
	})
	suite.NotNil(miramar)

	sanDiego := testdatagen.MakeDutyStation(suite.DB(), testdatagen.Assertions{
		DutyStation: models.DutyStation{
			Name: "USMC San Diego",
		},
	})
	suite.NotNil(sanDiego)

	_, conversionErr := converthelper.ConvertProfileOrdersToGHC(suite.DB(), move.ID)
	suite.FatalNoError(conversionErr)
	_, conversionErr = converthelper.ConvertFromPPMToGHC(suite.DB(), move.ID)
	suite.FatalNoError(conversionErr)
	var moveOrders []models.MoveOrder

	err := suite.DB().Where("customer_id = $1", sm.ID).All(&moveOrders)
	suite.FatalNoError(err)
	suite.Equal(1, len(moveOrders))

	var moveTaskOrders []models.MoveTaskOrder
	err = suite.DB().Where("move_order_id = $1", moveOrders[0].ID).All(&moveTaskOrders)
	suite.FatalNoError(err)
	suite.Equal(1, len(moveTaskOrders))
}