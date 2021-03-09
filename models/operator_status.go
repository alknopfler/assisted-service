// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OperatorStatus Represents the operator state.
//
// swagger:model operator-status
type OperatorStatus string

const (

	// OperatorStatusFailed captures enum value "failed"
	OperatorStatusFailed OperatorStatus = "failed"

	// OperatorStatusProgressing captures enum value "progressing"
	OperatorStatusProgressing OperatorStatus = "progressing"

	// OperatorStatusAvailable captures enum value "available"
	OperatorStatusAvailable OperatorStatus = "available"
)

// for schema
var operatorStatusEnum []interface{}

func init() {
	var res []OperatorStatus
	if err := json.Unmarshal([]byte(`["failed","progressing","available"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operatorStatusEnum = append(operatorStatusEnum, v)
	}
}

func (m OperatorStatus) validateOperatorStatusEnum(path, location string, value OperatorStatus) error {
	if err := validate.EnumCase(path, location, value, operatorStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this operator status
func (m OperatorStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOperatorStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}