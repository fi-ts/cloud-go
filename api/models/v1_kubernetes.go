// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Kubernetes v1 kubernetes
//
// swagger:model v1.Kubernetes
type V1Kubernetes struct {

	// expiration date
	// Required: true
	// Format: date-time
	ExpirationDate *strfmt.DateTime `json:"ExpirationDate"`

	// max pods per node
	MaxPodsPerNode int32 `json:"MaxPodsPerNode,omitempty"`

	// pod p i ds limit
	// Required: true
	PodPIDsLimit *int64 `json:"PodPIDsLimit"`

	// version
	// Required: true
	Version *string `json:"Version"`

	// default pod security standard
	// Required: true
	// Enum: ["","baseline","privileged","restricted"]
	DefaultPodSecurityStandard *string `json:"defaultPodSecurityStandard"`
}

// Validate validates this v1 kubernetes
func (m *V1Kubernetes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodPIDsLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultPodSecurityStandard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Kubernetes) validateExpirationDate(formats strfmt.Registry) error {

	if err := validate.Required("ExpirationDate", "body", m.ExpirationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("ExpirationDate", "body", "date-time", m.ExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Kubernetes) validatePodPIDsLimit(formats strfmt.Registry) error {

	if err := validate.Required("PodPIDsLimit", "body", m.PodPIDsLimit); err != nil {
		return err
	}

	return nil
}

func (m *V1Kubernetes) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("Version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

var v1KubernetesTypeDefaultPodSecurityStandardPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","baseline","privileged","restricted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1KubernetesTypeDefaultPodSecurityStandardPropEnum = append(v1KubernetesTypeDefaultPodSecurityStandardPropEnum, v)
	}
}

const (

	// V1KubernetesDefaultPodSecurityStandardEmpty captures enum value ""
	V1KubernetesDefaultPodSecurityStandardEmpty string = ""

	// V1KubernetesDefaultPodSecurityStandardBaseline captures enum value "baseline"
	V1KubernetesDefaultPodSecurityStandardBaseline string = "baseline"

	// V1KubernetesDefaultPodSecurityStandardPrivileged captures enum value "privileged"
	V1KubernetesDefaultPodSecurityStandardPrivileged string = "privileged"

	// V1KubernetesDefaultPodSecurityStandardRestricted captures enum value "restricted"
	V1KubernetesDefaultPodSecurityStandardRestricted string = "restricted"
)

// prop value enum
func (m *V1Kubernetes) validateDefaultPodSecurityStandardEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1KubernetesTypeDefaultPodSecurityStandardPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1Kubernetes) validateDefaultPodSecurityStandard(formats strfmt.Registry) error {

	if err := validate.Required("defaultPodSecurityStandard", "body", m.DefaultPodSecurityStandard); err != nil {
		return err
	}

	// value enum
	if err := m.validateDefaultPodSecurityStandardEnum("defaultPodSecurityStandard", "body", *m.DefaultPodSecurityStandard); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 kubernetes based on context it is used
func (m *V1Kubernetes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Kubernetes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Kubernetes) UnmarshalBinary(b []byte) error {
	var res V1Kubernetes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
