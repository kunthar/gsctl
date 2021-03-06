// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4ClusterDetailsResponseWorkersItemsMemory v4 cluster details response workers items memory
// swagger:model v4ClusterDetailsResponseWorkersItemsMemory
type V4ClusterDetailsResponseWorkersItemsMemory struct {

	// RAM size in GB. Can be an integer or float.
	SizeGb float64 `json:"size_gb,omitempty"`
}

// Validate validates this v4 cluster details response workers items memory
func (m *V4ClusterDetailsResponseWorkersItemsMemory) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4ClusterDetailsResponseWorkersItemsMemory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4ClusterDetailsResponseWorkersItemsMemory) UnmarshalBinary(b []byte) error {
	var res V4ClusterDetailsResponseWorkersItemsMemory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
