// Code generated by go-swagger; DO NOT EDIT.

package ghcoperations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/customer"
	"github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/move_task_order"
	"github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/mto_service_item"
	"github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/payment_requests"
)

// NewMymoveAPI creates a new Mymove instance
func NewMymoveAPI(spec *loads.Document) *MymoveAPI {
	return &MymoveAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		MtoServiceItemCreateMTOServiceItemHandler: mto_service_item.CreateMTOServiceItemHandlerFunc(func(params mto_service_item.CreateMTOServiceItemParams) middleware.Responder {
			return middleware.NotImplemented("operation MtoServiceItemCreateMTOServiceItem has not yet been implemented")
		}),
		MtoServiceItemDeleteMTOServiceItemHandler: mto_service_item.DeleteMTOServiceItemHandlerFunc(func(params mto_service_item.DeleteMTOServiceItemParams) middleware.Responder {
			return middleware.NotImplemented("operation MtoServiceItemDeleteMTOServiceItem has not yet been implemented")
		}),
		MoveTaskOrderDeleteMoveTaskOrderHandler: move_task_order.DeleteMoveTaskOrderHandlerFunc(func(params move_task_order.DeleteMoveTaskOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderDeleteMoveTaskOrder has not yet been implemented")
		}),
		CustomerGetAllCustomerMovesHandler: customer.GetAllCustomerMovesHandlerFunc(func(params customer.GetAllCustomerMovesParams) middleware.Responder {
			return middleware.NotImplemented("operation CustomerGetAllCustomerMoves has not yet been implemented")
		}),
		CustomerGetCustomerInfoHandler: customer.GetCustomerInfoHandlerFunc(func(params customer.GetCustomerInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation CustomerGetCustomerInfo has not yet been implemented")
		}),
		MoveTaskOrderGetEntitlementsHandler: move_task_order.GetEntitlementsHandlerFunc(func(params move_task_order.GetEntitlementsParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderGetEntitlements has not yet been implemented")
		}),
		MtoServiceItemGetMTOServiceItemHandler: mto_service_item.GetMTOServiceItemHandlerFunc(func(params mto_service_item.GetMTOServiceItemParams) middleware.Responder {
			return middleware.NotImplemented("operation MtoServiceItemGetMTOServiceItem has not yet been implemented")
		}),
		MoveTaskOrderGetMoveTaskOrderHandler: move_task_order.GetMoveTaskOrderHandlerFunc(func(params move_task_order.GetMoveTaskOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderGetMoveTaskOrder has not yet been implemented")
		}),
		PaymentRequestsGetPaymentRequestHandler: payment_requests.GetPaymentRequestHandlerFunc(func(params payment_requests.GetPaymentRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation PaymentRequestsGetPaymentRequest has not yet been implemented")
		}),
		MtoServiceItemListMTOServiceItemsHandler: mto_service_item.ListMTOServiceItemsHandlerFunc(func(params mto_service_item.ListMTOServiceItemsParams) middleware.Responder {
			return middleware.NotImplemented("operation MtoServiceItemListMTOServiceItems has not yet been implemented")
		}),
		MoveTaskOrderListMoveTaskOrdersHandler: move_task_order.ListMoveTaskOrdersHandlerFunc(func(params move_task_order.ListMoveTaskOrdersParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderListMoveTaskOrders has not yet been implemented")
		}),
		PaymentRequestsListPaymentRequestsHandler: payment_requests.ListPaymentRequestsHandlerFunc(func(params payment_requests.ListPaymentRequestsParams) middleware.Responder {
			return middleware.NotImplemented("operation PaymentRequestsListPaymentRequests has not yet been implemented")
		}),
		MtoServiceItemUpdateMTOServiceItemHandler: mto_service_item.UpdateMTOServiceItemHandlerFunc(func(params mto_service_item.UpdateMTOServiceItemParams) middleware.Responder {
			return middleware.NotImplemented("operation MtoServiceItemUpdateMTOServiceItem has not yet been implemented")
		}),
		MtoServiceItemUpdateMTOServiceItemstatusHandler: mto_service_item.UpdateMTOServiceItemstatusHandlerFunc(func(params mto_service_item.UpdateMTOServiceItemstatusParams) middleware.Responder {
			return middleware.NotImplemented("operation MtoServiceItemUpdateMTOServiceItemstatus has not yet been implemented")
		}),
		MoveTaskOrderUpdateMoveTaskOrderHandler: move_task_order.UpdateMoveTaskOrderHandlerFunc(func(params move_task_order.UpdateMoveTaskOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderUpdateMoveTaskOrder has not yet been implemented")
		}),
		MoveTaskOrderUpdateMoveTaskOrderStatusHandler: move_task_order.UpdateMoveTaskOrderStatusHandlerFunc(func(params move_task_order.UpdateMoveTaskOrderStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderUpdateMoveTaskOrderStatus has not yet been implemented")
		}),
		PaymentRequestsUpdatePaymentRequestHandler: payment_requests.UpdatePaymentRequestHandlerFunc(func(params payment_requests.UpdatePaymentRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation PaymentRequestsUpdatePaymentRequest has not yet been implemented")
		}),
		PaymentRequestsUpdatePaymentRequestStatusHandler: payment_requests.UpdatePaymentRequestStatusHandlerFunc(func(params payment_requests.UpdatePaymentRequestStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation PaymentRequestsUpdatePaymentRequestStatus has not yet been implemented")
		}),
	}
}

