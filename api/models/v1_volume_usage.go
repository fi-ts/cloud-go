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

// V1VolumeUsage v1 volume usage
//
// swagger:model v1.VolumeUsage
type V1VolumeUsage struct {

	// the capacity seconds of this volume (byte*s)
	// Required: true
	Capacityseconds *string `json:"capacityseconds"`

	// the class of this volume
	// Required: true
	Class *string `json:"class"`

	// the cluster id of this volume
	// Required: true
	Clusterid *string `json:"clusterid"`

	// the cluster name of this volume
	// Required: true
	Clustername *string `json:"clustername"`

	// the end time of this volume
	// Required: true
	// Format: date-time
	End *strfmt.DateTime `json:"end"`

	// the duration that this volume is running
	// Required: true
	Lifetime *int64 `json:"lifetime"`

	// the name of this volume
	// Required: true
	Name *string `json:"name"`

	// the partition of this volume
	// Required: true
	Partition *string `json:"partition"`

	// the project id of this volume
	// Required: true
	Projectid *string `json:"projectid"`

	// the project name of this volume
	// Required: true
	Projectname *string `json:"projectname"`

	// the start time of this volume
	// Required: true
	// Format: date-time
	Start *strfmt.DateTime `json:"start"`

	// the tenant of this volume
	// Required: true
	Tenant *string `json:"tenant"`

	// the type of this volume
	// Required: true
	Type *string `json:"type"`

	// warnings that occurred when calculating the usage of this volume
	// Required: true
	Warnings []string `json:"warnings"`
}

// Validate validates this v1 volume usage
func (m *V1VolumeUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacityseconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *V1VolumeUsage) validateCapacityseconds(formats strfmt.Registry) error {

	if err := validate.Required("capacityseconds", "body", m.Capacityseconds); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validateClass(formats strfmt.Registry) error {

	if err := validate.Required("class", "body", m.Class); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validateClusterid(formats strfmt.Registry) error {

	if err := validate.Required("clusterid", "body", m.Clusterid); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validateClustername(formats strfmt.Registry) error {

	if err := validate.Required("clustername", "body", m.Clustername); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validateLifetime(formats strfmt.Registry) error {

	if err := validate.Required("lifetime", "body", m.Lifetime); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validatePartition(formats strfmt.Registry) error {

	if err := validate.Required("partition", "body", m.Partition); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validateProjectname(formats strfmt.Registry) error {

	if err := validate.Required("projectname", "body", m.Projectname); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsage) validateWarnings(formats strfmt.Registry) error {

	if err := validate.Required("warnings", "body", m.Warnings); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 volume usage based on context it is used
func (m *V1VolumeUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VolumeUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VolumeUsage) UnmarshalBinary(b []byte) error {
	var res V1VolumeUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
