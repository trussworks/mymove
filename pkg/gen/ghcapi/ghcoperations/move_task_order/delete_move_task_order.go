// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteMoveTaskOrderHandlerFunc turns a function with the right signature into a delete move task order handler
type DeleteMoveTaskOrderHandlerFunc func(DeleteMoveTaskOrderParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMoveTaskOrderHandlerFunc) Handle(params DeleteMoveTaskOrderParams) middleware.Responder {
	return fn(params)
}

// DeleteMoveTaskOrderHandler interface for that can handle valid delete move task order params
type DeleteMoveTaskOrderHandler interface {
	Handle(DeleteMoveTaskOrderParams) middleware.Responder
}

// NewDeleteMoveTaskOrder creates a new http.Handler for the delete move task order operation
func NewDeleteMoveTaskOrder(ctx *middleware.Context, handler DeleteMoveTaskOrderHandler) *DeleteMoveTaskOrder {
	return &DeleteMoveTaskOrder{Context: ctx, Handler: handler}
}

/*DeleteMoveTaskOrder swagger:route DELETE /move-task-orders/{moveTaskOrderID} moveTaskOrder deleteMoveTaskOrder

Deletes a move order by ID

Deletes a move order by ID

*/
type DeleteMoveTaskOrder struct {
	Context *middleware.Context
	Handler DeleteMoveTaskOrderHandler
}

func (o *DeleteMoveTaskOrder) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteMoveTaskOrderParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
