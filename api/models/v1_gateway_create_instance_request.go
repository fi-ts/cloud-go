// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1GatewayCreateInstanceRequest v1 gateway create instance request
//
// swagger:model v1.GatewayCreateInstanceRequest
type V1GatewayCreateInstanceRequest struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// port
	Port int64 `json:"port,omitempty"`

	// project UID
	// Required: true
	ProjectUID *string `json:"projectUID"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this v1 gateway create instance request
func (m *V1GatewayCreateInstanceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GatewayCreateInstanceRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1GatewayCreateInstanceRequest) validateProjectUID(formats strfmt.Registry) error {

	if err := validate.Required("projectUID", "body", m.ProjectUID); err != nil {
		return err
	}

	return nil
}

func (m *V1GatewayCreateInstanceRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 gateway create instance request based on context it is used
func (m *V1GatewayCreateInstanceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GatewayCreateInstanceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GatewayCreateInstanceRequest) UnmarshalBinary(b []byte) error {
	var res V1GatewayCreateInstanceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}