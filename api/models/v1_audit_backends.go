// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AuditBackends v1 audit backends
//
// swagger:model v1.AuditBackends
type V1AuditBackends struct {

	// cluster forwarding
	ClusterForwarding *V1AuditBackendClusterForwarding `json:"clusterForwarding,omitempty"`

	// splunk
	Splunk *V1AuditBackendSplunk `json:"splunk,omitempty"`
}

// Validate validates this v1 audit backends
func (m *V1AuditBackends) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterForwarding(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSplunk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AuditBackends) validateClusterForwarding(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterForwarding) { // not required
		return nil
	}

	if m.ClusterForwarding != nil {
		if err := m.ClusterForwarding.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterForwarding")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterForwarding")
			}
			return err
		}
	}

	return nil
}

func (m *V1AuditBackends) validateSplunk(formats strfmt.Registry) error {
	if swag.IsZero(m.Splunk) { // not required
		return nil
	}

	if m.Splunk != nil {
		if err := m.Splunk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("splunk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("splunk")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 audit backends based on the context it is used
func (m *V1AuditBackends) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterForwarding(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSplunk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AuditBackends) contextValidateClusterForwarding(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterForwarding != nil {

		if swag.IsZero(m.ClusterForwarding) { // not required
			return nil
		}

		if err := m.ClusterForwarding.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterForwarding")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterForwarding")
			}
			return err
		}
	}

	return nil
}

func (m *V1AuditBackends) contextValidateSplunk(ctx context.Context, formats strfmt.Registry) error {

	if m.Splunk != nil {

		if swag.IsZero(m.Splunk) { // not required
			return nil
		}

		if err := m.Splunk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("splunk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("splunk")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AuditBackends) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AuditBackends) UnmarshalBinary(b []byte) error {
	var res V1AuditBackends
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
