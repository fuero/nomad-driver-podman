// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortMapping PortMapping is one or more ports that will be mapped into the container.
//
// swagger:model PortMapping
type PortMapping struct {

	// ContainerPort is the port number that will be exposed from the
	// container.
	// Mandatory.
	ContainerPort uint16 `json:"container_port,omitempty"`

	// HostIP is the IP that we will bind to on the host.
	// If unset, assumed to be 0.0.0.0 (all interfaces).
	HostIP string `json:"host_ip,omitempty"`

	// HostPort is the port number that will be forwarded from the host into
	// the container.
	// If omitted, will be assumed to be identical to
	HostPort uint16 `json:"host_port,omitempty"`

	// Protocol is the protocol forward.
	// Must be either "tcp", "udp", and "sctp", or some combination of these
	// separated by commas.
	// If unset, assumed to be TCP.
	Protocol string `json:"protocol,omitempty"`

	// Range is the number of ports that will be forwarded, starting at
	// HostPort and ContainerPort and counting up.
	// This is 1-indexed, so 1 is assumed to be a single port (only the
	// Hostport:Containerport mapping will be added), 2 is two ports (both
	// Hostport:Containerport and Hostport+1:Containerport+1), etc.
	// If unset, assumed to be 1 (a single port).
	// Both hostport + range and containerport + range must be less than
	// 65536.
	Range uint16 `json:"range,omitempty"`
}

// Validate validates this port mapping
func (m *PortMapping) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortMapping) UnmarshalBinary(b []byte) error {
	var res PortMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}