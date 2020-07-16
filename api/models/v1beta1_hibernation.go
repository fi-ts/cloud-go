// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1beta1Hibernation v1beta1 hibernation
// swagger:model v1beta1.Hibernation
type V1beta1Hibernation struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// schedules
	Schedules []*V1beta1HibernationSchedule `json:"schedules"`
}

// Validate validates this v1beta1 hibernation
func (m *V1beta1Hibernation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1Hibernation) validateSchedules(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedules) { // not required
		return nil
	}

	for i := 0; i < len(m.Schedules); i++ {
		if swag.IsZero(m.Schedules[i]) { // not required
			continue
		}

		if m.Schedules[i] != nil {
			if err := m.Schedules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1Hibernation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1Hibernation) UnmarshalBinary(b []byte) error {
	var res V1beta1Hibernation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}