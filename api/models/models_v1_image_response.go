// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsV1ImageResponse models v1 image response
//
// swagger:model models.V1ImageResponse
type ModelsV1ImageResponse struct {

	// changed
	Changed string `json:"changed,omitempty"`

	// classification
	Classification string `json:"classification,omitempty"`

	// created
	Created string `json:"created,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// expiration date
	// Required: true
	ExpirationDate *string `json:"expirationDate"`

	// features
	// Required: true
	Features []string `json:"features"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// usedby
	// Required: true
	Usedby []string `json:"usedby"`
}

// Validate validates this models v1 image response
func (m *ModelsV1ImageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedby(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1ImageResponse) validateExpirationDate(formats strfmt.Registry) error {

	if err := validate.Required("expirationDate", "body", m.ExpirationDate); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1ImageResponse) validateFeatures(formats strfmt.Registry) error {

	if err := validate.Required("features", "body", m.Features); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1ImageResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1ImageResponse) validateUsedby(formats strfmt.Registry) error {

	if err := validate.Required("usedby", "body", m.Usedby); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsV1ImageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsV1ImageResponse) UnmarshalBinary(b []byte) error {
	var res ModelsV1ImageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
