// Code generated by go-swagger; DO NOT EDIT.

package computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDeleteComputerParams creates a new DeleteComputerParams object
//
// There are no default values defined in the spec.
func NewDeleteComputerParams() DeleteComputerParams {

	return DeleteComputerParams{}
}

// DeleteComputerParams contains all the bound params for the delete computer operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteComputer
type DeleteComputerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*name of the computer to be removed from the service
	  Required: true
	  In: path
	*/
	ComputerName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteComputerParams() beforehand.
func (o *DeleteComputerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rComputerName, rhkComputerName, _ := route.Params.GetOK("computerName")
	if err := o.bindComputerName(rComputerName, rhkComputerName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindComputerName binds and validates parameter ComputerName from path.
func (o *DeleteComputerParams) bindComputerName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ComputerName = raw

	return nil
}