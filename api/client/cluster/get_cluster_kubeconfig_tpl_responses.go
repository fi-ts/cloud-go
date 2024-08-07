// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fi-ts/cloud-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
)

// GetClusterKubeconfigTplReader is a Reader for the GetClusterKubeconfigTpl structure.
type GetClusterKubeconfigTplReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterKubeconfigTplReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterKubeconfigTplOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClusterKubeconfigTplDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterKubeconfigTplOK creates a GetClusterKubeconfigTplOK with default headers values
func NewGetClusterKubeconfigTplOK() *GetClusterKubeconfigTplOK {
	return &GetClusterKubeconfigTplOK{}
}

/*
GetClusterKubeconfigTplOK describes a response with status code 200, with default header values.

OK
*/
type GetClusterKubeconfigTplOK struct {
	Payload *models.V1ClusterKubeconfigResponse
}

// IsSuccess returns true when this get cluster kubeconfig tpl o k response has a 2xx status code
func (o *GetClusterKubeconfigTplOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster kubeconfig tpl o k response has a 3xx status code
func (o *GetClusterKubeconfigTplOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster kubeconfig tpl o k response has a 4xx status code
func (o *GetClusterKubeconfigTplOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster kubeconfig tpl o k response has a 5xx status code
func (o *GetClusterKubeconfigTplOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster kubeconfig tpl o k response a status code equal to that given
func (o *GetClusterKubeconfigTplOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cluster kubeconfig tpl o k response
func (o *GetClusterKubeconfigTplOK) Code() int {
	return 200
}

func (o *GetClusterKubeconfigTplOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/cluster/{id}/kubeconfigtpl][%d] getClusterKubeconfigTplOK %s", 200, payload)
}

func (o *GetClusterKubeconfigTplOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/cluster/{id}/kubeconfigtpl][%d] getClusterKubeconfigTplOK %s", 200, payload)
}

func (o *GetClusterKubeconfigTplOK) GetPayload() *models.V1ClusterKubeconfigResponse {
	return o.Payload
}

func (o *GetClusterKubeconfigTplOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterKubeconfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterKubeconfigTplDefault creates a GetClusterKubeconfigTplDefault with default headers values
func NewGetClusterKubeconfigTplDefault(code int) *GetClusterKubeconfigTplDefault {
	return &GetClusterKubeconfigTplDefault{
		_statusCode: code,
	}
}

/*
GetClusterKubeconfigTplDefault describes a response with status code -1, with default header values.

Error
*/
type GetClusterKubeconfigTplDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this get cluster kubeconfig tpl default response has a 2xx status code
func (o *GetClusterKubeconfigTplDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster kubeconfig tpl default response has a 3xx status code
func (o *GetClusterKubeconfigTplDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster kubeconfig tpl default response has a 4xx status code
func (o *GetClusterKubeconfigTplDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster kubeconfig tpl default response has a 5xx status code
func (o *GetClusterKubeconfigTplDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster kubeconfig tpl default response a status code equal to that given
func (o *GetClusterKubeconfigTplDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get cluster kubeconfig tpl default response
func (o *GetClusterKubeconfigTplDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterKubeconfigTplDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/cluster/{id}/kubeconfigtpl][%d] getClusterKubeconfigTpl default %s", o._statusCode, payload)
}

func (o *GetClusterKubeconfigTplDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/cluster/{id}/kubeconfigtpl][%d] getClusterKubeconfigTpl default %s", o._statusCode, payload)
}

func (o *GetClusterKubeconfigTplDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetClusterKubeconfigTplDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
