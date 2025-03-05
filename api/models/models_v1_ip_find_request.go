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

// ModelsV1IPFindRequest models v1 IP find request
//
// swagger:model models.V1IPFindRequest
type ModelsV1IPFindRequest struct {

	// addressfamily
	Addressfamily string `json:"addressfamily,omitempty"`

	// allocationuuid
	Allocationuuid string `json:"allocationuuid,omitempty"`

	// ipaddress
	Ipaddress string `json:"ipaddress,omitempty"`

	// machineid
	Machineid string `json:"machineid,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// networkid
	Networkid string `json:"networkid,omitempty"`

	// networkprefix
	Networkprefix string `json:"networkprefix,omitempty"`

	// projectid
	Projectid string `json:"projectid,omitempty"`

	// tags
	// Required: true
	Tags []string `json:"tags"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this models v1 IP find request
func (m *ModelsV1IPFindRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1IPFindRequest) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models v1 IP find request based on context it is used
func (m *ModelsV1IPFindRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsV1IPFindRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsV1IPFindRequest) UnmarshalBinary(b []byte) error {
	var res ModelsV1IPFindRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
