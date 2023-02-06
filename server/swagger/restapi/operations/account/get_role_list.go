// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"quocbang/swagger-microservices/swagger/models"
)

// GetRoleListHandlerFunc turns a function with the right signature into a get role list handler
type GetRoleListHandlerFunc func(GetRoleListParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRoleListHandlerFunc) Handle(params GetRoleListParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetRoleListHandler interface for that can handle valid get role list params
type GetRoleListHandler interface {
	Handle(GetRoleListParams, *models.Principal) middleware.Responder
}

// NewGetRoleList creates a new http.Handler for the get role list operation
func NewGetRoleList(ctx *middleware.Context, handler GetRoleListHandler) *GetRoleList {
	return &GetRoleList{Context: ctx, Handler: handler}
}

/* GetRoleList swagger:route GET /account/role-list account getRoleList

取得角色清單

*/
type GetRoleList struct {
	Context *middleware.Context
	Handler GetRoleListHandler
}

func (o *GetRoleList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetRoleListParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetRoleListOKBody get role list o k body
//
// swagger:model GetRoleListOKBody
type GetRoleListOKBody struct {

	// data
	Data []*GetRoleListOKBodyDataItems0 `json:"data"`
}

// Validate validates this get role list o k body
func (o *GetRoleListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRoleListOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRoleListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get role list o k body based on the context it is used
func (o *GetRoleListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRoleListOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRoleListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRoleListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoleListOKBody) UnmarshalBinary(b []byte) error {
	var res GetRoleListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetRoleListOKBodyDataItems0 get role list o k body data items0
//
// swagger:model GetRoleListOKBodyDataItems0
type GetRoleListOKBodyDataItems0 struct {

	// ID
	ID models.Role `json:"ID"`

	// 角色名稱
	// Example: SCHEDULER
	Name string `json:"name,omitempty"`
}

// Validate validates this get role list o k body data items0
func (o *GetRoleListOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRoleListOKBodyDataItems0) validateID(formats strfmt.Registry) error {
	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := o.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ID")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get role list o k body data items0 based on the context it is used
func (o *GetRoleListOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRoleListOKBodyDataItems0) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := o.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ID")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRoleListOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoleListOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res GetRoleListOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}