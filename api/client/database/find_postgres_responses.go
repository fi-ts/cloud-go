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

// FindPostgresReader is a Reader for the FindPostgres structure.
type FindPostgresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPostgresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindPostgresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindPostgresDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindPostgresOK creates a FindPostgresOK with default headers values
func NewFindPostgresOK() *FindPostgresOK {
	return &FindPostgresOK{}
}

/*
FindPostgresOK describes a response with status code 200, with default header values.

OK
*/
type FindPostgresOK struct {
	Payload []*models.V1PostgresResponse
}

// IsSuccess returns true when this find postgres o k response has a 2xx status code
func (o *FindPostgresOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find postgres o k response has a 3xx status code
func (o *FindPostgresOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find postgres o k response has a 4xx status code
func (o *FindPostgresOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find postgres o k response has a 5xx status code
func (o *FindPostgresOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find postgres o k response a status code equal to that given
func (o *FindPostgresOK) IsCode(code int) bool {
	return code == 200
}

func (o *FindPostgresOK) Error() string {
	return fmt.Sprintf("[POST /v1/database/postgres/find][%d] findPostgresOK  %+v", 200, o.Payload)
}

func (o *FindPostgresOK) String() string {
	return fmt.Sprintf("[POST /v1/database/postgres/find][%d] findPostgresOK  %+v", 200, o.Payload)
}

func (o *FindPostgresOK) GetPayload() []*models.V1PostgresResponse {
	return o.Payload
}

func (o *FindPostgresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindPostgresDefault creates a FindPostgresDefault with default headers values
func NewFindPostgresDefault(code int) *FindPostgresDefault {
	return &FindPostgresDefault{
		_statusCode: code,
	}
}

/*
FindPostgresDefault describes a response with status code -1, with default header values.

Error
*/
type FindPostgresDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the find postgres default response
func (o *FindPostgresDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this find postgres default response has a 2xx status code
func (o *FindPostgresDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find postgres default response has a 3xx status code
func (o *FindPostgresDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find postgres default response has a 4xx status code
func (o *FindPostgresDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find postgres default response has a 5xx status code
func (o *FindPostgresDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find postgres default response a status code equal to that given
func (o *FindPostgresDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FindPostgresDefault) Error() string {
	return fmt.Sprintf("[POST /v1/database/postgres/find][%d] findPostgres default  %+v", o._statusCode, o.Payload)
}

func (o *FindPostgresDefault) String() string {
	return fmt.Sprintf("[POST /v1/database/postgres/find][%d] findPostgres default  %+v", o._statusCode, o.Payload)
}

func (o *FindPostgresDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindPostgresDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
