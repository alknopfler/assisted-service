// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListSupportedOpenshiftVersionsHandlerFunc turns a function with the right signature into a list supported openshift versions handler
type ListSupportedOpenshiftVersionsHandlerFunc func(ListSupportedOpenshiftVersionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListSupportedOpenshiftVersionsHandlerFunc) Handle(params ListSupportedOpenshiftVersionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListSupportedOpenshiftVersionsHandler interface for that can handle valid list supported openshift versions params
type ListSupportedOpenshiftVersionsHandler interface {
	Handle(ListSupportedOpenshiftVersionsParams, interface{}) middleware.Responder
}

// NewListSupportedOpenshiftVersions creates a new http.Handler for the list supported openshift versions operation
func NewListSupportedOpenshiftVersions(ctx *middleware.Context, handler ListSupportedOpenshiftVersionsHandler) *ListSupportedOpenshiftVersions {
	return &ListSupportedOpenshiftVersions{Context: ctx, Handler: handler}
}

/*ListSupportedOpenshiftVersions swagger:route GET /v1/openshift_versions versions listSupportedOpenshiftVersions

Retrieves the list of OpenShift supported versions.

*/
type ListSupportedOpenshiftVersions struct {
	Context *middleware.Context
	Handler ListSupportedOpenshiftVersionsHandler
}

func (o *ListSupportedOpenshiftVersions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListSupportedOpenshiftVersionsParams()

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
