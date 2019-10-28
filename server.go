// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
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

// Server Server
//
// HAProxy backend server configuration
// swagger:model server
type Server struct {

	// address
	// Pattern: ^[^\s]+$
	Address string `json:"address,omitempty"`

	// agent addr
	// Pattern: ^[^\s]+$
	AgentAddr string `json:"agent-addr,omitempty"`

	// agent check
	// Enum: [enabled disabled]
	AgentCheck string `json:"agent-check,omitempty"`

	// agent inter
	AgentInter *int64 `json:"agent-inter,omitempty"`

	// agent port
	// Maximum: 65535
	// Minimum: 1
	AgentPort *int64 `json:"agent-port,omitempty"`

	// agent send
	AgentSend string `json:"agent-send,omitempty"`

	// allow 0rtt
	Allow0rtt bool `json:"allow_0rtt,omitempty"`

	// backup
	// Enum: [enabled disabled]
	Backup string `json:"backup,omitempty"`

	// check
	// Enum: [enabled disabled]
	Check string `json:"check,omitempty"`

	// check ssl
	// Enum: [enabled disabled]
	CheckSsl string `json:"check-ssl,omitempty"`

	// cookie
	// Pattern: ^[^\s]+$
	Cookie string `json:"cookie,omitempty"`

	// downinter
	Downinter *int64 `json:"downinter,omitempty"`

	// fastinter
	Fastinter *int64 `json:"fastinter,omitempty"`

	// health check port
	// Maximum: 65535
	// Minimum: 1
	HealthCheckPort *int64 `json:"health_check_port,omitempty"`

	// inter
	Inter *int64 `json:"inter,omitempty"`

	// maintenance
	// Enum: [enabled disabled]
	Maintenance string `json:"maintenance,omitempty"`

	// maxconn
	Maxconn *int64 `json:"maxconn,omitempty"`

	// name
	// Required: true
	// Pattern: ^[^\s]+$
	Name string `json:"name"`

	// on error
	// Enum: [fastinter fail-check sudden-death mark-down]
	OnError string `json:"on-error,omitempty"`

	// on marked down
	// Enum: [shutdown-sessions]
	OnMarkedDown string `json:"on-marked-down,omitempty"`

	// on marked up
	// Enum: [shutdown-backup-sessions]
	OnMarkedUp string `json:"on-marked-up,omitempty"`

	// port
	// Maximum: 65535
	// Minimum: 1
	Port *int64 `json:"port,omitempty"`

	// send proxy
	// Enum: [enabled disabled]
	SendProxy string `json:"send-proxy,omitempty"`

	// send proxy v2
	// Enum: [enabled disabled]
	SendProxyV2 string `json:"send-proxy-v2,omitempty"`

	// ssl
	// Enum: [enabled disabled]
	Ssl string `json:"ssl,omitempty"`

	// ssl cafile
	// Pattern: ^[^\s]+$
	SslCafile string `json:"ssl_cafile,omitempty"`

	// ssl certificate
	// Pattern: ^[^\s]+$
	SslCertificate string `json:"ssl_certificate,omitempty"`

	// tls tickets
	// Enum: [enabled disabled]
	TLSTickets string `json:"tls_tickets,omitempty"`

	// verify
	// Enum: [none required]
	Verify string `json:"verify,omitempty"`

	// weight
	Weight *int64 `json:"weight,omitempty"`
}

// Validate validates this server
func (m *Server) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgentAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgentCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgentPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheckSsl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCookie(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthCheckPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnMarkedDown(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnMarkedUp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSendProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSendProxyV2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslCafile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLSTickets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerify(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Server) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := validate.Pattern("address", "body", string(m.Address), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateAgentAddr(formats strfmt.Registry) error {

	if swag.IsZero(m.AgentAddr) { // not required
		return nil
	}

	if err := validate.Pattern("agent-addr", "body", string(m.AgentAddr), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var serverTypeAgentCheckPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeAgentCheckPropEnum = append(serverTypeAgentCheckPropEnum, v)
	}
}

const (

	// ServerAgentCheckEnabled captures enum value "enabled"
	ServerAgentCheckEnabled string = "enabled"

	// ServerAgentCheckDisabled captures enum value "disabled"
	ServerAgentCheckDisabled string = "disabled"
)

// prop value enum
func (m *Server) validateAgentCheckEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeAgentCheckPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateAgentCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.AgentCheck) { // not required
		return nil
	}

	// value enum
	if err := m.validateAgentCheckEnum("agent-check", "body", m.AgentCheck); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateAgentPort(formats strfmt.Registry) error {

	if swag.IsZero(m.AgentPort) { // not required
		return nil
	}

	if err := validate.MinimumInt("agent-port", "body", int64(*m.AgentPort), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("agent-port", "body", int64(*m.AgentPort), 65535, false); err != nil {
		return err
	}

	return nil
}

var serverTypeBackupPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeBackupPropEnum = append(serverTypeBackupPropEnum, v)
	}
}

const (

	// ServerBackupEnabled captures enum value "enabled"
	ServerBackupEnabled string = "enabled"

	// ServerBackupDisabled captures enum value "disabled"
	ServerBackupDisabled string = "disabled"
)

// prop value enum
func (m *Server) validateBackupEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeBackupPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateBackup(formats strfmt.Registry) error {

	if swag.IsZero(m.Backup) { // not required
		return nil
	}

	// value enum
	if err := m.validateBackupEnum("backup", "body", m.Backup); err != nil {
		return err
	}

	return nil
}

var serverTypeCheckPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeCheckPropEnum = append(serverTypeCheckPropEnum, v)
	}
}

