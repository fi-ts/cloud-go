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

// V1PostgresBackupConfigResponse v1 postgres backup config response
//
// swagger:model v1.PostgresBackupConfigResponse
type V1PostgresBackupConfigResponse struct {

	// created by
	// Required: true
	CreatedBy *string `json:"createdBy"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	Name string `json:"name,omitempty"`

	// project ID
	ProjectID string `json:"projectID,omitempty"`

	// retention
	Retention int32 `json:"retention,omitempty"`

	// s3 bucket name
	S3BucketName string `json:"s3BucketName,omitempty"`

	// s3 endpoint
	S3Endpoint string `json:"s3Endpoint,omitempty"`

	// s3 region
	S3Region string `json:"s3Region,omitempty"`

	// schedule
	Schedule string `json:"schedule,omitempty"`

	// secret
	Secret *V1PostgresBackupSecret `json:"secret,omitempty"`

	// tenant
	Tenant string `json:"tenant,omitempty"`
}

// Validate validates this v1 postgres backup config response
func (m *V1PostgresBackupConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PostgresBackupConfigResponse) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("createdBy", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresBackupConfigResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresBackupConfigResponse) validateSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.Secret) { // not required
		return nil
	}

	if m.Secret != nil {
		if err := m.Secret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 postgres backup config response based on the context it is used
func (m *V1PostgresBackupConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PostgresBackupConfigResponse) contextValidateSecret(ctx context.Context, formats strfmt.Registry) error {

	if m.Secret != nil {
		if err := m.Secret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PostgresBackupConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PostgresBackupConfigResponse) UnmarshalBinary(b []byte) error {
	var res V1PostgresBackupConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}