// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CompleteInstallationHandlerFunc turns a function with the right signature into a complete installation handler
type CompleteInstallationHandlerFunc func(CompleteInstallationParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CompleteInstallationHandlerFunc) Handle(params CompleteInstallationParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CompleteInstallationHandler interface for that can handle valid complete installation params
type CompleteInstallationHandler interface {
	Handle(CompleteInstallationParams, interface{}) middleware.Responder
}

// NewCompleteInstallation creates a new http.Handler for the complete installation operation
func NewCompleteInstallation(ctx *middleware.Context, handler CompleteInstallationHandler) *CompleteInstallation {
	return &CompleteInstallation{Context: ctx, Handler: handler}
}

/* CompleteInstallation swagger:route POST /v1/clusters/{cluster_id}/actions/complete_installation installer completeInstallation

Agent API to mark a finalizing installation as complete.

*/
type CompleteInstallation struct {
	Context *middleware.Context
	Handler CompleteInstallationHandler
}

func (o *CompleteInstallation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCompleteInstallationParams()
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
