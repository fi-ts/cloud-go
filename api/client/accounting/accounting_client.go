// Code generated by go-swagger; DO NOT EDIT.

package accounting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new accounting API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for accounting API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ClusterUsage(params *ClusterUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ClusterUsageOK, error)

	ClusterUsageCSV(params *ClusterUsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ClusterUsageCSVOK, error)

	ContainerUsage(params *ContainerUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ContainerUsageOK, error)

	ContainerUsageCSV(params *ContainerUsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ContainerUsageCSVOK, error)

	IPUsage(params *IPUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IPUsageOK, error)

	IPUsageCSV(params *IPUsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IPUsageCSVOK, error)

	NetworkUsage(params *NetworkUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NetworkUsageOK, error)

	NetworkUsageCSV(params *NetworkUsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NetworkUsageCSVOK, error)

	PostgresUsage(params *PostgresUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostgresUsageOK, error)

	PostgresUsageCSV(params *PostgresUsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostgresUsageCSVOK, error)

	S3Usage(params *S3UsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*S3UsageOK, error)

	S3UsageCSV(params *S3UsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*S3UsageCSVOK, error)

	VolumeUsage(params *VolumeUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VolumeUsageOK, error)

	VolumeUsageCSV(params *VolumeUsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VolumeUsageCSVOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ClusterUsage finds cluster usage for given accounting query
*/
func (a *Client) ClusterUsage(params *ClusterUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ClusterUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClusterUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "clusterUsage",
		Method:             "POST",
		PathPattern:        "/v1/accounting/cluster-usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ClusterUsageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClusterUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ClusterUsageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ClusterUsageCSV finds cluster usage for given accounting query
*/
func (a *Client) ClusterUsageCSV(params *ClusterUsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ClusterUsageCSVOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClusterUsageCSVParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "clusterUsageCSV",
		Method:             "POST",
		PathPattern:        "/v1/accounting/cluster-usage-csv",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ClusterUsageCSVReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClusterUsageCSVOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ClusterUsageCSVDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ContainerUsage finds container usage for given accounting query
*/
func (a *Client) ContainerUsage(params *ContainerUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ContainerUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "containerUsage",
		Method:             "POST",
		PathPattern:        "/v1/accounting/container-usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContainerUsageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ContainerUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContainerUsageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ContainerUsageCSV finds container usage for given accounting query
*/
func (a *Client) ContainerUsageCSV(params *ContainerUsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ContainerUsageCSVOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerUsageCSVParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "containerUsageCSV",
		Method:             "POST",
		PathPattern:        "/v1/accounting/container-usage-csv",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContainerUsageCSVReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ContainerUsageCSVOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContainerUsageCSVDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IPUsage finds ip usage for given accounting query
*/
func (a *Client) IPUsage(params *IPUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IPUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipUsage",
		Method:             "POST",
		PathPattern:        "/v1/accounting/ip-usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPUsageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IPUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IPUsageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IPUsageCSV finds ip usage for given accounting query
*/
func (a *Client) IPUsageCSV(params *IPUsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IPUsageCSVOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPUsageCSVParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipUsageCSV",
		Method:             "POST",
		PathPattern:        "/v1/accounting/ip-usage-csv",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPUsageCSVReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IPUsageCSVOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IPUsageCSVDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NetworkUsage finds network usage for given accounting query
*/
func (a *Client) NetworkUsage(params *NetworkUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NetworkUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "networkUsage",
		Method:             "POST",
		PathPattern:        "/v1/accounting/network-usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NetworkUsageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NetworkUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NetworkUsageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NetworkUsageCSV finds network usage for given accounting query
*/
func (a *Client) NetworkUsageCSV(params *NetworkUsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NetworkUsageCSVOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkUsageCSVParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "networkUsageCSV",
		Method:             "POST",
		PathPattern:        "/v1/accounting/network-usage-csv",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NetworkUsageCSVReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NetworkUsageCSVOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NetworkUsageCSVDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostgresUsage finds postgres usage for given accounting query
*/
func (a *Client) PostgresUsage(params *PostgresUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostgresUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostgresUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postgresUsage",
		Method:             "POST",
		PathPattern:        "/v1/accounting/postgres-usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostgresUsageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostgresUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostgresUsageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostgresUsageCSV finds postgres usage for given accounting query
*/
func (a *Client) PostgresUsageCSV(params *PostgresUsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostgresUsageCSVOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostgresUsageCSVParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postgresUsageCSV",
		Method:             "POST",
		PathPattern:        "/v1/accounting/postgres-usage-csv",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostgresUsageCSVReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostgresUsageCSVOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostgresUsageCSVDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  S3Usage finds s3 bucket usage for given accounting query
*/
func (a *Client) S3Usage(params *S3UsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*S3UsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewS3UsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "s3Usage",
		Method:             "POST",
		PathPattern:        "/v1/accounting/s3-usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &S3UsageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*S3UsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*S3UsageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  S3UsageCSV finds s3 bucket usage for given accounting query
*/
func (a *Client) S3UsageCSV(params *S3UsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*S3UsageCSVOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewS3UsageCSVParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "s3UsageCSV",
		Method:             "POST",
		PathPattern:        "/v1/accounting/s3-usage-csv",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &S3UsageCSVReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*S3UsageCSVOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*S3UsageCSVDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  VolumeUsage finds volume usage for given accounting query
*/
func (a *Client) VolumeUsage(params *VolumeUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VolumeUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "volumeUsage",
		Method:             "POST",
		PathPattern:        "/v1/accounting/volume-usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VolumeUsageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumeUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VolumeUsageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  VolumeUsageCSV finds volume usage for given accounting query
*/
func (a *Client) VolumeUsageCSV(params *VolumeUsageCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VolumeUsageCSVOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeUsageCSVParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "volumeUsageCSV",
		Method:             "POST",
		PathPattern:        "/v1/accounting/volume-usage-csv",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VolumeUsageCSVReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumeUsageCSVOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VolumeUsageCSVDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
