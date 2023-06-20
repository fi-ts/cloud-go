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

// V1ClusterUpdateRequest v1 cluster update request
//
// swagger:model v1.ClusterUpdateRequest
type V1ClusterUpdateRequest struct {

	// additional networks
	// Required: true
	AdditionalNetworks []string `json:"AdditionalNetworks"`

	// audit
	// Required: true
	Audit *V1Audit `json:"Audit"`

	// cluster features
	// Required: true
	ClusterFeatures *V1ClusterFeatures `json:"ClusterFeatures"`

	// custom default storage class
	// Required: true
	CustomDefaultStorageClass *V1CustomDefaultStorageClass `json:"CustomDefaultStorageClass"`

	// egress rules
	// Required: true
	EgressRules []*V1EgressRule `json:"EgressRules"`

	// firewall controller version
	// Required: true
	FirewallControllerVersion *string `json:"FirewallControllerVersion"`

	// firewall image
	// Required: true
	FirewallImage *string `json:"FirewallImage"`

	// firewall size
	// Required: true
	FirewallSize *string `json:"FirewallSize"`

	// ID
	// Required: true
	ID *string `json:"ID"`

	// kubernetes
	// Required: true
	Kubernetes *V1Kubernetes `json:"Kubernetes"`

	// labels
	// Required: true
	Labels map[string]string `json:"Labels"`

	// maintenance
	// Required: true
	Maintenance *V1Maintenance `json:"Maintenance"`

	// purpose
	// Required: true
	Purpose *string `json:"Purpose"`

	// workers
	// Required: true
	Workers []*V1Worker `json:"Workers"`

	// seed name on which the cluster will be scheduled
	// Required: true
	SeedName *string `json:"seedName"`
}

// Validate validates this v1 cluster update request
func (m *V1ClusterUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAudit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomDefaultStorageClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEgressRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewallControllerVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewallImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewallSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurpose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeedName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterUpdateRequest) validateAdditionalNetworks(formats strfmt.Registry) error {

	if err := validate.Required("AdditionalNetworks", "body", m.AdditionalNetworks); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateAudit(formats strfmt.Registry) error {

	if err := validate.Required("Audit", "body", m.Audit); err != nil {
		return err
	}

	if m.Audit != nil {
		if err := m.Audit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Audit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Audit")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateClusterFeatures(formats strfmt.Registry) error {

	if err := validate.Required("ClusterFeatures", "body", m.ClusterFeatures); err != nil {
		return err
	}

	if m.ClusterFeatures != nil {
		if err := m.ClusterFeatures.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ClusterFeatures")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ClusterFeatures")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateCustomDefaultStorageClass(formats strfmt.Registry) error {

	if err := validate.Required("CustomDefaultStorageClass", "body", m.CustomDefaultStorageClass); err != nil {
		return err
	}

	if m.CustomDefaultStorageClass != nil {
		if err := m.CustomDefaultStorageClass.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CustomDefaultStorageClass")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CustomDefaultStorageClass")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateEgressRules(formats strfmt.Registry) error {

	if err := validate.Required("EgressRules", "body", m.EgressRules); err != nil {
		return err
	}

	for i := 0; i < len(m.EgressRules); i++ {
		if swag.IsZero(m.EgressRules[i]) { // not required
			continue
		}

		if m.EgressRules[i] != nil {
			if err := m.EgressRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EgressRules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("EgressRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateFirewallControllerVersion(formats strfmt.Registry) error {

	if err := validate.Required("FirewallControllerVersion", "body", m.FirewallControllerVersion); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateFirewallImage(formats strfmt.Registry) error {

	if err := validate.Required("FirewallImage", "body", m.FirewallImage); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateFirewallSize(formats strfmt.Registry) error {

	if err := validate.Required("FirewallSize", "body", m.FirewallSize); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateKubernetes(formats strfmt.Registry) error {

	if err := validate.Required("Kubernetes", "body", m.Kubernetes); err != nil {
		return err
	}

	if m.Kubernetes != nil {
		if err := m.Kubernetes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Kubernetes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("Labels", "body", m.Labels); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateMaintenance(formats strfmt.Registry) error {

	if err := validate.Required("Maintenance", "body", m.Maintenance); err != nil {
		return err
	}

	if m.Maintenance != nil {
		if err := m.Maintenance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Maintenance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Maintenance")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterUpdateRequest) validatePurpose(formats strfmt.Registry) error {

	if err := validate.Required("Purpose", "body", m.Purpose); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateWorkers(formats strfmt.Registry) error {

	if err := validate.Required("Workers", "body", m.Workers); err != nil {
		return err
	}

	for i := 0; i < len(m.Workers); i++ {
		if swag.IsZero(m.Workers[i]) { // not required
			continue
		}

		if m.Workers[i] != nil {
			if err := m.Workers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Workers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Workers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterUpdateRequest) validateSeedName(formats strfmt.Registry) error {

	if err := validate.Required("seedName", "body", m.SeedName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 cluster update request based on the context it is used
func (m *V1ClusterUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAudit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomDefaultStorageClass(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEgressRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubernetes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaintenance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterUpdateRequest) contextValidateAudit(ctx context.Context, formats strfmt.Registry) error {

	if m.Audit != nil {

		if err := m.Audit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Audit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Audit")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterUpdateRequest) contextValidateClusterFeatures(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterFeatures != nil {

		if err := m.ClusterFeatures.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ClusterFeatures")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ClusterFeatures")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterUpdateRequest) contextValidateCustomDefaultStorageClass(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomDefaultStorageClass != nil {

		if err := m.CustomDefaultStorageClass.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CustomDefaultStorageClass")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CustomDefaultStorageClass")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterUpdateRequest) contextValidateEgressRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EgressRules); i++ {

		if m.EgressRules[i] != nil {

			if swag.IsZero(m.EgressRules[i]) { // not required
				return nil
			}

			if err := m.EgressRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EgressRules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("EgressRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterUpdateRequest) contextValidateKubernetes(ctx context.Context, formats strfmt.Registry) error {

	if m.Kubernetes != nil {

		if err := m.Kubernetes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Kubernetes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterUpdateRequest) contextValidateMaintenance(ctx context.Context, formats strfmt.Registry) error {

	if m.Maintenance != nil {

		if err := m.Maintenance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Maintenance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Maintenance")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterUpdateRequest) contextValidateWorkers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Workers); i++ {

		if m.Workers[i] != nil {

			if swag.IsZero(m.Workers[i]) { // not required
				return nil
			}

			if err := m.Workers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Workers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Workers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterUpdateRequest) UnmarshalBinary(b []byte) error {
	var res V1ClusterUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
