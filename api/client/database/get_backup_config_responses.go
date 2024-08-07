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

// GetBackupConfigReader is a Reader for the GetBackupConfig structure.
type GetBackupConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBackupConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackupConfigOK creates a GetBackupConfigOK with default headers values
func NewGetBackupConfigOK() *GetBackupConfigOK {
	return &GetBackupConfigOK{}
}

/*
GetBackupConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetBackupConfigOK struct {
	Payload *models.V1PostgresBackupConfigResponse
}

// IsSuccess returns true when this get backup config o k response has a 2xx status code
func (o *GetBackupConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get backup config o k response has a 3xx status code
func (o *GetBackupConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup config o k response has a 4xx status code
func (o *GetBackupConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backup config o k response has a 5xx status code
func (o *GetBackupConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup config o k response a status code equal to that given
func (o *GetBackupConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get backup config o k response
func (o *GetBackupConfigOK) Code() int {
	return 200
}

func (o *GetBackupConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/database/postgres/backup-config/{id}][%d] getBackupConfigOK %s", 200, payload)
}

func (o *GetBackupConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/database/postgres/backup-config/{id}][%d] getBackupConfigOK %s", 200, payload)
}

func (o *GetBackupConfigOK) GetPayload() *models.V1PostgresBackupConfigResponse {
	return o.Payload
}

func (o *GetBackupConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PostgresBackupConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupConfigDefault creates a GetBackupConfigDefault with default headers values
func NewGetBackupConfigDefault(code int) *GetBackupConfigDefault {
	return &GetBackupConfigDefault{
		_statusCode: code,
	}
}

/*
GetBackupConfigDefault describes a response with status code -1, with default header values.

Error
*/
type GetBackupConfigDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this get backup config default response has a 2xx status code
func (o *GetBackupConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get backup config default response has a 3xx status code
func (o *GetBackupConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get backup config default response has a 4xx status code
func (o *GetBackupConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get backup config default response has a 5xx status code
func (o *GetBackupConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get backup config default response a status code equal to that given
func (o *GetBackupConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get backup config default response
func (o *GetBackupConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetBackupConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/database/postgres/backup-config/{id}][%d] getBackupConfig default %s", o._statusCode, payload)
}

func (o *GetBackupConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/database/postgres/backup-config/{id}][%d] getBackupConfig default %s", o._statusCode, payload)
}

func (o *GetBackupConfigDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetBackupConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
