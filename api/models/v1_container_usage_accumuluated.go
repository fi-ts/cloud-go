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

// V1ContainerUsageAccumuluated v1 container usage accumuluated
//
// swagger:model v1.ContainerUsageAccumuluated
type V1ContainerUsageAccumuluated struct {

	// the accumuluated cpu seconds of the containers in this response (s*s)
	// Required: true
	Cpuseconds *string `json:"cpuseconds"`

	// the duration that this container is running
	// Required: true
	Lifetime *int64 `json:"lifetime"`

	// the accumulated memory seconds of the containers in this response (byte*s)
	// Required: true
	Memoryseconds *string `json:"memoryseconds"`
}

// Validate validates this v1 container usage accumuluated
func (m *V1ContainerUsageAccumuluated) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCpuseconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryseconds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ContainerUsageAccumuluated) validateCpuseconds(formats strfmt.Registry) error {

	if err := validate.Required("cpuseconds", "body", m.Cpuseconds); err != nil {
		return err
	}

	return nil
}

func (m *V1ContainerUsageAccumuluated) validateLifetime(formats strfmt.Registry) error {

	if err := validate.Required("lifetime", "body", m.Lifetime); err != nil {
		return err
	}

	return nil
}

func (m *V1ContainerUsageAccumuluated) validateMemoryseconds(formats strfmt.Registry) error {

	if err := validate.Required("memoryseconds", "body", m.Memoryseconds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 container usage accumuluated based on context it is used
func (m *V1ContainerUsageAccumuluated) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ContainerUsageAccumuluated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ContainerUsageAccumuluated) UnmarshalBinary(b []byte) error {
	var res V1ContainerUsageAccumuluated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
