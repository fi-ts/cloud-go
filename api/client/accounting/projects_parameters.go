// Code generated by go-swagger; DO NOT EDIT.

package accounting

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

// NewProjectsParams creates a new ProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsParams() *ProjectsParams {
	return &ProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsParamsWithTimeout creates a new ProjectsParams object
// with the ability to set a timeout on a request.
func NewProjectsParamsWithTimeout(timeout time.Duration) *ProjectsParams {
	return &ProjectsParams{
		timeout: timeout,
	}
}

// NewProjectsParamsWithContext creates a new ProjectsParams object
// with the ability to set a context for a request.
func NewProjectsParamsWithContext(ctx context.Context) *ProjectsParams {
	return &ProjectsParams{
		Context: ctx,
	}
}

// NewProjectsParamsWithHTTPClient creates a new ProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsParamsWithHTTPClient(client *http.Client) *ProjectsParams {
	return &ProjectsParams{
		HTTPClient: client,
	}
}

/*
ProjectsParams contains all the parameters to send to the API endpoint

	for the projects operation.

	Typically these are written to a http.Request.
*/
type ProjectsParams struct {

	// Body.
	Body *models.V1ProjectInfoRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsParams) WithDefaults() *ProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects params
func (o *ProjectsParams) WithTimeout(timeout time.Duration) *ProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects params
func (o *ProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects params
func (o *ProjectsParams) WithContext(ctx context.Context) *ProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects params
func (o *ProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects params
func (o *ProjectsParams) WithHTTPClient(client *http.Client) *ProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects params
func (o *ProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the projects params
func (o *ProjectsParams) WithBody(body *models.V1ProjectInfoRequest) *ProjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the projects params
func (o *ProjectsParams) SetBody(body *models.V1ProjectInfoRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
