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

// V1ClusterWorker v1 cluster worker
//
// swagger:model v1.ClusterWorker
type V1ClusterWorker struct {

	// the machine count of this worker group
	// Required: true
	Machinecount *int64 `json:"machinecount"`

	// the machine type
	// Required: true
	Machinetype *string `json:"machinetype"`

	// name of the worker group
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this v1 cluster worker
func (m *V1ClusterWorker) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachinecount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachinetype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterWorker) validateMachinecount(formats strfmt.Registry) error {

	if err := validate.Required("machinecount", "body", m.Machinecount); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterWorker) validateMachinetype(formats strfmt.Registry) error {

	if err := validate.Required("machinetype", "body", m.Machinetype); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterWorker) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 cluster worker based on context it is used
func (m *V1ClusterWorker) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterWorker) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterWorker) UnmarshalBinary(b []byte) error {
	var res V1ClusterWorker
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
