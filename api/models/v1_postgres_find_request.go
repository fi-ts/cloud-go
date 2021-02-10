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

// V1PostgresFindRequest v1 postgres find request
//
// swagger:model v1.PostgresFindRequest
type V1PostgresFindRequest struct {

	// description
	// Required: true
	Description *string `json:"Description"`

	// ID
	// Required: true
	ID *string `json:"ID"`

	// partition ID
	// Required: true
	PartitionID *string `json:"PartitionID"`

	// project ID
	// Required: true
	ProjectID *string `json:"ProjectID"`

	// tenant
	// Required: true
	Tenant *string `json:"Tenant"`
}

// Validate validates this v1 postgres find request
func (m *V1PostgresFindRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
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

func (m *V1PostgresFindRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresFindRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresFindRequest) validatePartitionID(formats strfmt.Registry) error {

	if err := validate.Required("PartitionID", "body", m.PartitionID); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresFindRequest) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("ProjectID", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresFindRequest) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("Tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 postgres find request based on context it is used
func (m *V1PostgresFindRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PostgresFindRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PostgresFindRequest) UnmarshalBinary(b []byte) error {
	var res V1PostgresFindRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
