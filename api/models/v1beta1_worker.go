// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1beta1Worker v1beta1 worker
// swagger:model v1beta1.Worker
type V1beta1Worker struct {

	// annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// ca bundle
	CaBundle string `json:"caBundle,omitempty"`

	// cri
	Cri *V1beta1CRI `json:"cri,omitempty"`

	// data volumes
	DataVolumes []*V1beta1Volume `json:"dataVolumes"`

	// kubelet data volume name
	KubeletDataVolumeName string `json:"kubeletDataVolumeName,omitempty"`

	// kubernetes
	Kubernetes *V1beta1WorkerKubernetes `json:"kubernetes,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// machine
	// Required: true
	Machine *V1beta1Machine `json:"machine"`

	// max surge
	MaxSurge string `json:"maxSurge,omitempty"`

	// max unavailable
	MaxUnavailable string `json:"maxUnavailable,omitempty"`

	// maximum
	// Required: true
	Maximum *int32 `json:"maximum"`

	// minimum
	// Required: true
	Minimum *int32 `json:"minimum"`

	// name
	// Required: true
	Name *string `json:"name"`

	// provider config
	ProviderConfig string `json:"providerConfig,omitempty"`

	// taints
	Taints []*V1Taint `json:"taints"`

	// volume
	Volume *V1beta1Volume `json:"volume,omitempty"`

	// zones
	Zones []string `json:"zones"`
}

// Validate validates this v1beta1 worker
func (m *V1beta1Worker) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCri(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachine(formats); err != nil {
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

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1Worker) validateCri(formats strfmt.Registry) error {

	if swag.IsZero(m.Cri) { // not required
		return nil
	}

	if m.Cri != nil {
		if err := m.Cri.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cri")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Worker) validateDataVolumes(formats strfmt.Registry) error {

	if swag.IsZero(m.DataVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.DataVolumes); i++ {
		if swag.IsZero(m.DataVolumes[i]) { // not required
			continue
		}

		if m.DataVolumes[i] != nil {
			if err := m.DataVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataVolumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1beta1Worker) validateKubernetes(formats strfmt.Registry) error {

	if swag.IsZero(m.Kubernetes) { // not required
		return nil
	}

	if m.Kubernetes != nil {
		if err := m.Kubernetes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Worker) validateMachine(formats strfmt.Registry) error {

	if err := validate.Required("machine", "body", m.Machine); err != nil {
		return err
	}

	if m.Machine != nil {
		if err := m.Machine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machine")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Worker) validateMaximum(formats strfmt.Registry) error {

	if err := validate.Required("maximum", "body", m.Maximum); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Worker) validateMinimum(formats strfmt.Registry) error {

	if err := validate.Required("minimum", "body", m.Minimum); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Worker) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Worker) validateTaints(formats strfmt.Registry) error {

	if swag.IsZero(m.Taints) { // not required
		return nil
	}

	for i := 0; i < len(m.Taints); i++ {
		if swag.IsZero(m.Taints[i]) { // not required
			continue
		}

		if m.Taints[i] != nil {
			if err := m.Taints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1beta1Worker) validateVolume(formats strfmt.Registry) error {

	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1Worker) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1Worker) UnmarshalBinary(b []byte) error {
	var res V1beta1Worker
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}