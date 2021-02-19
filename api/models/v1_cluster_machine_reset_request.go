// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ClusterMachineResetRequest v1 cluster machine reset request
//
// swagger:model v1.ClusterMachineResetRequest
type V1ClusterMachineResetRequest struct {

	// uuid of the machine to power reset
	// Required: true
	Machineid *string `json:"machineid"`
}

// Validate validates this v1 cluster machine reset request
func (m *V1ClusterMachineResetRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachineid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterMachineResetRequest) validateMachineid(formats strfmt.Registry) error {

	if err := validate.Required("machineid", "body", m.Machineid); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterMachineResetRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterMachineResetRequest) UnmarshalBinary(b []byte) error {
	var res V1ClusterMachineResetRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}