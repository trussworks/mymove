// Code generated by go-swagger; DO NOT EDIT.

package m_t_oservice_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateMTOServiceItemHandlerFunc turns a function with the right signature into a update m t o service item handler
type UpdateMTOServiceItemHandlerFunc func(UpdateMTOServiceItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateMTOServiceItemHandlerFunc) Handle(params UpdateMTOServiceItemParams) middleware.Responder {
	return fn(params)
}

// UpdateMTOServiceItemHandler interface for that can handle valid update m t o service item params
type UpdateMTOServiceItemHandler interface {
	Handle(UpdateMTOServiceItemParams) middleware.Responder
}

// NewUpdateMTOServiceItem creates a new http.Handler for the update m t o service item operation
func NewUpdateMTOServiceItem(ctx *middleware.Context, handler UpdateMTOServiceItemHandler) *UpdateMTOServiceItem {
	return &UpdateMTOServiceItem{Context: ctx, Handler: handler}
}

/*UpdateMTOServiceItem swagger:route PATCH /move-task-orders/{moveTaskOrderID}/service-items/{mtoServiceItemID} MTOserviceItem updateMTOServiceItem

Updates a service item by ID for a move order by ID

Updates a service item by ID for a move order by ID

*/
type UpdateMTOServiceItem struct {
	Context *middleware.Context
	Handler UpdateMTOServiceItemHandler
}

func (o *UpdateMTOServiceItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateMTOServiceItemParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
