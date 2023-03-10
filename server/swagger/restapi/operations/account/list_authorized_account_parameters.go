// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewListAuthorizedAccountParams creates a new ListAuthorizedAccountParams object
//
// There are no default values defined in the spec.
func NewListAuthorizedAccountParams() ListAuthorizedAccountParams {

	return ListAuthorizedAccountParams{}
}

// ListAuthorizedAccountParams contains all the bound params for the list authorized account operation
// typically these are obtained from a http.Request
//
// swagger:parameters ListAuthorizedAccount
type ListAuthorizedAccountParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*部門OID
	  Required: true
	  In: path
	*/
	DepartmentOID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListAuthorizedAccountParams() beforehand.
func (o *ListAuthorizedAccountParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDepartmentOID, rhkDepartmentOID, _ := route.Params.GetOK("departmentOID")
	if err := o.bindDepartmentOID(rDepartmentOID, rhkDepartmentOID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDepartmentOID binds and validates parameter DepartmentOID from path.
func (o *ListAuthorizedAccountParams) bindDepartmentOID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.DepartmentOID = raw

	return nil
}
