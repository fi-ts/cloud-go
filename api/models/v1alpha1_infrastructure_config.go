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

// V1alpha1InfrastructureConfig TypeMeta describes an individual object in an API response or request with strings representing the type of the object and its API schema version. Structures that are versioned or persisted should inline TypeMeta.
// swagger:model v1alpha1.InfrastructureConfig
type V1alpha1InfrastructureConfig struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// firewall
	// Required: true
	Firewall *V1alpha1Firewall `json:"firewall"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// partition ID
	// Required: true
	PartitionID *string `json:"partitionID"`

	// project ID
	// Required: true
	ProjectID *string `json:"projectID"`
}

// Validate validates this v1alpha1 infrastructure config
func (m *V1alpha1InfrastructureConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirewall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1InfrastructureConfig) validateFirewall(formats strfmt.Registry) error {

	if err := validate.Required("firewall", "body", m.Firewall); err != nil {
		return err
	}

	if m.Firewall != nil {
		if err := m.Firewall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firewall")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1InfrastructureConfig) validatePartitionID(formats strfmt.Registry) error {

	if err := validate.Required("partitionID", "body", m.PartitionID); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1InfrastructureConfig) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectID", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1InfrastructureConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1InfrastructureConfig) UnmarshalBinary(b []byte) error {
	var res V1alpha1InfrastructureConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
