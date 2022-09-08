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

// ModelsV1MachineRecentProvisioningEvents models v1 machine recent provisioning events
//
// swagger:model models.V1MachineRecentProvisioningEvents
type ModelsV1MachineRecentProvisioningEvents struct {

	// crash loop
	// Required: true
	CrashLoop *bool `json:"crash_loop"`

	// failed machine reclaim
	// Required: true
	FailedMachineReclaim *bool `json:"failed_machine_reclaim"`

	// incomplete provisioning cycles
	// Required: true
	IncompleteProvisioningCycles *string `json:"incomplete_provisioning_cycles"`

	// last error event
	LastErrorEvent *ModelsV1MachineProvisioningEvent `json:"last_error_event,omitempty"`

	// last event time
	LastEventTime string `json:"last_event_time,omitempty"`

	// log
	// Required: true
	Log []*ModelsV1MachineProvisioningEvent `json:"log"`
}

// Validate validates this models v1 machine recent provisioning events
func (m *ModelsV1MachineRecentProvisioningEvents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrashLoop(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailedMachineReclaim(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncompleteProvisioningCycles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastErrorEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1MachineRecentProvisioningEvents) validateCrashLoop(formats strfmt.Registry) error {

	if err := validate.Required("crash_loop", "body", m.CrashLoop); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineRecentProvisioningEvents) validateFailedMachineReclaim(formats strfmt.Registry) error {

	if err := validate.Required("failed_machine_reclaim", "body", m.FailedMachineReclaim); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineRecentProvisioningEvents) validateIncompleteProvisioningCycles(formats strfmt.Registry) error {

	if err := validate.Required("incomplete_provisioning_cycles", "body", m.IncompleteProvisioningCycles); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineRecentProvisioningEvents) validateLastErrorEvent(formats strfmt.Registry) error {
	if swag.IsZero(m.LastErrorEvent) { // not required
		return nil
	}

	if m.LastErrorEvent != nil {
		if err := m.LastErrorEvent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_error_event")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_error_event")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsV1MachineRecentProvisioningEvents) validateLog(formats strfmt.Registry) error {

	if err := validate.Required("log", "body", m.Log); err != nil {
		return err
	}

	for i := 0; i < len(m.Log); i++ {
		if swag.IsZero(m.Log[i]) { // not required
			continue
		}

		if m.Log[i] != nil {
			if err := m.Log[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("log" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("log" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this models v1 machine recent provisioning events based on the context it is used
func (m *ModelsV1MachineRecentProvisioningEvents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastErrorEvent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1MachineRecentProvisioningEvents) contextValidateLastErrorEvent(ctx context.Context, formats strfmt.Registry) error {

	if m.LastErrorEvent != nil {
		if err := m.LastErrorEvent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_error_event")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_error_event")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsV1MachineRecentProvisioningEvents) contextValidateLog(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Log); i++ {

		if m.Log[i] != nil {
			if err := m.Log[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("log" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("log" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsV1MachineRecentProvisioningEvents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsV1MachineRecentProvisioningEvents) UnmarshalBinary(b []byte) error {
	var res ModelsV1MachineRecentProvisioningEvents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
