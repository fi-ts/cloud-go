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

	"github.com/fi-ts/cloud-go/api/models"
)

// NewRestorePostgresParams creates a new RestorePostgresParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestorePostgresParams() *RestorePostgresParams {
	return &RestorePostgresParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestorePostgresParamsWithTimeout creates a new RestorePostgresParams object
// with the ability to set a timeout on a request.
func NewRestorePostgresParamsWithTimeout(timeout time.Duration) *RestorePostgresParams {
	return &RestorePostgresParams{
		timeout: timeout,
	}
}

// NewRestorePostgresParamsWithContext creates a new RestorePostgresParams object
// with the ability to set a context for a request.
func NewRestorePostgresParamsWithContext(ctx context.Context) *RestorePostgresParams {
	return &RestorePostgresParams{
		Context: ctx,
	}
}

// NewRestorePostgresParamsWithHTTPClient creates a new RestorePostgresParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestorePostgresParamsWithHTTPClient(client *http.Client) *RestorePostgresParams {
	return &RestorePostgresParams{
		HTTPClient: client,
	}
}

/*
RestorePostgresParams contains all the parameters to send to the API endpoint

	for the restore postgres operation.

	Typically these are written to a http.Request.
*/
type RestorePostgresParams struct {

	// Body.
	Body *models.V1PostgresRestoreRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restore postgres params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestorePostgresParams) WithDefaults() *RestorePostgresParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restore postgres params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestorePostgresParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restore postgres params
func (o *RestorePostgresParams) WithTimeout(timeout time.Duration) *RestorePostgresParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore postgres params
func (o *RestorePostgresParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore postgres params
func (o *RestorePostgresParams) WithContext(ctx context.Context) *RestorePostgresParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore postgres params
func (o *RestorePostgresParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore postgres params
func (o *RestorePostgresParams) WithHTTPClient(client *http.Client) *RestorePostgresParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore postgres params
func (o *RestorePostgresParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the restore postgres params
func (o *RestorePostgresParams) WithBody(body *models.V1PostgresRestoreRequest) *RestorePostgresParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the restore postgres params
func (o *RestorePostgresParams) SetBody(body *models.V1PostgresRestoreRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RestorePostgresParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
