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

// V1MachineUsageRequest v1 machine usage request
//
// swagger:model v1.MachineUsageRequest
type V1MachineUsageRequest struct {

	// the cluster id to account for
	Clusterid string `json:"clusterid,omitempty"`

	// the start time in the accounting window to look at
	// Required: true
	// Format: date-time
	From *strfmt.DateTime `json:"from"`

	// the id of the machine
	ID string `json:"id,omitempty"`

	// the partition of the machines
	Partition string `json:"partition,omitempty"`

	// the project id to account for
	Projectid string `json:"projectid,omitempty"`

	// the size of the machines
	Sizeid string `json:"sizeid,omitempty"`

	// the tenant to get the container usage for (defaults to all tenants)
	Tenant string `json:"tenant,omitempty"`

	// the end time in the accounting window to look at (defaults to current system time)
	// Format: date-time
	To strfmt.DateTime `json:"to,omitempty"`
}

// Validate validates this v1 machine usage request
func (m *V1MachineUsageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineUsageRequest) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	if err := validate.FormatOf("from", "body", "date-time", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsageRequest) validateTo(formats strfmt.Registry) error {
	if swag.IsZero(m.To) { // not required
		return nil
	}

	if err := validate.FormatOf("to", "body", "date-time", m.To.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine usage request based on context it is used
func (m *V1MachineUsageRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineUsageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineUsageRequest) UnmarshalBinary(b []byte) error {
	var res V1MachineUsageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
