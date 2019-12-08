// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1Replica Distributed Replica specification
// swagger:model v1Replica
type V1Replica struct {

	// Optioanl container to run
	Container *V1Container `json:"container,omitempty"`

	// Optional environment section
	Environment *V1Environment `json:"environment,omitempty"`

	// Optional init section
	Init *V1Init `json:"init,omitempty"`

	// Optioanl mount section
	Mounts *V1Mounts `json:"mounts,omitempty"`

	// Number of replicas
	Replicas int32 `json:"replicas,omitempty"`

	// optional termination section
	Termination *V1Termination `json:"termination,omitempty"`
}

// Validate validates this v1 replica
func (m *V1Replica) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTermination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Replica) validateContainer(formats strfmt.Registry) error {

	if swag.IsZero(m.Container) { // not required
		return nil
	}

	if m.Container != nil {
		if err := m.Container.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container")
			}
			return err
		}
	}

	return nil
}

func (m *V1Replica) validateEnvironment(formats strfmt.Registry) error {

	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *V1Replica) validateInit(formats strfmt.Registry) error {

	if swag.IsZero(m.Init) { // not required
		return nil
	}

	if m.Init != nil {
		if err := m.Init.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("init")
			}
			return err
		}
	}

	return nil
}

func (m *V1Replica) validateMounts(formats strfmt.Registry) error {

	if swag.IsZero(m.Mounts) { // not required
		return nil
	}

	if m.Mounts != nil {
		if err := m.Mounts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mounts")
			}
			return err
		}
	}

	return nil
}

func (m *V1Replica) validateTermination(formats strfmt.Registry) error {

	if swag.IsZero(m.Termination) { // not required
		return nil
	}

	if m.Termination != nil {
		if err := m.Termination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("termination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Replica) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Replica) UnmarshalBinary(b []byte) error {
	var res V1Replica
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}