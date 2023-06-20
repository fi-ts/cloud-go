// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1beta1ShootSSHKeypairRotation v1beta1 shoot SSH keypair rotation
//
// swagger:model v1beta1.ShootSSHKeypairRotation
type V1beta1ShootSSHKeypairRotation struct {

	// last completion time
	LastCompletionTime string `json:"lastCompletionTime,omitempty"`

	// last initiation time
	LastInitiationTime string `json:"lastInitiationTime,omitempty"`
}

// Validate validates this v1beta1 shoot SSH keypair rotation
func (m *V1beta1ShootSSHKeypairRotation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1beta1 shoot SSH keypair rotation based on context it is used
func (m *V1beta1ShootSSHKeypairRotation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1ShootSSHKeypairRotation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1ShootSSHKeypairRotation) UnmarshalBinary(b []byte) error {
	var res V1beta1ShootSSHKeypairRotation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}