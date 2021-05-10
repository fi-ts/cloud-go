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

// V1GatewayCreatePipesRequest v1 gateway create pipes request
//
// swagger:model v1.GatewayCreatePipesRequest
type V1GatewayCreatePipesRequest struct {

	// client
	// Required: true
	Client *V1GatewayInstanceIdentity `json:"Client"`

	// server
	// Required: true
	Server *V1GatewayInstanceIdentity `json:"Server"`

	// pipes
	// Required: true
	Pipes []*V1PipeSpec `json:"pipes"`
}

// Validate validates this v1 gateway create pipes request
func (m *V1GatewayCreatePipesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GatewayCreatePipesRequest) validateClient(formats strfmt.Registry) error {

	if err := validate.Required("Client", "body", m.Client); err != nil {
		return err
	}

	if m.Client != nil {
		if err := m.Client.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Client")
			}
			return err
		}
	}

	return nil
}

func (m *V1GatewayCreatePipesRequest) validateServer(formats strfmt.Registry) error {

	if err := validate.Required("Server", "body", m.Server); err != nil {
		return err
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Server")
			}
			return err
		}
	}

	return nil
}

func (m *V1GatewayCreatePipesRequest) validatePipes(formats strfmt.Registry) error {

	if err := validate.Required("pipes", "body", m.Pipes); err != nil {
		return err
	}

	for i := 0; i < len(m.Pipes); i++ {
		if swag.IsZero(m.Pipes[i]) { // not required
			continue
		}

		if m.Pipes[i] != nil {
			if err := m.Pipes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 gateway create pipes request based on the context it is used
func (m *V1GatewayCreatePipesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePipes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GatewayCreatePipesRequest) contextValidateClient(ctx context.Context, formats strfmt.Registry) error {

	if m.Client != nil {
		if err := m.Client.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Client")
			}
			return err
		}
	}

	return nil
}

func (m *V1GatewayCreatePipesRequest) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

	if m.Server != nil {
		if err := m.Server.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Server")
			}
			return err
		}
	}

	return nil
}

func (m *V1GatewayCreatePipesRequest) contextValidatePipes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Pipes); i++ {

		if m.Pipes[i] != nil {
			if err := m.Pipes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GatewayCreatePipesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GatewayCreatePipesRequest) UnmarshalBinary(b []byte) error {
	var res V1GatewayCreatePipesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}