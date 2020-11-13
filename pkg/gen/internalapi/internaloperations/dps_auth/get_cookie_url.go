// Code generated by go-swagger; DO NOT EDIT.

package dps_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCookieURLHandlerFunc turns a function with the right signature into a get cookie URL handler
type GetCookieURLHandlerFunc func(GetCookieURLParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCookieURLHandlerFunc) Handle(params GetCookieURLParams) middleware.Responder {
	return fn(params)
}

// GetCookieURLHandler interface for that can handle valid get cookie URL params
type GetCookieURLHandler interface {
	Handle(GetCookieURLParams) middleware.Responder
}

// NewGetCookieURL creates a new http.Handler for the get cookie URL operation
func NewGetCookieURL(ctx *middleware.Context, handler GetCookieURLHandler) *GetCookieURL {
	return &GetCookieURL{Context: ctx, Handler: handler}
}

/*GetCookieURL swagger:route POST /dps_auth/cookie_url dps_auth getCookieUrl

Returns the URL to redirect to that begins DPS auth

Returns the URL to redirect to that begins DPS auth

*/
type GetCookieURL struct {
	Context *middleware.Context
	Handler GetCookieURLHandler
}

func (o *GetCookieURL) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCookieURLParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
