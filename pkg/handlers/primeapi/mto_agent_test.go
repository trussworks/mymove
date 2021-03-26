package primeapi

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/etag"
	mtoshipmentops "github.com/transcom/mymove/pkg/gen/primeapi/primeoperations/mto_shipment"
	"github.com/transcom/mymove/pkg/handlers/primeapi/payloads"
	movetaskorder "github.com/transcom/mymove/pkg/services/move_task_order"
	mtoagent "github.com/transcom/mymove/pkg/services/mto_agent"

	"github.com/transcom/mymove/pkg/handlers"
)

func (suite *HandlerSuite) TestUpdateMTOAgentHandler() {
	// Set up db objects
	agent := testdatagen.MakeMTOAgent(suite.DB(), testdatagen.Assertions{
		Move: testdatagen.MakeAvailableMove(suite.DB()),
	})

	firstName := "Carol"
	lastName := "Romilly"
	email := "carol.romilly@example.com"
	phone := "456-555-7890"

	newAgent := models.MTOAgent{
		FirstName: &firstName,
		LastName:  &lastName,
		Email:     &email,
		Phone:     &phone,
	}

	// Create handler and request
	handler := UpdateMTOAgentHandler{
		handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
		mtoagent.NewMTOAgentUpdater(suite.DB()),
	}
	req := httptest.NewRequest("PUT", fmt.Sprintf("/mto-shipments/%s/agents/%s", agent.MTOShipmentID.String(), agent.ID.String()), nil)

	eTag := etag.GenerateEtag(agent.UpdatedAt)
	var updatedETag string

	// Test a successful request + update
	suite.T().Run("200 - OK response", func(t *testing.T) {
		payload := payloads.MTOAgent(&newAgent)
		params := mtoshipmentops.UpdateMTOAgentParams{
			HTTPRequest:   req,
			AgentID:       *handlers.FmtUUID(agent.ID),
			MtoShipmentID: *handlers.FmtUUID(agent.MTOShipmentID),
			Body:          payload,
			IfMatch:       eTag,
		}
		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		// Run handler and check response
		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.UpdateMTOAgentOK{}, response)

		// Check values
		agentOK := response.(*mtoshipmentops.UpdateMTOAgentOK)
		suite.Equal(agentOK.Payload.ID.String(), agent.ID.String())
		suite.Equal(agentOK.Payload.MtoShipmentID.String(), agent.MTOShipmentID.String())
		suite.Equal(string(agentOK.Payload.AgentType), string(agent.MTOAgentType)) // wasn't updated, should be original value
		suite.Equal(agentOK.Payload.FirstName, newAgent.FirstName)
		suite.Equal(agentOK.Payload.LastName, newAgent.LastName)
		suite.Equal(agentOK.Payload.Email, newAgent.Email)
		suite.Equal(agentOK.Payload.Phone, newAgent.Phone)
		updatedETag = agentOK.Payload.ETag
	})

	// Test stale eTag
	suite.T().Run("412 - Precondition failed response", func(t *testing.T) {
		// Let's test with the same valid values, but the original, expired eTag:
		payload := payloads.MTOAgent(&newAgent)
		params := mtoshipmentops.UpdateMTOAgentParams{
			HTTPRequest:   req,
			AgentID:       *handlers.FmtUUID(agent.ID),
			MtoShipmentID: *handlers.FmtUUID(agent.MTOShipmentID),
			Body:          payload,
			IfMatch:       eTag, // stale
		}
		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		// Run handler and check response
		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.UpdateMTOAgentPreconditionFailed{}, response)
	})

	// Test invalid IDs in the body vs. path values
	suite.T().Run("422 - Unprocessable response for bad ID values", func(t *testing.T) {
		fakeUUID := uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001")

		badAgent := newAgent
		badAgent.ID = fakeUUID
		badAgent.MTOShipmentID = fakeUUID

		payload := payloads.MTOAgent(&badAgent)
		params := mtoshipmentops.UpdateMTOAgentParams{
			HTTPRequest:   req,
			AgentID:       *handlers.FmtUUID(agent.ID),
			MtoShipmentID: *handlers.FmtUUID(agent.MTOShipmentID),
			Body:          payload,
			IfMatch:       updatedETag,
		}
		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		// Run handler and check response
		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.UpdateMTOAgentUnprocessableEntity{}, response)

		// Check error message for the invalid fields
		agentUnprocessable := response.(*mtoshipmentops.UpdateMTOAgentUnprocessableEntity)
		_, okID := agentUnprocessable.Payload.InvalidFields["id"]
		_, okMTOShipmentID := agentUnprocessable.Payload.InvalidFields["mtoShipmentID"]
		suite.True(okID)
		suite.True(okMTOShipmentID)
	})

	// Test invalid input
	suite.T().Run("422 - Unprocessable response for invalid input", func(t *testing.T) {
		empty := ""

		payload := payloads.MTOAgent(&newAgent)
		payload.FirstName = &empty
		payload.Email = &empty
		payload.Phone = &empty

		params := mtoshipmentops.UpdateMTOAgentParams{
			HTTPRequest:   req,
			AgentID:       *handlers.FmtUUID(agent.ID),
			MtoShipmentID: *handlers.FmtUUID(agent.MTOShipmentID),
			Body:          payload,
			IfMatch:       updatedETag,
		}
		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		// Run handler and check response
		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.UpdateMTOAgentUnprocessableEntity{}, response)

		// Check error message for the invalid fields
		agentUnprocessable := response.(*mtoshipmentops.UpdateMTOAgentUnprocessableEntity)
		_, okFirstName := agentUnprocessable.Payload.InvalidFields["firstName"]
		_, okContactInfo := agentUnprocessable.Payload.InvalidFields["contactInfo"]
		suite.True(okFirstName)
		suite.True(okContactInfo)
	})

	// Test not found response
	suite.T().Run("404 - Not found response", func(t *testing.T) {
		payload := payloads.MTOAgent(&newAgent)
		params := mtoshipmentops.UpdateMTOAgentParams{
			HTTPRequest:   req,
			AgentID:       *handlers.FmtUUID(agent.MTOShipmentID), // instead of agent.ID
			MtoShipmentID: *handlers.FmtUUID(agent.MTOShipmentID),
			Body:          payload,
			IfMatch:       updatedETag,
		}
		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		// Run handler and check response
		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.UpdateMTOAgentNotFound{}, response)

		// Check error message for the incorrect ID
		agentNotFound := response.(*mtoshipmentops.UpdateMTOAgentNotFound)
		suite.Contains(*agentNotFound.Payload.Detail, agent.MTOShipmentID.String())
	})

	// Test not Prime-available (not found response)
	suite.T().Run("404 - Not available response", func(t *testing.T) {
		unavailableAgent := testdatagen.MakeDefaultMTOAgent(suite.DB()) // default is not available to Prime

		payload := payloads.MTOAgent(&unavailableAgent)
		params := mtoshipmentops.UpdateMTOAgentParams{
			HTTPRequest:   req,
			AgentID:       *handlers.FmtUUID(unavailableAgent.ID),
			MtoShipmentID: *handlers.FmtUUID(unavailableAgent.MTOShipmentID),
			Body:          payload,
			IfMatch:       etag.GenerateEtag(unavailableAgent.UpdatedAt),
		}
		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		// Run handler and check response
		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.UpdateMTOAgentNotFound{}, response)

		// Check error message for the unavailable ID
		agentNotFound := response.(*mtoshipmentops.UpdateMTOAgentNotFound)
		suite.Contains(*agentNotFound.Payload.Detail, unavailableAgent.ID.String())
	})
}

