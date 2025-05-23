// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateServiceAccountForm update service account form
//
// swagger:model UpdateServiceAccountForm
type UpdateServiceAccountForm struct {

	// is disabled
	IsDisabled *bool `json:"isDisabled,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// role
	// Enum: [None Viewer Editor Admin]
	Role string `json:"role,omitempty"`

	// service account Id
	ServiceAccountID int64 `json:"serviceAccountId,omitempty"`
}

// Validate validates this update service account form
func (m *UpdateServiceAccountForm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateServiceAccountFormTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Viewer","Editor","Admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateServiceAccountFormTypeRolePropEnum = append(updateServiceAccountFormTypeRolePropEnum, v)
	}
}

const (

	// UpdateServiceAccountFormRoleNone captures enum value "None"
	UpdateServiceAccountFormRoleNone string = "None"

	// UpdateServiceAccountFormRoleViewer captures enum value "Viewer"
	UpdateServiceAccountFormRoleViewer string = "Viewer"

	// UpdateServiceAccountFormRoleEditor captures enum value "Editor"
	UpdateServiceAccountFormRoleEditor string = "Editor"

	// UpdateServiceAccountFormRoleAdmin captures enum value "Admin"
	UpdateServiceAccountFormRoleAdmin string = "Admin"
)

// prop value enum
func (m *UpdateServiceAccountForm) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateServiceAccountFormTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateServiceAccountForm) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update service account form based on context it is used
func (m *UpdateServiceAccountForm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateServiceAccountForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateServiceAccountForm) UnmarshalBinary(b []byte) error {
	var res UpdateServiceAccountForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
