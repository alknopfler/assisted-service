// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ResetHostValidationHandlerFunc turns a function with the right signature into a reset host validation handler
type ResetHostValidationHandlerFunc func(ResetHostValidationParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ResetHostValidationHandlerFunc) Handle(params ResetHostValidationParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ResetHostValidationHandler interface for that can handle valid reset host validation params
type ResetHostValidationHandler interface {
	Handle(ResetHostValidationParams, interface{}) middleware.Responder
}

// NewResetHostValidation creates a new http.Handler for the reset host validation operation
func NewResetHostValidation(ctx *middleware.Context, handler ResetHostValidationHandler) *ResetHostValidation {
	return &ResetHostValidation{Context: ctx, Handler: handler}
}

/* ResetHostValidation swagger:route PATCH /v1/clusters/{cluster_id}/hosts/{host_id}/actions/reset-validation/{validation_id} installer resetHostValidation

Reset failed host validation.

Reset failed host validation.  It may be performed on any host validation with persistent validation result.

*/
type ResetHostValidation struct {
	Context *middleware.Context
	Handler ResetHostValidationHandler
}

func (o *ResetHostValidation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewResetHostValidationParams()
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
