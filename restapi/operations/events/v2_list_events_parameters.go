// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewV2ListEventsParams creates a new V2ListEventsParams object
//
// There are no default values defined in the spec.
func NewV2ListEventsParams() V2ListEventsParams {

	return V2ListEventsParams{}
}

// V2ListEventsParams contains all the bound params for the v2 list events operation
// typically these are obtained from a http.Request
//
// swagger:parameters v2ListEvents
type V2ListEventsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A comma-separated list of event categories.
	  In: query
	*/
	Categories []string
	/*The cluster to return events for.
	  In: query
	*/
	ClusterID *strfmt.UUID
	/*A host in the specified cluster to return events for.
	  In: query
	*/
	HostID *strfmt.UUID
	/*The infra env to return events for.
	  In: query
	*/
	InfraEnvID *strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewV2ListEventsParams() beforehand.
func (o *V2ListEventsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCategories, qhkCategories, _ := qs.GetOK("categories")
	if err := o.bindCategories(qCategories, qhkCategories, route.Formats); err != nil {
		res = append(res, err)
	}

	qClusterID, qhkClusterID, _ := qs.GetOK("cluster_id")
	if err := o.bindClusterID(qClusterID, qhkClusterID, route.Formats); err != nil {
		res = append(res, err)
	}

	qHostID, qhkHostID, _ := qs.GetOK("host_id")
	if err := o.bindHostID(qHostID, qhkHostID, route.Formats); err != nil {
		res = append(res, err)
	}

	qInfraEnvID, qhkInfraEnvID, _ := qs.GetOK("infra_env_id")
	if err := o.bindInfraEnvID(qInfraEnvID, qhkInfraEnvID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCategories binds and validates array parameter Categories from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *V2ListEventsParams) bindCategories(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvCategories string
	if len(rawData) > 0 {
		qvCategories = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	categoriesIC := swag.SplitByFormat(qvCategories, "")
	if len(categoriesIC) == 0 {
		return nil
	}

	var categoriesIR []string
	for _, categoriesIV := range categoriesIC {
		categoriesI := categoriesIV

		categoriesIR = append(categoriesIR, categoriesI)
	}

	o.Categories = categoriesIR

	return nil
}

// bindClusterID binds and validates parameter ClusterID from query.
func (o *V2ListEventsParams) bindClusterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("cluster_id", "query", "strfmt.UUID", raw)
	}
	o.ClusterID = (value.(*strfmt.UUID))

	if err := o.validateClusterID(formats); err != nil {
		return err
	}

	return nil
}

// validateClusterID carries on validations for parameter ClusterID
func (o *V2ListEventsParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.FormatOf("cluster_id", "query", "uuid", o.ClusterID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindHostID binds and validates parameter HostID from query.
func (o *V2ListEventsParams) bindHostID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("host_id", "query", "strfmt.UUID", raw)
	}
	o.HostID = (value.(*strfmt.UUID))

	if err := o.validateHostID(formats); err != nil {
		return err
	}

	return nil
}

// validateHostID carries on validations for parameter HostID
func (o *V2ListEventsParams) validateHostID(formats strfmt.Registry) error {

	if err := validate.FormatOf("host_id", "query", "uuid", o.HostID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindInfraEnvID binds and validates parameter InfraEnvID from query.
func (o *V2ListEventsParams) bindInfraEnvID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("infra_env_id", "query", "strfmt.UUID", raw)
	}
	o.InfraEnvID = (value.(*strfmt.UUID))

	if err := o.validateInfraEnvID(formats); err != nil {
		return err
	}

	return nil
}

// validateInfraEnvID carries on validations for parameter InfraEnvID
func (o *V2ListEventsParams) validateInfraEnvID(formats strfmt.Registry) error {

	if err := validate.FormatOf("infra_env_id", "query", "uuid", o.InfraEnvID.String(), formats); err != nil {
		return err
	}
	return nil
}
