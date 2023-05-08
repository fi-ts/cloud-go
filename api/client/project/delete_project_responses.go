// Code generated by go-swagger; DO NOT EDIT.

package project

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

// DeleteProjectReader is a Reader for the DeleteProject structure.
type DeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProjectOK creates a DeleteProjectOK with default headers values
func NewDeleteProjectOK() *DeleteProjectOK {
	return &DeleteProjectOK{}
}

/*
DeleteProjectOK describes a response with status code 200, with default header values.

OK
*/
type DeleteProjectOK struct {
	Payload *models.V1ProjectResponse
}

// IsSuccess returns true when this delete project o k response has a 2xx status code
func (o *DeleteProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project o k response has a 3xx status code
func (o *DeleteProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project o k response has a 4xx status code
func (o *DeleteProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project o k response has a 5xx status code
func (o *DeleteProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project o k response a status code equal to that given
func (o *DeleteProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete project o k response
func (o *DeleteProjectOK) Code() int {
	return 200
}

func (o *DeleteProjectOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/project/{id}][%d] deleteProjectOK  %+v", 200, o.Payload)
}

func (o *DeleteProjectOK) String() string {
	return fmt.Sprintf("[DELETE /v1/project/{id}][%d] deleteProjectOK  %+v", 200, o.Payload)
}

func (o *DeleteProjectOK) GetPayload() *models.V1ProjectResponse {
	return o.Payload
}

func (o *DeleteProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ProjectResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectDefault creates a DeleteProjectDefault with default headers values
func NewDeleteProjectDefault(code int) *DeleteProjectDefault {
	return &DeleteProjectDefault{
		_statusCode: code,
	}
}

/*
DeleteProjectDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteProjectDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this delete project default response has a 2xx status code
func (o *DeleteProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete project default response has a 3xx status code
func (o *DeleteProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete project default response has a 4xx status code
func (o *DeleteProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete project default response has a 5xx status code
func (o *DeleteProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete project default response a status code equal to that given
func (o *DeleteProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete project default response
func (o *DeleteProjectDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProjectDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/project/{id}][%d] deleteProject default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/project/{id}][%d] deleteProject default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *DeleteProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
