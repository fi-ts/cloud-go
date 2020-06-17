// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1beta1Maintenance v1beta1 maintenance
// swagger:model v1beta1.Maintenance
type V1beta1Maintenance struct {

	// auto update
	AutoUpdate *V1beta1MaintenanceAutoUpdate `json:"autoUpdate,omitempty"`

	// confine spec update rollout
	ConfineSpecUpdateRollout bool `json:"confineSpecUpdateRollout,omitempty"`

	// time window
	TimeWindow *V1beta1MaintenanceTimeWindow `json:"timeWindow,omitempty"`
}

// Validate validates this v1beta1 maintenance
func (m *V1beta1Maintenance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeWindow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1Maintenance) validateAutoUpdate(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoUpdate) { // not required
		return nil
	}

	if m.AutoUpdate != nil {
		if err := m.AutoUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoUpdate")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Maintenance) validateTimeWindow(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeWindow) { // not required
		return nil
	}

	if m.TimeWindow != nil {
		if err := m.TimeWindow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeWindow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1Maintenance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1Maintenance) UnmarshalBinary(b []byte) error {
	var res V1beta1Maintenance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
