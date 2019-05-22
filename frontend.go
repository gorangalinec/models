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
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Frontend Frontend
//
// HAProxy frontend configuration
// swagger:model frontend
type Frontend struct {

	// clflog
	Clflog bool `json:"clflog,omitempty"`

	// client timeout
	ClientTimeout *int64 `json:"client_timeout,omitempty"`

	// clitcpka
	// Enum: [enabled disabled]
	Clitcpka string `json:"clitcpka,omitempty"`

	// contstats
	// Enum: [enabled disabled]
	Contstats string `json:"contstats,omitempty"`

	// default backend
	// Pattern: ^[A-Za-z0-9-_.:]+$
	DefaultBackend string `json:"default_backend,omitempty"`

	// dontlognull
	// Enum: [enabled disabled]
	Dontlognull string `json:"dontlognull,omitempty"`

	// http use htx
	// Enum: [enabled disabled]
	HTTPUseHtx string `json:"http-use-htx,omitempty"`

	// http connection mode
	// Enum: [http-tunnel httpclose forceclose http-server-close http-keep-alive]
	HTTPConnectionMode string `json:"http_connection_mode,omitempty"`

	// http keep alive timeout
	HTTPKeepAliveTimeout *int64 `json:"http_keep_alive_timeout,omitempty"`

	// http pretend keepalive
	// Enum: [enabled disabled]
	HTTPPretendKeepalive string `json:"http_pretend_keepalive,omitempty"`

	// http request timeout
	HTTPRequestTimeout *int64 `json:"http_request_timeout,omitempty"`

	// httplog
	Httplog bool `json:"httplog,omitempty"`

	// log format
	LogFormat string `json:"log_format,omitempty"`

	// log format sd
	LogFormatSd string `json:"log_format_sd,omitempty"`

	// log separate errors
	// Enum: [enabled disabled]
	LogSeparateErrors string `json:"log_separate_errors,omitempty"`

	// log tag
	// Pattern: ^[A-Za-z0-9-_.:]+$
	LogTag string `json:"log_tag,omitempty"`

	// maxconn
	Maxconn *int64 `json:"maxconn,omitempty"`

	// mode
	// Enum: [http tcp]
	Mode string `json:"mode,omitempty"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	Name string `json:"name"`

	// tcplog
	Tcplog bool `json:"tcplog,omitempty"`
}

// Validate validates this frontend
func (m *Frontend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClitcpka(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContstats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultBackend(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDontlognull(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPUseHtx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPConnectionMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPPretendKeepalive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogSeparateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var frontendTypeClitcpkaPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeClitcpkaPropEnum = append(frontendTypeClitcpkaPropEnum, v)
	}
}

const (

	// FrontendClitcpkaEnabled captures enum value "enabled"
	FrontendClitcpkaEnabled string = "enabled"

	// FrontendClitcpkaDisabled captures enum value "disabled"
	FrontendClitcpkaDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateClitcpkaEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, frontendTypeClitcpkaPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateClitcpka(formats strfmt.Registry) error {

	if swag.IsZero(m.Clitcpka) { // not required
		return nil
	}

	// value enum
	if err := m.validateClitcpkaEnum("clitcpka", "body", m.Clitcpka); err != nil {
		return err
	}

	return nil
}

var frontendTypeContstatsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeContstatsPropEnum = append(frontendTypeContstatsPropEnum, v)
	}
}

const (

	// FrontendContstatsEnabled captures enum value "enabled"
	FrontendContstatsEnabled string = "enabled"

	// FrontendContstatsDisabled captures enum value "disabled"
	FrontendContstatsDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateContstatsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, frontendTypeContstatsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateContstats(formats strfmt.Registry) error {

	if swag.IsZero(m.Contstats) { // not required
		return nil
	}

	// value enum
	if err := m.validateContstatsEnum("contstats", "body", m.Contstats); err != nil {
		return err
	}

	return nil
}

func (m *Frontend) validateDefaultBackend(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultBackend) { // not required
		return nil
	}

	if err := validate.Pattern("default_backend", "body", string(m.DefaultBackend), `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

var frontendTypeDontlognullPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeDontlognullPropEnum = append(frontendTypeDontlognullPropEnum, v)
	}
}

