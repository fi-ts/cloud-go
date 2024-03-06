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

// ModelsV1FirewallIngressRule models v1 firewall ingress rule
//
// swagger:model models.V1FirewallIngressRule
type ModelsV1FirewallIngressRule struct {

	// comment
	Comment string `json:"comment,omitempty"`

	// from
	// Required: true
	From []string `json:"from"`

	// ports
	// Required: true
	Ports []int32 `json:"ports"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// to
	// Required: true
	To []string `json:"to"`
}

// Validate validates this models v1 firewall ingress rule
func (m *ModelsV1FirewallIngressRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1FirewallIngressRule) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1FirewallIngressRule) validatePorts(formats strfmt.Registry) error {

	if err := validate.Required("ports", "body", m.Ports); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1FirewallIngressRule) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models v1 firewall ingress rule based on context it is used
func (m *ModelsV1FirewallIngressRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsV1FirewallIngressRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsV1FirewallIngressRule) UnmarshalBinary(b []byte) error {
	var res ModelsV1FirewallIngressRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
