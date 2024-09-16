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

// V1MachineReservationCreateRequest v1 machine reservation create request
//
// swagger:model v1.MachineReservationCreateRequest
type V1MachineReservationCreateRequest struct {

	// amount
	// Required: true
	Amount *int32 `json:"amount"`

	// description
	// Required: true
	Description *string `json:"description"`

	// partitionids
	// Required: true
	Partitionids []string `json:"partitionids"`

	// projectid
	// Required: true
	Projectid *string `json:"projectid"`

	// sizeid
	// Required: true
	Sizeid *string `json:"sizeid"`
}

// Validate validates this v1 machine reservation create request
func (m *V1MachineReservationCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineReservationCreateRequest) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineReservationCreateRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineReservationCreateRequest) validatePartitionids(formats strfmt.Registry) error {

	if err := validate.Required("partitionids", "body", m.Partitionids); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineReservationCreateRequest) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineReservationCreateRequest) validateSizeid(formats strfmt.Registry) error {

	if err := validate.Required("sizeid", "body", m.Sizeid); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine reservation create request based on context it is used
func (m *V1MachineReservationCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineReservationCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineReservationCreateRequest) UnmarshalBinary(b []byte) error {
	var res V1MachineReservationCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}