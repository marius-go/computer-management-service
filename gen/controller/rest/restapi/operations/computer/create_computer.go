// Code generated by go-swagger; DO NOT EDIT.

package computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateComputerHandlerFunc turns a function with the right signature into a create computer handler
type CreateComputerHandlerFunc func(CreateComputerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateComputerHandlerFunc) Handle(params CreateComputerParams) middleware.Responder {
	return fn(params)
}

// CreateComputerHandler interface for that can handle valid create computer params
type CreateComputerHandler interface {
	Handle(CreateComputerParams) middleware.Responder
}

// NewCreateComputer creates a new http.Handler for the create computer operation
func NewCreateComputer(ctx *middleware.Context, handler CreateComputerHandler) *CreateComputer {
	return &CreateComputer{Context: ctx, Handler: handler}
}

/*
	CreateComputer swagger:route POST /computer computer createComputer

create a new computer entry
*/
type CreateComputer struct {
	Context *middleware.Context
	Handler CreateComputerHandler
}

func (o *CreateComputer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateComputerParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}