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

// NewAcceptPostgresRestoreParams creates a new AcceptPostgresRestoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAcceptPostgresRestoreParams() *AcceptPostgresRestoreParams {
	return &AcceptPostgresRestoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptPostgresRestoreParamsWithTimeout creates a new AcceptPostgresRestoreParams object
// with the ability to set a timeout on a request.
func NewAcceptPostgresRestoreParamsWithTimeout(timeout time.Duration) *AcceptPostgresRestoreParams {
	return &AcceptPostgresRestoreParams{
		timeout: timeout,
	}
}

// NewAcceptPostgresRestoreParamsWithContext creates a new AcceptPostgresRestoreParams object
// with the ability to set a context for a request.
func NewAcceptPostgresRestoreParamsWithContext(ctx context.Context) *AcceptPostgresRestoreParams {
	return &AcceptPostgresRestoreParams{
		Context: ctx,
	}
}

// NewAcceptPostgresRestoreParamsWithHTTPClient creates a new AcceptPostgresRestoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewAcceptPostgresRestoreParamsWithHTTPClient(client *http.Client) *AcceptPostgresRestoreParams {
	return &AcceptPostgresRestoreParams{
		HTTPClient: client,
	}
}

/*
AcceptPostgresRestoreParams contains all the parameters to send to the API endpoint

	for the accept postgres restore operation.

	Typically these are written to a http.Request.
*/
type AcceptPostgresRestoreParams struct {

	/* ID.

	   identifier of the postgres
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accept postgres restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptPostgresRestoreParams) WithDefaults() *AcceptPostgresRestoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accept postgres restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptPostgresRestoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the accept postgres restore params
func (o *AcceptPostgresRestoreParams) WithTimeout(timeout time.Duration) *AcceptPostgresRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept postgres restore params
func (o *AcceptPostgresRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept postgres restore params
func (o *AcceptPostgresRestoreParams) WithContext(ctx context.Context) *AcceptPostgresRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept postgres restore params
func (o *AcceptPostgresRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept postgres restore params
func (o *AcceptPostgresRestoreParams) WithHTTPClient(client *http.Client) *AcceptPostgresRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept postgres restore params
func (o *AcceptPostgresRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the accept postgres restore params
func (o *AcceptPostgresRestoreParams) WithID(id string) *AcceptPostgresRestoreParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the accept postgres restore params
func (o *AcceptPostgresRestoreParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptPostgresRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
