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

	"github.com/RossMerr/couchdb_go/models"
)

// NewServerReplicationParams creates a new ServerReplicationParams object
// with the default values initialized.
func NewServerReplicationParams() *ServerReplicationParams {
	var ()
	return &ServerReplicationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServerReplicationParamsWithTimeout creates a new ServerReplicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServerReplicationParamsWithTimeout(timeout time.Duration) *ServerReplicationParams {
	var ()
	return &ServerReplicationParams{

		timeout: timeout,
	}
}

// NewServerReplicationParamsWithContext creates a new ServerReplicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewServerReplicationParamsWithContext(ctx context.Context) *ServerReplicationParams {
	var ()
	return &ServerReplicationParams{

		Context: ctx,
	}
}

// NewServerReplicationParamsWithHTTPClient creates a new ServerReplicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServerReplicationParamsWithHTTPClient(client *http.Client) *ServerReplicationParams {
	var ()
	return &ServerReplicationParams{
		HTTPClient: client,
	}
}

/*ServerReplicationParams contains all the parameters to send to the API endpoint
for the server replication operation typically these are written to a http.Request
*/
type ServerReplicationParams struct {

	/*Body*/
	Body *models.Replicate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the server replication params
func (o *ServerReplicationParams) WithTimeout(timeout time.Duration) *ServerReplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the server replication params
func (o *ServerReplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the server replication params
func (o *ServerReplicationParams) WithContext(ctx context.Context) *ServerReplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the server replication params
func (o *ServerReplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the server replication params
func (o *ServerReplicationParams) WithHTTPClient(client *http.Client) *ServerReplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the server replication params
func (o *ServerReplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the server replication params
func (o *ServerReplicationParams) WithBody(body *models.Replicate) *ServerReplicationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the server replication params
func (o *ServerReplicationParams) SetBody(body *models.Replicate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ServerReplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
