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

// CreatePostgresStandbyReader is a Reader for the CreatePostgresStandby structure.
type CreatePostgresStandbyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePostgresStandbyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePostgresStandbyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreatePostgresStandbyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePostgresStandbyCreated creates a CreatePostgresStandbyCreated with default headers values
func NewCreatePostgresStandbyCreated() *CreatePostgresStandbyCreated {
	return &CreatePostgresStandbyCreated{}
}

/* CreatePostgresStandbyCreated describes a response with status code 201, with default header values.

Created
*/
type CreatePostgresStandbyCreated struct {
	Payload *models.V1PostgresResponse
}

func (o *CreatePostgresStandbyCreated) Error() string {
	return fmt.Sprintf("[PUT /v1/database/postgres/standby][%d] createPostgresStandbyCreated  %+v", 201, o.Payload)
}
func (o *CreatePostgresStandbyCreated) GetPayload() *models.V1PostgresResponse {
	return o.Payload
}

func (o *CreatePostgresStandbyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PostgresResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePostgresStandbyDefault creates a CreatePostgresStandbyDefault with default headers values
func NewCreatePostgresStandbyDefault(code int) *CreatePostgresStandbyDefault {
	return &CreatePostgresStandbyDefault{
		_statusCode: code,
	}
}

/* CreatePostgresStandbyDefault describes a response with status code -1, with default header values.

Error
*/
type CreatePostgresStandbyDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the create postgres standby default response
func (o *CreatePostgresStandbyDefault) Code() int {
	return o._statusCode
}

func (o *CreatePostgresStandbyDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/database/postgres/standby][%d] createPostgresStandby default  %+v", o._statusCode, o.Payload)
}
func (o *CreatePostgresStandbyDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CreatePostgresStandbyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
