// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1MasterdataLookupRequest v1 masterdata lookup request
//
// swagger:model v1.MasterdataLookupRequest
type V1MasterdataLookupRequest struct {

	// lookup by clusterID as returned by cloud-api (e.g. 345abc12-3321-4dbc-8d17-55c6ea4fcb38)
	ClusterID string `json:"cluster_id,omitempty"`

	// lookup by clustername and shoot-project
	ClusterNameProject *V1ClusterNameProject `json:"cluster_name_project,omitempty"`

	// lookup at some point in time by projectID as returned by cloud-api (e.g. 10241dd7-a8de-4856-8ac0-b55830b22036)
	ProjectIDTime *V1ProjectIDTime `json:"project_id_time,omitempty"`
}

// Validate validates this v1 masterdata lookup request
func (m *V1MasterdataLookupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterNameProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectIDTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MasterdataLookupRequest) validateClusterNameProject(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterNameProject) { // not required
		return nil
	}

	if m.ClusterNameProject != nil {
		if err := m.ClusterNameProject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_name_project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_name_project")
			}
			return err
		}
	}

	return nil
}

func (m *V1MasterdataLookupRequest) validateProjectIDTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectIDTime) { // not required
		return nil
	}

	if m.ProjectIDTime != nil {
		if err := m.ProjectIDTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project_id_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project_id_time")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 masterdata lookup request based on the context it is used
func (m *V1MasterdataLookupRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterNameProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjectIDTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MasterdataLookupRequest) contextValidateClusterNameProject(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterNameProject != nil {

		if swag.IsZero(m.ClusterNameProject) { // not required
			return nil
		}

		if err := m.ClusterNameProject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_name_project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_name_project")
			}
			return err
		}
	}

	return nil
}

func (m *V1MasterdataLookupRequest) contextValidateProjectIDTime(ctx context.Context, formats strfmt.Registry) error {

	if m.ProjectIDTime != nil {

		if swag.IsZero(m.ProjectIDTime) { // not required
			return nil
		}

		if err := m.ProjectIDTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project_id_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project_id_time")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MasterdataLookupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MasterdataLookupRequest) UnmarshalBinary(b []byte) error {
	var res V1MasterdataLookupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
