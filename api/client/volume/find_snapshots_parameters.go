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

	"github.com/fi-ts/cloud-go/api/models"
)

// NewFindSnapshotsParams creates a new FindSnapshotsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindSnapshotsParams() *FindSnapshotsParams {
	return &FindSnapshotsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindSnapshotsParamsWithTimeout creates a new FindSnapshotsParams object
// with the ability to set a timeout on a request.
func NewFindSnapshotsParamsWithTimeout(timeout time.Duration) *FindSnapshotsParams {
	return &FindSnapshotsParams{
		timeout: timeout,
	}
}

// NewFindSnapshotsParamsWithContext creates a new FindSnapshotsParams object
// with the ability to set a context for a request.
func NewFindSnapshotsParamsWithContext(ctx context.Context) *FindSnapshotsParams {
	return &FindSnapshotsParams{
		Context: ctx,
	}
}

// NewFindSnapshotsParamsWithHTTPClient creates a new FindSnapshotsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindSnapshotsParamsWithHTTPClient(client *http.Client) *FindSnapshotsParams {
	return &FindSnapshotsParams{
		HTTPClient: client,
	}
}

/*
FindSnapshotsParams contains all the parameters to send to the API endpoint

	for the find snapshots operation.

	Typically these are written to a http.Request.
*/
type FindSnapshotsParams struct {

	// Body.
	Body *models.V1SnapshotFindRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindSnapshotsParams) WithDefaults() *FindSnapshotsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindSnapshotsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find snapshots params
func (o *FindSnapshotsParams) WithTimeout(timeout time.Duration) *FindSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find snapshots params
func (o *FindSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find snapshots params
func (o *FindSnapshotsParams) WithContext(ctx context.Context) *FindSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find snapshots params
func (o *FindSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find snapshots params
func (o *FindSnapshotsParams) WithHTTPClient(client *http.Client) *FindSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find snapshots params
func (o *FindSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the find snapshots params
func (o *FindSnapshotsParams) WithBody(body *models.V1SnapshotFindRequest) *FindSnapshotsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the find snapshots params
func (o *FindSnapshotsParams) SetBody(body *models.V1SnapshotFindRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FindSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
