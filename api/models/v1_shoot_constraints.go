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

// V1ShootConstraints v1 shoot constraints
//
// swagger:model v1.ShootConstraints
type V1ShootConstraints struct {

	// the list of available firewall controller version
	// Required: true
	FirewallControllerVersions []*V1FirewallControllerVersion `json:"firewall_controller_versions"`

	// the list of available firewall images
	// Required: true
	FirewallImages []string `json:"firewall_images"`

	// the list of available firewall types
	// Required: true
	FirewallTypes []string `json:"firewall_types"`

	// the list of available kubernetes versions
	// Required: true
	KubernetesVersions []string `json:"kubernetes_versions"`

	// the list of available machine images
	// Required: true
	MachineImages []*V1MachineImage `json:"machine_images"`

	// the list of available machine types
	// Required: true
	MachineTypes []string `json:"machine_types"`

	// network restrictions by partition
	// Required: true
	NetworkAccessRestrictions map[string]V1NetworkAccessRestrictions `json:"network_access_restrictions"`

	// the list of available networks for cluster creation
	// Required: true
	Networks []*V1Network `json:"networks"`

	// the list of available partitions
	// Required: true
	Partitions []string `json:"partitions"`

	// the available seeds by partition id
	// Required: true
	Seeds map[string][]string `json:"seeds"`
}

// Validate validates this v1 shoot constraints
func (m *V1ShootConstraints) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirewallControllerVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewallImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewallTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkAccessRestrictions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeeds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ShootConstraints) validateFirewallControllerVersions(formats strfmt.Registry) error {

	if err := validate.Required("firewall_controller_versions", "body", m.FirewallControllerVersions); err != nil {
		return err
	}

	for i := 0; i < len(m.FirewallControllerVersions); i++ {
		if swag.IsZero(m.FirewallControllerVersions[i]) { // not required
			continue
		}

		if m.FirewallControllerVersions[i] != nil {
			if err := m.FirewallControllerVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("firewall_controller_versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("firewall_controller_versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ShootConstraints) validateFirewallImages(formats strfmt.Registry) error {

	if err := validate.Required("firewall_images", "body", m.FirewallImages); err != nil {
		return err
	}

	return nil
}

func (m *V1ShootConstraints) validateFirewallTypes(formats strfmt.Registry) error {

	if err := validate.Required("firewall_types", "body", m.FirewallTypes); err != nil {
		return err
	}

	return nil
}

func (m *V1ShootConstraints) validateKubernetesVersions(formats strfmt.Registry) error {

	if err := validate.Required("kubernetes_versions", "body", m.KubernetesVersions); err != nil {
		return err
	}

	return nil
}

func (m *V1ShootConstraints) validateMachineImages(formats strfmt.Registry) error {

	if err := validate.Required("machine_images", "body", m.MachineImages); err != nil {
		return err
	}

	for i := 0; i < len(m.MachineImages); i++ {
		if swag.IsZero(m.MachineImages[i]) { // not required
			continue
		}

		if m.MachineImages[i] != nil {
			if err := m.MachineImages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machine_images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machine_images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ShootConstraints) validateMachineTypes(formats strfmt.Registry) error {

	if err := validate.Required("machine_types", "body", m.MachineTypes); err != nil {
		return err
	}

	return nil
}

func (m *V1ShootConstraints) validateNetworkAccessRestrictions(formats strfmt.Registry) error {

	if err := validate.Required("network_access_restrictions", "body", m.NetworkAccessRestrictions); err != nil {
		return err
	}

	for k := range m.NetworkAccessRestrictions {

		if err := validate.Required("network_access_restrictions"+"."+k, "body", m.NetworkAccessRestrictions[k]); err != nil {
			return err
		}
		if val, ok := m.NetworkAccessRestrictions[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("network_access_restrictions" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("network_access_restrictions" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ShootConstraints) validateNetworks(formats strfmt.Registry) error {

	if err := validate.Required("networks", "body", m.Networks); err != nil {
		return err
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ShootConstraints) validatePartitions(formats strfmt.Registry) error {

	if err := validate.Required("partitions", "body", m.Partitions); err != nil {
		return err
	}

	return nil
}

func (m *V1ShootConstraints) validateSeeds(formats strfmt.Registry) error {

	if err := validate.Required("seeds", "body", m.Seeds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 shoot constraints based on the context it is used
func (m *V1ShootConstraints) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFirewallControllerVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachineImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkAccessRestrictions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ShootConstraints) contextValidateFirewallControllerVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FirewallControllerVersions); i++ {

		if m.FirewallControllerVersions[i] != nil {

			if swag.IsZero(m.FirewallControllerVersions[i]) { // not required
				return nil
			}

			if err := m.FirewallControllerVersions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("firewall_controller_versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("firewall_controller_versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ShootConstraints) contextValidateMachineImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MachineImages); i++ {

		if m.MachineImages[i] != nil {

			if swag.IsZero(m.MachineImages[i]) { // not required
				return nil
			}

			if err := m.MachineImages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machine_images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machine_images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ShootConstraints) contextValidateNetworkAccessRestrictions(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("network_access_restrictions", "body", m.NetworkAccessRestrictions); err != nil {
		return err
	}

	for k := range m.NetworkAccessRestrictions {

		if val, ok := m.NetworkAccessRestrictions[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V1ShootConstraints) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Networks); i++ {

		if m.Networks[i] != nil {

			if swag.IsZero(m.Networks[i]) { // not required
				return nil
			}

			if err := m.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ShootConstraints) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ShootConstraints) UnmarshalBinary(b []byte) error {
	var res V1ShootConstraints
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
