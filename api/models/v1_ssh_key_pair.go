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

// V1SSHKeyPair v1 SSH key pair
//
// swagger:model v1.SSHKeyPair
type V1SSHKeyPair struct {

	// private key
	// Required: true
	PrivateKey *string `json:"PrivateKey"`

	// public key
	// Required: true
	PublicKey *string `json:"PublicKey"`
}

// Validate validates this v1 SSH key pair
func (m *V1SSHKeyPair) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SSHKeyPair) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("PrivateKey", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *V1SSHKeyPair) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("PublicKey", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 SSH key pair based on context it is used
func (m *V1SSHKeyPair) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SSHKeyPair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SSHKeyPair) UnmarshalBinary(b []byte) error {
	var res V1SSHKeyPair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
