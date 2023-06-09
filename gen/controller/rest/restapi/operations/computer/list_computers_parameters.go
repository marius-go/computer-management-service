// Code generated by go-swagger; DO NOT EDIT.

package computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewListComputersParams creates a new ListComputersParams object
//
// There are no default values defined in the spec.
func NewListComputersParams() ListComputersParams {

	return ListComputersParams{}
}

// ListComputersParams contains all the bound params for the list computers operation
// typically these are obtained from a http.Request
//
// swagger:parameters listComputers
type ListComputersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*filter computer by employee abbreviation
	  In: query
	*/
	EmployeeAbbreviation *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListComputersParams() beforehand.
func (o *ListComputersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qEmployeeAbbreviation, qhkEmployeeAbbreviation, _ := qs.GetOK("employeeAbbreviation")
	if err := o.bindEmployeeAbbreviation(qEmployeeAbbreviation, qhkEmployeeAbbreviation, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEmployeeAbbreviation binds and validates parameter EmployeeAbbreviation from query.
func (o *ListComputersParams) bindEmployeeAbbreviation(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.EmployeeAbbreviation = &raw

	return nil
}
