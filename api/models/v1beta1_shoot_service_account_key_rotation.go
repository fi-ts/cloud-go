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

// V1beta1ShootServiceAccountKeyRotation v1beta1 shoot service account key rotation
//
// swagger:model v1beta1.ShootServiceAccountKeyRotation
type V1beta1ShootServiceAccountKeyRotation struct {

	// last completion time
	LastCompletionTime string `json:"lastCompletionTime,omitempty"`

	// last initiation time
	LastInitiationTime string `json:"lastInitiationTime,omitempty"`

	// phase
	// Required: true
	Phase *string `json:"phase"`
}

// Validate validates this v1beta1 shoot service account key rotation
func (m *V1beta1ShootServiceAccountKeyRotation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1ShootServiceAccountKeyRotation) validatePhase(formats strfmt.Registry) error {

	if err := validate.Required("phase", "body", m.Phase); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1beta1 shoot service account key rotation based on context it is used
func (m *V1beta1ShootServiceAccountKeyRotation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1ShootServiceAccountKeyRotation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1ShootServiceAccountKeyRotation) UnmarshalBinary(b []byte) error {
	var res V1beta1ShootServiceAccountKeyRotation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}