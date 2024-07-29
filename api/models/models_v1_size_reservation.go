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

// ModelsV1SizeReservation models v1 size reservation
//
// swagger:model models.V1SizeReservation
type ModelsV1SizeReservation struct {

	// amount
	// Required: true
	Amount *int32 `json:"amount"`

	// description
	Description string `json:"description,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// partitionids
	// Required: true
	Partitionids []string `json:"partitionids"`

	// projectid
	// Required: true
	Projectid *string `json:"projectid"`
}

// Validate validates this models v1 size reservation
func (m *ModelsV1SizeReservation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitionids(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1SizeReservation) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1SizeReservation) validatePartitionids(formats strfmt.Registry) error {

	if err := validate.Required("partitionids", "body", m.Partitionids); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1SizeReservation) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models v1 size reservation based on context it is used
func (m *ModelsV1SizeReservation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsV1SizeReservation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsV1SizeReservation) UnmarshalBinary(b []byte) error {
	var res ModelsV1SizeReservation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
