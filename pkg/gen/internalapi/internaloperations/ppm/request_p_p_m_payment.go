// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RequestPPMPaymentHandlerFunc turns a function with the right signature into a request p p m payment handler
type RequestPPMPaymentHandlerFunc func(RequestPPMPaymentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RequestPPMPaymentHandlerFunc) Handle(params RequestPPMPaymentParams) middleware.Responder {
	return fn(params)
}

// RequestPPMPaymentHandler interface for that can handle valid request p p m payment params
type RequestPPMPaymentHandler interface {
	Handle(RequestPPMPaymentParams) middleware.Responder
}

// NewRequestPPMPayment creates a new http.Handler for the request p p m payment operation
func NewRequestPPMPayment(ctx *middleware.Context, handler RequestPPMPaymentHandler) *RequestPPMPayment {
	return &RequestPPMPayment{Context: ctx, Handler: handler}
}

/*RequestPPMPayment swagger:route POST /personally_procured_move/{personallyProcuredMoveId}/request_payment ppm requestPPMPayment

Moves the PPM and the move into the PAYMENT_REQUESTED state

Moves the PPM and the move into the PAYMENT_REQUESTED state

*/
type RequestPPMPayment struct {
	Context *middleware.Context
	Handler RequestPPMPaymentHandler
}

func (o *RequestPPMPayment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRequestPPMPaymentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
