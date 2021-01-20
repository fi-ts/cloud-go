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

// V1Kubernetes v1 kubernetes
//
// swagger:model v1.Kubernetes
type V1Kubernetes struct {

	// allow privileged containers
	// Required: true
	AllowPrivilegedContainers *bool `json:"AllowPrivilegedContainers"`

	// expiration date
	// Required: true
	// Format: date-time
	ExpirationDate *strfmt.DateTime `json:"ExpirationDate"`

	// version
	// Required: true
	Version *string `json:"Version"`
}

// Validate validates this v1 kubernetes
func (m *V1Kubernetes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowPrivilegedContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Kubernetes) validateAllowPrivilegedContainers(formats strfmt.Registry) error {

	if err := validate.Required("AllowPrivilegedContainers", "body", m.AllowPrivilegedContainers); err != nil {
		return err
	}

	return nil
}

func (m *V1Kubernetes) validateExpirationDate(formats strfmt.Registry) error {

	if err := validate.Required("ExpirationDate", "body", m.ExpirationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("ExpirationDate", "body", "date-time", m.ExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Kubernetes) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("Version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 kubernetes based on context it is used
func (m *V1Kubernetes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Kubernetes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Kubernetes) UnmarshalBinary(b []byte) error {
	var res V1Kubernetes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
