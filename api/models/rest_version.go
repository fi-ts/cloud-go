// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestVersion rest version
//
// swagger:model rest.version
type RestVersion struct {

	// builddate
	// Required: true
	Builddate *string `json:"builddate"`

	// gitsha1
	// Required: true
	Gitsha1 *string `json:"gitsha1"`

	// min client version
	// Required: true
	MinClientVersion *string `json:"min_client_version"`

	// name
	// Required: true
	Name *string `json:"name"`

	// release version
	ReleaseVersion string `json:"release_version,omitempty"`

	// revision
	// Required: true
	Revision *string `json:"revision"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this rest version
func (m *RestVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuilddate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitsha1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinClientVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestVersion) validateBuilddate(formats strfmt.Registry) error {

	if err := validate.Required("builddate", "body", m.Builddate); err != nil {
		return err
	}

	return nil
}

func (m *RestVersion) validateGitsha1(formats strfmt.Registry) error {

	if err := validate.Required("gitsha1", "body", m.Gitsha1); err != nil {
		return err
	}

	return nil
}

func (m *RestVersion) validateMinClientVersion(formats strfmt.Registry) error {

	if err := validate.Required("min_client_version", "body", m.MinClientVersion); err != nil {
		return err
	}

	return nil
}

func (m *RestVersion) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RestVersion) validateRevision(formats strfmt.Registry) error {

	if err := validate.Required("revision", "body", m.Revision); err != nil {
		return err
	}

	return nil
}

func (m *RestVersion) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rest version based on context it is used
func (m *RestVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestVersion) UnmarshalBinary(b []byte) error {
	var res RestVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
