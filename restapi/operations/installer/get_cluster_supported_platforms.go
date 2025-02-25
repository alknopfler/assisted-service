// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetClusterSupportedPlatformsHandlerFunc turns a function with the right signature into a get cluster supported platforms handler
type GetClusterSupportedPlatformsHandlerFunc func(GetClusterSupportedPlatformsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClusterSupportedPlatformsHandlerFunc) Handle(params GetClusterSupportedPlatformsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetClusterSupportedPlatformsHandler interface for that can handle valid get cluster supported platforms params
type GetClusterSupportedPlatformsHandler interface {
	Handle(GetClusterSupportedPlatformsParams, interface{}) middleware.Responder
}

// NewGetClusterSupportedPlatforms creates a new http.Handler for the get cluster supported platforms operation
func NewGetClusterSupportedPlatforms(ctx *middleware.Context, handler GetClusterSupportedPlatformsHandler) *GetClusterSupportedPlatforms {
	return &GetClusterSupportedPlatforms{Context: ctx, Handler: handler}
}

/* GetClusterSupportedPlatforms swagger:route GET /v2/clusters/{cluster_id}/supported-platforms installer getClusterSupportedPlatforms

A list of platforms that this cluster can support in its current configuration.

*/
type GetClusterSupportedPlatforms struct {
	Context *middleware.Context
	Handler GetClusterSupportedPlatformsHandler
}

func (o *GetClusterSupportedPlatforms) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetClusterSupportedPlatformsParams()
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
