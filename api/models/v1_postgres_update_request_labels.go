// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// V1PostgresUpdateRequestLabels v1 postgres update request labels
//
// swagger:model v1.PostgresUpdateRequest.labels
type V1PostgresUpdateRequestLabels map[string]string

// Validate validates this v1 postgres update request labels
func (m V1PostgresUpdateRequestLabels) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 postgres update request labels based on context it is used
func (m V1PostgresUpdateRequestLabels) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
