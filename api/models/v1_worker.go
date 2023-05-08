// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Worker v1 worker
//
// swagger:model v1.Worker
type V1Worker struct {

	// annotations
	// Required: true
	Annotations map[string]string `json:"Annotations"`

	// c r i
	// Required: true
	CRI *string `json:"CRI"`

	// drain timeout
	DrainTimeout int64 `json:"DrainTimeout,omitempty"`

	// health timeout
	HealthTimeout int64 `json:"HealthTimeout,omitempty"`

	// kubernetes version
	// Required: true
	KubernetesVersion *string `json:"KubernetesVersion"`

	// labels
	// Required: true
	Labels map[string]string `json:"Labels"`

	// machine image
	// Required: true
	MachineImage *V1MachineImage `json:"MachineImage"`

	// machine type
	// Required: true
	MachineType *string `json:"MachineType"`

	// max surge
	// Required: true
	MaxSurge *string `json:"MaxSurge"`

	// max unavailable
	// Required: true
	MaxUnavailable *string `json:"MaxUnavailable"`

	// maximum
	// Required: true
	Maximum *int32 `json:"Maximum"`

	// minimum
	// Required: true
	Minimum *int32 `json:"Minimum"`

	// name
	// Required: true
	Name *string `json:"Name"`

	// taints
	// Required: true
	Taints []*V1Taint `json:"Taints"`
}

// Validate validates this v1 worker
func (m *V1Worker) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCRI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxSurge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxUnavailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaximum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinimum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Worker) validateAnnotations(formats strfmt.Registry) error {

	if err := validate.Required("Annotations", "body", m.Annotations); err != nil {
		return err
	}

	return nil
}

func (m *V1Worker) validateCRI(formats strfmt.Registry) error {

	if err := validate.Required("CRI", "body", m.CRI); err != nil {
		return err
	}

	return nil
}

func (m *V1Worker) validateKubernetesVersion(formats strfmt.Registry) error {

	if err := validate.Required("KubernetesVersion", "body", m.KubernetesVersion); err != nil {
		return err
	}

	return nil
}

func (m *V1Worker) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("Labels", "body", m.Labels); err != nil {
		return err
	}

	return nil
}

func (m *V1Worker) validateMachineImage(formats strfmt.Registry) error {

	if err := validate.Required("MachineImage", "body", m.MachineImage); err != nil {
		return err
	}

	if m.MachineImage != nil {
		if err := m.MachineImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MachineImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MachineImage")
			}
			return err
		}
	}

	return nil
}

func (m *V1Worker) validateMachineType(formats strfmt.Registry) error {

	if err := validate.Required("MachineType", "body", m.MachineType); err != nil {
		return err
	}

	return nil
}

func (m *V1Worker) validateMaxSurge(formats strfmt.Registry) error {

	if err := validate.Required("MaxSurge", "body", m.MaxSurge); err != nil {
		return err
	}

	return nil
}

func (m *V1Worker) validateMaxUnavailable(formats strfmt.Registry) error {

	if err := validate.Required("MaxUnavailable", "body", m.MaxUnavailable); err != nil {
		return err
	}

	return nil
}

func (m *V1Worker) validateMaximum(formats strfmt.Registry) error {

	if err := validate.Required("Maximum", "body", m.Maximum); err != nil {
		return err
	}

	return nil
}

func (m *V1Worker) validateMinimum(formats strfmt.Registry) error {

	if err := validate.Required("Minimum", "body", m.Minimum); err != nil {
		return err
	}

	return nil
}

func (m *V1Worker) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1Worker) validateTaints(formats strfmt.Registry) error {

	if err := validate.Required("Taints", "body", m.Taints); err != nil {
		return err
	}

	for i := 0; i < len(m.Taints); i++ {
		if swag.IsZero(m.Taints[i]) { // not required
			continue
		}

		if m.Taints[i] != nil {
			if err := m.Taints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Taints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Taints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 worker based on the context it is used
func (m *V1Worker) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMachineImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Worker) contextValidateMachineImage(ctx context.Context, formats strfmt.Registry) error {

	if m.MachineImage != nil {
		if err := m.MachineImage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MachineImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MachineImage")
			}
			return err
		}
	}

	return nil
}

func (m *V1Worker) contextValidateTaints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Taints); i++ {

		if m.Taints[i] != nil {
			if err := m.Taints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Taints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Taints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Worker) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Worker) UnmarshalBinary(b []byte) error {
	var res V1Worker
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
