// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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

func (o *GetClusterKubeconfigTplOK) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/{id}/kubeconfigtpl][%d] getClusterKubeconfigTplOK  %+v", 200, o.Payload)
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

// Code gets the status code for the get cluster kubeconfig tpl default response
func (o *GetClusterKubeconfigTplDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterKubeconfigTplDefault) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/{id}/kubeconfigtpl][%d] getClusterKubeconfigTpl default  %+v", o._statusCode, o.Payload)
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
