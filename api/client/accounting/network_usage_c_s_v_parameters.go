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

// NewNetworkUsageCSVParams creates a new NetworkUsageCSVParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkUsageCSVParams() *NetworkUsageCSVParams {
	return &NetworkUsageCSVParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkUsageCSVParamsWithTimeout creates a new NetworkUsageCSVParams object
// with the ability to set a timeout on a request.
func NewNetworkUsageCSVParamsWithTimeout(timeout time.Duration) *NetworkUsageCSVParams {
	return &NetworkUsageCSVParams{
		timeout: timeout,
	}
}

// NewNetworkUsageCSVParamsWithContext creates a new NetworkUsageCSVParams object
// with the ability to set a context for a request.
func NewNetworkUsageCSVParamsWithContext(ctx context.Context) *NetworkUsageCSVParams {
	return &NetworkUsageCSVParams{
		Context: ctx,
	}
}

// NewNetworkUsageCSVParamsWithHTTPClient creates a new NetworkUsageCSVParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkUsageCSVParamsWithHTTPClient(client *http.Client) *NetworkUsageCSVParams {
	return &NetworkUsageCSVParams{
		HTTPClient: client,
	}
}

/*
NetworkUsageCSVParams contains all the parameters to send to the API endpoint

	for the network usage c s v operation.

	Typically these are written to a http.Request.
*/
type NetworkUsageCSVParams struct {

	// Body.
	Body *models.V1NetworkUsageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network usage c s v params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkUsageCSVParams) WithDefaults() *NetworkUsageCSVParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network usage c s v params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkUsageCSVParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network usage c s v params
func (o *NetworkUsageCSVParams) WithTimeout(timeout time.Duration) *NetworkUsageCSVParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network usage c s v params
func (o *NetworkUsageCSVParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network usage c s v params
func (o *NetworkUsageCSVParams) WithContext(ctx context.Context) *NetworkUsageCSVParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network usage c s v params
func (o *NetworkUsageCSVParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network usage c s v params
func (o *NetworkUsageCSVParams) WithHTTPClient(client *http.Client) *NetworkUsageCSVParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network usage c s v params
func (o *NetworkUsageCSVParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the network usage c s v params
func (o *NetworkUsageCSVParams) WithBody(body *models.V1NetworkUsageRequest) *NetworkUsageCSVParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the network usage c s v params
func (o *NetworkUsageCSVParams) SetBody(body *models.V1NetworkUsageRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkUsageCSVParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
