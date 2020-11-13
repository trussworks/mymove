// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMTOServiceItemHandlerFunc turns a function with the right signature into a get m t o service item handler
type GetMTOServiceItemHandlerFunc func(GetMTOServiceItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMTOServiceItemHandlerFunc) Handle(params GetMTOServiceItemParams) middleware.Responder {
	return fn(params)
}

// GetMTOServiceItemHandler interface for that can handle valid get m t o service item params
type GetMTOServiceItemHandler interface {
	Handle(GetMTOServiceItemParams) middleware.Responder
}

// NewGetMTOServiceItem creates a new http.Handler for the get m t o service item operation
func NewGetMTOServiceItem(ctx *middleware.Context, handler GetMTOServiceItemHandler) *GetMTOServiceItem {
	return &GetMTOServiceItem{Context: ctx, Handler: handler}
}

/*GetMTOServiceItem swagger:route GET /move-task-orders/{moveTaskOrderID}/service-items/{mtoServiceItemID} mtoServiceItem getMTOServiceItem

Gets a line item by ID for a move order by ID

Gets a line item by ID for a move order by ID

*/
type GetMTOServiceItem struct {
	Context *middleware.Context
	Handler GetMTOServiceItemHandler
}

func (o *GetMTOServiceItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMTOServiceItemParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
