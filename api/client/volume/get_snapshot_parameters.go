// Code generated by go-swagger; DO NOT EDIT.

package volume

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

// NewGetSnapshotParams creates a new GetSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSnapshotParams() *GetSnapshotParams {
	return &GetSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnapshotParamsWithTimeout creates a new GetSnapshotParams object
// with the ability to set a timeout on a request.
func NewGetSnapshotParamsWithTimeout(timeout time.Duration) *GetSnapshotParams {
	return &GetSnapshotParams{
		timeout: timeout,
	}
}

// NewGetSnapshotParamsWithContext creates a new GetSnapshotParams object
// with the ability to set a context for a request.
func NewGetSnapshotParamsWithContext(ctx context.Context) *GetSnapshotParams {
	return &GetSnapshotParams{
		Context: ctx,
	}
}

// NewGetSnapshotParamsWithHTTPClient creates a new GetSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSnapshotParamsWithHTTPClient(client *http.Client) *GetSnapshotParams {
	return &GetSnapshotParams{
		HTTPClient: client,
	}
}

/*
GetSnapshotParams contains all the parameters to send to the API endpoint

	for the get snapshot operation.

	Typically these are written to a http.Request.
*/
type GetSnapshotParams struct {

	/* ID.

	   identifier of the snapshot
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnapshotParams) WithDefaults() *GetSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get snapshot params
func (o *GetSnapshotParams) WithTimeout(timeout time.Duration) *GetSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snapshot params
func (o *GetSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snapshot params
func (o *GetSnapshotParams) WithContext(ctx context.Context) *GetSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snapshot params
func (o *GetSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snapshot params
func (o *GetSnapshotParams) WithHTTPClient(client *http.Client) *GetSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snapshot params
func (o *GetSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get snapshot params
func (o *GetSnapshotParams) WithID(id string) *GetSnapshotParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get snapshot params
func (o *GetSnapshotParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
