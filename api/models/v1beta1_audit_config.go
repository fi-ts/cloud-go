// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1beta1AuditConfig v1beta1 audit config
// swagger:model v1beta1.AuditConfig
type V1beta1AuditConfig struct {

	// audit policy
	AuditPolicy *V1beta1AuditPolicy `json:"auditPolicy,omitempty"`
}

// Validate validates this v1beta1 audit config
func (m *V1beta1AuditConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1AuditConfig) validateAuditPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditPolicy) { // not required
		return nil
	}

	if m.AuditPolicy != nil {
		if err := m.AuditPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auditPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1AuditConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1AuditConfig) UnmarshalBinary(b []byte) error {
	var res V1beta1AuditConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}