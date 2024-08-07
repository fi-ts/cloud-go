// Code generated by go-swagger; DO NOT EDIT.

package database

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

// GetPostgresSecretsReader is a Reader for the GetPostgresSecrets structure.
type GetPostgresSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPostgresSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPostgresSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPostgresSecretsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPostgresSecretsOK creates a GetPostgresSecretsOK with default headers values
func NewGetPostgresSecretsOK() *GetPostgresSecretsOK {
	return &GetPostgresSecretsOK{}
}

/*
GetPostgresSecretsOK describes a response with status code 200, with default header values.

OK
*/
type GetPostgresSecretsOK struct {
	Payload *models.V1PostgresSecretsResponse
}

// IsSuccess returns true when this get postgres secrets o k response has a 2xx status code
func (o *GetPostgresSecretsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get postgres secrets o k response has a 3xx status code
func (o *GetPostgresSecretsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get postgres secrets o k response has a 4xx status code
func (o *GetPostgresSecretsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get postgres secrets o k response has a 5xx status code
func (o *GetPostgresSecretsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get postgres secrets o k response a status code equal to that given
func (o *GetPostgresSecretsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get postgres secrets o k response
func (o *GetPostgresSecretsOK) Code() int {
	return 200
}

func (o *GetPostgresSecretsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/database/postgres/{id}/secrets][%d] getPostgresSecretsOK %s", 200, payload)
}

func (o *GetPostgresSecretsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/database/postgres/{id}/secrets][%d] getPostgresSecretsOK %s", 200, payload)
}

func (o *GetPostgresSecretsOK) GetPayload() *models.V1PostgresSecretsResponse {
	return o.Payload
}

func (o *GetPostgresSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PostgresSecretsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPostgresSecretsDefault creates a GetPostgresSecretsDefault with default headers values
func NewGetPostgresSecretsDefault(code int) *GetPostgresSecretsDefault {
	return &GetPostgresSecretsDefault{
		_statusCode: code,
	}
}

/*
GetPostgresSecretsDefault describes a response with status code -1, with default header values.

Error
*/
type GetPostgresSecretsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this get postgres secrets default response has a 2xx status code
func (o *GetPostgresSecretsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get postgres secrets default response has a 3xx status code
func (o *GetPostgresSecretsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get postgres secrets default response has a 4xx status code
func (o *GetPostgresSecretsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get postgres secrets default response has a 5xx status code
func (o *GetPostgresSecretsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get postgres secrets default response a status code equal to that given
func (o *GetPostgresSecretsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get postgres secrets default response
func (o *GetPostgresSecretsDefault) Code() int {
	return o._statusCode
}

func (o *GetPostgresSecretsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/database/postgres/{id}/secrets][%d] getPostgresSecrets default %s", o._statusCode, payload)
}

func (o *GetPostgresSecretsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/database/postgres/{id}/secrets][%d] getPostgresSecrets default %s", o._statusCode, payload)
}

func (o *GetPostgresSecretsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetPostgresSecretsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
