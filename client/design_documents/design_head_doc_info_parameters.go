// Code generated by go-swagger; DO NOT EDIT.

package design_documents

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

// NewDesignHeadDocInfoParams creates a new DesignHeadDocInfoParams object
// with the default values initialized.
func NewDesignHeadDocInfoParams() *DesignHeadDocInfoParams {
	var ()
	return &DesignHeadDocInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDesignHeadDocInfoParamsWithTimeout creates a new DesignHeadDocInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDesignHeadDocInfoParamsWithTimeout(timeout time.Duration) *DesignHeadDocInfoParams {
	var ()
	return &DesignHeadDocInfoParams{

		timeout: timeout,
	}
}

// NewDesignHeadDocInfoParamsWithContext creates a new DesignHeadDocInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewDesignHeadDocInfoParamsWithContext(ctx context.Context) *DesignHeadDocInfoParams {
	var ()
	return &DesignHeadDocInfoParams{

		Context: ctx,
	}
}

// NewDesignHeadDocInfoParamsWithHTTPClient creates a new DesignHeadDocInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDesignHeadDocInfoParamsWithHTTPClient(client *http.Client) *DesignHeadDocInfoParams {
	var ()
	return &DesignHeadDocInfoParams{
		HTTPClient: client,
	}
}

/*DesignHeadDocInfoParams contains all the parameters to send to the API endpoint
for the design head doc info operation typically these are written to a http.Request
*/
type DesignHeadDocInfoParams struct {

	/*Db
	  Database name

	*/
	Db string
	/*Ddoc
	  Design document id

	*/
	Ddoc string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the design head doc info params
func (o *DesignHeadDocInfoParams) WithTimeout(timeout time.Duration) *DesignHeadDocInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the design head doc info params
func (o *DesignHeadDocInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the design head doc info params
func (o *DesignHeadDocInfoParams) WithContext(ctx context.Context) *DesignHeadDocInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the design head doc info params
func (o *DesignHeadDocInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the design head doc info params
func (o *DesignHeadDocInfoParams) WithHTTPClient(client *http.Client) *DesignHeadDocInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the design head doc info params
func (o *DesignHeadDocInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDb adds the db to the design head doc info params
func (o *DesignHeadDocInfoParams) WithDb(db string) *DesignHeadDocInfoParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the design head doc info params
func (o *DesignHeadDocInfoParams) SetDb(db string) {
	o.Db = db
}

// WithDdoc adds the ddoc to the design head doc info params
func (o *DesignHeadDocInfoParams) WithDdoc(ddoc string) *DesignHeadDocInfoParams {
	o.SetDdoc(ddoc)
	return o
}

// SetDdoc adds the ddoc to the design head doc info params
func (o *DesignHeadDocInfoParams) SetDdoc(ddoc string) {
	o.Ddoc = ddoc
}

// WriteToRequest writes these params to a swagger request
func (o *DesignHeadDocInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param db
	if err := r.SetPathParam("db", o.Db); err != nil {
		return err
	}

	// path param ddoc
	if err := r.SetPathParam("ddoc", o.Ddoc); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}