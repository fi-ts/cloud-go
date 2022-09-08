// Code generated by go-swagger; DO NOT EDIT.

package database

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

// DeletePostgresBackupConfigReader is a Reader for the DeletePostgresBackupConfig structure.
type DeletePostgresBackupConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePostgresBackupConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePostgresBackupConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeletePostgresBackupConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePostgresBackupConfigOK creates a DeletePostgresBackupConfigOK with default headers values
func NewDeletePostgresBackupConfigOK() *DeletePostgresBackupConfigOK {
	return &DeletePostgresBackupConfigOK{}
}

/*
DeletePostgresBackupConfigOK describes a response with status code 200, with default header values.

OK
*/
type DeletePostgresBackupConfigOK struct {
	Payload *models.V1PostgresBackupConfigResponse
}

// IsSuccess returns true when this delete postgres backup config o k response has a 2xx status code
func (o *DeletePostgresBackupConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete postgres backup config o k response has a 3xx status code
func (o *DeletePostgresBackupConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete postgres backup config o k response has a 4xx status code
func (o *DeletePostgresBackupConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete postgres backup config o k response has a 5xx status code
func (o *DeletePostgresBackupConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete postgres backup config o k response a status code equal to that given
func (o *DeletePostgresBackupConfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeletePostgresBackupConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/database/postgres/backup-config/{id}][%d] deletePostgresBackupConfigOK  %+v", 200, o.Payload)
}

func (o *DeletePostgresBackupConfigOK) String() string {
	return fmt.Sprintf("[DELETE /v1/database/postgres/backup-config/{id}][%d] deletePostgresBackupConfigOK  %+v", 200, o.Payload)
}

func (o *DeletePostgresBackupConfigOK) GetPayload() *models.V1PostgresBackupConfigResponse {
	return o.Payload
}

func (o *DeletePostgresBackupConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PostgresBackupConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePostgresBackupConfigDefault creates a DeletePostgresBackupConfigDefault with default headers values
func NewDeletePostgresBackupConfigDefault(code int) *DeletePostgresBackupConfigDefault {
	return &DeletePostgresBackupConfigDefault{
		_statusCode: code,
	}
}

/*
DeletePostgresBackupConfigDefault describes a response with status code -1, with default header values.

Error
*/
type DeletePostgresBackupConfigDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the delete postgres backup config default response
func (o *DeletePostgresBackupConfigDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete postgres backup config default response has a 2xx status code
func (o *DeletePostgresBackupConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete postgres backup config default response has a 3xx status code
func (o *DeletePostgresBackupConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete postgres backup config default response has a 4xx status code
func (o *DeletePostgresBackupConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete postgres backup config default response has a 5xx status code
func (o *DeletePostgresBackupConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete postgres backup config default response a status code equal to that given
func (o *DeletePostgresBackupConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeletePostgresBackupConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/database/postgres/backup-config/{id}][%d] deletePostgresBackupConfig default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePostgresBackupConfigDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/database/postgres/backup-config/{id}][%d] deletePostgresBackupConfig default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePostgresBackupConfigDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *DeletePostgresBackupConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
