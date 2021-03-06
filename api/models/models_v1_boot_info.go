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

// ModelsV1BootInfo models v1 boot info
//
// swagger:model models.V1BootInfo
type ModelsV1BootInfo struct {

	// bootloaderid
	// Required: true
	Bootloaderid *string `json:"bootloaderid"`

	// cmdline
	// Required: true
	Cmdline *string `json:"cmdline"`

	// image id
	// Required: true
	ImageID *string `json:"image_id"`

	// initrd
	// Required: true
	Initrd *string `json:"initrd"`

	// kernel
	// Required: true
	Kernel *string `json:"kernel"`

	// os partition
	// Required: true
	OsPartition *string `json:"os_partition"`

	// primary disk
	// Required: true
	PrimaryDisk *string `json:"primary_disk"`
}

// Validate validates this models v1 boot info
func (m *ModelsV1BootInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootloaderid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCmdline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitrd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKernel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsPartition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryDisk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1BootInfo) validateBootloaderid(formats strfmt.Registry) error {

	if err := validate.Required("bootloaderid", "body", m.Bootloaderid); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1BootInfo) validateCmdline(formats strfmt.Registry) error {

	if err := validate.Required("cmdline", "body", m.Cmdline); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1BootInfo) validateImageID(formats strfmt.Registry) error {

	if err := validate.Required("image_id", "body", m.ImageID); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1BootInfo) validateInitrd(formats strfmt.Registry) error {

	if err := validate.Required("initrd", "body", m.Initrd); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1BootInfo) validateKernel(formats strfmt.Registry) error {

	if err := validate.Required("kernel", "body", m.Kernel); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1BootInfo) validateOsPartition(formats strfmt.Registry) error {

	if err := validate.Required("os_partition", "body", m.OsPartition); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1BootInfo) validatePrimaryDisk(formats strfmt.Registry) error {

	if err := validate.Required("primary_disk", "body", m.PrimaryDisk); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models v1 boot info based on context it is used
func (m *ModelsV1BootInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsV1BootInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsV1BootInfo) UnmarshalBinary(b []byte) error {
	var res ModelsV1BootInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
