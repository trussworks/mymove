// Code generated by go-swagger; DO NOT EDIT.

package payment_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPaymentRequestEDIHandlerFunc turns a function with the right signature into a get payment request e d i handler
type GetPaymentRequestEDIHandlerFunc func(GetPaymentRequestEDIParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPaymentRequestEDIHandlerFunc) Handle(params GetPaymentRequestEDIParams) middleware.Responder {
	return fn(params)
}

// GetPaymentRequestEDIHandler interface for that can handle valid get payment request e d i params
type GetPaymentRequestEDIHandler interface {
	Handle(GetPaymentRequestEDIParams) middleware.Responder
}

// NewGetPaymentRequestEDI creates a new http.Handler for the get payment request e d i operation
func NewGetPaymentRequestEDI(ctx *middleware.Context, handler GetPaymentRequestEDIHandler) *GetPaymentRequestEDI {
	return &GetPaymentRequestEDI{Context: ctx, Handler: handler}
}

/*GetPaymentRequestEDI swagger:route GET /payment-requests/{paymentRequestID}/edi paymentRequest getPaymentRequestEDI

getPaymentRequestEDI

Returns the EDI (Electronic Data Interchange) message for the payment request identified
by the given payment request ID. Note that the EDI returned in the JSON payload will have \n where there
would normally be line breaks (due to JSON not allowing line breaks in a string).

This is a support endpoint and will not be available in production.


*/
type GetPaymentRequestEDI struct {
	Context *middleware.Context
	Handler GetPaymentRequestEDIHandler
}

func (o *GetPaymentRequestEDI) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPaymentRequestEDIParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
