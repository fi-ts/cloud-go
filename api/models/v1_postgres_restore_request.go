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

// V1PostgresRestoreRequest v1 postgres restore request
//
// swagger:model v1.PostgresRestoreRequest
type V1PostgresRestoreRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// maintenance
	Maintenance []string `json:"maintenance"`

	// partition ID
	PartitionID string `json:"partitionID,omitempty"`

	// source Id
	// Required: true
	SourceID *string `json:"sourceId"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 postgres restore request
func (m *V1PostgresRestoreRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PostgresRestoreRequest) validateSourceID(formats strfmt.Registry) error {

	if err := validate.Required("sourceId", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 postgres restore request based on context it is used
func (m *V1PostgresRestoreRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PostgresRestoreRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PostgresRestoreRequest) UnmarshalBinary(b []byte) error {
	var res V1PostgresRestoreRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
