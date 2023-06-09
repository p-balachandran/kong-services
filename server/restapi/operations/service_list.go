// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ServiceListHandlerFunc turns a function with the right signature into a service list handler
type ServiceListHandlerFunc func(ServiceListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ServiceListHandlerFunc) Handle(params ServiceListParams) middleware.Responder {
	return fn(params)
}

// ServiceListHandler interface for that can handle valid service list params
type ServiceListHandler interface {
	Handle(ServiceListParams) middleware.Responder
}

// NewServiceList creates a new http.Handler for the service list operation
func NewServiceList(ctx *middleware.Context, handler ServiceListHandler) *ServiceList {
	return &ServiceList{Context: ctx, Handler: handler}
}

/*
	ServiceList swagger:route GET /services serviceList

# Get all services

It returns a list of services
*/
type ServiceList struct {
	Context *middleware.Context
	Handler ServiceListHandler
}

func (o *ServiceList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewServiceListParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
