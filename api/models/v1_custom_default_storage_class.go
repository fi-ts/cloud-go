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

// V1CustomDefaultStorageClass v1 custom default storage class
//
// swagger:model v1.CustomDefaultStorageClass
type V1CustomDefaultStorageClass struct {

	// class name
	// Required: true
	ClassName *string `json:"ClassName"`
}

// Validate validates this v1 custom default storage class
func (m *V1CustomDefaultStorageClass) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CustomDefaultStorageClass) validateClassName(formats strfmt.Registry) error {

	if err := validate.Required("ClassName", "body", m.ClassName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 custom default storage class based on context it is used
func (m *V1CustomDefaultStorageClass) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CustomDefaultStorageClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CustomDefaultStorageClass) UnmarshalBinary(b []byte) error {
	var res V1CustomDefaultStorageClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
