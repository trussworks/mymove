// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new mto shipment API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mto shipment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateMTOShipment creates m t o shipment

Creates a MTO shipment for the specified Move Task Order.
Required fields include:
* Shipment Type
* Customer requested pick-up date
* Pick-up Address
* Delivery Address
* Releasing / Receiving agents

Optional fields include:
* Customer Remarks
* Releasing / Receiving agents
* An array of optional accessorial service item codes

*/
func (a *Client) CreateMTOShipment(params *CreateMTOShipmentParams) (*CreateMTOShipmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMTOShipmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createMTOShipment",
		Method:             "POST",
		PathPattern:        "/mto-shipments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateMTOShipmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMTOShipmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMTOShipment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateMTOAgent updates m t o agent

### Functionality
This endpoint is used to **update** the agents for an MTO Shipment. Only the fields being modified need to be sent in the request body.

### Errors:
The agent must always have a name and at least one method of contact (either `email` or `phone`).

The agent must be associated with the MTO shipment passed in the url.

The shipment should be associated with an MTO that is available to the Prime.
If the caller requests an update to an agent, and the shipment is not on an available MTO, the caller will receive a **NotFound** response.

*/
func (a *Client) UpdateMTOAgent(params *UpdateMTOAgentParams) (*UpdateMTOAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMTOAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMTOAgent",
		Method:             "PUT",
		PathPattern:        "/mto-shipments/{mtoShipmentID}/agents/{agentID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateMTOAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateMTOAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateMTOAgent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateMTOShipment updates m t o shipment

Updates an existing shipment for a Move Task Order (MTO). Only the following fields can be updated using this endpoint:

* `scheduledPickupDate`
* `actualPickupDate`
* `firstAvailableDeliveryDate`
* `destinationAddress`
* `pickupAddress`
* `secondaryDeliveryAddress`
* `secondaryPickupAddress`
* `primeEstimatedWeight`
* `primeActualWeight`
* `shipmentType`
* `agents` - all subfields except `mtoShipmentID`, `createdAt`, `updatedAt`. You cannot add new agents to a shipment.

Note that some fields cannot be manually changed but will still be updated automatically, such as `primeEstimatedWeightRecordedDate` and `requiredDeliveryDate`.

*/
func (a *Client) UpdateMTOShipment(params *UpdateMTOShipmentParams) (*UpdateMTOShipmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMTOShipmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMTOShipment",
		Method:             "PUT",
		PathPattern:        "/mto-shipments/{mtoShipmentID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateMTOShipmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateMTOShipmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateMTOShipment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateMTOShipmentAddress updates m t o shipment address

### Functionality
This endpoint is used to **update** the addresses on an MTO Shipment. The address details completely replace the original, except for the UUID.
Therefore a complete address should be sent in the request.

This endpoint **cannot create** an address.
To create an address on an MTO shipment, the caller must use [updateMTOShipment](#operation/updateMTOShipment) as the parent shipment has to be updated with the appropriate link to the address.

### Errors
The address must be associated with the mtoShipment passed in the url.
In other words, it should be listed as pickupAddress, destinationAddress, secondaryPickupAddress or secondaryDeliveryAddress on the mtoShipment provided.
If it is not, caller will receive a **Conflict** Error.

The mtoShipment should be associated with an MTO that is available to prime.
If the caller requests an update to an address, and the shipment is not on an available MTO, the caller will receive a **NotFound** Error.

*/
func (a *Client) UpdateMTOShipmentAddress(params *UpdateMTOShipmentAddressParams) (*UpdateMTOShipmentAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMTOShipmentAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMTOShipmentAddress",
		Method:             "PUT",
		PathPattern:        "/mto-shipments/{mtoShipmentID}/addresses/{addressID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateMTOShipmentAddressReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateMTOShipmentAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateMTOShipmentAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
