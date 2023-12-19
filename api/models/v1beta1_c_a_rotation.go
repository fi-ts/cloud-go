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

// V1beta1CARotation v1beta1 c a rotation
//
// swagger:model v1beta1.CARotation
type V1beta1CARotation struct {

	// last completion time
	LastCompletionTime string `json:"lastCompletionTime,omitempty"`

	// last completion triggered time
	LastCompletionTriggeredTime string `json:"lastCompletionTriggeredTime,omitempty"`

	// last initiation finished time
	LastInitiationFinishedTime string `json:"lastInitiationFinishedTime,omitempty"`

	// last initiation time
	LastInitiationTime string `json:"lastInitiationTime,omitempty"`

	// phase
	// Required: true
	Phase *string `json:"phase"`
}

// Validate validates this v1beta1 c a rotation
func (m *V1beta1CARotation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1CARotation) validatePhase(formats strfmt.Registry) error {

	if err := validate.Required("phase", "body", m.Phase); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1beta1 c a rotation based on context it is used
func (m *V1beta1CARotation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1CARotation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1CARotation) UnmarshalBinary(b []byte) error {
	var res V1beta1CARotation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}