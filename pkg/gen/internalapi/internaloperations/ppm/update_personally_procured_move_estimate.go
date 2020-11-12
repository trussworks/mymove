// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdatePersonallyProcuredMoveEstimateHandlerFunc turns a function with the right signature into a update personally procured move estimate handler
type UpdatePersonallyProcuredMoveEstimateHandlerFunc func(UpdatePersonallyProcuredMoveEstimateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdatePersonallyProcuredMoveEstimateHandlerFunc) Handle(params UpdatePersonallyProcuredMoveEstimateParams) middleware.Responder {
	return fn(params)
}

// UpdatePersonallyProcuredMoveEstimateHandler interface for that can handle valid update personally procured move estimate params
type UpdatePersonallyProcuredMoveEstimateHandler interface {
	Handle(UpdatePersonallyProcuredMoveEstimateParams) middleware.Responder
}

// NewUpdatePersonallyProcuredMoveEstimate creates a new http.Handler for the update personally procured move estimate operation
func NewUpdatePersonallyProcuredMoveEstimate(ctx *middleware.Context, handler UpdatePersonallyProcuredMoveEstimateHandler) *UpdatePersonallyProcuredMoveEstimate {
	return &UpdatePersonallyProcuredMoveEstimate{Context: ctx, Handler: handler}
}

/*UpdatePersonallyProcuredMoveEstimate swagger:route PATCH /moves/{moveId}/personally_procured_move/{personallyProcuredMoveId}/estimate ppm updatePersonallyProcuredMoveEstimate

Calculates the estimated incentive of a PPM

This request calculates the estimated incentive of a PPM and attaches this range to the PPM

*/
type UpdatePersonallyProcuredMoveEstimate struct {
	Context *middleware.Context
	Handler UpdatePersonallyProcuredMoveEstimateHandler
}

func (o *UpdatePersonallyProcuredMoveEstimate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdatePersonallyProcuredMoveEstimateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
