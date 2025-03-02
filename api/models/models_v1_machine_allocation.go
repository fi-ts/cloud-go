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

// ModelsV1MachineAllocation models v1 machine allocation
//
// swagger:model models.V1MachineAllocation
type ModelsV1MachineAllocation struct {

	// allocationuuid
	// Required: true
	Allocationuuid *string `json:"allocationuuid"`

	// boot info
	BootInfo *ModelsV1BootInfo `json:"boot_info,omitempty"`

	// created
	// Required: true
	Created *string `json:"created"`

	// creator
	// Required: true
	Creator *string `json:"creator"`

	// description
	Description string `json:"description,omitempty"`

	// dns servers
	// Required: true
	DNSServers []*ModelsV1DNSServer `json:"dns_servers"`

	// filesystemlayout
	Filesystemlayout *ModelsV1FilesystemLayoutResponse `json:"filesystemlayout,omitempty"`

	// firewall rules
	FirewallRules *ModelsV1FirewallRules `json:"firewall_rules,omitempty"`

	// hostname
	// Required: true
	Hostname *string `json:"hostname"`

	// image
	Image *ModelsV1ImageResponse `json:"image,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// networks
	// Required: true
	Networks []*ModelsV1MachineNetwork `json:"networks"`

	// ntp servers
	// Required: true
	NtpServers []*ModelsV1NTPServer `json:"ntp_servers"`

	// project
	// Required: true
	Project *string `json:"project"`

	// reinstall
	// Required: true
	Reinstall *bool `json:"reinstall"`

	// role
	// Required: true
	Role *string `json:"role"`

	// ssh pub keys
	// Required: true
	SSHPubKeys []string `json:"ssh_pub_keys"`

	// succeeded
	// Required: true
	Succeeded *bool `json:"succeeded"`

	// user data
	UserData string `json:"user_data,omitempty"`

	// vpn
	Vpn *ModelsV1MachineVPN `json:"vpn,omitempty"`
}

// Validate validates this models v1 machine allocation
func (m *ModelsV1MachineAllocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocationuuid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilesystemlayout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewallRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNtpServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReinstall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHPubKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSucceeded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1MachineAllocation) validateAllocationuuid(formats strfmt.Registry) error {

	if err := validate.Required("allocationuuid", "body", m.Allocationuuid); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateBootInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.BootInfo) { // not required
		return nil
	}

	if m.BootInfo != nil {
		if err := m.BootInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("boot_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("boot_info")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateCreator(formats strfmt.Registry) error {

	if err := validate.Required("creator", "body", m.Creator); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateDNSServers(formats strfmt.Registry) error {

	if err := validate.Required("dns_servers", "body", m.DNSServers); err != nil {
		return err
	}

	for i := 0; i < len(m.DNSServers); i++ {
		if swag.IsZero(m.DNSServers[i]) { // not required
			continue
		}

		if m.DNSServers[i] != nil {
			if err := m.DNSServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dns_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateFilesystemlayout(formats strfmt.Registry) error {
	if swag.IsZero(m.Filesystemlayout) { // not required
		return nil
	}

	if m.Filesystemlayout != nil {
		if err := m.Filesystemlayout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filesystemlayout")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filesystemlayout")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateFirewallRules(formats strfmt.Registry) error {
	if swag.IsZero(m.FirewallRules) { // not required
		return nil
	}

	if m.FirewallRules != nil {
		if err := m.FirewallRules.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firewall_rules")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firewall_rules")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateImage(formats strfmt.Registry) error {
	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateNetworks(formats strfmt.Registry) error {

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

func (m *ModelsV1MachineAllocation) validateNtpServers(formats strfmt.Registry) error {

	if err := validate.Required("ntp_servers", "body", m.NtpServers); err != nil {
		return err
	}

	for i := 0; i < len(m.NtpServers); i++ {
		if swag.IsZero(m.NtpServers[i]) { // not required
			continue
		}

		if m.NtpServers[i] != nil {
			if err := m.NtpServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ntp_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ntp_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateReinstall(formats strfmt.Registry) error {

	if err := validate.Required("reinstall", "body", m.Reinstall); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateSSHPubKeys(formats strfmt.Registry) error {

	if err := validate.Required("ssh_pub_keys", "body", m.SSHPubKeys); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateSucceeded(formats strfmt.Registry) error {

	if err := validate.Required("succeeded", "body", m.Succeeded); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineAllocation) validateVpn(formats strfmt.Registry) error {
	if swag.IsZero(m.Vpn) { // not required
		return nil
	}

	if m.Vpn != nil {
		if err := m.Vpn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpn")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpn")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this models v1 machine allocation based on the context it is used
func (m *ModelsV1MachineAllocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBootInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNSServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilesystemlayout(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirewallRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNtpServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1MachineAllocation) contextValidateBootInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.BootInfo != nil {

		if swag.IsZero(m.BootInfo) { // not required
			return nil
		}

		if err := m.BootInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("boot_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("boot_info")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsV1MachineAllocation) contextValidateDNSServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNSServers); i++ {

		if m.DNSServers[i] != nil {

			if swag.IsZero(m.DNSServers[i]) { // not required
				return nil
			}

			if err := m.DNSServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dns_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsV1MachineAllocation) contextValidateFilesystemlayout(ctx context.Context, formats strfmt.Registry) error {

	if m.Filesystemlayout != nil {

		if swag.IsZero(m.Filesystemlayout) { // not required
			return nil
		}

		if err := m.Filesystemlayout.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filesystemlayout")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filesystemlayout")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsV1MachineAllocation) contextValidateFirewallRules(ctx context.Context, formats strfmt.Registry) error {

	if m.FirewallRules != nil {

		if swag.IsZero(m.FirewallRules) { // not required
			return nil
		}

		if err := m.FirewallRules.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firewall_rules")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firewall_rules")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsV1MachineAllocation) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

	if m.Image != nil {

		if swag.IsZero(m.Image) { // not required
			return nil
		}

		if err := m.Image.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsV1MachineAllocation) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ModelsV1MachineAllocation) contextValidateNtpServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NtpServers); i++ {

		if m.NtpServers[i] != nil {

			if swag.IsZero(m.NtpServers[i]) { // not required
				return nil
			}

			if err := m.NtpServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ntp_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ntp_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsV1MachineAllocation) contextValidateVpn(ctx context.Context, formats strfmt.Registry) error {

	if m.Vpn != nil {

		if swag.IsZero(m.Vpn) { // not required
			return nil
		}

		if err := m.Vpn.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpn")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpn")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsV1MachineAllocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsV1MachineAllocation) UnmarshalBinary(b []byte) error {
	var res ModelsV1MachineAllocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
