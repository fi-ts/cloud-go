// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewFindProjectsParams creates a new FindProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindProjectsParams() *FindProjectsParams {
	return &FindProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindProjectsParamsWithTimeout creates a new FindProjectsParams object
// with the ability to set a timeout on a request.
func NewFindProjectsParamsWithTimeout(timeout time.Duration) *FindProjectsParams {
	return &FindProjectsParams{
		timeout: timeout,
	}
}

// NewFindProjectsParamsWithContext creates a new FindProjectsParams object
// with the ability to set a context for a request.
func NewFindProjectsParamsWithContext(ctx context.Context) *FindProjectsParams {
	return &FindProjectsParams{
		Context: ctx,
	}
}

// NewFindProjectsParamsWithHTTPClient creates a new FindProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindProjectsParamsWithHTTPClient(client *http.Client) *FindProjectsParams {
	return &FindProjectsParams{
		HTTPClient: client,
	}
}

/*
FindProjectsParams contains all the parameters to send to the API endpoint

	for the find projects operation.

	Typically these are written to a http.Request.
*/
type FindProjectsParams struct {

	// Body.
	Body *models.V1ProjectFindRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindProjectsParams) WithDefaults() *FindProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindProjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find projects params
func (o *FindProjectsParams) WithTimeout(timeout time.Duration) *FindProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find projects params
func (o *FindProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find projects params
func (o *FindProjectsParams) WithContext(ctx context.Context) *FindProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find projects params
func (o *FindProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find projects params
func (o *FindProjectsParams) WithHTTPClient(client *http.Client) *FindProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find projects params
func (o *FindProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the find projects params
func (o *FindProjectsParams) WithBody(body *models.V1ProjectFindRequest) *FindProjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the find projects params
func (o *FindProjectsParams) SetBody(body *models.V1ProjectFindRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FindProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
