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

// ModelsV1VolumeGroup models v1 volume group
//
// swagger:model models.V1VolumeGroup
type ModelsV1VolumeGroup struct {

	// devices
	// Required: true
	Devices []string `json:"devices"`

	// name
	// Required: true
	Name *string `json:"name"`

	// tags
	// Required: true
	Tags []string `json:"tags"`
}

// Validate validates this models v1 volume group
func (m *ModelsV1VolumeGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1VolumeGroup) validateDevices(formats strfmt.Registry) error {

	if err := validate.Required("devices", "body", m.Devices); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1VolumeGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1VolumeGroup) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models v1 volume group based on context it is used
func (m *ModelsV1VolumeGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsV1VolumeGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsV1VolumeGroup) UnmarshalBinary(b []byte) error {
	var res ModelsV1VolumeGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}