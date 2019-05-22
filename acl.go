// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this files except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ACL ACL Lines
//
// The use of Access Control Lists (ACL) provides a flexible solution to perform
// content switching and generally to take decisions based on content extracted
// from the request, the response or any environmental status.
//
// swagger:model acl
type ACL struct {

	// acl name
	// Required: true
	// Pattern: ^[^\s]+$
	ACLName string `json:"acl_name"`

	// criterion
	// Required: true
	// Pattern: ^[^\s]+$
	Criterion string `json:"criterion"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// value
	// Required: true
	Value string `json:"value"`
}

// Validate validates this acl
func (m *ACL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateACLName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCriterion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ACL) validateACLName(formats strfmt.Registry) error {

	if err := validate.RequiredString("acl_name", "body", string(m.ACLName)); err != nil {
		return err
	}

	if err := validate.Pattern("acl_name", "body", string(m.ACLName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ACL) validateCriterion(formats strfmt.Registry) error {

	if err := validate.RequiredString("criterion", "body", string(m.Criterion)); err != nil {
		return err
	}

	if err := validate.Pattern("criterion", "body", string(m.Criterion), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ACL) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ACL) validateValue(formats strfmt.Registry) error {

	if err := validate.RequiredString("value", "body", string(m.Value)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ACL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ACL) UnmarshalBinary(b []byte) error {
	var res ACL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
