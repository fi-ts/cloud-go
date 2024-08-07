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

// AcceptPostgresRestoreReader is a Reader for the AcceptPostgresRestore structure.
type AcceptPostgresRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcceptPostgresRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAcceptPostgresRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAcceptPostgresRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAcceptPostgresRestoreOK creates a AcceptPostgresRestoreOK with default headers values
func NewAcceptPostgresRestoreOK() *AcceptPostgresRestoreOK {
	return &AcceptPostgresRestoreOK{}
}

/*
AcceptPostgresRestoreOK describes a response with status code 200, with default header values.

OK
*/
type AcceptPostgresRestoreOK struct {
	Payload *models.V1PostgresResponse
}

// IsSuccess returns true when this accept postgres restore o k response has a 2xx status code
func (o *AcceptPostgresRestoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this accept postgres restore o k response has a 3xx status code
func (o *AcceptPostgresRestoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this accept postgres restore o k response has a 4xx status code
func (o *AcceptPostgresRestoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this accept postgres restore o k response has a 5xx status code
func (o *AcceptPostgresRestoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this accept postgres restore o k response a status code equal to that given
func (o *AcceptPostgresRestoreOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the accept postgres restore o k response
func (o *AcceptPostgresRestoreOK) Code() int {
	return 200
}

func (o *AcceptPostgresRestoreOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/database/postgres/restore/{id}][%d] acceptPostgresRestoreOK %s", 200, payload)
}

func (o *AcceptPostgresRestoreOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/database/postgres/restore/{id}][%d] acceptPostgresRestoreOK %s", 200, payload)
}

func (o *AcceptPostgresRestoreOK) GetPayload() *models.V1PostgresResponse {
	return o.Payload
}

func (o *AcceptPostgresRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PostgresResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptPostgresRestoreDefault creates a AcceptPostgresRestoreDefault with default headers values
func NewAcceptPostgresRestoreDefault(code int) *AcceptPostgresRestoreDefault {
	return &AcceptPostgresRestoreDefault{
		_statusCode: code,
	}
}

/*
AcceptPostgresRestoreDefault describes a response with status code -1, with default header values.

Error
*/
type AcceptPostgresRestoreDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this accept postgres restore default response has a 2xx status code
func (o *AcceptPostgresRestoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this accept postgres restore default response has a 3xx status code
func (o *AcceptPostgresRestoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this accept postgres restore default response has a 4xx status code
func (o *AcceptPostgresRestoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this accept postgres restore default response has a 5xx status code
func (o *AcceptPostgresRestoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this accept postgres restore default response a status code equal to that given
func (o *AcceptPostgresRestoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the accept postgres restore default response
func (o *AcceptPostgresRestoreDefault) Code() int {
	return o._statusCode
}

func (o *AcceptPostgresRestoreDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/database/postgres/restore/{id}][%d] acceptPostgresRestore default %s", o._statusCode, payload)
}

func (o *AcceptPostgresRestoreDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/database/postgres/restore/{id}][%d] acceptPostgresRestore default %s", o._statusCode, payload)
}

func (o *AcceptPostgresRestoreDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *AcceptPostgresRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
