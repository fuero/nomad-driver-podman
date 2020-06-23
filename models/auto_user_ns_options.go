// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AutoUserNsOptions AutoUserNsOptions defines how to automatically create a user namespace.
//
// swagger:model AutoUserNsOptions
type AutoUserNsOptions struct {

	// AdditionalGIDMappings specified additional GID mappings to include in
	// the generated user namespace.
	AdditionalGIDMappings []*IDMap `json:"AdditionalGIDMappings"`

	// AdditionalUIDMappings specified additional UID mappings to include in
	// the generated user namespace.
	AdditionalUIDMappings []*IDMap `json:"AdditionalUIDMappings"`

	// GroupFile to use if the container uses a volume.
	GroupFile string `json:"GroupFile,omitempty"`

	// InitialSize defines the minimum size for the user namespace.
	// The created user namespace will have at least this size.
	InitialSize uint32 `json:"InitialSize,omitempty"`

	// PasswdFile to use if the container uses a volume.
	PasswdFile string `json:"PasswdFile,omitempty"`

	// Size defines the size for the user namespace.  If it is set to a
	// value bigger than 0, the user namespace will have exactly this size.
	// If it is not set, some heuristics will be used to find its size.
	Size uint32 `json:"Size,omitempty"`
}

// Validate validates this auto user ns options
func (m *AutoUserNsOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalGIDMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdditionalUIDMappings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUserNsOptions) validateAdditionalGIDMappings(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalGIDMappings) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalGIDMappings); i++ {
		if swag.IsZero(m.AdditionalGIDMappings[i]) { // not required
			continue
		}

		if m.AdditionalGIDMappings[i] != nil {
			if err := m.AdditionalGIDMappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AdditionalGIDMappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AutoUserNsOptions) validateAdditionalUIDMappings(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalUIDMappings) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalUIDMappings); i++ {
		if swag.IsZero(m.AdditionalUIDMappings[i]) { // not required
			continue
		}

		if m.AdditionalUIDMappings[i] != nil {
			if err := m.AdditionalUIDMappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AdditionalUIDMappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoUserNsOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoUserNsOptions) UnmarshalBinary(b []byte) error {
	var res AutoUserNsOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}