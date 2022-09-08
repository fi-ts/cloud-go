// Code generated by go-swagger; DO NOT EDIT.

package s3

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

// Gets3Reader is a Reader for the Gets3 structure.
type Gets3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Gets3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGets3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGets3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGets3OK creates a Gets3OK with default headers values
func NewGets3OK() *Gets3OK {
	return &Gets3OK{}
}

/*
Gets3OK describes a response with status code 200, with default header values.

OK
*/
type Gets3OK struct {
	Payload *models.V1S3CredentialsResponse
}

// IsSuccess returns true when this gets3 o k response has a 2xx status code
func (o *Gets3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gets3 o k response has a 3xx status code
func (o *Gets3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gets3 o k response has a 4xx status code
func (o *Gets3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this gets3 o k response has a 5xx status code
func (o *Gets3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this gets3 o k response a status code equal to that given
func (o *Gets3OK) IsCode(code int) bool {
	return code == 200
}

func (o *Gets3OK) Error() string {
	return fmt.Sprintf("[GET /v1/s3][%d] gets3OK  %+v", 200, o.Payload)
}

func (o *Gets3OK) String() string {
	return fmt.Sprintf("[GET /v1/s3][%d] gets3OK  %+v", 200, o.Payload)
}

func (o *Gets3OK) GetPayload() *models.V1S3CredentialsResponse {
	return o.Payload
}

func (o *Gets3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1S3CredentialsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGets3Default creates a Gets3Default with default headers values
func NewGets3Default(code int) *Gets3Default {
	return &Gets3Default{
		_statusCode: code,
	}
}

/*
Gets3Default describes a response with status code -1, with default header values.

Error
*/
type Gets3Default struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the gets3 default response
func (o *Gets3Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this gets3 default response has a 2xx status code
func (o *Gets3Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this gets3 default response has a 3xx status code
func (o *Gets3Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this gets3 default response has a 4xx status code
func (o *Gets3Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this gets3 default response has a 5xx status code
func (o *Gets3Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this gets3 default response a status code equal to that given
func (o *Gets3Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *Gets3Default) Error() string {
	return fmt.Sprintf("[GET /v1/s3][%d] gets3 default  %+v", o._statusCode, o.Payload)
}

func (o *Gets3Default) String() string {
	return fmt.Sprintf("[GET /v1/s3][%d] gets3 default  %+v", o._statusCode, o.Payload)
}

func (o *Gets3Default) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *Gets3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
