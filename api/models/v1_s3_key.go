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

// V1S3Key v1 s3 key
//
// swagger:model v1.S3Key
type V1S3Key struct {

	// access key
	// Required: true
	AccessKey *string `json:"access_key"`

	// secret key
	// Required: true
	SecretKey *string `json:"secret_key"`
}

// Validate validates this v1 s3 key
func (m *V1S3Key) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1S3Key) validateAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("access_key", "body", m.AccessKey); err != nil {
		return err
	}

	return nil
}

func (m *V1S3Key) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("secret_key", "body", m.SecretKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 s3 key based on context it is used
func (m *V1S3Key) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1S3Key) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1S3Key) UnmarshalBinary(b []byte) error {
	var res V1S3Key
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
