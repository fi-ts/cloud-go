// Code generated by go-swagger; DO NOT EDIT.

package gateway

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

// NewCreateClientParams creates a new CreateClientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateClientParams() *CreateClientParams {
	return &CreateClientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClientParamsWithTimeout creates a new CreateClientParams object
// with the ability to set a timeout on a request.
func NewCreateClientParamsWithTimeout(timeout time.Duration) *CreateClientParams {
	return &CreateClientParams{
		timeout: timeout,
	}
}

// NewCreateClientParamsWithContext creates a new CreateClientParams object
// with the ability to set a context for a request.
func NewCreateClientParamsWithContext(ctx context.Context) *CreateClientParams {
	return &CreateClientParams{
		Context: ctx,
	}
}

// NewCreateClientParamsWithHTTPClient creates a new CreateClientParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateClientParamsWithHTTPClient(client *http.Client) *CreateClientParams {
	return &CreateClientParams{
		HTTPClient: client,
	}
}

/* CreateClientParams contains all the parameters to send to the API endpoint
   for the create client operation.

   Typically these are written to a http.Request.
*/
type CreateClientParams struct {

	// Body.
	Body *models.V1GatewayInstanceIdentity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClientParams) WithDefaults() *CreateClientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClientParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create client params
func (o *CreateClientParams) WithTimeout(timeout time.Duration) *CreateClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create client params
func (o *CreateClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create client params
func (o *CreateClientParams) WithContext(ctx context.Context) *CreateClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create client params
func (o *CreateClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create client params
func (o *CreateClientParams) WithHTTPClient(client *http.Client) *CreateClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create client params
func (o *CreateClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create client params
func (o *CreateClientParams) WithBody(body *models.V1GatewayInstanceIdentity) *CreateClientParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create client params
func (o *CreateClientParams) SetBody(body *models.V1GatewayInstanceIdentity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