const (

	// ServerCheckEnabled captures enum value "enabled"
	ServerCheckEnabled string = "enabled"

	// ServerCheckDisabled captures enum value "disabled"
	ServerCheckDisabled string = "disabled"
)

// prop value enum
func (m *Server) validateCheckEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeCheckPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.Check) { // not required
		return nil
	}

	// value enum
	if err := m.validateCheckEnum("check", "body", m.Check); err != nil {
		return err
	}

	return nil
}

var serverTypeCheckSslPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeCheckSslPropEnum = append(serverTypeCheckSslPropEnum, v)
	}
}

const (

	// ServerCheckSslEnabled captures enum value "enabled"
	ServerCheckSslEnabled string = "enabled"

	// ServerCheckSslDisabled captures enum value "disabled"
	ServerCheckSslDisabled string = "disabled"
)

// prop value enum
func (m *Server) validateCheckSslEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeCheckSslPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateCheckSsl(formats strfmt.Registry) error {

	if swag.IsZero(m.CheckSsl) { // not required
		return nil
	}

	// value enum
	if err := m.validateCheckSslEnum("check-ssl", "body", m.CheckSsl); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateCookie(formats strfmt.Registry) error {

	if swag.IsZero(m.Cookie) { // not required
		return nil
	}

	if err := validate.Pattern("cookie", "body", string(m.Cookie), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateHealthCheckPort(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthCheckPort) { // not required
		return nil
	}

	if err := validate.MinimumInt("health_check_port", "body", int64(*m.HealthCheckPort), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("health_check_port", "body", int64(*m.HealthCheckPort), 65535, false); err != nil {
		return err
	}

	return nil
}

var serverTypeMaintenancePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeMaintenancePropEnum = append(serverTypeMaintenancePropEnum, v)
	}
}

const (

	// ServerMaintenanceEnabled captures enum value "enabled"
	ServerMaintenanceEnabled string = "enabled"

	// ServerMaintenanceDisabled captures enum value "disabled"
	ServerMaintenanceDisabled string = "disabled"
)

// prop value enum
func (m *Server) validateMaintenanceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeMaintenancePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateMaintenance(formats strfmt.Registry) error {

	if swag.IsZero(m.Maintenance) { // not required
		return nil
	}

	// value enum
	if err := m.validateMaintenanceEnum("maintenance", "body", m.Maintenance); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var serverTypeOnErrorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fastinter","fail-check","sudden-death","mark-down"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeOnErrorPropEnum = append(serverTypeOnErrorPropEnum, v)
	}
}

const (

	// ServerOnErrorFastinter captures enum value "fastinter"
	ServerOnErrorFastinter string = "fastinter"

	// ServerOnErrorFailCheck captures enum value "fail-check"
	ServerOnErrorFailCheck string = "fail-check"

	// ServerOnErrorSuddenDeath captures enum value "sudden-death"
	ServerOnErrorSuddenDeath string = "sudden-death"

	// ServerOnErrorMarkDown captures enum value "mark-down"
	ServerOnErrorMarkDown string = "mark-down"
)

// prop value enum
func (m *Server) validateOnErrorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeOnErrorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateOnError(formats strfmt.Registry) error {

	if swag.IsZero(m.OnError) { // not required
		return nil
	}

	// value enum
	if err := m.validateOnErrorEnum("on-error", "body", m.OnError); err != nil {
		return err
	}

	return nil
}

var serverTypeOnMarkedDownPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["shutdown-sessions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeOnMarkedDownPropEnum = append(serverTypeOnMarkedDownPropEnum, v)
	}
}

const (

	// ServerOnMarkedDownShutdownSessions captures enum value "shutdown-sessions"
	ServerOnMarkedDownShutdownSessions string = "shutdown-sessions"
)

// prop value enum
func (m *Server) validateOnMarkedDownEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeOnMarkedDownPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateOnMarkedDown(formats strfmt.Registry) error {

	if swag.IsZero(m.OnMarkedDown) { // not required
		return nil
	}

	// value enum
	if err := m.validateOnMarkedDownEnum("on-marked-down", "body", m.OnMarkedDown); err != nil {
		return err
	}

	return nil
}

var serverTypeOnMarkedUpPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["shutdown-backup-sessions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeOnMarkedUpPropEnum = append(serverTypeOnMarkedUpPropEnum, v)
	}
}

const (

	// ServerOnMarkedUpShutdownBackupSessions captures enum value "shutdown-backup-sessions"
	ServerOnMarkedUpShutdownBackupSessions string = "shutdown-backup-sessions"
)

// prop value enum
func (m *Server) validateOnMarkedUpEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeOnMarkedUpPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateOnMarkedUp(formats strfmt.Registry) error {

	if swag.IsZero(m.OnMarkedUp) { // not required
		return nil
	}

	// value enum
	if err := m.validateOnMarkedUpEnum("on-marked-up", "body", m.OnMarkedUp); err != nil {
		return err
	}

	return nil
}

func (m *Server) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

var serverTypeSendProxyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeSendProxyPropEnum = append(serverTypeSendProxyPropEnum, v)
	}
}

const (

	// ServerSendProxyEnabled captures enum value "enabled"
	ServerSendProxyEnabled string = "enabled"

	// ServerSendProxyDisabled captures enum value "disabled"
	ServerSendProxyDisabled string = "disabled"
)

// prop value enum
func (m *Server) validateSendProxyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeSendProxyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateSendProxy(formats strfmt.Registry) error {

	if swag.IsZero(m.SendProxy) { // not required
		return nil
	}

	// value enum
	if err := m.validateSendProxyEnum("send-proxy", "body", m.SendProxy); err != nil {
		return err
	}

	return nil
}

var serverTypeSendProxyV2PropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeSendProxyV2PropEnum = append(serverTypeSendProxyV2PropEnum, v)
	}
}

const (

	// ServerSendProxyV2Enabled captures enum value "enabled"
	ServerSendProxyV2Enabled string = "enabled"

	// ServerSendProxyV2Disabled captures enum value "disabled"
	ServerSendProxyV2Disabled string = "disabled"
)

// prop value enum
func (m *Server) validateSendProxyV2Enum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeSendProxyV2PropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateSendProxyV2(formats strfmt.Registry) error {

	if swag.IsZero(m.SendProxyV2) { // not required
		return nil
	}

	// value enum
	if err := m.validateSendProxyV2Enum("send-proxy-v2", "body", m.SendProxyV2); err != nil {
		return err
	}

	return nil
}

var serverTypeSslPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeSslPropEnum = append(serverTypeSslPropEnum, v)
	}
}

const (

	// ServerSslEnabled captures enum value "enabled"
	ServerSslEnabled string = "enabled"

	// ServerSslDisabled captures enum value "disabled"
	ServerSslDisabled string = "disabled"
)

// prop value enum
func (m *Server) validateSslEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeSslPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateSsl(formats strfmt.Registry) error {

	if swag.IsZero(m.Ssl) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslEnum("ssl", "body", m.Ssl); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateSslCafile(formats strfmt.Registry) error {

	if swag.IsZero(m.SslCafile) { // not required
		return nil
	}

	if err := validate.Pattern("ssl_cafile", "body", string(m.SslCafile), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateSslCertificate(formats strfmt.Registry) error {

	if swag.IsZero(m.SslCertificate) { // not required
		return nil
	}

	if err := validate.Pattern("ssl_certificate", "body", string(m.SslCertificate), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var serverTypeTLSTicketsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeTLSTicketsPropEnum = append(serverTypeTLSTicketsPropEnum, v)
	}
}

const (

	// ServerTLSTicketsEnabled captures enum value "enabled"
	ServerTLSTicketsEnabled string = "enabled"

	// ServerTLSTicketsDisabled captures enum value "disabled"
	ServerTLSTicketsDisabled string = "disabled"
)

// prop value enum
func (m *Server) validateTLSTicketsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeTLSTicketsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateTLSTickets(formats strfmt.Registry) error {

	if swag.IsZero(m.TLSTickets) { // not required
		return nil
	}

	// value enum
	if err := m.validateTLSTicketsEnum("tls_tickets", "body", m.TLSTickets); err != nil {
		return err
	}

	return nil
}

var serverTypeVerifyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","required"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeVerifyPropEnum = append(serverTypeVerifyPropEnum, v)
	}
}

const (

	// ServerVerifyNone captures enum value "none"
	ServerVerifyNone string = "none"

	// ServerVerifyRequired captures enum value "required"
	ServerVerifyRequired string = "required"
)

// prop value enum
func (m *Server) validateVerifyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeVerifyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateVerify(formats strfmt.Registry) error {

	if swag.IsZero(m.Verify) { // not required
		return nil
	}

	// value enum
	if err := m.validateVerifyEnum("verify", "body", m.Verify); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Server) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Server) UnmarshalBinary(b []byte) error {
	var res Server
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
