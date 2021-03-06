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

// DeletePostgresReader is a Reader for the DeletePostgres structure.
type DeletePostgresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePostgresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePostgresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeletePostgresDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePostgresOK creates a DeletePostgresOK with default headers values
func NewDeletePostgresOK() *DeletePostgresOK {
	return &DeletePostgresOK{}
}

/* DeletePostgresOK describes a response with status code 200, with default header values.

OK
*/
type DeletePostgresOK struct {
	Payload *models.V1PostgresResponse
}

func (o *DeletePostgresOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/database/postgres/{id}][%d] deletePostgresOK  %+v", 200, o.Payload)
}
func (o *DeletePostgresOK) GetPayload() *models.V1PostgresResponse {
	return o.Payload
}

func (o *DeletePostgresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PostgresResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePostgresDefault creates a DeletePostgresDefault with default headers values
func NewDeletePostgresDefault(code int) *DeletePostgresDefault {
	return &DeletePostgresDefault{
		_statusCode: code,
	}
}

/* DeletePostgresDefault describes a response with status code -1, with default header values.

Error
*/
type DeletePostgresDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the delete postgres default response
func (o *DeletePostgresDefault) Code() int {
	return o._statusCode
}

func (o *DeletePostgresDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/database/postgres/{id}][%d] deletePostgres default  %+v", o._statusCode, o.Payload)
}
func (o *DeletePostgresDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *DeletePostgresDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
