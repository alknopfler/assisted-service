// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateHostIgnitionHandlerFunc turns a function with the right signature into a update host ignition handler
type UpdateHostIgnitionHandlerFunc func(UpdateHostIgnitionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateHostIgnitionHandlerFunc) Handle(params UpdateHostIgnitionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateHostIgnitionHandler interface for that can handle valid update host ignition params
type UpdateHostIgnitionHandler interface {
	Handle(UpdateHostIgnitionParams, interface{}) middleware.Responder
}

// NewUpdateHostIgnition creates a new http.Handler for the update host ignition operation
func NewUpdateHostIgnition(ctx *middleware.Context, handler UpdateHostIgnitionHandler) *UpdateHostIgnition {
	return &UpdateHostIgnition{Context: ctx, Handler: handler}
}

/* UpdateHostIgnition swagger:route PATCH /v1/clusters/{cluster_id}/hosts/{host_id}/ignition installer updateHostIgnition

Patch the ignition file for this host

*/
type UpdateHostIgnition struct {
	Context *middleware.Context
	Handler UpdateHostIgnitionHandler
}

func (o *UpdateHostIgnition) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateHostIgnitionParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
