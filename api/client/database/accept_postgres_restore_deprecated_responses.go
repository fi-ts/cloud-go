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

// AcceptPostgresRestoreDeprecatedReader is a Reader for the AcceptPostgresRestoreDeprecated structure.
type AcceptPostgresRestoreDeprecatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcceptPostgresRestoreDeprecatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAcceptPostgresRestoreDeprecatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAcceptPostgresRestoreDeprecatedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAcceptPostgresRestoreDeprecatedOK creates a AcceptPostgresRestoreDeprecatedOK with default headers values
func NewAcceptPostgresRestoreDeprecatedOK() *AcceptPostgresRestoreDeprecatedOK {
	return &AcceptPostgresRestoreDeprecatedOK{}
}

/*
AcceptPostgresRestoreDeprecatedOK describes a response with status code 200, with default header values.

OK
*/
type AcceptPostgresRestoreDeprecatedOK struct {
	Payload *models.V1PostgresResponse
}

// IsSuccess returns true when this accept postgres restore deprecated o k response has a 2xx status code
func (o *AcceptPostgresRestoreDeprecatedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this accept postgres restore deprecated o k response has a 3xx status code
func (o *AcceptPostgresRestoreDeprecatedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this accept postgres restore deprecated o k response has a 4xx status code
func (o *AcceptPostgresRestoreDeprecatedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this accept postgres restore deprecated o k response has a 5xx status code
func (o *AcceptPostgresRestoreDeprecatedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this accept postgres restore deprecated o k response a status code equal to that given
func (o *AcceptPostgresRestoreDeprecatedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the accept postgres restore deprecated o k response
func (o *AcceptPostgresRestoreDeprecatedOK) Code() int {
	return 200
}

func (o *AcceptPostgresRestoreDeprecatedOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/database/postgres/restore/{id}][%d] acceptPostgresRestoreDeprecatedOK %s", 200, payload)
}

func (o *AcceptPostgresRestoreDeprecatedOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/database/postgres/restore/{id}][%d] acceptPostgresRestoreDeprecatedOK %s", 200, payload)
}

func (o *AcceptPostgresRestoreDeprecatedOK) GetPayload() *models.V1PostgresResponse {
	return o.Payload
}

func (o *AcceptPostgresRestoreDeprecatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PostgresResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptPostgresRestoreDeprecatedDefault creates a AcceptPostgresRestoreDeprecatedDefault with default headers values
func NewAcceptPostgresRestoreDeprecatedDefault(code int) *AcceptPostgresRestoreDeprecatedDefault {
	return &AcceptPostgresRestoreDeprecatedDefault{
		_statusCode: code,
	}
}

/*
AcceptPostgresRestoreDeprecatedDefault describes a response with status code -1, with default header values.

Error
*/
type AcceptPostgresRestoreDeprecatedDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this accept postgres restore deprecated default response has a 2xx status code
func (o *AcceptPostgresRestoreDeprecatedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this accept postgres restore deprecated default response has a 3xx status code
func (o *AcceptPostgresRestoreDeprecatedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this accept postgres restore deprecated default response has a 4xx status code
func (o *AcceptPostgresRestoreDeprecatedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this accept postgres restore deprecated default response has a 5xx status code
func (o *AcceptPostgresRestoreDeprecatedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this accept postgres restore deprecated default response a status code equal to that given
func (o *AcceptPostgresRestoreDeprecatedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the accept postgres restore deprecated default response
func (o *AcceptPostgresRestoreDeprecatedDefault) Code() int {
	return o._statusCode
}

func (o *AcceptPostgresRestoreDeprecatedDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/database/postgres/restore/{id}][%d] acceptPostgresRestoreDeprecated default %s", o._statusCode, payload)
}

func (o *AcceptPostgresRestoreDeprecatedDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/database/postgres/restore/{id}][%d] acceptPostgresRestoreDeprecated default %s", o._statusCode, payload)
}

func (o *AcceptPostgresRestoreDeprecatedDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *AcceptPostgresRestoreDeprecatedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
