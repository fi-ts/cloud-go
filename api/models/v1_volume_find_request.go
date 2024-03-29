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

// V1VolumeFindRequest v1 volume find request
//
// swagger:model v1.VolumeFindRequest
type V1VolumeFindRequest struct {

	// partition ID
	// Required: true
	PartitionID *string `json:"PartitionID"`

	// project ID
	// Required: true
	ProjectID *string `json:"ProjectID"`

	// tenant ID
	// Required: true
	TenantID *string `json:"TenantID"`

	// volume ID
	// Required: true
	VolumeID *string `json:"VolumeID"`
}

// Validate validates this v1 volume find request
func (m *V1VolumeFindRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePartitionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VolumeFindRequest) validatePartitionID(formats strfmt.Registry) error {

	if err := validate.Required("PartitionID", "body", m.PartitionID); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeFindRequest) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("ProjectID", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeFindRequest) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("TenantID", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeFindRequest) validateVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("VolumeID", "body", m.VolumeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 volume find request based on context it is used
func (m *V1VolumeFindRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VolumeFindRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VolumeFindRequest) UnmarshalBinary(b []byte) error {
	var res V1VolumeFindRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
