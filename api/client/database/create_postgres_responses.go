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

// CreatePostgresReader is a Reader for the CreatePostgres structure.
type CreatePostgresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePostgresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePostgresCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreatePostgresConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreatePostgresDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePostgresCreated creates a CreatePostgresCreated with default headers values
func NewCreatePostgresCreated() *CreatePostgresCreated {
	return &CreatePostgresCreated{}
}

/* CreatePostgresCreated describes a response with status code 201, with default header values.

Created
*/
type CreatePostgresCreated struct {
	Payload *models.V1PostgresResponse
}

func (o *CreatePostgresCreated) Error() string {
	return fmt.Sprintf("[PUT /v1/database/postgres][%d] createPostgresCreated  %+v", 201, o.Payload)
}
func (o *CreatePostgresCreated) GetPayload() *models.V1PostgresResponse {
	return o.Payload
}

func (o *CreatePostgresCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PostgresResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePostgresConflict creates a CreatePostgresConflict with default headers values
func NewCreatePostgresConflict() *CreatePostgresConflict {
	return &CreatePostgresConflict{}
}

/* CreatePostgresConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreatePostgresConflict struct {
	Payload *httperrors.HTTPErrorResponse
}

func (o *CreatePostgresConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/database/postgres][%d] createPostgresConflict  %+v", 409, o.Payload)
}
func (o *CreatePostgresConflict) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CreatePostgresConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePostgresDefault creates a CreatePostgresDefault with default headers values
func NewCreatePostgresDefault(code int) *CreatePostgresDefault {
	return &CreatePostgresDefault{
		_statusCode: code,
	}
}

/* CreatePostgresDefault describes a response with status code -1, with default header values.

Error
*/
type CreatePostgresDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the create postgres default response
func (o *CreatePostgresDefault) Code() int {
	return o._statusCode
}

func (o *CreatePostgresDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/database/postgres][%d] createPostgres default  %+v", o._statusCode, o.Payload)
}
func (o *CreatePostgresDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CreatePostgresDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
