// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewResetMachineParams creates a new ResetMachineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResetMachineParams() *ResetMachineParams {
	return &ResetMachineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResetMachineParamsWithTimeout creates a new ResetMachineParams object
// with the ability to set a timeout on a request.
func NewResetMachineParamsWithTimeout(timeout time.Duration) *ResetMachineParams {
	return &ResetMachineParams{
		timeout: timeout,
	}
}

// NewResetMachineParamsWithContext creates a new ResetMachineParams object
// with the ability to set a context for a request.
func NewResetMachineParamsWithContext(ctx context.Context) *ResetMachineParams {
	return &ResetMachineParams{
		Context: ctx,
	}
}

// NewResetMachineParamsWithHTTPClient creates a new ResetMachineParams object
// with the ability to set a custom HTTPClient for a request.
func NewResetMachineParamsWithHTTPClient(client *http.Client) *ResetMachineParams {
	return &ResetMachineParams{
		HTTPClient: client,
	}
}

/* ResetMachineParams contains all the parameters to send to the API endpoint
   for the reset machine operation.

   Typically these are written to a http.Request.
*/
type ResetMachineParams struct {

	// Body.
	Body *models.V1ClusterMachineResetRequest

	/* ID.

	   identifier of the cluster
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reset machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetMachineParams) WithDefaults() *ResetMachineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reset machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetMachineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reset machine params
func (o *ResetMachineParams) WithTimeout(timeout time.Duration) *ResetMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset machine params
func (o *ResetMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset machine params
func (o *ResetMachineParams) WithContext(ctx context.Context) *ResetMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset machine params
func (o *ResetMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset machine params
func (o *ResetMachineParams) WithHTTPClient(client *http.Client) *ResetMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset machine params
func (o *ResetMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the reset machine params
func (o *ResetMachineParams) WithBody(body *models.V1ClusterMachineResetRequest) *ResetMachineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reset machine params
func (o *ResetMachineParams) SetBody(body *models.V1ClusterMachineResetRequest) {
	o.Body = body
}

// WithID adds the id to the reset machine params
func (o *ResetMachineParams) WithID(id string) *ResetMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the reset machine params
func (o *ResetMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResetMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
