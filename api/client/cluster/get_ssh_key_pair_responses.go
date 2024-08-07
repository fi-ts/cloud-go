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

// GetSSHKeyPairReader is a Reader for the GetSSHKeyPair structure.
type GetSSHKeyPairReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSSHKeyPairReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSSHKeyPairOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSSHKeyPairDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSSHKeyPairOK creates a GetSSHKeyPairOK with default headers values
func NewGetSSHKeyPairOK() *GetSSHKeyPairOK {
	return &GetSSHKeyPairOK{}
}

/*
GetSSHKeyPairOK describes a response with status code 200, with default header values.

OK
*/
type GetSSHKeyPairOK struct {
	Payload *models.V1ClusterCredentialsResponse
}

// IsSuccess returns true when this get Ssh key pair o k response has a 2xx status code
func (o *GetSSHKeyPairOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Ssh key pair o k response has a 3xx status code
func (o *GetSSHKeyPairOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Ssh key pair o k response has a 4xx status code
func (o *GetSSHKeyPairOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Ssh key pair o k response has a 5xx status code
func (o *GetSSHKeyPairOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Ssh key pair o k response a status code equal to that given
func (o *GetSSHKeyPairOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Ssh key pair o k response
func (o *GetSSHKeyPairOK) Code() int {
	return 200
}

func (o *GetSSHKeyPairOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/cluster/{id}/sshkeypair][%d] getSshKeyPairOK %s", 200, payload)
}

func (o *GetSSHKeyPairOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/cluster/{id}/sshkeypair][%d] getSshKeyPairOK %s", 200, payload)
}

func (o *GetSSHKeyPairOK) GetPayload() *models.V1ClusterCredentialsResponse {
	return o.Payload
}

func (o *GetSSHKeyPairOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterCredentialsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSSHKeyPairDefault creates a GetSSHKeyPairDefault with default headers values
func NewGetSSHKeyPairDefault(code int) *GetSSHKeyPairDefault {
	return &GetSSHKeyPairDefault{
		_statusCode: code,
	}
}

/*
GetSSHKeyPairDefault describes a response with status code -1, with default header values.

Error
*/
type GetSSHKeyPairDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this get SSH key pair default response has a 2xx status code
func (o *GetSSHKeyPairDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get SSH key pair default response has a 3xx status code
func (o *GetSSHKeyPairDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get SSH key pair default response has a 4xx status code
func (o *GetSSHKeyPairDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get SSH key pair default response has a 5xx status code
func (o *GetSSHKeyPairDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get SSH key pair default response a status code equal to that given
func (o *GetSSHKeyPairDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get SSH key pair default response
func (o *GetSSHKeyPairDefault) Code() int {
	return o._statusCode
}

func (o *GetSSHKeyPairDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/cluster/{id}/sshkeypair][%d] getSSHKeyPair default %s", o._statusCode, payload)
}

func (o *GetSSHKeyPairDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/cluster/{id}/sshkeypair][%d] getSSHKeyPair default %s", o._statusCode, payload)
}

func (o *GetSSHKeyPairDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetSSHKeyPairDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
