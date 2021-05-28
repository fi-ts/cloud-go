// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PostgresBackupSecret v1 postgres backup secret
//
// swagger:model v1.PostgresBackupSecret
type V1PostgresBackupSecret struct {

	// accesskey
	Accesskey string `json:"accesskey,omitempty"`

	// s3encryptionkey
	S3encryptionkey string `json:"s3encryptionkey,omitempty"`

	// secretkey
	Secretkey string `json:"secretkey,omitempty"`
}

// Validate validates this v1 postgres backup secret
func (m *V1PostgresBackupSecret) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 postgres backup secret based on context it is used
func (m *V1PostgresBackupSecret) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PostgresBackupSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PostgresBackupSecret) UnmarshalBinary(b []byte) error {
	var res V1PostgresBackupSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}