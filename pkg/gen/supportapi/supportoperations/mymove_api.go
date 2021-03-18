// Code generated by go-swagger; DO NOT EDIT.

package supportoperations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/move_task_order"
	"github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/mto_service_item"
	"github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/mto_shipment"
	"github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/payment_request"
	"github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/webhook"
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
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		MoveTaskOrderCreateMoveTaskOrderHandler: move_task_order.CreateMoveTaskOrderHandlerFunc(func(params move_task_order.CreateMoveTaskOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation move_task_order.CreateMoveTaskOrder has not yet been implemented")
		}),
		WebhookCreateWebhookNotificationHandler: webhook.CreateWebhookNotificationHandlerFunc(func(params webhook.CreateWebhookNotificationParams) middleware.Responder {
			return middleware.NotImplemented("operation webhook.CreateWebhookNotification has not yet been implemented")
		}),
		MoveTaskOrderGetMoveTaskOrderHandler: move_task_order.GetMoveTaskOrderHandlerFunc(func(params move_task_order.GetMoveTaskOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation move_task_order.GetMoveTaskOrder has not yet been implemented")
		}),
		PaymentRequestGetPaymentRequestEDIHandler: payment_request.GetPaymentRequestEDIHandlerFunc(func(params payment_request.GetPaymentRequestEDIParams) middleware.Responder {
			return middleware.NotImplemented("operation payment_request.GetPaymentRequestEDI has not yet been implemented")
		}),
		MoveTaskOrderHideNonFakeMoveTaskOrdersHandler: move_task_order.HideNonFakeMoveTaskOrdersHandlerFunc(func(params move_task_order.HideNonFakeMoveTaskOrdersParams) middleware.Responder {
			return middleware.NotImplemented("operation move_task_order.HideNonFakeMoveTaskOrders has not yet been implemented")
		}),
		PaymentRequestListMTOPaymentRequestsHandler: payment_request.ListMTOPaymentRequestsHandlerFunc(func(params payment_request.ListMTOPaymentRequestsParams) middleware.Responder {
			return middleware.NotImplemented("operation payment_request.ListMTOPaymentRequests has not yet been implemented")
		}),
		MoveTaskOrderListMTOsHandler: move_task_order.ListMTOsHandlerFunc(func(params move_task_order.ListMTOsParams) middleware.Responder {
			return middleware.NotImplemented("operation move_task_order.ListMTOs has not yet been implemented")
		}),
		MoveTaskOrderMakeMoveTaskOrderAvailableHandler: move_task_order.MakeMoveTaskOrderAvailableHandlerFunc(func(params move_task_order.MakeMoveTaskOrderAvailableParams) middleware.Responder {
			return middleware.NotImplemented("operation move_task_order.MakeMoveTaskOrderAvailable has not yet been implemented")
		}),
		PaymentRequestProcessReviewedPaymentRequestsHandler: payment_request.ProcessReviewedPaymentRequestsHandlerFunc(func(params payment_request.ProcessReviewedPaymentRequestsParams) middleware.Responder {
			return middleware.NotImplemented("operation payment_request.ProcessReviewedPaymentRequests has not yet been implemented")
		}),
		WebhookReceiveWebhookNotificationHandler: webhook.ReceiveWebhookNotificationHandlerFunc(func(params webhook.ReceiveWebhookNotificationParams) middleware.Responder {
			return middleware.NotImplemented("operation webhook.ReceiveWebhookNotification has not yet been implemented")
		}),
		MtoServiceItemUpdateMTOServiceItemStatusHandler: mto_service_item.UpdateMTOServiceItemStatusHandlerFunc(func(params mto_service_item.UpdateMTOServiceItemStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation mto_service_item.UpdateMTOServiceItemStatus has not yet been implemented")
		}),
		MtoShipmentUpdateMTOShipmentStatusHandler: mto_shipment.UpdateMTOShipmentStatusHandlerFunc(func(params mto_shipment.UpdateMTOShipmentStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation mto_shipment.UpdateMTOShipmentStatus has not yet been implemented")
		}),
		PaymentRequestUpdatePaymentRequestStatusHandler: payment_request.UpdatePaymentRequestStatusHandlerFunc(func(params payment_request.UpdatePaymentRequestStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation payment_request.UpdatePaymentRequestStatus has not yet been implemented")
		}),
	}
}