func (suite *HandlerSuite) TestCreateMTOAgentHandler() {
	// Create new mtoShipment with no agents
	move := testdatagen.MakeAvailableMove(suite.DB())
	mtoShipment := testdatagen.MakeMTOShipmentMinimal(suite.DB(), testdatagen.Assertions{
		Move: move,
	})

	const RECEIVING_AGENT = "RECEIVING_AGENT"
	const RELEASING_AGENT = "RELEASING_AGENT"

	receivingAgent := &models.MTOAgent{

		FirstName:     swag.String("Riley"),
		LastName:      swag.String("Baker"),
		MTOAgentType:  RECEIVING_AGENT,
		Email:         swag.String("rileybaker@example.com"),
		Phone:         swag.String("555-555-5555"),
		MTOShipmentID: mtoShipment.ID,
	}

	releasingAgent := &models.MTOAgent{

		FirstName:     swag.String("Jason"),
		LastName:      swag.String("Ash"),
		MTOAgentType:  RELEASING_AGENT,
		Email:         swag.String("jasonash@example.com"),
		Phone:         swag.String("555-555-5555"),
		MTOShipmentID: mtoShipment.ID,
	}

	// Create Handler
	handler := CreateMTOAgentHandler{
		handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
		mtoagent.NewMTOAgentCreator(suite.DB(), movetaskorder.NewMoveTaskOrderChecker(suite.DB())),
	}
	req := httptest.NewRequest("POST", fmt.Sprintf("/mto-shipments/%s/agents", mtoShipment.ID), nil)

	suite.T().Run("200 - OK response Receiving Agent", func(t *testing.T) {
		payload := payloads.MTOAgent(receivingAgent)
		params := mtoshipmentops.CreateMTOAgentParams{
			HTTPRequest:   req,
			MtoShipmentID: *handlers.FmtUUID(mtoShipment.ID),
			Body:          payload,
		}

		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		// Run handler and check response
		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.CreateMTOAgentOK{}, response)

		// Check Values
		agentOK := response.(*mtoshipmentops.CreateMTOAgentOK)
		suite.Equal(agentOK.Payload.MtoShipmentID.String(), receivingAgent.MTOShipmentID.String())
		suite.Equal(string(agentOK.Payload.AgentType), string(receivingAgent.MTOAgentType)) // wasn't updated, should be original value
		suite.Equal(agentOK.Payload.FirstName, receivingAgent.FirstName)
		suite.Equal(agentOK.Payload.LastName, receivingAgent.LastName)
		suite.Equal(agentOK.Payload.Email, receivingAgent.Email)
		suite.Equal(agentOK.Payload.Phone, receivingAgent.Phone)

	})

	suite.T().Run("200 - OK response Releasing Agent", func(t *testing.T) {
		payload := payloads.MTOAgent(releasingAgent)
		params := mtoshipmentops.CreateMTOAgentParams{
			HTTPRequest:   req,
			MtoShipmentID: *handlers.FmtUUID(mtoShipment.ID),
			Body:          payload,
		}

		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		// Run handler and check response
		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.CreateMTOAgentOK{}, response)

		// Check Values
		agentOK := response.(*mtoshipmentops.CreateMTOAgentOK)
		suite.Equal(agentOK.Payload.MtoShipmentID.String(), releasingAgent.MTOShipmentID.String())
		suite.Equal(string(agentOK.Payload.AgentType), string(releasingAgent.MTOAgentType)) // wasn't updated, should be original value
		suite.Equal(agentOK.Payload.FirstName, releasingAgent.FirstName)
		suite.Equal(agentOK.Payload.LastName, releasingAgent.LastName)
		suite.Equal(agentOK.Payload.Email, releasingAgent.Email)
		suite.Equal(agentOK.Payload.Phone, releasingAgent.Phone)

	})

	suite.T().Run("404 - Not Found response", func(t *testing.T) {
		releasingAgent.MTOShipmentID = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001")
		payload := payloads.MTOAgent(releasingAgent)
		params := mtoshipmentops.CreateMTOAgentParams{
			HTTPRequest:   req,
			MtoShipmentID: *handlers.FmtUUID(releasingAgent.MTOShipmentID),
			Body:          payload,
		}

		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		// Run handler and check response
		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.CreateMTOAgentNotFound{}, response)

	})

	suite.T().Run("409 - Conflict response", func(t *testing.T) {
		// Try to add Receiving Agent again, when one already exists
		payload := payloads.MTOAgent(receivingAgent)
		params := mtoshipmentops.CreateMTOAgentParams{
			HTTPRequest:   req,
			MtoShipmentID: *handlers.FmtUUID(mtoShipment.ID),
			Body:          payload,
		}

		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		// Run handler and check response
		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.CreateMTOAgentConflict{}, response)
	})

	// Test invalid input
	suite.T().Run("422 - Unprocessable response for invalid input", func(t *testing.T) {
		newMTOShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			Move: move,
		})
		releasingAgent.MTOShipmentID = newMTOShipment.ID
		empty := ""

		payload := payloads.MTOAgent(releasingAgent)
		payload.FirstName = &empty
		payload.Email = &empty
		payload.Phone = &empty

		params := mtoshipmentops.CreateMTOAgentParams{
			HTTPRequest:   req,
			MtoShipmentID: *handlers.FmtUUID(newMTOShipment.ID),
			Body:          payload,
		}
		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		// Run handler and check response
		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.CreateMTOAgentUnprocessableEntity{}, response)

		// Check error message for the invalid fields
		// agentUnprocessable := response.(*mtoshipmentops.CreateMTOAgentUnprocessableEntity)
		// _, okFirstName := agentUnprocessable.Payload.InvalidFields["firstName"]
		// // _, okContactInfo := agentUnprocessable.Payload.InvalidFields["contactInfo"]
		// suite.True(okFirstName)
		// suite.True(okContactInfo)
	})
}
