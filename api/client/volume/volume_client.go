// Code generated by go-swagger; DO NOT EDIT.

package volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new volume API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for volume API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteVolume(params *DeleteVolumeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVolumeOK, error)

	FindVolumes(params *FindVolumesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindVolumesOK, error)

	ListVolumes(params *ListVolumesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVolumesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteVolume deletes a volume including all data
*/
func (a *Client) DeleteVolume(params *DeleteVolumeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVolumeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVolumeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteVolume",
		Method:             "DELETE",
		PathPattern:        "/v1/volume/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteVolumeReader{formats: a.formats},
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
	success, ok := result.(*DeleteVolumeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteVolumeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindVolumes finds volumes by multiple criteria
*/
func (a *Client) FindVolumes(params *FindVolumesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindVolumesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindVolumesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findVolumes",
		Method:             "POST",
		PathPattern:        "/v1/volume/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindVolumesReader{formats: a.formats},
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
	success, ok := result.(*FindVolumesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindVolumesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListVolumes gets all volumes
*/
func (a *Client) ListVolumes(params *ListVolumesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVolumesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVolumesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVolumes",
		Method:             "GET",
		PathPattern:        "/v1/volume",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListVolumesReader{formats: a.formats},
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
	success, ok := result.(*ListVolumesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVolumesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
