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
	ClusterInfo(params *ClusterInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ClusterInfoOK, error)

	DeleteSnapshot(params *DeleteSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSnapshotOK, error)

	DeleteVolume(params *DeleteVolumeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVolumeOK, error)

	FindSnapshots(params *FindSnapshotsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSnapshotsOK, error)

	FindVolumes(params *FindVolumesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindVolumesOK, error)

	GetSnapshot(params *GetSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSnapshotOK, error)

	GetVolume(params *GetVolumeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVolumeOK, error)

	ListSnapshots(params *ListSnapshotsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSnapshotsOK, error)

	ListVolumes(params *ListVolumesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVolumesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ClusterInfo returns info and status to connected storage clusters
*/
func (a *Client) ClusterInfo(params *ClusterInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ClusterInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClusterInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "clusterInfo",
		Method:             "GET",
		PathPattern:        "/v1/volume/clusterinfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ClusterInfoReader{formats: a.formats},
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
	success, ok := result.(*ClusterInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ClusterInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteSnapshot deletes a snapshot including all data
*/
func (a *Client) DeleteSnapshot(params *DeleteSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSnapshot",
		Method:             "DELETE",
		PathPattern:        "/v1/volume/snapshot/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSnapshotReader{formats: a.formats},
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
	success, ok := result.(*DeleteSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSnapshotDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
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
FindSnapshots finds snapshots by multiple criteria
*/
func (a *Client) FindSnapshots(params *FindSnapshotsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSnapshotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSnapshotsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findSnapshots",
		Method:             "POST",
		PathPattern:        "/v1/volume/snapshot/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindSnapshotsReader{formats: a.formats},
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
	success, ok := result.(*FindSnapshotsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindSnapshotsDefault)
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
GetSnapshot gets a snapshot
*/
func (a *Client) GetSnapshot(params *GetSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSnapshot",
		Method:             "GET",
		PathPattern:        "/v1/volume/snapshot/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSnapshotReader{formats: a.formats},
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
	success, ok := result.(*GetSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSnapshotDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVolume gets a volume
*/
func (a *Client) GetVolume(params *GetVolumeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVolumeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVolumeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVolume",
		Method:             "GET",
		PathPattern:        "/v1/volume/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVolumeReader{formats: a.formats},
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
	success, ok := result.(*GetVolumeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVolumeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListSnapshots gets all snapshots
*/
func (a *Client) ListSnapshots(params *ListSnapshotsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSnapshotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSnapshotsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSnapshots",
		Method:             "GET",
		PathPattern:        "/v1/volume/snapshot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSnapshotsReader{formats: a.formats},
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
	success, ok := result.(*ListSnapshotsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListSnapshotsDefault)
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
