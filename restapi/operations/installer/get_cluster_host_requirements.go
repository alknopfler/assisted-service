// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetClusterHostRequirementsHandlerFunc turns a function with the right signature into a get cluster host requirements handler
type GetClusterHostRequirementsHandlerFunc func(GetClusterHostRequirementsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClusterHostRequirementsHandlerFunc) Handle(params GetClusterHostRequirementsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetClusterHostRequirementsHandler interface for that can handle valid get cluster host requirements params
type GetClusterHostRequirementsHandler interface {
	Handle(GetClusterHostRequirementsParams, interface{}) middleware.Responder
}

// NewGetClusterHostRequirements creates a new http.Handler for the get cluster host requirements operation
func NewGetClusterHostRequirements(ctx *middleware.Context, handler GetClusterHostRequirementsHandler) *GetClusterHostRequirements {
	return &GetClusterHostRequirements{Context: ctx, Handler: handler}
}

/*GetClusterHostRequirements swagger:route GET /v1/clusters/{cluster_id}/host-requirements installer getClusterHostRequirements

Get host requirements of a cluster.

*/
type GetClusterHostRequirements struct {
	Context *middleware.Context
	Handler GetClusterHostRequirementsHandler
}

func (o *GetClusterHostRequirements) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetClusterHostRequirementsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
