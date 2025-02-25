// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDownloadMinimalInitrdParams creates a new DownloadMinimalInitrdParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadMinimalInitrdParams() *DownloadMinimalInitrdParams {
	return &DownloadMinimalInitrdParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadMinimalInitrdParamsWithTimeout creates a new DownloadMinimalInitrdParams object
// with the ability to set a timeout on a request.
func NewDownloadMinimalInitrdParamsWithTimeout(timeout time.Duration) *DownloadMinimalInitrdParams {
	return &DownloadMinimalInitrdParams{
		timeout: timeout,
	}
}

// NewDownloadMinimalInitrdParamsWithContext creates a new DownloadMinimalInitrdParams object
// with the ability to set a context for a request.
func NewDownloadMinimalInitrdParamsWithContext(ctx context.Context) *DownloadMinimalInitrdParams {
	return &DownloadMinimalInitrdParams{
		Context: ctx,
	}
}

// NewDownloadMinimalInitrdParamsWithHTTPClient creates a new DownloadMinimalInitrdParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadMinimalInitrdParamsWithHTTPClient(client *http.Client) *DownloadMinimalInitrdParams {
	return &DownloadMinimalInitrdParams{
		HTTPClient: client,
	}
}

/* DownloadMinimalInitrdParams contains all the parameters to send to the API endpoint
   for the download minimal initrd operation.

   Typically these are written to a http.Request.
*/
type DownloadMinimalInitrdParams struct {

	/* InfraEnvID.

	   The infra env of the host that should be retrieved.

	   Format: uuid
	*/
	InfraEnvID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download minimal initrd params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadMinimalInitrdParams) WithDefaults() *DownloadMinimalInitrdParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download minimal initrd params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadMinimalInitrdParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download minimal initrd params
func (o *DownloadMinimalInitrdParams) WithTimeout(timeout time.Duration) *DownloadMinimalInitrdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download minimal initrd params
func (o *DownloadMinimalInitrdParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download minimal initrd params
func (o *DownloadMinimalInitrdParams) WithContext(ctx context.Context) *DownloadMinimalInitrdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download minimal initrd params
func (o *DownloadMinimalInitrdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download minimal initrd params
func (o *DownloadMinimalInitrdParams) WithHTTPClient(client *http.Client) *DownloadMinimalInitrdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download minimal initrd params
func (o *DownloadMinimalInitrdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfraEnvID adds the infraEnvID to the download minimal initrd params
func (o *DownloadMinimalInitrdParams) WithInfraEnvID(infraEnvID strfmt.UUID) *DownloadMinimalInitrdParams {
	o.SetInfraEnvID(infraEnvID)
	return o
}

// SetInfraEnvID adds the infraEnvId to the download minimal initrd params
func (o *DownloadMinimalInitrdParams) SetInfraEnvID(infraEnvID strfmt.UUID) {
	o.InfraEnvID = infraEnvID
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadMinimalInitrdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param infra_env_id
	if err := r.SetPathParam("infra_env_id", o.InfraEnvID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
