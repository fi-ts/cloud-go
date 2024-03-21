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

// V1VolumeSetQoSPolicyRequest v1 volume set qo s policy request
//
// swagger:model v1.VolumeSetQoSPolicyRequest
type V1VolumeSetQoSPolicyRequest struct {

	// qo s policy ID
	// Required: true
	QoSPolicyID *string `json:"QoSPolicyID"`

	// qo s policy name
	// Required: true
	QoSPolicyName *string `json:"QoSPolicyName"`
}

// Validate validates this v1 volume set qo s policy request
func (m *V1VolumeSetQoSPolicyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQoSPolicyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQoSPolicyName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VolumeSetQoSPolicyRequest) validateQoSPolicyID(formats strfmt.Registry) error {

	if err := validate.Required("QoSPolicyID", "body", m.QoSPolicyID); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeSetQoSPolicyRequest) validateQoSPolicyName(formats strfmt.Registry) error {

	if err := validate.Required("QoSPolicyName", "body", m.QoSPolicyName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 volume set qo s policy request based on context it is used
func (m *V1VolumeSetQoSPolicyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VolumeSetQoSPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VolumeSetQoSPolicyRequest) UnmarshalBinary(b []byte) error {
	var res V1VolumeSetQoSPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
