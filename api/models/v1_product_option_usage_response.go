// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ProductOptionUsageResponse v1 product option usage response
//
// swagger:model v1.ProductOptionUsageResponse
type V1ProductOptionUsageResponse struct {

	// just the usage data of the individual product options summed up
	// Required: true
	Accumulatedusage *V1ProductOptionUsageAccumuluated `json:"accumulatedusage"`

	// the start time in the accounting window to look at
	// Required: true
	// Format: date-time
	From *strfmt.DateTime `json:"from"`

	// the end time in the accounting window to look at (defaults to current system time)
	// Format: date-time
	To strfmt.DateTime `json:"to,omitempty"`

	// the usage data of the individual product options
	// Required: true
	Usage []*V1ProductOptionUsage `json:"usage"`
}

// Validate validates this v1 product option usage response
func (m *V1ProductOptionUsageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccumulatedusage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProductOptionUsageResponse) validateAccumulatedusage(formats strfmt.Registry) error {

	if err := validate.Required("accumulatedusage", "body", m.Accumulatedusage); err != nil {
		return err
	}

	if m.Accumulatedusage != nil {
		if err := m.Accumulatedusage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accumulatedusage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accumulatedusage")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProductOptionUsageResponse) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	if err := validate.FormatOf("from", "body", "date-time", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ProductOptionUsageResponse) validateTo(formats strfmt.Registry) error {
	if swag.IsZero(m.To) { // not required
		return nil
	}

	if err := validate.FormatOf("to", "body", "date-time", m.To.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ProductOptionUsageResponse) validateUsage(formats strfmt.Registry) error {

	if err := validate.Required("usage", "body", m.Usage); err != nil {
		return err
	}

	for i := 0; i < len(m.Usage); i++ {
		if swag.IsZero(m.Usage[i]) { // not required
			continue
		}

		if m.Usage[i] != nil {
			if err := m.Usage[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usage" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 product option usage response based on the context it is used
func (m *V1ProductOptionUsageResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccumulatedusage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProductOptionUsageResponse) contextValidateAccumulatedusage(ctx context.Context, formats strfmt.Registry) error {

	if m.Accumulatedusage != nil {

		if err := m.Accumulatedusage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accumulatedusage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accumulatedusage")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProductOptionUsageResponse) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Usage); i++ {

		if m.Usage[i] != nil {

			if swag.IsZero(m.Usage[i]) { // not required
				return nil
			}

			if err := m.Usage[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usage" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ProductOptionUsageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProductOptionUsageResponse) UnmarshalBinary(b []byte) error {
	var res V1ProductOptionUsageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}