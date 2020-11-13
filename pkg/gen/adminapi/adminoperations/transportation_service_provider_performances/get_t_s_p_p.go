// Code generated by go-swagger; DO NOT EDIT.

package transportation_service_provider_performances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTSPPHandlerFunc turns a function with the right signature into a get t s p p handler
type GetTSPPHandlerFunc func(GetTSPPParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTSPPHandlerFunc) Handle(params GetTSPPParams) middleware.Responder {
	return fn(params)
}

// GetTSPPHandler interface for that can handle valid get t s p p params
type GetTSPPHandler interface {
	Handle(GetTSPPParams) middleware.Responder
}

// NewGetTSPP creates a new http.Handler for the get t s p p operation
func NewGetTSPP(ctx *middleware.Context, handler GetTSPPHandler) *GetTSPP {
	return &GetTSPP{Context: ctx, Handler: handler}
}

/*GetTSPP swagger:route GET /transportation_service_provider_performances/{tsppId} transportation_service_provider_performances getTSPP

Fetch a specific tspp

Returns a single tspp

*/
type GetTSPP struct {
	Context *middleware.Context
	Handler GetTSPPHandler
}

func (o *GetTSPP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTSPPParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
