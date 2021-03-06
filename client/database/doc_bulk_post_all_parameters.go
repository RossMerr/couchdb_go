// Code generated by go-swagger; DO NOT EDIT.

package database

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

// NewDocBulkPostAllParams creates a new DocBulkPostAllParams object
// with the default values initialized.
func NewDocBulkPostAllParams() *DocBulkPostAllParams {
	var ()
	return &DocBulkPostAllParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDocBulkPostAllParamsWithTimeout creates a new DocBulkPostAllParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDocBulkPostAllParamsWithTimeout(timeout time.Duration) *DocBulkPostAllParams {
	var ()
	return &DocBulkPostAllParams{

		timeout: timeout,
	}
}

// NewDocBulkPostAllParamsWithContext creates a new DocBulkPostAllParams object
// with the default values initialized, and the ability to set a context for a request
func NewDocBulkPostAllParamsWithContext(ctx context.Context) *DocBulkPostAllParams {
	var ()
	return &DocBulkPostAllParams{

		Context: ctx,
	}
}

// NewDocBulkPostAllParamsWithHTTPClient creates a new DocBulkPostAllParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDocBulkPostAllParamsWithHTTPClient(client *http.Client) *DocBulkPostAllParams {
	var ()
	return &DocBulkPostAllParams{
		HTTPClient: client,
	}
}

/*DocBulkPostAllParams contains all the parameters to send to the API endpoint
for the doc bulk post all operation typically these are written to a http.Request
*/
type DocBulkPostAllParams struct {

	/*Body*/
	Body DocBulkPostAllBody
	/*Db
	  Database name

	*/
	Db string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the doc bulk post all params
func (o *DocBulkPostAllParams) WithTimeout(timeout time.Duration) *DocBulkPostAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the doc bulk post all params
func (o *DocBulkPostAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the doc bulk post all params
func (o *DocBulkPostAllParams) WithContext(ctx context.Context) *DocBulkPostAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the doc bulk post all params
func (o *DocBulkPostAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the doc bulk post all params
func (o *DocBulkPostAllParams) WithHTTPClient(client *http.Client) *DocBulkPostAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the doc bulk post all params
func (o *DocBulkPostAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the doc bulk post all params
func (o *DocBulkPostAllParams) WithBody(body DocBulkPostAllBody) *DocBulkPostAllParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the doc bulk post all params
func (o *DocBulkPostAllParams) SetBody(body DocBulkPostAllBody) {
	o.Body = body
}

// WithDb adds the db to the doc bulk post all params
func (o *DocBulkPostAllParams) WithDb(db string) *DocBulkPostAllParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the doc bulk post all params
func (o *DocBulkPostAllParams) SetDb(db string) {
	o.Db = db
}

// WriteToRequest writes these params to a swagger request
func (o *DocBulkPostAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param db
	if err := r.SetPathParam("db", o.Db); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
