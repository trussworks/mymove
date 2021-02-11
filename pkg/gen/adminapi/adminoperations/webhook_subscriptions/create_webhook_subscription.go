// Code generated by go-swagger; DO NOT EDIT.

package webhook_subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateWebhookSubscriptionHandlerFunc turns a function with the right signature into a create webhook subscription handler
type CreateWebhookSubscriptionHandlerFunc func(CreateWebhookSubscriptionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateWebhookSubscriptionHandlerFunc) Handle(params CreateWebhookSubscriptionParams) middleware.Responder {
	return fn(params)
}

// CreateWebhookSubscriptionHandler interface for that can handle valid create webhook subscription params
type CreateWebhookSubscriptionHandler interface {
	Handle(CreateWebhookSubscriptionParams) middleware.Responder
}

// NewCreateWebhookSubscription creates a new http.Handler for the create webhook subscription operation
func NewCreateWebhookSubscription(ctx *middleware.Context, handler CreateWebhookSubscriptionHandler) *CreateWebhookSubscription {
	return &CreateWebhookSubscription{Context: ctx, Handler: handler}
}

/*CreateWebhookSubscription swagger:route POST /webhook_subscriptions webhook_subscriptions createWebhookSubscription

create a webhook subscription

creates and returns a webhook subscription

*/
type CreateWebhookSubscription struct {
	Context *middleware.Context
	Handler CreateWebhookSubscriptionHandler
}

func (o *CreateWebhookSubscription) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateWebhookSubscriptionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
