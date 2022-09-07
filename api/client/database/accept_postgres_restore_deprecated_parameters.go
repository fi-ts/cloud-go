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

// NewAcceptPostgresRestoreDeprecatedParams creates a new AcceptPostgresRestoreDeprecatedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAcceptPostgresRestoreDeprecatedParams() *AcceptPostgresRestoreDeprecatedParams {
	return &AcceptPostgresRestoreDeprecatedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptPostgresRestoreDeprecatedParamsWithTimeout creates a new AcceptPostgresRestoreDeprecatedParams object
// with the ability to set a timeout on a request.
func NewAcceptPostgresRestoreDeprecatedParamsWithTimeout(timeout time.Duration) *AcceptPostgresRestoreDeprecatedParams {
	return &AcceptPostgresRestoreDeprecatedParams{
		timeout: timeout,
	}
}

// NewAcceptPostgresRestoreDeprecatedParamsWithContext creates a new AcceptPostgresRestoreDeprecatedParams object
// with the ability to set a context for a request.
func NewAcceptPostgresRestoreDeprecatedParamsWithContext(ctx context.Context) *AcceptPostgresRestoreDeprecatedParams {
	return &AcceptPostgresRestoreDeprecatedParams{
		Context: ctx,
	}
}

// NewAcceptPostgresRestoreDeprecatedParamsWithHTTPClient creates a new AcceptPostgresRestoreDeprecatedParams object
// with the ability to set a custom HTTPClient for a request.
func NewAcceptPostgresRestoreDeprecatedParamsWithHTTPClient(client *http.Client) *AcceptPostgresRestoreDeprecatedParams {
	return &AcceptPostgresRestoreDeprecatedParams{
		HTTPClient: client,
	}
}

/*
AcceptPostgresRestoreDeprecatedParams contains all the parameters to send to the API endpoint

	for the accept postgres restore deprecated operation.

	Typically these are written to a http.Request.
*/
type AcceptPostgresRestoreDeprecatedParams struct {

	/* ID.

	   identifier of the postgres
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accept postgres restore deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptPostgresRestoreDeprecatedParams) WithDefaults() *AcceptPostgresRestoreDeprecatedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accept postgres restore deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptPostgresRestoreDeprecatedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the accept postgres restore deprecated params
func (o *AcceptPostgresRestoreDeprecatedParams) WithTimeout(timeout time.Duration) *AcceptPostgresRestoreDeprecatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept postgres restore deprecated params
func (o *AcceptPostgresRestoreDeprecatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept postgres restore deprecated params
func (o *AcceptPostgresRestoreDeprecatedParams) WithContext(ctx context.Context) *AcceptPostgresRestoreDeprecatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept postgres restore deprecated params
func (o *AcceptPostgresRestoreDeprecatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept postgres restore deprecated params
func (o *AcceptPostgresRestoreDeprecatedParams) WithHTTPClient(client *http.Client) *AcceptPostgresRestoreDeprecatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept postgres restore deprecated params
func (o *AcceptPostgresRestoreDeprecatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the accept postgres restore deprecated params
func (o *AcceptPostgresRestoreDeprecatedParams) WithID(id string) *AcceptPostgresRestoreDeprecatedParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the accept postgres restore deprecated params
func (o *AcceptPostgresRestoreDeprecatedParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptPostgresRestoreDeprecatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
