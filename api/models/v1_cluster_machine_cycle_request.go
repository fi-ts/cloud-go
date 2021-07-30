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

// V1ClusterMachineCycleRequest v1 cluster machine cycle request
//
// swagger:model v1.ClusterMachineCycleRequest
type V1ClusterMachineCycleRequest struct {

	// uuid of the machine to power cycle
	// Required: true
	Machineid *string `json:"machineid"`
}

// Validate validates this v1 cluster machine cycle request
func (m *V1ClusterMachineCycleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachineid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterMachineCycleRequest) validateMachineid(formats strfmt.Registry) error {

	if err := validate.Required("machineid", "body", m.Machineid); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 cluster machine cycle request based on context it is used
func (m *V1ClusterMachineCycleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterMachineCycleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterMachineCycleRequest) UnmarshalBinary(b []byte) error {
	var res V1ClusterMachineCycleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}