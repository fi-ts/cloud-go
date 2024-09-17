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

// V1MachineReservationResponse v1 machine reservation response
//
// swagger:model v1.MachineReservationResponse
type V1MachineReservationResponse struct {

	// amount
	// Required: true
	Amount *int32 `json:"amount"`

	// description
	Description string `json:"description,omitempty"`

	// labels
	// Required: true
	Labels map[string]string `json:"labels"`

	// partitionids
	// Required: true
	Partitionids []string `json:"partitionids"`

	// projectid
	// Required: true
	Projectid *string `json:"projectid"`

	// sizeid
	// Required: true
	Sizeid *string `json:"sizeid"`

	// tenant
	// Required: true
	Tenant *string `json:"tenant"`
}

// Validate validates this v1 machine reservation response
func (m *V1MachineReservationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitionids(formats); err != nil {
		res = append(res, err)
	}

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

func (m *V1MachineReservationResponse) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineReservationResponse) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("labels", "body", m.Labels); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineReservationResponse) validatePartitionids(formats strfmt.Registry) error {

	if err := validate.Required("partitionids", "body", m.Partitionids); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineReservationResponse) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineReservationResponse) validateSizeid(formats strfmt.Registry) error {

	if err := validate.Required("sizeid", "body", m.Sizeid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineReservationResponse) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine reservation response based on context it is used
func (m *V1MachineReservationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineReservationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineReservationResponse) UnmarshalBinary(b []byte) error {
	var res V1MachineReservationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
