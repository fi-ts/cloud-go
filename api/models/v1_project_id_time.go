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

// V1ProjectIDTime v1 project ID time
// swagger:model v1.ProjectIDTime
type V1ProjectIDTime struct {

	// projectID as returned by cloud-api (e.g. 10241dd7-a8de-4856-8ac0-b55830b22036)
	ProjectID string `json:"project_id,omitempty"`

	// point in time
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`
}

// Validate validates this v1 project ID time
func (m *V1ProjectIDTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectIDTime) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectIDTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectIDTime) UnmarshalBinary(b []byte) error {
	var res V1ProjectIDTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}