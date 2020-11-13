// Code generated by go-swagger; DO NOT EDIT.

package access_codes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// IndexAccessCodesHandlerFunc turns a function with the right signature into a index access codes handler
type IndexAccessCodesHandlerFunc func(IndexAccessCodesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn IndexAccessCodesHandlerFunc) Handle(params IndexAccessCodesParams) middleware.Responder {
	return fn(params)
}

// IndexAccessCodesHandler interface for that can handle valid index access codes params
type IndexAccessCodesHandler interface {
	Handle(IndexAccessCodesParams) middleware.Responder
}

// NewIndexAccessCodes creates a new http.Handler for the index access codes operation
func NewIndexAccessCodes(ctx *middleware.Context, handler IndexAccessCodesHandler) *IndexAccessCodes {
	return &IndexAccessCodes{Context: ctx, Handler: handler}
}

/*IndexAccessCodes swagger:route GET /access_codes access_codes indexAccessCodes

List access codes

Returns a list of access codes

*/
type IndexAccessCodes struct {
	Context *middleware.Context
	Handler IndexAccessCodesHandler
}

func (o *IndexAccessCodes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewIndexAccessCodesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
