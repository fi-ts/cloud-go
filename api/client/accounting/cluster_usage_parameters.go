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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-stack/cloud-go/api/models"
)

// NewClusterUsageParams creates a new ClusterUsageParams object
// with the default values initialized.
func NewClusterUsageParams() *ClusterUsageParams {
	var ()
	return &ClusterUsageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewClusterUsageParamsWithTimeout creates a new ClusterUsageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewClusterUsageParamsWithTimeout(timeout time.Duration) *ClusterUsageParams {
	var ()
	return &ClusterUsageParams{

		timeout: timeout,
	}
}

// NewClusterUsageParamsWithContext creates a new ClusterUsageParams object
// with the default values initialized, and the ability to set a context for a request
func NewClusterUsageParamsWithContext(ctx context.Context) *ClusterUsageParams {
	var ()
	return &ClusterUsageParams{

		Context: ctx,
	}
}

// NewClusterUsageParamsWithHTTPClient creates a new ClusterUsageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewClusterUsageParamsWithHTTPClient(client *http.Client) *ClusterUsageParams {
	var ()
	return &ClusterUsageParams{
		HTTPClient: client,
	}
}

/*ClusterUsageParams contains all the parameters to send to the API endpoint
for the cluster usage operation typically these are written to a http.Request
*/
type ClusterUsageParams struct {

	/*Body*/
	Body *models.V1ClusterUsageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cluster usage params
func (o *ClusterUsageParams) WithTimeout(timeout time.Duration) *ClusterUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster usage params
func (o *ClusterUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster usage params
func (o *ClusterUsageParams) WithContext(ctx context.Context) *ClusterUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster usage params
func (o *ClusterUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster usage params
func (o *ClusterUsageParams) WithHTTPClient(client *http.Client) *ClusterUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster usage params
func (o *ClusterUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cluster usage params
func (o *ClusterUsageParams) WithBody(body *models.V1ClusterUsageRequest) *ClusterUsageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cluster usage params
func (o *ClusterUsageParams) SetBody(body *models.V1ClusterUsageRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
