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

// ListPostgresBackupConfigsReader is a Reader for the ListPostgresBackupConfigs structure.
type ListPostgresBackupConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPostgresBackupConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPostgresBackupConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListPostgresBackupConfigsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPostgresBackupConfigsOK creates a ListPostgresBackupConfigsOK with default headers values
func NewListPostgresBackupConfigsOK() *ListPostgresBackupConfigsOK {
	return &ListPostgresBackupConfigsOK{}
}

/* ListPostgresBackupConfigsOK describes a response with status code 200, with default header values.

OK
*/
type ListPostgresBackupConfigsOK struct {
	Payload []*models.V1PostgresBackupConfigResponse
}

func (o *ListPostgresBackupConfigsOK) Error() string {
	return fmt.Sprintf("[GET /v1/database/postgres/backup-config][%d] listPostgresBackupConfigsOK  %+v", 200, o.Payload)
}
func (o *ListPostgresBackupConfigsOK) GetPayload() []*models.V1PostgresBackupConfigResponse {
	return o.Payload
}

func (o *ListPostgresBackupConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPostgresBackupConfigsDefault creates a ListPostgresBackupConfigsDefault with default headers values
func NewListPostgresBackupConfigsDefault(code int) *ListPostgresBackupConfigsDefault {
	return &ListPostgresBackupConfigsDefault{
		_statusCode: code,
	}
}

/* ListPostgresBackupConfigsDefault describes a response with status code -1, with default header values.

Error
*/
type ListPostgresBackupConfigsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the list postgres backup configs default response
func (o *ListPostgresBackupConfigsDefault) Code() int {
	return o._statusCode
}

func (o *ListPostgresBackupConfigsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/database/postgres/backup-config][%d] listPostgresBackupConfigs default  %+v", o._statusCode, o.Payload)
}
func (o *ListPostgresBackupConfigsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ListPostgresBackupConfigsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
