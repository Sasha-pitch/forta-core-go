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

// Alert alert
//
// swagger:model Alert
type Alert struct {

	// Addresses involved in the source of this alert
	// Example: ["0x98883145049dec03c00cb7708cbc938058802520","0x1fFa3471A45C22B1284fE5a251eD74F40580a1E3"]
	Addresses []string `json:"addresses"`

	// alert Id
	// Example: OZ-GNOSIS-EVENTS
	AlertID string `json:"alertId,omitempty"`

	// Timestamp (RFC3339Nano)
	// Example: 2022-03-01T12:24:33.379756298Z
	CreatedAt string `json:"createdAt,omitempty"`

	// description
	// Example: Detected Transfer event
	Description string `json:"description,omitempty"`

	// finding type
	// Enum: [UNKNOWN_TYPE EXPLOIT SUSPICIOUS DEGRADED INFORMATION]
	FindingType string `json:"findingType,omitempty"`

	// Deterministic alert hash
	// Example: 0xe9cfda18f167de5cdd63c101e38ec0d4cb0a1c2dea80921ecc4405c2b010855f
	Hash string `json:"hash,omitempty"`

	// An associative array of string values
	// Example: {"contractAddress":"0x98883145049dec03c00cb7708cbc938058802520","operator":"0x1fFa3471A45C22B1284fE5a251eD74F40580a1E3"}
	Metadata interface{} `json:"metadata,omitempty"`

	// name
	// Example: Transfer Event
	Name string `json:"name,omitempty"`

	// protocol
	// Example: ethereum
	Protocol string `json:"protocol,omitempty"`

	// severity
	// Enum: [UNKNOWN INFO LOW MEDIUM HIGH CRITICAL]
	Severity string `json:"severity,omitempty"`

	// source
	Source *AlertSource `json:"source,omitempty"`
}

// Validate validates this alert
func (m *Alert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFindingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var alertTypeFindingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN_TYPE","EXPLOIT","SUSPICIOUS","DEGRADED","INFORMATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertTypeFindingTypePropEnum = append(alertTypeFindingTypePropEnum, v)
	}
}

const (

	// AlertFindingTypeUNKNOWNTYPE captures enum value "UNKNOWN_TYPE"
	AlertFindingTypeUNKNOWNTYPE string = "UNKNOWN_TYPE"

	// AlertFindingTypeEXPLOIT captures enum value "EXPLOIT"
	AlertFindingTypeEXPLOIT string = "EXPLOIT"

	// AlertFindingTypeSUSPICIOUS captures enum value "SUSPICIOUS"
	AlertFindingTypeSUSPICIOUS string = "SUSPICIOUS"

	// AlertFindingTypeDEGRADED captures enum value "DEGRADED"
	AlertFindingTypeDEGRADED string = "DEGRADED"

	// AlertFindingTypeINFORMATION captures enum value "INFORMATION"
	AlertFindingTypeINFORMATION string = "INFORMATION"
)

// prop value enum
func (m *Alert) validateFindingTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, alertTypeFindingTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Alert) validateFindingType(formats strfmt.Registry) error {
	if swag.IsZero(m.FindingType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFindingTypeEnum("findingType", "body", m.FindingType); err != nil {
		return err
	}

	return nil
}

var alertTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","INFO","LOW","MEDIUM","HIGH","CRITICAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertTypeSeverityPropEnum = append(alertTypeSeverityPropEnum, v)
	}
}

const (

	// AlertSeverityUNKNOWN captures enum value "UNKNOWN"
	AlertSeverityUNKNOWN string = "UNKNOWN"

	// AlertSeverityINFO captures enum value "INFO"
	AlertSeverityINFO string = "INFO"

	// AlertSeverityLOW captures enum value "LOW"
	AlertSeverityLOW string = "LOW"

	// AlertSeverityMEDIUM captures enum value "MEDIUM"
	AlertSeverityMEDIUM string = "MEDIUM"

	// AlertSeverityHIGH captures enum value "HIGH"
	AlertSeverityHIGH string = "HIGH"

	// AlertSeverityCRITICAL captures enum value "CRITICAL"
	AlertSeverityCRITICAL string = "CRITICAL"
)

// prop value enum
func (m *Alert) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, alertTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Alert) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

func (m *Alert) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this alert based on the context it is used
func (m *Alert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Alert) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Alert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Alert) UnmarshalBinary(b []byte) error {
	var res Alert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