const (

	// FrontendDontlognullEnabled captures enum value "enabled"
	FrontendDontlognullEnabled string = "enabled"

	// FrontendDontlognullDisabled captures enum value "disabled"
	FrontendDontlognullDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateDontlognullEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, frontendTypeDontlognullPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateDontlognull(formats strfmt.Registry) error {

	if swag.IsZero(m.Dontlognull) { // not required
		return nil
	}

	// value enum
	if err := m.validateDontlognullEnum("dontlognull", "body", m.Dontlognull); err != nil {
		return err
	}

	return nil
}

var frontendTypeHTTPUseHtxPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeHTTPUseHtxPropEnum = append(frontendTypeHTTPUseHtxPropEnum, v)
	}
}

const (

	// FrontendHTTPUseHtxEnabled captures enum value "enabled"
	FrontendHTTPUseHtxEnabled string = "enabled"

	// FrontendHTTPUseHtxDisabled captures enum value "disabled"
	FrontendHTTPUseHtxDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateHTTPUseHtxEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, frontendTypeHTTPUseHtxPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateHTTPUseHtx(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPUseHtx) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPUseHtxEnum("http-use-htx", "body", m.HTTPUseHtx); err != nil {
		return err
	}

	return nil
}

var frontendTypeHTTPConnectionModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http-tunnel","httpclose","forceclose","http-server-close","http-keep-alive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeHTTPConnectionModePropEnum = append(frontendTypeHTTPConnectionModePropEnum, v)
	}
}

const (

	// FrontendHTTPConnectionModeHTTPTunnel captures enum value "http-tunnel"
	FrontendHTTPConnectionModeHTTPTunnel string = "http-tunnel"

	// FrontendHTTPConnectionModeHttpclose captures enum value "httpclose"
	FrontendHTTPConnectionModeHttpclose string = "httpclose"

	// FrontendHTTPConnectionModeForceclose captures enum value "forceclose"
	FrontendHTTPConnectionModeForceclose string = "forceclose"

	// FrontendHTTPConnectionModeHTTPServerClose captures enum value "http-server-close"
	FrontendHTTPConnectionModeHTTPServerClose string = "http-server-close"

	// FrontendHTTPConnectionModeHTTPKeepAlive captures enum value "http-keep-alive"
	FrontendHTTPConnectionModeHTTPKeepAlive string = "http-keep-alive"
)

// prop value enum
func (m *Frontend) validateHTTPConnectionModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, frontendTypeHTTPConnectionModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateHTTPConnectionMode(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPConnectionMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPConnectionModeEnum("http_connection_mode", "body", m.HTTPConnectionMode); err != nil {
		return err
	}

	return nil
}

var frontendTypeHTTPPretendKeepalivePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeHTTPPretendKeepalivePropEnum = append(frontendTypeHTTPPretendKeepalivePropEnum, v)
	}
}

const (

	// FrontendHTTPPretendKeepaliveEnabled captures enum value "enabled"
	FrontendHTTPPretendKeepaliveEnabled string = "enabled"

	// FrontendHTTPPretendKeepaliveDisabled captures enum value "disabled"
	FrontendHTTPPretendKeepaliveDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateHTTPPretendKeepaliveEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, frontendTypeHTTPPretendKeepalivePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateHTTPPretendKeepalive(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPPretendKeepalive) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPPretendKeepaliveEnum("http_pretend_keepalive", "body", m.HTTPPretendKeepalive); err != nil {
		return err
	}

	return nil
}

var frontendTypeLogSeparateErrorsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeLogSeparateErrorsPropEnum = append(frontendTypeLogSeparateErrorsPropEnum, v)
	}
}

const (

	// FrontendLogSeparateErrorsEnabled captures enum value "enabled"
	FrontendLogSeparateErrorsEnabled string = "enabled"

	// FrontendLogSeparateErrorsDisabled captures enum value "disabled"
	FrontendLogSeparateErrorsDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateLogSeparateErrorsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, frontendTypeLogSeparateErrorsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateLogSeparateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.LogSeparateErrors) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogSeparateErrorsEnum("log_separate_errors", "body", m.LogSeparateErrors); err != nil {
		return err
	}

	return nil
}

func (m *Frontend) validateLogTag(formats strfmt.Registry) error {

	if swag.IsZero(m.LogTag) { // not required
		return nil
	}

	if err := validate.Pattern("log_tag", "body", string(m.LogTag), `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

var frontendTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","tcp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeModePropEnum = append(frontendTypeModePropEnum, v)
	}
}

const (

	// FrontendModeHTTP captures enum value "http"
	FrontendModeHTTP string = "http"

	// FrontendModeTCP captures enum value "tcp"
	FrontendModeTCP string = "tcp"
)

// prop value enum
func (m *Frontend) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, frontendTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *Frontend) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Frontend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Frontend) UnmarshalBinary(b []byte) error {
	var res Frontend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
