// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Paging v1 paging
//
// swagger:model v1.Paging
type V1Paging struct {

	// count
	Count int64 `json:"count,omitempty"`

	// page
	Page int64 `json:"page,omitempty"`
}

// Validate validates this v1 paging
func (m *V1Paging) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 paging based on context it is used
func (m *V1Paging) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Paging) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Paging) UnmarshalBinary(b []byte) error {
	var res V1Paging
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}