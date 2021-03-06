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

// CreatePostgresBackupConfigReader is a Reader for the CreatePostgresBackupConfig structure.
type CreatePostgresBackupConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePostgresBackupConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePostgresBackupConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreatePostgresBackupConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePostgresBackupConfigCreated creates a CreatePostgresBackupConfigCreated with default headers values
func NewCreatePostgresBackupConfigCreated() *CreatePostgresBackupConfigCreated {
	return &CreatePostgresBackupConfigCreated{}
}

/* CreatePostgresBackupConfigCreated describes a response with status code 201, with default header values.

Created
*/
type CreatePostgresBackupConfigCreated struct {
	Payload *models.V1PostgresBackupConfigResponse
}

func (o *CreatePostgresBackupConfigCreated) Error() string {
	return fmt.Sprintf("[PUT /v1/database/postgres/backup-config][%d] createPostgresBackupConfigCreated  %+v", 201, o.Payload)
}
func (o *CreatePostgresBackupConfigCreated) GetPayload() *models.V1PostgresBackupConfigResponse {
	return o.Payload
}

func (o *CreatePostgresBackupConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PostgresBackupConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePostgresBackupConfigDefault creates a CreatePostgresBackupConfigDefault with default headers values
func NewCreatePostgresBackupConfigDefault(code int) *CreatePostgresBackupConfigDefault {
	return &CreatePostgresBackupConfigDefault{
		_statusCode: code,
	}
}

/* CreatePostgresBackupConfigDefault describes a response with status code -1, with default header values.

Error
*/
type CreatePostgresBackupConfigDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the create postgres backup config default response
func (o *CreatePostgresBackupConfigDefault) Code() int {
	return o._statusCode
}

func (o *CreatePostgresBackupConfigDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/database/postgres/backup-config][%d] createPostgresBackupConfig default  %+v", o._statusCode, o.Payload)
}
func (o *CreatePostgresBackupConfigDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CreatePostgresBackupConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
