// Code generated by go-swagger; DO NOT EDIT.

package payment_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListMTOPaymentRequestsHandlerFunc turns a function with the right signature into a list m t o payment requests handler
type ListMTOPaymentRequestsHandlerFunc func(ListMTOPaymentRequestsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListMTOPaymentRequestsHandlerFunc) Handle(params ListMTOPaymentRequestsParams) middleware.Responder {
	return fn(params)
}

// ListMTOPaymentRequestsHandler interface for that can handle valid list m t o payment requests params
type ListMTOPaymentRequestsHandler interface {
	Handle(ListMTOPaymentRequestsParams) middleware.Responder
}

// NewListMTOPaymentRequests creates a new http.Handler for the list m t o payment requests operation
func NewListMTOPaymentRequests(ctx *middleware.Context, handler ListMTOPaymentRequestsHandler) *ListMTOPaymentRequests {
	return &ListMTOPaymentRequests{Context: ctx, Handler: handler}
}

/*ListMTOPaymentRequests swagger:route GET /move-task-orders/{moveTaskOrderID}/payment-requests paymentRequest listMTOPaymentRequests

listMTOPaymentRequests

### Functionality

This endpoint lists all PaymentRequests associated with a given MoveTaskOrder.

This is a support endpoint and is not available in production.


*/
type ListMTOPaymentRequests struct {
	Context *middleware.Context
	Handler ListMTOPaymentRequestsHandler
}

func (o *ListMTOPaymentRequests) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListMTOPaymentRequestsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
