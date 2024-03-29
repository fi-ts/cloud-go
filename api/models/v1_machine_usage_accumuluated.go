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

// V1MachineUsageAccumuluated v1 machine usage accumuluated
//
// swagger:model v1.MachineUsageAccumuluated
type V1MachineUsageAccumuluated struct {

	// the duration that the machines is running
	// Required: true
	Lifetime *int64 `json:"lifetime"`
}

// Validate validates this v1 machine usage accumuluated
func (m *V1MachineUsageAccumuluated) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLifetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineUsageAccumuluated) validateLifetime(formats strfmt.Registry) error {

	if err := validate.Required("lifetime", "body", m.Lifetime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine usage accumuluated based on context it is used
func (m *V1MachineUsageAccumuluated) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineUsageAccumuluated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineUsageAccumuluated) UnmarshalBinary(b []byte) error {
	var res V1MachineUsageAccumuluated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
