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

// V1ClusterCreateRequest v1 cluster create request
//
// swagger:model v1.ClusterCreateRequest
type V1ClusterCreateRequest struct {

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

	// description
	// Required: true
	Description *string `json:"Description"`

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

	// kubernetes
	// Required: true
	Kubernetes *V1Kubernetes `json:"Kubernetes"`

	// labels
	// Required: true
	Labels map[string]string `json:"Labels"`

	// maintenance
	// Required: true
	Maintenance *V1Maintenance `json:"Maintenance"`

	// name
	// Required: true
	Name *string `json:"Name"`

	// partition ID
	// Required: true
	PartitionID *string `json:"PartitionID"`

	// project ID
	// Required: true
	ProjectID *string `json:"ProjectID"`

	// purpose
	// Required: true
	Purpose *string `json:"Purpose"`

	// system components
	// Required: true
	SystemComponents *V1SystemComponents `json:"SystemComponents"`

	// tenant
	// Required: true
	Tenant *string `json:"Tenant"`

	// workers
	// Required: true
	Workers []*V1Worker `json:"Workers"`

	// cni
	Cni string `json:"cni,omitempty"`

	// seed name on which the cluster will be scheduled
	SeedName string `json:"seedName,omitempty"`
}

// Validate validates this v1 cluster create request
func (m *V1ClusterCreateRequest) Validate(formats strfmt.Registry) error {
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

	if err := m.validateDescription(formats); err != nil {
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

	if err := m.validateKubernetes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurpose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterCreateRequest) validateAdditionalNetworks(formats strfmt.Registry) error {

	if err := validate.Required("AdditionalNetworks", "body", m.AdditionalNetworks); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterCreateRequest) validateAudit(formats strfmt.Registry) error {

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

func (m *V1ClusterCreateRequest) validateClusterFeatures(formats strfmt.Registry) error {

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

func (m *V1ClusterCreateRequest) validateCustomDefaultStorageClass(formats strfmt.Registry) error {

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

func (m *V1ClusterCreateRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterCreateRequest) validateEgressRules(formats strfmt.Registry) error {

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

func (m *V1ClusterCreateRequest) validateFirewallControllerVersion(formats strfmt.Registry) error {

	if err := validate.Required("FirewallControllerVersion", "body", m.FirewallControllerVersion); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterCreateRequest) validateFirewallImage(formats strfmt.Registry) error {

	if err := validate.Required("FirewallImage", "body", m.FirewallImage); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterCreateRequest) validateFirewallSize(formats strfmt.Registry) error {

	if err := validate.Required("FirewallSize", "body", m.FirewallSize); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterCreateRequest) validateKubernetes(formats strfmt.Registry) error {

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

func (m *V1ClusterCreateRequest) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("Labels", "body", m.Labels); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterCreateRequest) validateMaintenance(formats strfmt.Registry) error {

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

func (m *V1ClusterCreateRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterCreateRequest) validatePartitionID(formats strfmt.Registry) error {

	if err := validate.Required("PartitionID", "body", m.PartitionID); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterCreateRequest) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("ProjectID", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterCreateRequest) validatePurpose(formats strfmt.Registry) error {

	if err := validate.Required("Purpose", "body", m.Purpose); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterCreateRequest) validateSystemComponents(formats strfmt.Registry) error {

	if err := validate.Required("SystemComponents", "body", m.SystemComponents); err != nil {
		return err
	}

	if m.SystemComponents != nil {
		if err := m.SystemComponents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SystemComponents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SystemComponents")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterCreateRequest) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("Tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterCreateRequest) validateWorkers(formats strfmt.Registry) error {

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

// ContextValidate validate this v1 cluster create request based on the context it is used
func (m *V1ClusterCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateSystemComponents(ctx, formats); err != nil {
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

func (m *V1ClusterCreateRequest) contextValidateAudit(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1ClusterCreateRequest) contextValidateClusterFeatures(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1ClusterCreateRequest) contextValidateCustomDefaultStorageClass(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1ClusterCreateRequest) contextValidateEgressRules(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1ClusterCreateRequest) contextValidateKubernetes(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1ClusterCreateRequest) contextValidateMaintenance(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1ClusterCreateRequest) contextValidateSystemComponents(ctx context.Context, formats strfmt.Registry) error {

	if m.SystemComponents != nil {

		if err := m.SystemComponents.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SystemComponents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SystemComponents")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterCreateRequest) contextValidateWorkers(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1ClusterCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterCreateRequest) UnmarshalBinary(b []byte) error {
	var res V1ClusterCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
