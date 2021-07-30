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

// V1FirewallControllerVersion v1 firewall controller version
//
// swagger:model v1.FirewallControllerVersion
type V1FirewallControllerVersion struct {

	// URL
	// Required: true
	URL *string `json:"URL"`

	// version
	// Required: true
	Version *string `json:"Version"`
}

// Validate validates this v1 firewall controller version
func (m *V1FirewallControllerVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1FirewallControllerVersion) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("URL", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *V1FirewallControllerVersion) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("Version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 firewall controller version based on context it is used
func (m *V1FirewallControllerVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1FirewallControllerVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FirewallControllerVersion) UnmarshalBinary(b []byte) error {
	var res V1FirewallControllerVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}