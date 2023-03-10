// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// LoginType 登入類別:
//   * 0 - MES
//   * 1 - Windows
//
//
// swagger:model LoginType
type LoginType int64

// for schema
var loginTypeEnum []interface{}

func init() {
	var res []LoginType
	if err := json.Unmarshal([]byte(`[0,1]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		loginTypeEnum = append(loginTypeEnum, v)
	}
}

func (m LoginType) validateLoginTypeEnum(path, location string, value LoginType) error {
	if err := validate.EnumCase(path, location, value, loginTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this login type
func (m LoginType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLoginTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this login type based on context it is used
func (m LoginType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
