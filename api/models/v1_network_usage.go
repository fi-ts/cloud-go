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

// V1NetworkUsage v1 network usage
//
// swagger:model v1.NetworkUsage
type V1NetworkUsage struct {

	// the cluster id of this network device
	// Required: true
	Clusterid *string `json:"clusterid"`

	// the cluster name of this network device
	// Required: true
	Clustername *string `json:"clustername"`

	// the device name of this network device
	// Required: true
	Device *string `json:"device"`

	// the ingoing traffic of this network device (byte)
	// Required: true
	In *string `json:"in"`

	// the duration that this network device is running
	// Required: true
	Lifetime *int64 `json:"lifetime"`

	// the outgoing traffic of this network device (byte)
	// Required: true
	Out *string `json:"out"`

	// the partition of this network device
	// Required: true
	Partition *string `json:"partition"`

	// the project id of this network device
	// Required: true
	Projectid *string `json:"projectid"`

	// the project name of this network device
	// Required: true
	Projectname *string `json:"projectname"`

	// the tenant of this network device
	// Required: true
	Tenant *string `json:"tenant"`

	// the total traffic of this network device (byte)
	// Required: true
	Total *string `json:"total"`

	// warnings that occurred when calculating the usage of this device's network traffic
	// Required: true
	Warnings []string `json:"warnings"`
}

// Validate validates this v1 network usage
func (m *V1NetworkUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOut(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarnings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NetworkUsage) validateClusterid(formats strfmt.Registry) error {

	if err := validate.Required("clusterid", "body", m.Clusterid); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateClustername(formats strfmt.Registry) error {

	if err := validate.Required("clustername", "body", m.Clustername); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateIn(formats strfmt.Registry) error {

	if err := validate.Required("in", "body", m.In); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateLifetime(formats strfmt.Registry) error {

	if err := validate.Required("lifetime", "body", m.Lifetime); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateOut(formats strfmt.Registry) error {

	if err := validate.Required("out", "body", m.Out); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validatePartition(formats strfmt.Registry) error {

	if err := validate.Required("partition", "body", m.Partition); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateProjectname(formats strfmt.Registry) error {

	if err := validate.Required("projectname", "body", m.Projectname); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateWarnings(formats strfmt.Registry) error {

	if err := validate.Required("warnings", "body", m.Warnings); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 network usage based on context it is used
func (m *V1NetworkUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NetworkUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NetworkUsage) UnmarshalBinary(b []byte) error {
	var res V1NetworkUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
