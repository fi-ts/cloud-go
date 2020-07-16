// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsV1MachineDiskPartition models v1 machine disk partition
// swagger:model models.V1MachineDiskPartition
type ModelsV1MachineDiskPartition struct {

	// device
	// Required: true
	Device *string `json:"device"`

	// filesystem
	// Required: true
	Filesystem *string `json:"filesystem"`

	// gptguid
	// Required: true
	Gptguid *string `json:"gptguid"`

	// gpttyoe
	// Required: true
	Gpttyoe *string `json:"gpttyoe"`

	// label
	// Required: true
	Label *string `json:"label"`

	// mountoptions
	// Required: true
	Mountoptions []string `json:"mountoptions"`

	// mountpoint
	// Required: true
	Mountpoint *string `json:"mountpoint"`

	// number
	// Required: true
	Number *int64 `json:"number"`

	// properties
	// Required: true
	Properties map[string]string `json:"properties"`

	// size
	// Required: true
	Size *int64 `json:"size"`
}

// Validate validates this models v1 machine disk partition
func (m *ModelsV1MachineDiskPartition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilesystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGptguid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpttyoe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountoptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1MachineDiskPartition) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineDiskPartition) validateFilesystem(formats strfmt.Registry) error {

	if err := validate.Required("filesystem", "body", m.Filesystem); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineDiskPartition) validateGptguid(formats strfmt.Registry) error {

	if err := validate.Required("gptguid", "body", m.Gptguid); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineDiskPartition) validateGpttyoe(formats strfmt.Registry) error {

	if err := validate.Required("gpttyoe", "body", m.Gpttyoe); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineDiskPartition) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineDiskPartition) validateMountoptions(formats strfmt.Registry) error {

	if err := validate.Required("mountoptions", "body", m.Mountoptions); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineDiskPartition) validateMountpoint(formats strfmt.Registry) error {

	if err := validate.Required("mountpoint", "body", m.Mountpoint); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineDiskPartition) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineDiskPartition) validateProperties(formats strfmt.Registry) error {

	return nil
}

func (m *ModelsV1MachineDiskPartition) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsV1MachineDiskPartition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsV1MachineDiskPartition) UnmarshalBinary(b []byte) error {
	var res ModelsV1MachineDiskPartition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}