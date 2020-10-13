// Code generated by go-swagger; DO NOT EDIT.

package server

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
	"github.com/go-openapi/swag"
)

// NewServerClusterSetupParams creates a new ServerClusterSetupParams object
// with the default values initialized.
func NewServerClusterSetupParams() *ServerClusterSetupParams {
	var ()
	return &ServerClusterSetupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServerClusterSetupParamsWithTimeout creates a new ServerClusterSetupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServerClusterSetupParamsWithTimeout(timeout time.Duration) *ServerClusterSetupParams {
	var ()
	return &ServerClusterSetupParams{

		timeout: timeout,
	}
}

// NewServerClusterSetupParamsWithContext creates a new ServerClusterSetupParams object
// with the default values initialized, and the ability to set a context for a request
func NewServerClusterSetupParamsWithContext(ctx context.Context) *ServerClusterSetupParams {
	var ()
	return &ServerClusterSetupParams{

		Context: ctx,
	}
}

// NewServerClusterSetupParamsWithHTTPClient creates a new ServerClusterSetupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServerClusterSetupParamsWithHTTPClient(client *http.Client) *ServerClusterSetupParams {
	var ()
	return &ServerClusterSetupParams{
		HTTPClient: client,
	}
}

/*ServerClusterSetupParams contains all the parameters to send to the API endpoint
for the server cluster setup operation typically these are written to a http.Request
*/
type ServerClusterSetupParams struct {

	/*EnsureDbsExist
	  List of system databases to ensure exist on the node/cluster. Defaults to ["_users","_replicator"].

	*/
	EnsureDbsExist []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the server cluster setup params
func (o *ServerClusterSetupParams) WithTimeout(timeout time.Duration) *ServerClusterSetupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the server cluster setup params
func (o *ServerClusterSetupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the server cluster setup params
func (o *ServerClusterSetupParams) WithContext(ctx context.Context) *ServerClusterSetupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the server cluster setup params
func (o *ServerClusterSetupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the server cluster setup params
func (o *ServerClusterSetupParams) WithHTTPClient(client *http.Client) *ServerClusterSetupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the server cluster setup params
func (o *ServerClusterSetupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnsureDbsExist adds the ensureDbsExist to the server cluster setup params
func (o *ServerClusterSetupParams) WithEnsureDbsExist(ensureDbsExist []string) *ServerClusterSetupParams {
	o.SetEnsureDbsExist(ensureDbsExist)
	return o
}

// SetEnsureDbsExist adds the ensureDbsExist to the server cluster setup params
func (o *ServerClusterSetupParams) SetEnsureDbsExist(ensureDbsExist []string) {
	o.EnsureDbsExist = ensureDbsExist
}

// WriteToRequest writes these params to a swagger request
func (o *ServerClusterSetupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesEnsureDbsExist := o.EnsureDbsExist

	joinedEnsureDbsExist := swag.JoinByFormat(valuesEnsureDbsExist, "multi")
	// query array param ensure_dbs_exist
	if err := r.SetQueryParam("ensure_dbs_exist", joinedEnsureDbsExist...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
