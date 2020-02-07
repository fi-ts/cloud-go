// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1OwnerReference OwnerReference contains enough information to let you identify an owning object. An owning object must be in the same namespace as the dependent, or be cluster-scoped, so there is no namespace field.
// swagger:model v1.OwnerReference
type V1OwnerReference struct {

	// API version of the referent.
	// Required: true
	APIVersion *string `json:"apiVersion"`

	// If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs "delete" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.
	BlockOwnerDeletion bool `json:"blockOwnerDeletion,omitempty"`

	// If true, this reference points to the managing controller.
	Controller bool `json:"controller,omitempty"`

	// Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// Required: true
	Kind *string `json:"kind"`

	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// Required: true
	Name *string `json:"name"`

	// UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	// Required: true
	UID *string `json:"uid"`
}

// Validate validates this v1 owner reference
func (m *V1OwnerReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OwnerReference) validateAPIVersion(formats strfmt.Registry) error {

	if err := validate.Required("apiVersion", "body", m.APIVersion); err != nil {
		return err
	}

	return nil
}

func (m *V1OwnerReference) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *V1OwnerReference) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1OwnerReference) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OwnerReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OwnerReference) UnmarshalBinary(b []byte) error {
	var res V1OwnerReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