/*MymoveAPI The API for move.mil */
type MymoveAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// MtoServiceItemCreateMTOServiceItemHandler sets the operation handler for the create m t o service item operation
	MtoServiceItemCreateMTOServiceItemHandler mto_service_item.CreateMTOServiceItemHandler
	// MtoServiceItemDeleteMTOServiceItemHandler sets the operation handler for the delete m t o service item operation
	MtoServiceItemDeleteMTOServiceItemHandler mto_service_item.DeleteMTOServiceItemHandler
	// MoveTaskOrderDeleteMoveTaskOrderHandler sets the operation handler for the delete move task order operation
	MoveTaskOrderDeleteMoveTaskOrderHandler move_task_order.DeleteMoveTaskOrderHandler
	// CustomerGetAllCustomerMovesHandler sets the operation handler for the get all customer moves operation
	CustomerGetAllCustomerMovesHandler customer.GetAllCustomerMovesHandler
	// CustomerGetCustomerInfoHandler sets the operation handler for the get customer info operation
	CustomerGetCustomerInfoHandler customer.GetCustomerInfoHandler
	// MoveTaskOrderGetEntitlementsHandler sets the operation handler for the get entitlements operation
	MoveTaskOrderGetEntitlementsHandler move_task_order.GetEntitlementsHandler
	// MtoServiceItemGetMTOServiceItemHandler sets the operation handler for the get m t o service item operation
	MtoServiceItemGetMTOServiceItemHandler mto_service_item.GetMTOServiceItemHandler
	// MoveTaskOrderGetMoveTaskOrderHandler sets the operation handler for the get move task order operation
	MoveTaskOrderGetMoveTaskOrderHandler move_task_order.GetMoveTaskOrderHandler
	// PaymentRequestsGetPaymentRequestHandler sets the operation handler for the get payment request operation
	PaymentRequestsGetPaymentRequestHandler payment_requests.GetPaymentRequestHandler
	// MtoServiceItemListMTOServiceItemsHandler sets the operation handler for the list m t o service items operation
	MtoServiceItemListMTOServiceItemsHandler mto_service_item.ListMTOServiceItemsHandler
	// MoveTaskOrderListMoveTaskOrdersHandler sets the operation handler for the list move task orders operation
	MoveTaskOrderListMoveTaskOrdersHandler move_task_order.ListMoveTaskOrdersHandler
	// PaymentRequestsListPaymentRequestsHandler sets the operation handler for the list payment requests operation
	PaymentRequestsListPaymentRequestsHandler payment_requests.ListPaymentRequestsHandler
	// MtoServiceItemUpdateMTOServiceItemHandler sets the operation handler for the update m t o service item operation
	MtoServiceItemUpdateMTOServiceItemHandler mto_service_item.UpdateMTOServiceItemHandler
	// MtoServiceItemUpdateMTOServiceItemstatusHandler sets the operation handler for the update m t o service itemstatus operation
	MtoServiceItemUpdateMTOServiceItemstatusHandler mto_service_item.UpdateMTOServiceItemstatusHandler
	// MoveTaskOrderUpdateMoveTaskOrderHandler sets the operation handler for the update move task order operation
	MoveTaskOrderUpdateMoveTaskOrderHandler move_task_order.UpdateMoveTaskOrderHandler
	// MoveTaskOrderUpdateMoveTaskOrderStatusHandler sets the operation handler for the update move task order status operation
	MoveTaskOrderUpdateMoveTaskOrderStatusHandler move_task_order.UpdateMoveTaskOrderStatusHandler
	// PaymentRequestsUpdatePaymentRequestHandler sets the operation handler for the update payment request operation
	PaymentRequestsUpdatePaymentRequestHandler payment_requests.UpdatePaymentRequestHandler
	// PaymentRequestsUpdatePaymentRequestStatusHandler sets the operation handler for the update payment request status operation
	PaymentRequestsUpdatePaymentRequestStatusHandler payment_requests.UpdatePaymentRequestStatusHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *MymoveAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MymoveAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MymoveAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MymoveAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MymoveAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MymoveAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MymoveAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MymoveAPI
