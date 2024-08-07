// Code generated by go-swagger; DO NOT EDIT.

package ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new ip API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new ip API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new ip API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for ip API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AllocateIP(params *AllocateIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AllocateIPCreated, error)

	FindIPs(params *FindIPsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindIPsOK, error)

	FreeIP(params *FreeIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FreeIPOK, error)

	GetIP(params *GetIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIPOK, error)

	ListIPs(params *ListIPsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListIPsOK, error)

	UpdateIP(params *UpdateIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateIPOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AllocateIP allocates an ip in a given network
*/
func (a *Client) AllocateIP(params *AllocateIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AllocateIPCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocateIPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "allocateIP",
		Method:             "POST",
		PathPattern:        "/v1/ip/allocate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocateIPReader{formats: a.formats},
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
	success, ok := result.(*AllocateIPCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AllocateIPDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
FindIPs finds ips by multiple criteria
*/
func (a *Client) FindIPs(params *FindIPsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindIPsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindIPsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findIPs",
		Method:             "POST",
		PathPattern:        "/v1/ip/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindIPsReader{formats: a.formats},
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
	success, ok := result.(*FindIPsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindIPsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
FreeIP frees an ip and returns the freed entity
*/
func (a *Client) FreeIP(params *FreeIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FreeIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFreeIPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "freeIP",
		Method:             "DELETE",
		PathPattern:        "/v1/ip/{ip}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FreeIPReader{formats: a.formats},
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
	success, ok := result.(*FreeIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FreeIPDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIP gets ip by address
*/
func (a *Client) GetIP(params *GetIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getIP",
		Method:             "GET",
		PathPattern:        "/v1/ip/{ip}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIPReader{formats: a.formats},
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
	success, ok := result.(*GetIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIPDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListIPs gets all ips
*/
func (a *Client) ListIPs(params *ListIPsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListIPsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListIPsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listIPs",
		Method:             "GET",
		PathPattern:        "/v1/ip",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListIPsReader{formats: a.formats},
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
	success, ok := result.(*ListIPsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListIPsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateIP updates an ip
*/
func (a *Client) UpdateIP(params *UpdateIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateIPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateIP",
		Method:             "POST",
		PathPattern:        "/v1/ip",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateIPReader{formats: a.formats},
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
	success, ok := result.(*UpdateIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateIPDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
