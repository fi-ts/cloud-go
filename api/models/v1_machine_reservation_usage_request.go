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

// V1MachineReservationUsageRequest v1 machine reservation usage request
//
// swagger:model v1.MachineReservationUsageRequest
type V1MachineReservationUsageRequest struct {

	// the project id of this size reservation
	// Required: true
	Projectid *string `json:"projectid"`

	// the size id of this size reservation
	// Required: true
	Sizeid *string `json:"sizeid"`

	// the tenant of this size reservation
	// Required: true
	Tenant *string `json:"tenant"`
}

// Validate validates this v1 machine reservation usage request
func (m *V1MachineReservationUsageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSizeid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineReservationUsageRequest) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineReservationUsageRequest) validateSizeid(formats strfmt.Registry) error {

	if err := validate.Required("sizeid", "body", m.Sizeid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineReservationUsageRequest) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine reservation usage request based on context it is used
func (m *V1MachineReservationUsageRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineReservationUsageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineReservationUsageRequest) UnmarshalBinary(b []byte) error {
	var res V1MachineReservationUsageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}