// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// IsLoggedInUserHandlerFunc turns a function with the right signature into a is logged in user handler
type IsLoggedInUserHandlerFunc func(IsLoggedInUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn IsLoggedInUserHandlerFunc) Handle(params IsLoggedInUserParams) middleware.Responder {
	return fn(params)
}

// IsLoggedInUserHandler interface for that can handle valid is logged in user params
type IsLoggedInUserHandler interface {
	Handle(IsLoggedInUserParams) middleware.Responder
}

// NewIsLoggedInUser creates a new http.Handler for the is logged in user operation
func NewIsLoggedInUser(ctx *middleware.Context, handler IsLoggedInUserHandler) *IsLoggedInUser {
	return &IsLoggedInUser{Context: ctx, Handler: handler}
}

/*IsLoggedInUser swagger:route GET /users/is_logged_in users isLoggedInUser

Returns boolean as to whether the user is logged in

Returns boolean as to whether the user is logged in

*/
type IsLoggedInUser struct {
	Context *middleware.Context
	Handler IsLoggedInUserHandler
}

func (o *IsLoggedInUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewIsLoggedInUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
