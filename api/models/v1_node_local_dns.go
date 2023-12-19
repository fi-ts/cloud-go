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

// V1NodeLocalDNS v1 node local DNS
//
// swagger:model v1.NodeLocalDNS
type V1NodeLocalDNS struct {

	// disable forward to upstream DNS
	// Required: true
	DisableForwardToUpstreamDNS *bool `json:"DisableForwardToUpstreamDNS"`

	// enabled
	// Required: true
	Enabled *bool `json:"Enabled"`
}

// Validate validates this v1 node local DNS
func (m *V1NodeLocalDNS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisableForwardToUpstreamDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NodeLocalDNS) validateDisableForwardToUpstreamDNS(formats strfmt.Registry) error {

	if err := validate.Required("DisableForwardToUpstreamDNS", "body", m.DisableForwardToUpstreamDNS); err != nil {
		return err
	}

	return nil
}

func (m *V1NodeLocalDNS) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("Enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 node local DNS based on context it is used
func (m *V1NodeLocalDNS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NodeLocalDNS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NodeLocalDNS) UnmarshalBinary(b []byte) error {
	var res V1NodeLocalDNS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}