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

// NewMachineReservationUsageParams creates a new MachineReservationUsageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMachineReservationUsageParams() *MachineReservationUsageParams {
	return &MachineReservationUsageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMachineReservationUsageParamsWithTimeout creates a new MachineReservationUsageParams object
// with the ability to set a timeout on a request.
func NewMachineReservationUsageParamsWithTimeout(timeout time.Duration) *MachineReservationUsageParams {
	return &MachineReservationUsageParams{
		timeout: timeout,
	}
}

// NewMachineReservationUsageParamsWithContext creates a new MachineReservationUsageParams object
// with the ability to set a context for a request.
func NewMachineReservationUsageParamsWithContext(ctx context.Context) *MachineReservationUsageParams {
	return &MachineReservationUsageParams{
		Context: ctx,
	}
}

// NewMachineReservationUsageParamsWithHTTPClient creates a new MachineReservationUsageParams object
// with the ability to set a custom HTTPClient for a request.
func NewMachineReservationUsageParamsWithHTTPClient(client *http.Client) *MachineReservationUsageParams {
	return &MachineReservationUsageParams{
		HTTPClient: client,
	}
}

/*
MachineReservationUsageParams contains all the parameters to send to the API endpoint

	for the machine reservation usage operation.

	Typically these are written to a http.Request.
*/
type MachineReservationUsageParams struct {

	// Body.
	Body *models.V1MachineReservationUsageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the machine reservation usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachineReservationUsageParams) WithDefaults() *MachineReservationUsageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the machine reservation usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachineReservationUsageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the machine reservation usage params
func (o *MachineReservationUsageParams) WithTimeout(timeout time.Duration) *MachineReservationUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the machine reservation usage params
func (o *MachineReservationUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the machine reservation usage params
func (o *MachineReservationUsageParams) WithContext(ctx context.Context) *MachineReservationUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the machine reservation usage params
func (o *MachineReservationUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the machine reservation usage params
func (o *MachineReservationUsageParams) WithHTTPClient(client *http.Client) *MachineReservationUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the machine reservation usage params
func (o *MachineReservationUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the machine reservation usage params
func (o *MachineReservationUsageParams) WithBody(body *models.V1MachineReservationUsageRequest) *MachineReservationUsageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the machine reservation usage params
func (o *MachineReservationUsageParams) SetBody(body *models.V1MachineReservationUsageRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *MachineReservationUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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