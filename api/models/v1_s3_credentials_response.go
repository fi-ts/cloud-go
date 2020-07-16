// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1S3CredentialsResponse v1 s3 credentials response
// swagger:model v1.S3CredentialsResponse
type V1S3CredentialsResponse struct {

	// endpoint
	// Required: true
	Endpoint *string `json:"endpoint"`

	// id
	// Required: true
	ID *string `json:"id"`

	// keys
	// Required: true
	Keys []*V1S3Key `json:"keys"`

	// max buckets
	// Required: true
	MaxBuckets *int64 `json:"max_buckets"`

	// name
	// Required: true
	Name *string `json:"name"`

	// partition
	// Required: true
	Partition *string `json:"partition"`

	// project
	// Required: true
	Project *string `json:"project"`

	// tenant
	// Required: true
	Tenant *string `json:"tenant"`
}

// Validate validates this v1 s3 credentials response
func (m *V1S3CredentialsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBuckets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1S3CredentialsResponse) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	return nil
}

func (m *V1S3CredentialsResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1S3CredentialsResponse) validateKeys(formats strfmt.Registry) error {

	if err := validate.Required("keys", "body", m.Keys); err != nil {
		return err
	}

	for i := 0; i < len(m.Keys); i++ {
		if swag.IsZero(m.Keys[i]) { // not required
			continue
		}

		if m.Keys[i] != nil {
			if err := m.Keys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1S3CredentialsResponse) validateMaxBuckets(formats strfmt.Registry) error {

	if err := validate.Required("max_buckets", "body", m.MaxBuckets); err != nil {
		return err
	}

	return nil
}

func (m *V1S3CredentialsResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1S3CredentialsResponse) validatePartition(formats strfmt.Registry) error {

	if err := validate.Required("partition", "body", m.Partition); err != nil {
		return err
	}

	return nil
}

func (m *V1S3CredentialsResponse) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	return nil
}

func (m *V1S3CredentialsResponse) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1S3CredentialsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1S3CredentialsResponse) UnmarshalBinary(b []byte) error {
	var res V1S3CredentialsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}