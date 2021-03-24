// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1IAMConfig v1 i a m config
//
// swagger:model v1.IAMConfig
type V1IAMConfig struct {

	// idm config
	IdmConfig *V1IDMConfig `json:"idm_config,omitempty"`

	// issuer config
	IssuerConfig *V1IssuerConfig `json:"issuer_config,omitempty"`
}

// Validate validates this v1 i a m config
func (m *V1IAMConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdmConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuerConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IAMConfig) validateIdmConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.IdmConfig) { // not required
		return nil
	}

	if m.IdmConfig != nil {
		if err := m.IdmConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("idm_config")
			}
			return err
		}
	}

	return nil
}

func (m *V1IAMConfig) validateIssuerConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.IssuerConfig) { // not required
		return nil
	}

	if m.IssuerConfig != nil {
		if err := m.IssuerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issuer_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IAMConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IAMConfig) UnmarshalBinary(b []byte) error {
	var res V1IAMConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
