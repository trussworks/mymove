// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// MakeMoveTaskOrderAvailableHandlerFunc turns a function with the right signature into a make move task order available handler
type MakeMoveTaskOrderAvailableHandlerFunc func(MakeMoveTaskOrderAvailableParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MakeMoveTaskOrderAvailableHandlerFunc) Handle(params MakeMoveTaskOrderAvailableParams) middleware.Responder {
	return fn(params)
}

// MakeMoveTaskOrderAvailableHandler interface for that can handle valid make move task order available params
type MakeMoveTaskOrderAvailableHandler interface {
	Handle(MakeMoveTaskOrderAvailableParams) middleware.Responder
}

// NewMakeMoveTaskOrderAvailable creates a new http.Handler for the make move task order available operation
func NewMakeMoveTaskOrderAvailable(ctx *middleware.Context, handler MakeMoveTaskOrderAvailableHandler) *MakeMoveTaskOrderAvailable {
	return &MakeMoveTaskOrderAvailable{Context: ctx, Handler: handler}
}

/*MakeMoveTaskOrderAvailable swagger:route PATCH /move-task-orders/{moveTaskOrderID}/available-to-prime moveTaskOrder makeMoveTaskOrderAvailable

makeMoveTaskOrderAvailable

Updates move task order `isAvailableToPrime` to TRUE to make it available to prime. <br />
<br />
This is a support endpoint and will not be available in production.


*/
type MakeMoveTaskOrderAvailable struct {
	Context *middleware.Context
	Handler MakeMoveTaskOrderAvailableHandler
}

func (o *MakeMoveTaskOrderAvailable) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewMakeMoveTaskOrderAvailableParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