/*MymoveAPI The Milmove Support API gives you programmatic access to support functionality useful for testing and debug.

This API is not available in production.

All endpoints are located at `primelocal/support/v1/`.
*/
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

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// MoveTaskOrderCreateMoveTaskOrderHandler sets the operation handler for the create move task order operation
	MoveTaskOrderCreateMoveTaskOrderHandler move_task_order.CreateMoveTaskOrderHandler
	// WebhookCreateWebhookNotificationHandler sets the operation handler for the create webhook notification operation
	WebhookCreateWebhookNotificationHandler webhook.CreateWebhookNotificationHandler
	// MoveTaskOrderGetMoveTaskOrderHandler sets the operation handler for the get move task order operation
	MoveTaskOrderGetMoveTaskOrderHandler move_task_order.GetMoveTaskOrderHandler
	// PaymentRequestGetPaymentRequestEDIHandler sets the operation handler for the get payment request e d i operation
	PaymentRequestGetPaymentRequestEDIHandler payment_request.GetPaymentRequestEDIHandler
	// MoveTaskOrderHideNonFakeMoveTaskOrdersHandler sets the operation handler for the hide non fake move task orders operation
	MoveTaskOrderHideNonFakeMoveTaskOrdersHandler move_task_order.HideNonFakeMoveTaskOrdersHandler
	// PaymentRequestListMTOPaymentRequestsHandler sets the operation handler for the list m t o payment requests operation
	PaymentRequestListMTOPaymentRequestsHandler payment_request.ListMTOPaymentRequestsHandler
	// MoveTaskOrderListMTOsHandler sets the operation handler for the list m t os operation
	MoveTaskOrderListMTOsHandler move_task_order.ListMTOsHandler
	// MoveTaskOrderMakeMoveTaskOrderAvailableHandler sets the operation handler for the make move task order available operation
	MoveTaskOrderMakeMoveTaskOrderAvailableHandler move_task_order.MakeMoveTaskOrderAvailableHandler
	// PaymentRequestProcessReviewedPaymentRequestsHandler sets the operation handler for the process reviewed payment requests operation
	PaymentRequestProcessReviewedPaymentRequestsHandler payment_request.ProcessReviewedPaymentRequestsHandler
	// WebhookReceiveWebhookNotificationHandler sets the operation handler for the receive webhook notification operation
	WebhookReceiveWebhookNotificationHandler webhook.ReceiveWebhookNotificationHandler
	// MtoServiceItemUpdateMTOServiceItemStatusHandler sets the operation handler for the update m t o service item status operation
	MtoServiceItemUpdateMTOServiceItemStatusHandler mto_service_item.UpdateMTOServiceItemStatusHandler
	// MtoShipmentUpdateMTOShipmentStatusHandler sets the operation handler for the update m t o shipment status operation
	MtoShipmentUpdateMTOShipmentStatusHandler mto_shipment.UpdateMTOShipmentStatusHandler
	// PaymentRequestUpdatePaymentRequestStatusHandler sets the operation handler for the update payment request status operation
	PaymentRequestUpdatePaymentRequestStatusHandler payment_request.UpdatePaymentRequestStatusHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

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

	if o.MoveTaskOrderCreateMoveTaskOrderHandler == nil {
		unregistered = append(unregistered, "move_task_order.CreateMoveTaskOrderHandler")
	}
	if o.WebhookCreateWebhookNotificationHandler == nil {
		unregistered = append(unregistered, "webhook.CreateWebhookNotificationHandler")
	}
	if o.MoveTaskOrderGetMoveTaskOrderHandler == nil {
		unregistered = append(unregistered, "move_task_order.GetMoveTaskOrderHandler")
	}
	if o.PaymentRequestGetPaymentRequestEDIHandler == nil {
		unregistered = append(unregistered, "payment_request.GetPaymentRequestEDIHandler")
	}
	if o.MoveTaskOrderHideNonFakeMoveTaskOrdersHandler == nil {
		unregistered = append(unregistered, "move_task_order.HideNonFakeMoveTaskOrdersHandler")
	}
	if o.PaymentRequestListMTOPaymentRequestsHandler == nil {
		unregistered = append(unregistered, "payment_request.ListMTOPaymentRequestsHandler")
	}
	if o.MoveTaskOrderListMTOsHandler == nil {
		unregistered = append(unregistered, "move_task_order.ListMTOsHandler")
	}
	if o.MoveTaskOrderMakeMoveTaskOrderAvailableHandler == nil {
		unregistered = append(unregistered, "move_task_order.MakeMoveTaskOrderAvailableHandler")
	}
	if o.PaymentRequestProcessReviewedPaymentRequestsHandler == nil {
		unregistered = append(unregistered, "payment_request.ProcessReviewedPaymentRequestsHandler")
	}
	if o.WebhookReceiveWebhookNotificationHandler == nil {
		unregistered = append(unregistered, "webhook.ReceiveWebhookNotificationHandler")
	}
	if o.MtoServiceItemUpdateMTOServiceItemStatusHandler == nil {
		unregistered = append(unregistered, "mto_service_item.UpdateMTOServiceItemStatusHandler")
	}
	if o.MtoShipmentUpdateMTOShipmentStatusHandler == nil {
		unregistered = append(unregistered, "mto_shipment.UpdateMTOShipmentStatusHandler")
	}
	if o.PaymentRequestUpdatePaymentRequestStatusHandler == nil {
		unregistered = append(unregistered, "payment_request.UpdatePaymentRequestStatusHandler")
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

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *MymoveAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
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

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *MymoveAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
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
	o.handlers["POST"]["/move-task-orders"] = move_task_order.NewCreateMoveTaskOrder(o.context, o.MoveTaskOrderCreateMoveTaskOrderHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/webhook-notifications"] = webhook.NewCreateWebhookNotification(o.context, o.WebhookCreateWebhookNotificationHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders/{moveTaskOrderID}"] = move_task_order.NewGetMoveTaskOrder(o.context, o.MoveTaskOrderGetMoveTaskOrderHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/payment-requests/{paymentRequestID}/edi"] = payment_request.NewGetPaymentRequestEDI(o.context, o.PaymentRequestGetPaymentRequestEDIHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/move-task-orders/hide"] = move_task_order.NewHideNonFakeMoveTaskOrders(o.context, o.MoveTaskOrderHideNonFakeMoveTaskOrdersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders/{moveTaskOrderID}/payment-requests"] = payment_request.NewListMTOPaymentRequests(o.context, o.PaymentRequestListMTOPaymentRequestsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders"] = move_task_order.NewListMTOs(o.context, o.MoveTaskOrderListMTOsHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/move-task-orders/{moveTaskOrderID}/available-to-prime"] = move_task_order.NewMakeMoveTaskOrderAvailable(o.context, o.MoveTaskOrderMakeMoveTaskOrderAvailableHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/payment-requests/process-reviewed"] = payment_request.NewProcessReviewedPaymentRequests(o.context, o.PaymentRequestProcessReviewedPaymentRequestsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/webhook-notify"] = webhook.NewReceiveWebhookNotification(o.context, o.WebhookReceiveWebhookNotificationHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/mto-service-items/{mtoServiceItemID}/status"] = mto_service_item.NewUpdateMTOServiceItemStatus(o.context, o.MtoServiceItemUpdateMTOServiceItemStatusHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/mto-shipments/{mtoShipmentID}/status"] = mto_shipment.NewUpdateMTOShipmentStatus(o.context, o.MtoShipmentUpdateMTOShipmentStatusHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/payment-requests/{paymentRequestID}/status"] = payment_request.NewUpdatePaymentRequestStatus(o.context, o.PaymentRequestUpdatePaymentRequestStatusHandler)
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

// AddMiddlewareFor adds a http middleware to existing handler
func (o *MymoveAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
