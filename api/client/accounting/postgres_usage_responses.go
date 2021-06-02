// Code generated by go-swagger; DO NOT EDIT.

package accounting

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

// PostgresUsageReader is a Reader for the PostgresUsage structure.
type PostgresUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostgresUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostgresUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostgresUsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostgresUsageOK creates a PostgresUsageOK with default headers values
func NewPostgresUsageOK() *PostgresUsageOK {
	return &PostgresUsageOK{}
}

/* PostgresUsageOK describes a response with status code 200, with default header values.

OK
*/
type PostgresUsageOK struct {
	Payload *models.V1PostgresUsageResponse
}

func (o *PostgresUsageOK) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/postgres-usage][%d] postgresUsageOK  %+v", 200, o.Payload)
}
func (o *PostgresUsageOK) GetPayload() *models.V1PostgresUsageResponse {
	return o.Payload
}

func (o *PostgresUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PostgresUsageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostgresUsageDefault creates a PostgresUsageDefault with default headers values
func NewPostgresUsageDefault(code int) *PostgresUsageDefault {
	return &PostgresUsageDefault{
		_statusCode: code,
	}
}

/* PostgresUsageDefault describes a response with status code -1, with default header values.

Error
*/
type PostgresUsageDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the postgres usage default response
func (o *PostgresUsageDefault) Code() int {
	return o._statusCode
}

func (o *PostgresUsageDefault) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/postgres-usage][%d] postgresUsage default  %+v", o._statusCode, o.Payload)
}
func (o *PostgresUsageDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *PostgresUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