func (o *MymoveAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.MtoServiceItemCreateMTOServiceItemHandler == nil {
		unregistered = append(unregistered, "mto_service_item.CreateMTOServiceItemHandler")
	}

	if o.MtoServiceItemDeleteMTOServiceItemHandler == nil {
		unregistered = append(unregistered, "mto_service_item.DeleteMTOServiceItemHandler")
	}

	if o.MoveTaskOrderDeleteMoveTaskOrderHandler == nil {
		unregistered = append(unregistered, "move_task_order.DeleteMoveTaskOrderHandler")
	}

	if o.CustomerGetAllCustomerMovesHandler == nil {
		unregistered = append(unregistered, "customer.GetAllCustomerMovesHandler")
	}

	if o.CustomerGetCustomerInfoHandler == nil {
		unregistered = append(unregistered, "customer.GetCustomerInfoHandler")
	}

	if o.MoveTaskOrderGetEntitlementsHandler == nil {
		unregistered = append(unregistered, "move_task_order.GetEntitlementsHandler")
	}

	if o.MtoServiceItemGetMTOServiceItemHandler == nil {
		unregistered = append(unregistered, "mto_service_item.GetMTOServiceItemHandler")
	}

	if o.MoveTaskOrderGetMoveTaskOrderHandler == nil {
		unregistered = append(unregistered, "move_task_order.GetMoveTaskOrderHandler")
	}

	if o.PaymentRequestsGetPaymentRequestHandler == nil {
		unregistered = append(unregistered, "payment_requests.GetPaymentRequestHandler")
	}

	if o.MtoServiceItemListMTOServiceItemsHandler == nil {
		unregistered = append(unregistered, "mto_service_item.ListMTOServiceItemsHandler")
	}

	if o.MoveTaskOrderListMoveTaskOrdersHandler == nil {
		unregistered = append(unregistered, "move_task_order.ListMoveTaskOrdersHandler")
	}

	if o.PaymentRequestsListPaymentRequestsHandler == nil {
		unregistered = append(unregistered, "payment_requests.ListPaymentRequestsHandler")
	}

	if o.MtoServiceItemUpdateMTOServiceItemHandler == nil {
		unregistered = append(unregistered, "mto_service_item.UpdateMTOServiceItemHandler")
	}

	if o.MtoServiceItemUpdateMTOServiceItemstatusHandler == nil {
		unregistered = append(unregistered, "mto_service_item.UpdateMTOServiceItemstatusHandler")
	}

	if o.MoveTaskOrderUpdateMoveTaskOrderHandler == nil {
		unregistered = append(unregistered, "move_task_order.UpdateMoveTaskOrderHandler")
	}

	if o.MoveTaskOrderUpdateMoveTaskOrderStatusHandler == nil {
		unregistered = append(unregistered, "move_task_order.UpdateMoveTaskOrderStatusHandler")
	}

	if o.PaymentRequestsUpdatePaymentRequestHandler == nil {
		unregistered = append(unregistered, "payment_requests.UpdatePaymentRequestHandler")
	}

	if o.PaymentRequestsUpdatePaymentRequestStatusHandler == nil {
		unregistered = append(unregistered, "payment_requests.UpdatePaymentRequestStatusHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MymoveAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MymoveAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *MymoveAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *MymoveAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *MymoveAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *MymoveAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the mymove API
func (o *MymoveAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MymoveAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/move_task_orders/{moveTaskOrderID}/mto_service_items"] = mto_service_item.NewCreateMTOServiceItem(o.context, o.MtoServiceItemCreateMTOServiceItemHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/move-task-orders/{moveTaskOrderID}/service-items/{mtoServiceItemID}"] = mto_service_item.NewDeleteMTOServiceItem(o.context, o.MtoServiceItemDeleteMTOServiceItemHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/move-task-orders/{moveTaskOrderID}"] = move_task_order.NewDeleteMoveTaskOrder(o.context, o.MoveTaskOrderDeleteMoveTaskOrderHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/customer"] = customer.NewGetAllCustomerMoves(o.context, o.CustomerGetAllCustomerMovesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/customer/{customerID}"] = customer.NewGetCustomerInfo(o.context, o.CustomerGetCustomerInfoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders/{moveTaskOrderID}/entitlements"] = move_task_order.NewGetEntitlements(o.context, o.MoveTaskOrderGetEntitlementsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders/{moveTaskOrderID}/service-items/{mtoServiceItemID}"] = mto_service_item.NewGetMTOServiceItem(o.context, o.MtoServiceItemGetMTOServiceItemHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders/{moveTaskOrderID}"] = move_task_order.NewGetMoveTaskOrder(o.context, o.MoveTaskOrderGetMoveTaskOrderHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/payment-requests/{paymentRequestID}"] = payment_requests.NewGetPaymentRequest(o.context, o.PaymentRequestsGetPaymentRequestHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move_task_orders/{moveTaskOrderID}/mto_service_items"] = mto_service_item.NewListMTOServiceItems(o.context, o.MtoServiceItemListMTOServiceItemsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders"] = move_task_order.NewListMoveTaskOrders(o.context, o.MoveTaskOrderListMoveTaskOrdersHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/payment-requests"] = payment_requests.NewListPaymentRequests(o.context, o.PaymentRequestsListPaymentRequestsHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/move-task-orders/{moveTaskOrderID}/service-items/{mtoServiceItemID}"] = mto_service_item.NewUpdateMTOServiceItem(o.context, o.MtoServiceItemUpdateMTOServiceItemHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/move-task-orders/{moveTaskOrderID}/service-items/{mtoServiceItemID}/status"] = mto_service_item.NewUpdateMTOServiceItemstatus(o.context, o.MtoServiceItemUpdateMTOServiceItemstatusHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/move-task-orders/{moveTaskOrderID}"] = move_task_order.NewUpdateMoveTaskOrder(o.context, o.MoveTaskOrderUpdateMoveTaskOrderHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/move-task-orders/{moveTaskOrderID}/status"] = move_task_order.NewUpdateMoveTaskOrderStatus(o.context, o.MoveTaskOrderUpdateMoveTaskOrderStatusHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/payment-requests/{paymentRequestID}"] = payment_requests.NewUpdatePaymentRequest(o.context, o.PaymentRequestsUpdatePaymentRequestHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/payment-requests/{paymentRequestID}/status"] = payment_requests.NewUpdatePaymentRequestStatus(o.context, o.PaymentRequestsUpdatePaymentRequestStatusHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MymoveAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *MymoveAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MymoveAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MymoveAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
