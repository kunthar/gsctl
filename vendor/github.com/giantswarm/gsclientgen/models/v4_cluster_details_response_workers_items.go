// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V4ClusterDetailsResponseWorkersItems v4 cluster details response workers items
// swagger:model v4ClusterDetailsResponseWorkersItems
type V4ClusterDetailsResponseWorkersItems struct {

	// aws
	Aws *V4ClusterDetailsResponseWorkersItemsAws `json:"aws,omitempty"`

	// azure
	Azure *V4ClusterDetailsResponseWorkersItemsAzure `json:"azure,omitempty"`

	// cpu
	CPU *V4ClusterDetailsResponseWorkersItemsCPU `json:"cpu,omitempty"`

	// labels
	Labels interface{} `json:"labels,omitempty"`

	// memory
	Memory *V4ClusterDetailsResponseWorkersItemsMemory `json:"memory,omitempty"`

	// storage
	Storage *V4ClusterDetailsResponseWorkersItemsStorage `json:"storage,omitempty"`
}

// Validate validates this v4 cluster details response workers items
func (m *V4ClusterDetailsResponseWorkersItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V4ClusterDetailsResponseWorkersItems) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *V4ClusterDetailsResponseWorkersItems) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *V4ClusterDetailsResponseWorkersItems) validateCPU(formats strfmt.Registry) error {

	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if m.CPU != nil {
		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *V4ClusterDetailsResponseWorkersItems) validateMemory(formats strfmt.Registry) error {

	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if m.Memory != nil {
		if err := m.Memory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory")
			}
			return err
		}
	}

	return nil
}

func (m *V4ClusterDetailsResponseWorkersItems) validateStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.Storage) { // not required
		return nil
	}

	if m.Storage != nil {
		if err := m.Storage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4ClusterDetailsResponseWorkersItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4ClusterDetailsResponseWorkersItems) UnmarshalBinary(b []byte) error {
	var res V4ClusterDetailsResponseWorkersItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}