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
	"github.com/go-openapi/swag"

	"github.com/fi-ts/cloud-go/api/models"
)

// NewFindClustersParams creates a new FindClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindClustersParams() *FindClustersParams {
	return &FindClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindClustersParamsWithTimeout creates a new FindClustersParams object
// with the ability to set a timeout on a request.
func NewFindClustersParamsWithTimeout(timeout time.Duration) *FindClustersParams {
	return &FindClustersParams{
		timeout: timeout,
	}
}

// NewFindClustersParamsWithContext creates a new FindClustersParams object
// with the ability to set a context for a request.
func NewFindClustersParamsWithContext(ctx context.Context) *FindClustersParams {
	return &FindClustersParams{
		Context: ctx,
	}
}

// NewFindClustersParamsWithHTTPClient creates a new FindClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindClustersParamsWithHTTPClient(client *http.Client) *FindClustersParams {
	return &FindClustersParams{
		HTTPClient: client,
	}
}

/*
FindClustersParams contains all the parameters to send to the API endpoint

	for the find clusters operation.

	Typically these are written to a http.Request.
*/
type FindClustersParams struct {

	// Body.
	Body *models.V1ClusterFindRequest

	/* ReturnMachines.

	   returns machines in the response, which is disabled by default for list responses because it makes the request slower
	*/
	ReturnMachines *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindClustersParams) WithDefaults() *FindClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindClustersParams) SetDefaults() {
	var (
		returnMachinesDefault = bool(false)
	)

	val := FindClustersParams{
		ReturnMachines: &returnMachinesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the find clusters params
func (o *FindClustersParams) WithTimeout(timeout time.Duration) *FindClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find clusters params
func (o *FindClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find clusters params
func (o *FindClustersParams) WithContext(ctx context.Context) *FindClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find clusters params
func (o *FindClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find clusters params
func (o *FindClustersParams) WithHTTPClient(client *http.Client) *FindClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find clusters params
func (o *FindClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the find clusters params
func (o *FindClustersParams) WithBody(body *models.V1ClusterFindRequest) *FindClustersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the find clusters params
func (o *FindClustersParams) SetBody(body *models.V1ClusterFindRequest) {
	o.Body = body
}

// WithReturnMachines adds the returnMachines to the find clusters params
func (o *FindClustersParams) WithReturnMachines(returnMachines *bool) *FindClustersParams {
	o.SetReturnMachines(returnMachines)
	return o
}

// SetReturnMachines adds the returnMachines to the find clusters params
func (o *FindClustersParams) SetReturnMachines(returnMachines *bool) {
	o.ReturnMachines = returnMachines
}

// WriteToRequest writes these params to a swagger request
func (o *FindClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ReturnMachines != nil {

		// query param return-machines
		var qrReturnMachines bool

		if o.ReturnMachines != nil {
			qrReturnMachines = *o.ReturnMachines
		}
		qReturnMachines := swag.FormatBool(qrReturnMachines)
		if qReturnMachines != "" {

			if err := r.SetQueryParam("return-machines", qReturnMachines); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
