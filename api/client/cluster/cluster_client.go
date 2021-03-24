// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCluster(params *CreateClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateClusterCreated, error)

	DeleteCluster(params *DeleteClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteClusterOK, error)

	FindCluster(params *FindClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindClusterOK, error)

	FindClusters(params *FindClustersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindClustersOK, error)

	GetClusterKubeconfigTpl(params *GetClusterKubeconfigTplParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterKubeconfigTplOK, error)

	GetSSHKeyPair(params *GetSSHKeyPairParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSSHKeyPairOK, error)

	ListClusters(params *ListClustersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListClustersOK, error)

	ListConstraints(params *ListConstraintsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListConstraintsOK, error)

	ReconcileCluster(params *ReconcileClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReconcileClusterOK, error)

	UpdateCluster(params *UpdateClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateClusterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCluster creates a cluster if the given ID already exists a conflict is returned
*/
func (a *Client) CreateCluster(params *CreateClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateClusterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createCluster",
		Method:             "PUT",
		PathPattern:        "/v1/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateClusterReader{formats: a.formats},
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
	success, ok := result.(*CreateClusterCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteCluster deletes an cluster and returns the deleted entity
*/
func (a *Client) DeleteCluster(params *DeleteClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteCluster",
		Method:             "DELETE",
		PathPattern:        "/v1/cluster/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterReader{formats: a.formats},
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
	success, ok := result.(*DeleteClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindCluster gets cluster by id
*/
func (a *Client) FindCluster(params *FindClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findCluster",
		Method:             "GET",
		PathPattern:        "/v1/cluster/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindClusterReader{formats: a.formats},
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
	success, ok := result.(*FindClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindClusters finds clusters by multiple criteria
*/
func (a *Client) FindClusters(params *FindClustersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindClustersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findClusters",
		Method:             "POST",
		PathPattern:        "/v1/cluster/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindClustersReader{formats: a.formats},
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
	success, ok := result.(*FindClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindClustersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetClusterKubeconfigTpl gets the kubeconfig template just with cluster infos for the cluster
*/
func (a *Client) GetClusterKubeconfigTpl(params *GetClusterKubeconfigTplParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterKubeconfigTplOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterKubeconfigTplParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClusterKubeconfigTpl",
		Method:             "GET",
		PathPattern:        "/v1/cluster/{id}/kubeconfigtpl",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterKubeconfigTplReader{formats: a.formats},
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
	success, ok := result.(*GetClusterKubeconfigTplOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClusterKubeconfigTplDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSSHKeyPair gets all the ssh keypairs of the cluster
*/
func (a *Client) GetSSHKeyPair(params *GetSSHKeyPairParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSSHKeyPairOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSSHKeyPairParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSSHKeyPair",
		Method:             "GET",
		PathPattern:        "/v1/cluster/{id}/sshkeypair",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSSHKeyPairReader{formats: a.formats},
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
	success, ok := result.(*GetSSHKeyPairOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSSHKeyPairDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListClusters gets all clusters
*/
func (a *Client) ListClusters(params *ListClustersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListClustersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listClusters",
		Method:             "GET",
		PathPattern:        "/v1/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListClustersReader{formats: a.formats},
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
	success, ok := result.(*ListClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListClustersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListConstraints gets constraints for cluster create
*/
func (a *Client) ListConstraints(params *ListConstraintsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListConstraintsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListConstraintsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listConstraints",
		Method:             "GET",
		PathPattern:        "/v1/cluster/constraints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListConstraintsReader{formats: a.formats},
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
	success, ok := result.(*ListConstraintsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListConstraintsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReconcileCluster triggers cluster reconcilation
*/
func (a *Client) ReconcileCluster(params *ReconcileClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReconcileClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReconcileClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "reconcileCluster",
		Method:             "POST",
		PathPattern:        "/v1/cluster/{id}/reconcile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReconcileClusterReader{formats: a.formats},
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
	success, ok := result.(*ReconcileClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReconcileClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateCluster updates a cluster if the cluster was changed since this one was read a conflict is returned
*/
func (a *Client) UpdateCluster(params *UpdateClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateCluster",
		Method:             "POST",
		PathPattern:        "/v1/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateClusterReader{formats: a.formats},
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
	success, ok := result.(*UpdateClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
