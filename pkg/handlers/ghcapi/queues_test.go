package ghcapi

import (
	"net/http/httptest"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"

	"github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/queues"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/models/roles"
	moveorder "github.com/transcom/mymove/pkg/services/move_order"
	officeuser "github.com/transcom/mymove/pkg/services/office_user"
	"github.com/transcom/mymove/pkg/services/query"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *HandlerSuite) TestGetMoveQueuesHandler() {
	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
	officeUser.User.Roles = append(officeUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeTOO,
	})

	hhgMoveType := models.SelectedMoveTypeHHG
	// Default Origin Duty Station GBLOC is LKNQ
	hhgMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
			Status:           models.MoveStatusSUBMITTED,
		},
	})

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: hhgMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	// Create a shipment on hhgMove that has Rejected status
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: hhgMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusRejected,
		},
	})

	// Create an order with an origin duty station outside of office user GBLOC
	transportationOffice := testdatagen.MakeTransportationOffice(suite.DB(), testdatagen.Assertions{
		TransportationOffice: models.TransportationOffice{
			Name:  "Fort Punxsutawney",
			Gbloc: "AGFM",
		},
	})

	dutyStation := testdatagen.MakeDutyStation(suite.DB(), testdatagen.Assertions{
		DutyStation: models.DutyStation{
			TransportationOffice:   transportationOffice,
			TransportationOfficeID: &transportationOffice.ID,
		},
	})

	excludedOrder := testdatagen.MakeOrder(suite.DB(), testdatagen.Assertions{
		OriginDutyStation: dutyStation,
	})

	excludedMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
		},
		Order: excludedOrder,
	})

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: excludedMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	request := httptest.NewRequest("GET", "/move-task-orders/{moveTaskOrderID}", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetMovesQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		officeuser.NewOfficeUserFetcher(query.NewQueryBuilder(context.DB())),
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueOK{}, response)
	payload := response.(*queues.GetMovesQueueOK).Payload

	order := hhgMove.Orders
	result := payload.QueueMoves[0]

	suite.Len(payload.QueueMoves, 1)
	suite.Equal(order.ServiceMember.ID.String(), result.Customer.ID.String())
	suite.Equal(*order.DepartmentIndicator, string(result.DepartmentIndicator))
	suite.Equal(order.OriginDutyStation.TransportationOffice.Gbloc, string(result.OriginGBLOC))
	suite.Equal(order.NewDutyStation.ID.String(), result.DestinationDutyStation.ID.String())
	suite.Equal(hhgMove.Locator, result.Locator)
	suite.Equal(int64(1), result.ShipmentsCount)

}

func (suite *HandlerSuite) TestGetMoveQueuesHandlerStatuses() {
	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
	officeUser.User.Roles = append(officeUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeTOO,
	})

	hhgMoveType := models.SelectedMoveTypeHHG
	// Default Origin Duty Station GBLOC is LKNQ
	hhgMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
			Status:           models.MoveStatusSUBMITTED,
		},
	})

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: hhgMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	// Create a shipment on hhgMove that has Rejected status
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: hhgMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusRejected,
		},
	})

	// Create an order with an origin duty station outside of office user GBLOC
	transportationOffice := testdatagen.MakeTransportationOffice(suite.DB(), testdatagen.Assertions{
		TransportationOffice: models.TransportationOffice{
			Name:  "Fort Punxsutawney",
			Gbloc: "AGFM",
		},
	})

	dutyStation := testdatagen.MakeDutyStation(suite.DB(), testdatagen.Assertions{
		DutyStation: models.DutyStation{
			TransportationOffice:   transportationOffice,
			TransportationOfficeID: &transportationOffice.ID,
		},
	})

	excludedOrder := testdatagen.MakeOrder(suite.DB(), testdatagen.Assertions{
		OriginDutyStation: dutyStation,
	})

	excludedMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
		},
		Order: excludedOrder,
	})

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: excludedMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	request := httptest.NewRequest("GET", "/move-task-orders/{moveTaskOrderID}", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetMovesQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		officeuser.NewOfficeUserFetcher(query.NewQueryBuilder(context.DB())),
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	payload := response.(*queues.GetMovesQueueOK).Payload
	result := payload.QueueMoves[0]

	suite.Equal(ghcmessages.QueueMoveStatus("New move"), result.Status)

	// let's test for the Move approved status
	hhgMove.Status = models.MoveStatusAPPROVED
	_, _ = suite.DB().ValidateAndSave(&hhgMove)

	response = handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueOK{}, response)
	payload = response.(*queues.GetMovesQueueOK).Payload

	result = payload.QueueMoves[0]

	suite.Equal(ghcmessages.QueueMoveStatus("Move approved"), result.Status)

	// Now let's test Approvals requested
	testdatagen.MakeMTOServiceItem(suite.DB(), testdatagen.Assertions{
		MTOServiceItem: models.MTOServiceItem{
			Status: models.MTOServiceItemStatusSubmitted,
		},
		Move: hhgMove,
	})

	response = handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueOK{}, response)
	payload = response.(*queues.GetMovesQueueOK).Payload

	result = payload.QueueMoves[0]

	suite.Equal(ghcmessages.QueueMoveStatus("Approvals requested"), result.Status)

}

func (suite *HandlerSuite) TestGetMoveQueuesHandlerUnauthorizedRole() {
	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
	officeUser.User.Roles = append(officeUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeTIO,
	})

	request := httptest.NewRequest("GET", "/move-task-orders/{moveTaskOrderID}", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetMovesQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		officeuser.NewOfficeUserFetcher(query.NewQueryBuilder(context.DB())),
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueForbidden{}, response)
}

func (suite *HandlerSuite) TestGetMoveQueuesHandlerUnauthorizedUser() {
	serviceUser := testdatagen.MakeDefaultServiceMember(suite.DB())
	serviceUser.User.Roles = append(serviceUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeCustomer,
	})

	request := httptest.NewRequest("GET", "/move-task-orders/{moveTaskOrderID}", nil)
	request = suite.AuthenticateRequest(request, serviceUser)
	params := queues.GetMovesQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		officeuser.NewOfficeUserFetcher(query.NewQueryBuilder(context.DB())),
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueForbidden{}, response)
}

func (suite *HandlerSuite) TestGetMoveQueuesHandlerEmptyResults() {
	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
	officeUser.User.Roles = append(officeUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeTOO,
	})

	// Create an order with an origin duty station outside of office user GBLOC
	transportationOffice := testdatagen.MakeTransportationOffice(suite.DB(), testdatagen.Assertions{
		TransportationOffice: models.TransportationOffice{
			Name:  "Fort Punxsutawney",
			Gbloc: "AGFM",
		},
	})

	dutyStation := testdatagen.MakeDutyStation(suite.DB(), testdatagen.Assertions{
		DutyStation: models.DutyStation{
			TransportationOffice:   transportationOffice,
			TransportationOfficeID: &transportationOffice.ID,
		},
	})

	excludedOrder := testdatagen.MakeOrder(suite.DB(), testdatagen.Assertions{
		OriginDutyStation: dutyStation,
	})

	hhgMoveType := models.SelectedMoveTypeHHG
	excludedMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
		},
		Order: excludedOrder,
	})

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: excludedMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	request := httptest.NewRequest("GET", "/move-task-orders/{moveTaskOrderID}", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetMovesQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		officeuser.NewOfficeUserFetcher(query.NewQueryBuilder(context.DB())),
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueOK{}, response)
	payload := response.(*queues.GetMovesQueueOK).Payload

	suite.Len(payload.QueueMoves, 0)
}
