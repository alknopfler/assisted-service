// Code generated by go-swagger; DO NOT EDIT.

package operators

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

// NewListOfClusterOperatorsParams creates a new ListOfClusterOperatorsParams object
// with the default values initialized.
func NewListOfClusterOperatorsParams() *ListOfClusterOperatorsParams {
	var ()
	return &ListOfClusterOperatorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOfClusterOperatorsParamsWithTimeout creates a new ListOfClusterOperatorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOfClusterOperatorsParamsWithTimeout(timeout time.Duration) *ListOfClusterOperatorsParams {
	var ()
	return &ListOfClusterOperatorsParams{

		timeout: timeout,
	}
}

// NewListOfClusterOperatorsParamsWithContext creates a new ListOfClusterOperatorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOfClusterOperatorsParamsWithContext(ctx context.Context) *ListOfClusterOperatorsParams {
	var ()
	return &ListOfClusterOperatorsParams{

		Context: ctx,
	}
}

// NewListOfClusterOperatorsParamsWithHTTPClient creates a new ListOfClusterOperatorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOfClusterOperatorsParamsWithHTTPClient(client *http.Client) *ListOfClusterOperatorsParams {
	var ()
	return &ListOfClusterOperatorsParams{
		HTTPClient: client,
	}
}

/*ListOfClusterOperatorsParams contains all the parameters to send to the API endpoint
for the list of cluster operators operation typically these are written to a http.Request
*/
type ListOfClusterOperatorsParams struct {

	/*ClusterID
	  The cluster to return operators for.

	*/
	ClusterID strfmt.UUID
	/*OperatorName
	  An operator in the specified cluster to return its data.

	*/
	OperatorName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list of cluster operators params
func (o *ListOfClusterOperatorsParams) WithTimeout(timeout time.Duration) *ListOfClusterOperatorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list of cluster operators params
func (o *ListOfClusterOperatorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list of cluster operators params
func (o *ListOfClusterOperatorsParams) WithContext(ctx context.Context) *ListOfClusterOperatorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list of cluster operators params
func (o *ListOfClusterOperatorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list of cluster operators params
func (o *ListOfClusterOperatorsParams) WithHTTPClient(client *http.Client) *ListOfClusterOperatorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list of cluster operators params
func (o *ListOfClusterOperatorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list of cluster operators params
func (o *ListOfClusterOperatorsParams) WithClusterID(clusterID strfmt.UUID) *ListOfClusterOperatorsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list of cluster operators params
func (o *ListOfClusterOperatorsParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithOperatorName adds the operatorName to the list of cluster operators params
func (o *ListOfClusterOperatorsParams) WithOperatorName(operatorName *string) *ListOfClusterOperatorsParams {
	o.SetOperatorName(operatorName)
	return o
}

// SetOperatorName adds the operatorName to the list of cluster operators params
func (o *ListOfClusterOperatorsParams) SetOperatorName(operatorName *string) {
	o.OperatorName = operatorName
}

// WriteToRequest writes these params to a swagger request
func (o *ListOfClusterOperatorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if o.OperatorName != nil {

		// query param operator_name
		var qrOperatorName string
		if o.OperatorName != nil {
			qrOperatorName = *o.OperatorName
		}
		qOperatorName := qrOperatorName
		if qOperatorName != "" {
			if err := r.SetQueryParam("operator_name", qOperatorName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}