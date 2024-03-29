// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new health API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for health API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Health(params *HealthParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HealthOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Health performs a healthcheck
*/
func (a *Client) Health(params *HealthParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HealthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHealthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "health",
		Method:             "GET",
		PathPattern:        "/v1/health",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HealthReader{formats: a.formats},
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
	success, ok := result.(*HealthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HealthDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
