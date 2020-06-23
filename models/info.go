// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Info Info is the overall struct that describes the host system
// running libpod/podman
//
// swagger:model Info
type Info struct {

	// registries
	Registries map[string]interface{} `json:"registries,omitempty"`

	// host
	Host *HostInfo `json:"host,omitempty"`

	// store
	Store *StoreInfo `json:"store,omitempty"`

	// version
	Version *Version `json:"version,omitempty"`
}

// Validate validates this info
func (m *Info) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Info) validateHost(formats strfmt.Registry) error {

	if swag.IsZero(m.Host) { // not required
		return nil
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *Info) validateStore(formats strfmt.Registry) error {

	if swag.IsZero(m.Store) { // not required
		return nil
	}

	if m.Store != nil {
		if err := m.Store.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("store")
			}
			return err
		}
	}

	return nil
}

func (m *Info) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Info) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Info) UnmarshalBinary(b []byte) error {
	var res Info
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}