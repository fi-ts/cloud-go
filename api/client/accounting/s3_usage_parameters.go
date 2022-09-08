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

// NewS3UsageParams creates a new S3UsageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3UsageParams() *S3UsageParams {
	return &S3UsageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3UsageParamsWithTimeout creates a new S3UsageParams object
// with the ability to set a timeout on a request.
func NewS3UsageParamsWithTimeout(timeout time.Duration) *S3UsageParams {
	return &S3UsageParams{
		timeout: timeout,
	}
}

// NewS3UsageParamsWithContext creates a new S3UsageParams object
// with the ability to set a context for a request.
func NewS3UsageParamsWithContext(ctx context.Context) *S3UsageParams {
	return &S3UsageParams{
		Context: ctx,
	}
}

// NewS3UsageParamsWithHTTPClient creates a new S3UsageParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3UsageParamsWithHTTPClient(client *http.Client) *S3UsageParams {
	return &S3UsageParams{
		HTTPClient: client,
	}
}

/*
S3UsageParams contains all the parameters to send to the API endpoint

	for the s3 usage operation.

	Typically these are written to a http.Request.
*/
type S3UsageParams struct {

	// Body.
	Body *models.V1S3UsageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3UsageParams) WithDefaults() *S3UsageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3UsageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s3 usage params
func (o *S3UsageParams) WithTimeout(timeout time.Duration) *S3UsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 usage params
func (o *S3UsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 usage params
func (o *S3UsageParams) WithContext(ctx context.Context) *S3UsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 usage params
func (o *S3UsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 usage params
func (o *S3UsageParams) WithHTTPClient(client *http.Client) *S3UsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 usage params
func (o *S3UsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the s3 usage params
func (o *S3UsageParams) WithBody(body *models.V1S3UsageRequest) *S3UsageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the s3 usage params
func (o *S3UsageParams) SetBody(body *models.V1S3UsageRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *S3UsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
